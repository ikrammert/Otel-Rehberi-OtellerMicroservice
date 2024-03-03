package controllers

import (
	"context"
	"fmt"
	"net/http"
	"oteller-microservice/database"
	"oteller-microservice/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var raporCollection *mongo.Collection = database.OpenCollection(database.Client, "rapors")

func CreateRaporByKonum() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("CreateRaporByKonum")
		konum := c.Param("konum")

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var rapor models.Rapor

		rapor.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		rapor.UUID = primitive.NewObjectID()
		rapor.Rapor_durumu = "Hazırlanıyor"
		rapor.Konum = konum

		// RabbitMQ'ya bağlanma
		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		if err != nil {
			msg := fmt.Sprintf("RabbitMQ sunucusuna bağlanılamadı: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer conn.Close()

		ch, err := conn.Channel()
		if err != nil {
			msg := fmt.Sprintf("Kanal oluşturulamadı: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer ch.Close()

		kuyruk, err := ch.QueueDeclare(
			"rapor_kuyruk",
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			msg := fmt.Sprintf("Kuyruk oluşturulamadı: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		// Rapor talebini RabbitMQ'ya gönderme
		err = ch.Publish(
			"",          // Exchange
			kuyruk.Name, // Kuyruk adı
			false,       // Mandatory
			false,       // Immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(rapor.UUID.Hex() + "**" + konum),
			},
		)
		if err != nil {
			msg := fmt.Sprintf("Mesaj gönderilemedi: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		result, insertErr := raporCollection.InsertOne(ctx, rapor)
		if insertErr != nil {
			msg := fmt.Sprintf("Rapor İsteği Eklenemedi")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		response := map[string]interface{}{
			"message":    "Rapor talebi başarıyla gönderildi",
			"rapor_id":   rapor.UUID.Hex(),
			"InsertedID": result.InsertedID,
		}

		// Başarılı bir yanıt dönme
		c.JSON(http.StatusOK, response)
	}
}

func ListRapors() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("ListRapors")
		cursor, err := raporCollection.Find(context.Background(), bson.M{}, options.Find().SetProjection(bson.M{"rapor_durumu": 1, "konum": 1, "uuid": 1, "created_at": 1}))

		if err != nil {
			msg := fmt.Sprintf("Raporlar alınamadı: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cursor.Close(context.Background())

		var raporlar []models.Rapor
		if err := cursor.All(context.Background(), &raporlar); err != nil {
			msg := fmt.Sprintf("Raporlar decode edilemedi: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, raporlar)
	}
}

func GetRaporById() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("GetRaporById")
		raporID := c.Param("rapor_id")

		raporObjId, err := primitive.ObjectIDFromHex(raporID)
		if err != nil {
			msg := fmt.Sprintf("ObjectID'ye dönüştürme hatası: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		filter := bson.M{"uuid": raporObjId}

		var rapor models.Rapor

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = raporCollection.FindOne(ctx, filter).Decode(&rapor)
		if err != nil {
			msg := fmt.Sprintf("Rapor bulunamadı: %+v", err)
			c.JSON(http.StatusNotFound, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, rapor)
	}
}
