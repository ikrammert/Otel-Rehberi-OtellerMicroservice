package main

import (
	"io"
	"log"
	"os"

	"oteller-microservice/routes"
	"oteller-microservice/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func init() {
	//Logging (ELK) - Set For Format
	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyLevel: "severity",
		},
	})
	logrus.SetLevel(logrus.TraceLevel)

	//File Writer for logs
	filePath := "./elk-stack/ingest_data/out.log"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Panicf("open file error: %v", err)
	}
	// write log to terminal and file
	logrus.SetOutput(io.MultiWriter(file, os.Stdout))
	defer file.Close()
}

func main() {
	//Servis
	go services.StartRabbitMQWorker() // Goroutine olarak çağrılıyor

	// logging
	refId := logrus.Fields{
		"ref-id": uuid.NewString(), //For main
	}

	logrus.WithFields(refId).Info("Logging Başlıyor...")
	// logrus.WithFields(refId).Warn("Uyarı Bildirimi")
	// logrus.WithFields(refId).Error("Logging Error!")

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.OtelRoutes(router)
	routes.CommunicationRoutes(router)
	routes.RaporRoutes(router)

	router.Run(":" + port)

}
