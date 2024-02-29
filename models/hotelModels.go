package models

// Otel modeli
type Hotel struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Location    string    `json:"location"`
	ContactInfo []Contact `json:"contact_info"`
}

// İletişim bilgisi modeli
type Contact struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
