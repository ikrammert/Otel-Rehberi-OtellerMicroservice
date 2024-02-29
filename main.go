package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// HTTP
	router := mux.NewRouter()

	// Rotaları tanımlıyorum
	router.HandleFunc("/hotels", CreateHotel).Methods("POST")
	router.HandleFunc("/hotels/{id}", GetHotel).Methods("GET")
	router.HandleFunc("/hotels/{id}", UpdateHotel).Methods("PUT")
	router.HandleFunc("/hotels/{id}", DeleteHotel).Methods("DELETE")

	// HTTP sunucu yapılandırması
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Varsayılan port
	}

	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("HTTP server listening on port %s...\n", port)

	// Sunucuyu başlat
	log.Fatal(http.ListenAndServe(addr, router))
}

// Örnek handler fonksiyonları
func CreateHotel(w http.ResponseWriter, r *http.Request) {
	// Otel oluşturma işlevi
}

func GetHotel(w http.ResponseWriter, r *http.Request) {
	// Otel getirme işlevi
}

func UpdateHotel(w http.ResponseWriter, r *http.Request) {
	// Otel güncelleme işlevi
}

func DeleteHotel(w http.ResponseWriter, r *http.Request) {
	// Otel silme işlevi
}
