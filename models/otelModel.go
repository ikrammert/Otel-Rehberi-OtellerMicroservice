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
	UUID             primitive.ObjectID `json:"uuid"`
	Yetkililer       []Yetkili          `json:"yetkililer"`
	Firma_unvan      string             `json:"firma_unvan"`
	Iletisim_bilgisi []IletisimBilgisi  `json:"iletisim_bilgisi"`
	Otel_id          string             `json:"otel_id"`
	Created_at       time.Time          `json:"created_at"`
	Konum            string             `json:"konum"`
}

type Yetkili struct {
	Yetkili_ad    string `json:"yetkili_ad"`
	Yetkili_soyad string `json:"yetkili_soyad"`
}

type IletisimBilgisi struct {
	Conn_id       string `json:"conn_id"`
	Bilgi_tipi    string `json:"bilgi_tipi"`    // Telefon Numarası, E-mail Adresi, Konum
	Bilgi_icerigi string `json:"bilgi_icerigi"` // Bilgi İçeriği
}
