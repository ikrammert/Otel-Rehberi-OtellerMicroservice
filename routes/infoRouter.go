package routes

import (
	controller "oteller-microservice/controllers"

	"github.com/gin-gonic/gin"
)

func InfoRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/info/:otel_id", controller.CreateInfo()) // otel iletişim bilgisi ekleme

	incomingRoutes.DELETE("/info/:otel_id", controller.DeleteInfo()) //Otel Kaldırma

}
