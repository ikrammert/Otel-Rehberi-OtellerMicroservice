package main

import (
	"log"
	"os"
	"path/filepath"

	"oteller-microservice/routes"
	"oteller-microservice/services"

	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"

	"github.com/gin-gonic/gin"
)

func main() {
	//Add log
	forLogs()
	//Servis
	logrus.Info("Servis Basliyor")
	go services.StartRabbitMQWorker() // Goroutine olarak çağrılıyor

	port := os.Getenv("PORT")
	if port == "" {
		port = "8182"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.OtelRoutes(router)
	routes.CommunicationRoutes(router)
	routes.RaporRoutes(router)

	router.Run(":" + port)

}

func forLogs() {
	logrus.SetFormatter(&ecslogrus.Formatter{})
	logrus.SetLevel(logrus.TraceLevel)

	logFilePath := "./logs/go.log"
	// Klasörün varlığını kontrol ediyoruz
	dir := filepath.Dir(logFilePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// Klasör yoksa, klasörü oluştur
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("Klasör oluşturulamadı: %v", err)
		}
	}
	file, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Panicf("\nHata: %+v", err)
	} else {
		log.SetOutput(file)
	}
	defer file.Close()
}
