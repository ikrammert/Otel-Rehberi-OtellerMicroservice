package database

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Bu satırı ekleyin
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = DBinstance()

func migrateDatabase() error {
	// Migrate işlemleri için veritabanı URL'sini belirtin
	// Migrate işlemleri için source dizinini belirtin
	migrationURL := "mongodb://localhost:27017/oteller"
	sourceURL := "file://db"

	// Migrate instance'ını oluşturun
	// p := &mongodb.Mongo{}
	// d, err := p.Open(migrationURL)
	// if err != nil {
	// 	return err
	// }
	log.Print("1")
	m, err := migrate.New(sourceURL, migrationURL)
	// m, err := migrate.NewWithDatabaseInstance(sourceURL, "oteller", d)
	if err != nil {
		log.Print("2")
		return err
	}
	log.Print("burası")
	// Migrate işlemlerini gerçekleştirin
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Print("3")
		return err
	}

	fmt.Println("Migration işlemi başarıyla tamamlandı")
	return nil
}

func DBinstance() *mongo.Client {
	MongoDB := "mongodb://localhost:27017" // MongoDB bağlantı URL'si burada
	fmt.Println("Connecting to MongoDB:", MongoDB)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MongoDB))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	// Migration işlemlerini gerçekleştirin
	err := migrateDatabase()
	if err != nil {
		log.Fatal("Migration işlemi sırasında hata oluştu:", err)
	}

	// Belirtilen koleksiyonu açın
	collection := client.Database("oteller").Collection(collectionName)
	return collection
}
