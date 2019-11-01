package Repositories

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"ExampleGoWebApp/WebService/DataModels"
	"ExampleGoWebApp/WebService/Singletons"
)

func CreateUser(user DataModels.DB_User) DataModels.CreatedRecordResult {
	client := Singletons.GetClientOptions()

	collection := client.Database("WebServiceDB").Collection("Users")
	res, err := collection.InsertOne(context.TODO(), bson.D{
		{"MealPlanId", user.MealPlanId},
		{"MarketId", user.MarketId},
		{"Name", user.Name},
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted: ", res.InsertedID)

	insertedId := res.InsertedID.(primitive.ObjectID)
	createdUserResult := DataModels.CreatedRecordResult {
		Id: insertedId,
	}

	return createdUserResult
}

func GetUserById(id primitive.ObjectID) DataModels.DB_User {
	client := Singletons.GetClientOptions()
	collection := client.Database("WebServiceDB").Collection("Users")

	filter := bson.M{"_id": id}

	var userResult DataModels.DB_User = DataModels.DB_User{}

	result := collection.FindOne(context.TODO(), filter)
	result.Decode(&userResult)

	return userResult
}

func UpdateUser(user DataModels.DB_User) DataModels.NumRecordsAffected {
	client := Singletons.GetClientOptions()

	collection := client.Database("WebServiceDB").Collection("Users")

	filter := bson.D{{"_id", user.Id}}

	update := bson.D{
		{"$set", bson.D{
			{"_id", user.Id},
			{"Name", user.Name},
			{"MealPlanId", user.MealPlanId},
			{"MarketId", user.MarketId},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	numRecordsAffected := DataModels.NumRecordsAffected{
		NumRecordsAffected: updateResult.ModifiedCount,
		TableName: "Users",
		Action: "Update",
	}
	return numRecordsAffected
}

func GetAllUsers() []DataModels.DB_User {
	client := Singletons.GetClientOptions()
	collection := client.Database("WebServiceDB").Collection("Users")
	filter := bson.M{}

	var results []DataModels.DB_User

	cur, _ := collection.Find(context.TODO(), filter)

	for cur.Next(context.TODO()) {
		var user DataModels.DB_User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, user)
	}
	return results
}

func DeleteUserById(id primitive.ObjectID) DataModels.NumRecordsAffected {
	client := Singletons.GetClientOptions()
	collection := client.Database("WebServiceDB").Collection("Users")

	filter := bson.M{"_id": id}	

	deleteResult, _ := collection.DeleteOne(context.TODO(), filter) 

	numRecordsAffected := DataModels.NumRecordsAffected{
		NumRecordsAffected: deleteResult.DeletedCount,
		TableName: "Users",
		Action: "Delete",
	}

	return numRecordsAffected
}