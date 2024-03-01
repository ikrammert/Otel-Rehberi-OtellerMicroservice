package routes

import (
	controller "oteller-microservice/controllers"

	"github.com/gin-gonic/gin"
)

func PhoneRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("/phone/:otel_id", controller.CreateInfo()) // otel iletişim bilgisi ekleme

	incomingRoutes.DELETE("/phone/:otel_id", controller.DeleteInfo()) //otel iletişim bilgisi  Kaldırma

}
