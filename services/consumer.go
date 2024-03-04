package services

import (
	"context"
	"log"
	"strconv"
	"strings"

	"oteller-microservice/database"
	"oteller-microservice/models"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var raporCollection *mongo.Collection = database.OpenCollection(database.Client, "rapors")
var otelCollection *mongo.Collection = database.OpenCollection(database.Client, "otel")

func StartRabbitMQWorker() {
	//RabbitMQ Sunucumuza bağlanıyoruz
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Printf("RabbitMQ bağlantı hatası %+v", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Print(err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"rapor_kuyruk", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		log.Fatalln(err)
	}

	msgs, err := ch.Consume(
		"rapor_kuyruk", //
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	if err != nil {
		log.Print(err)
	}
	//Burada goroutine ile çalışan fonksiyonumuz
	//çalışırken programın kapanmaması için
	//kanal oluşturduk
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Alınan mesaj: %s", d.Body)
			IdAndKonum := strings.Split(string(d.Body), "**")
			raporId := IdAndKonum[0]
			konum := IdAndKonum[1]

			// otelSayisi, err := otelCollection.CountDocuments(context.Background(), bson.M{"konum": konum})
			// if err != nil {
			// 	log.Printf("Otel sayısı alınamadı: %v", err)
			// 	continue
			// }

			cursor, err := otelCollection.Find(context.Background(), bson.M{"konum": konum})
			if err != nil {
				log.Printf("Oteller bulunamadı: %v", err)
				continue
			}

			otelSayisi := 0
			telefonSayisi := 0

			for cursor.Next(context.Background()) {
				var otel models.Otel
				if err := cursor.Decode(&otel); err != nil {
					log.Printf("Otel bilgileri alınamadı: %v", err)
					continue
				}

				// Her bir otelin iletişim bilgilerini kontrol et
				for _, iletisimBilgisi := range otel.Iletisim_bilgisi {
					if iletisimBilgisi.Bilgi_tipi == "Telefon" {
						telefonSayisi++
					}
				}
				// Otel sayısını arttır
				otelSayisi++
			}

			log.Printf("\nOtel sayısı: %+v, numara Sayısı:%+v", otelSayisi, telefonSayisi)
			log.Printf("filter id: %+v", raporId)
			// Güncelleme işlemi

			filterObjID, err := primitive.ObjectIDFromHex(raporId)
			if err != nil {
				log.Printf("ObjectID'ye dönüştürme hatası: %v", err)
				continue
			}
			filter := bson.M{"uuid": filterObjID}
			update := bson.D{{
				Key: "$set",
				Value: bson.M{
					"otel_sayisi":   strconv.Itoa(otelSayisi),
					"numara_sayisi": strconv.Itoa(telefonSayisi),
					"rapor_durumu":  "Tamamlandı",
				},
			}}

			upsert := false
			opt := options.UpdateOptions{
				Upsert: &upsert,
			}
			result, err := raporCollection.UpdateOne(context.Background(), filter, update, &opt)
			if err != nil {
				log.Printf("Rapor güncellenemedi: %v", err)
				continue
			}

			if result.ModifiedCount == 0 {
				log.Printf("Güncellenecek rapor bulunamadı: %v", err)
				continue
			}

			cursor.Close(context.Background()) // cursor'ı kapat
		}
	}()

	log.Printf(" rapor istekleri dinleniyor...")

	//program kapanmayacak ve sürekli olarak kuyruktaki mesajları çekecektir.
	<-forever
}
