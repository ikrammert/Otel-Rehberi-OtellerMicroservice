package routes

import (
	controller "oteller-microservice/controllers"

	"github.com/gin-gonic/gin"
)

func CommunicationRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("/conn/otel/:otel_id/info/:conn_id", controller.CreateCommInfo()) // otel iletişim bilgisi ekleme

	incomingRoutes.DELETE("/conn/otel/:otel_id/info/:conn_id", controller.DeleteCommInfo()) //otel iletişim bilgisi  Kaldırma

}
