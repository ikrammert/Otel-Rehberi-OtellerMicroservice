package controllers

import (
	"context"
	"fmt"
	"net/http"
	"oteller-microservice/database"
	"oteller-microservice/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var raporCollection *mongo.Collection = database.OpenCollection(database.Client, "rapors")

func CreateRaporByKonum() gin.HandlerFunc {
	return func(c *gin.Context) {
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

	}
}

func GetRaporById() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
