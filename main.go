package main

import (
	"os"

	"oteller-microservice/database"
	"oteller-microservice/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var otelCollection *mongo.Collection = database.OpenCollection(database.Client, "otel")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.OtelRoutes(router)
	routes.CommunicationRoutes(router)

	router.Run(":" + port)
}
