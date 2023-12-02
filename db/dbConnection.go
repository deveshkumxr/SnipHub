package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collectionName = "snippets"
)

var Collection *mongo.Collection

func Connect() {

	godotenv.Load()

	dbName := os.Getenv("DB_NAME")
	mongoURI := os.Getenv("MONGO_URI") + dbName
	
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Unable to connect to MongoDB:", err)
	}

	Collection = client.Database(dbName).Collection(collectionName)
	log.Println("Connected to MongoDB")
}
