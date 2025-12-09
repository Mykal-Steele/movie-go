package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// setup mongo client
func DBinstance() *mongo.Client {

	// load env
	err := godotenv.Load()
	if err != nil {
		log.Print("Dotenv file does not exist")
	}

	// get env
	MongoURI := os.Getenv("MONGODB_URI")
	if MongoURI == "" {
		log.Fatal("Cannot get MONGODB_URI")
	}

	fmt.Printf("Database name: %v\n", MongoURI)
	clientOption := options.Client().ApplyURI(MongoURI)
	client, err := mongo.Connect(clientOption)

	if err != nil {
		log.Fatal("Cannot connect to MongoDB: ", err)
	}
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(collectionName string) *mongo.Collection {
	err := godotenv.Load()
	if err != nil {
		log.Println("Dotenv file not found. ")
	}

	Database := os.Getenv("DATABASE_NAME")
	fmt.Print("Database: ", Database)
	if Database == "" {
		log.Fatal("Cannot get Database name. ")
	}
	collection := Client.Database(Database).Collection(collectionName)
	if collection == nil {
		return nil
	}
	return collection
}
