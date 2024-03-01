package routes

import (
	controller "oteller-microservice/controllers"

	"github.com/gin-gonic/gin"
)

func OtelRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/otels", controller.CreateOtel()) // Otel Oluşturma

	incomingRoutes.DELETE("/otels/:otel_id", controller.DeleteOtel()) //Otel Kaldırma

	incomingRoutes.GET("/otels/:otel_id", controller.GetOtel()) //otel bilgilerinin getirilmesi

	incomingRoutes.GET("/otels/:otel_id/owners", controller.GetOwners()) // Otel yetkililerinin listelenmesi

}
