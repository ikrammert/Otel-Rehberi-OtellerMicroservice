package controllers

import (
	"context"
	"net/http"
	"time"

	"oteller-microservice/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateCommInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("CreateConnInfo")
		otelID := c.Param("otel_id")
		connID := c.Param("conn_id")

		// İstekte gelen iletişim bilgilerini al
		var iletişim models.IletisimBilgisi
		iletişim.Conn_id = connID
		if err := c.BindJSON(&iletişim); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Belirli bir oteli bulmak için kullanılacak filtre
		filter := bson.M{"otel_id": otelID}

		// İletişim bilgisini eklemek için kullanılacak güncelleme
		update := bson.M{
			"$push": bson.M{"iletisim_bilgisi": iletişim},
		}

		// Belirli bir oteli güncelle
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		result, err := otelCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "İletişim bilgisi eklenemedi"})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func DeleteCommInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("DeleteCommInfo")
		otelID := c.Param("otel_id")
		connID := c.Param("conn_id") // Silinecek olan iletişim bilgisinin benzersiz kimliği

		// Belirli bir oteli bulmak için kullanılacak filtre
		filter := bson.M{"otel_id": otelID}

		// İletişim bilgisini silmek için kullanılacak güncelleme
		update := bson.M{
			"$pull": bson.M{"iletisim_bilgisi": bson.M{"conn_id": connID}},
		}

		// Belirli bir oteli güncelle
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		result, err := otelCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "İletişim bilgisi kaldırılamadı"})
			return
		}

		if result.ModifiedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "İletişim bilgisi bulunamadı"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "İletişim bilgisi başarıyla kaldırıldı"})
	}
}
