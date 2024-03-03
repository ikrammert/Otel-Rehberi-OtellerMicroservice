package database

import (
	"context"
	"fmt"
	"log"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = DBinstance()

func DBinstance() *mongo.Client {
	MongoDB := "mongodb://localhost:27017" // MongoDB bağlantı URL'si burada
	fmt.Println("Connecting to MongoDB:", MongoDB)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MongoDB))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("oteller")
	migrate.SetDatabase(db)
	if err := migrate.Up(migrate.AllAvailable); err != nil {
		log.Fatalf("Hata mig:%+v", err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	collection := client.Database("oteller").Collection(collectionName)
	return collection
}
