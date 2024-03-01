package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type Otel struct {
// 	ID         primitive.ObjectID `bson:"_id"`
// 	Name       *string            `json:"name"`
// 	Price      *float64           `json:"price"`
// 	Address    *string            `json:"adress"`
// 	Phone      *string            `json:"phone"`
// 	Created_at time.Time          `json:"created_at"`
// 	Otel_id    string             `json:"otel_id"`
// }

type Otel struct {
	UUID            primitive.ObjectID `json:"uuid"`
	YetkiliAd       string             `json:"yetkili_ad"`
	YetkiliSoyad    string             `json:"yetkili_soyad"`
	FirmaUnvan      string             `json:"firma_unvan"`
	IletisimBilgisi []IletisimBilgisi  `json:"iletisim_bilgisi"`
	OtelId          string             `json:"otel_id"`
	CreatedAt       time.Time          `json:"created_at"`
}

type IletisimBilgisi struct {
	BilgiTipi    string `json:"bilgi_tipi"`    // Telefon Numarası, E-mail Adresi, Konum
	BilgiIcerigi string `json:"bilgi_icerigi"` // Bilgi İçeriği
}