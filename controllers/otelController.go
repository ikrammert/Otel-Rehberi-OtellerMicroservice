package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"oteller-microservice/database"
	"oteller-microservice/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var otelCollection *mongo.Collection = database.OpenCollection(database.Client, "otel")

func CreateOtel() gin.HandlerFunc {
	return func(c *gin.Context) {
		var otel models.Otel

		if err := c.BindJSON(&otel); err != nil {
			logrus.WithFields(RefId).Errorf("error:%v", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		otel.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		otel.UUID = primitive.NewObjectID()

		// UUID değerini hex formatında alarak otel ID'sini oluştur
		otel.Otel_id = otel.UUID.Hex()

		result, insertErr := otelCollection.InsertOne(ctx, otel)
		if insertErr != nil {
			msg := fmt.Sprintf("Yeni Otel Eklenemedi")
			logrus.WithFields(RefId).Error(msg)
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		response := map[string]interface{}{
			"InsertedID": result.InsertedID,
			"OtelID":     otel.Otel_id,
		}

		logrus.WithFields(RefId).Info("İşlem Başarılı @CreateOtel")
		c.JSON(http.StatusOK, response)
	}
}

func DeleteOtel() gin.HandlerFunc {
	return func(c *gin.Context) {
		otelID := c.Param("otel_id")

		// Otel belgesini silmek için kullanılacak filtre
		filter := bson.M{"otel_id": otelID}

		// Belirli bir oteli sil
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		result, err := otelCollection.DeleteOne(ctx, filter)
		if err != nil {
			logrus.WithFields(RefId).Errorf("Otel Silinemedi")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Otel silinemedi"})
			return
		}

		if result.DeletedCount == 0 {
			logrus.WithFields(RefId).Error("Otel Bulunamadı")
			c.JSON(http.StatusNotFound, gin.H{"error": "Otel bulunamadı"})
			return
		}

		logrus.WithFields(RefId).Info("İşlem Başarılı @DeleteOtel")
		c.JSON(http.StatusOK, gin.H{"message": "Otel başarıyla silindi"})
	}
}

func GetOtel() gin.HandlerFunc {
	return func(c *gin.Context) {
		otelID := c.Param("otel_id")

		// Belirli bir oteli bulmak için kullanılacak filtre
		filter := bson.M{"otel_id": otelID}

		var otel models.Otel

		log.Print(otelID)

		// Belirli bir oteli veritabanından getir
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := otelCollection.FindOne(ctx, filter).Decode(&otel)
		if err != nil {
			logrus.WithFields(RefId).Error("Otel Bulunamadı")
			c.JSON(http.StatusNotFound, gin.H{"error": "Otel bulunamadı"})
			return
		}

		logrus.WithFields(RefId).Info("İşlem Başarılı @GetOtel")
		c.JSON(http.StatusOK, otel)
	}
}

func GetOwners() gin.HandlerFunc {
	return func(c *gin.Context) {
		otelID := c.Param("otel_id")

		// Belirli bir otelin yetkililerini bulmak için kullanılacak filtre
		filter := bson.M{"otel_id": otelID}

		var otel models.Otel

		// Belirli bir oteli veritabanından getir
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := otelCollection.FindOne(ctx, filter).Decode(&otel)
		if err != nil {
			logrus.WithFields(RefId).Error("Otel Bulunamadı")
			c.JSON(http.StatusNotFound, gin.H{"error": "Otel bulunamadı"})
			return
		}

		logrus.WithFields(RefId).Info("İşlem Başarılı @GetOwners")
		// Otelin yetkililerini istemciye döndür
		c.JSON(http.StatusOK, otel.Yetkililer)
	}
}
