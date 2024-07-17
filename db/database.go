package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Unable to load .env file")
		log.Fatal(err)
	}

	mongoURL := os.Getenv("MONGO_URL")

	dbName := os.Getenv("DB_NAME")

	collectionName := os.Getenv("USER_COLLECTION")

	client, connerr := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURL))

	if connerr != nil {
		fmt.Println("MongoDB connection ERROR")
		log.Fatal(connerr)
	}

	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("DB & Collection Ready")
}

func GetCollection() *mongo.Collection {
	return collection
}
