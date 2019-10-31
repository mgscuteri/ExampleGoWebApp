package main

import (
	"context"
	"fmt"
	"log"

	"GoodUncleWebService/WebService/Singletons"
)

func main() {

	var client = Singletons.GetClientOptions()

	collection := client.Database("WebServiceDB").Collection("MealPlans")
	print(collection)

	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
