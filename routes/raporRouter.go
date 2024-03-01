package routes

import (
	controller "oteller-microservice/controllers"

	"github.com/gin-gonic/gin"
)

func RaporRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/rapor/:konum", controller.CreateRaporByKonum()) // Konuma Göre Rapor Oluşturma isteği

	incomingRoutes.GET("/rapor/rapors", controller.ListRapors()) // Oluşan Raporları Listeleme

	incomingRoutes.GET("/rapor/:rapor_id", controller.GetRaporById()) // Oluşan Raporun Detay Bilgisini Alma

}
