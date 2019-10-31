package Singletons

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

var connectionOpen bool = false

func GetClientOptions() *mongo.Client {

	if !connectionOpen {
		var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")

		// Connect to MongoDB
		newClient, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		// Check the connection
		err = newClient.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connected to MongoDB Instance")
		connectionOpen = true
		client = newClient
		return client
	} else {
		return client
	}
}
