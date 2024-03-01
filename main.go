package main

import (
	"os"

	"oteller-microservice/routes"
	"oteller-microservice/services"

	"github.com/gin-gonic/gin"
)

func main() {
	//Servis
	go services.StartRabbitMQWorker() // Goroutine olarak çağrılıyor

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
