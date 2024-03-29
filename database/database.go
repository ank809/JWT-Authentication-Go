package database

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = DatabaseInstance()

func DatabaseInstance() *mongo.Client {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}
	url := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		fmt.Println(err)
	}
	return client

}

func OpenCollection(client *mongo.Client, coll_name string) *mongo.Collection {
	collection := client.Database("jwt-go").Collection(coll_name)
	return collection
}
