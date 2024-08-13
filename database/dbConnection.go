package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error in loading env file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongo_url := os.Getenv("MONGO_URL")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongo_url))

	if err != nil {
		log.Fatal("database connection error", err.Error())
	}
	fmt.Println("connected to database")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(cient *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = Client.Database("restaurant").Collection(collectionName)

	return collection
}
