package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	MongoDB := "mongodb://localhost:27017" //burayı daha sonra ayarlayacağız
	fmt.Print(MongoDB)

	//mongo.NewClient(options.Client().ApplyURI(MongoDB))
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MongoDB))
	if err != nil {
		log.Fatal(err)
	}

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// //err = client.Connect(ctx)
	fmt.Println("Connected to MongoDB")
	return client

}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var Collection *mongo.Collection = client.Database("oteller").Collection(collectionName)

	return Collection
}
