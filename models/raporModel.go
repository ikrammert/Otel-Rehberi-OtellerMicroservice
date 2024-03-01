package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rapor struct {
	UUID          primitive.ObjectID `json:"uuid"`
	Created_at    time.Time          `json:"created_at"`
	Konum         string             `json:"konum"`
	Otel_sayisi   string             `json:"otel_sayisi,omitempty"`
	Numara_sayisi string             `json:"numara_sayisi,omitempty"`
	Rapor_durumu  string             `json:"rapor_durumu"`
}
