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

func GetMealPlanById(id primitive.ObjectID) DataModels.DB_MealPlan {
	client := Singletons.GetClientOptions()
	collection := client.Database("WebServiceDB").Collection("MealPlans")

	filter := bson.M{"_id": id}

	var mealPlanResult DataModels.DB_MealPlan = DataModels.DB_MealPlan{}

	result := collection.FindOne(context.TODO(), filter)
	result.Decode(&mealPlanResult)

	return mealPlanResult
}

func CreateMealPlan(mealPlan DataModels.DB_MealPlan) DataModels.CreatedRecordResult {
	client := Singletons.GetClientOptions()

	collection := client.Database("WebServiceDB").Collection("MealPlans")
	res, err := collection.InsertOne(context.TODO(), bson.D{
		{"Name", mealPlan.Name},
		{"WeeklyCost", mealPlan.WeeklyCost},
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted: ", res.InsertedID)

	insertedId := res.InsertedID.(primitive.ObjectID)
	createdMealPlanResult := DataModels.CreatedRecordResult {
		Id: insertedId,
	}

	return createdMealPlanResult
}

func UpdateMealPlan(mealPlan DataModels.DB_MealPlan) DataModels.NumRecordsAffected {
	client := Singletons.GetClientOptions()

	collection := client.Database("WebServiceDB").Collection("MealPlans")

	filter := bson.D{{"_id", mealPlan.Id}}

	update := bson.D{
		{"$set", bson.D{
			{"_id", mealPlan.Id},
			{"Name", mealPlan.Name},
			{"WeeklyCost", mealPlan.WeeklyCost},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	numRecordsAffected := DataModels.NumRecordsAffected{
		NumRecordsAffected: updateResult.ModifiedCount,
		TableName: "MealPlans",
		Action: "Update",
	}
	return numRecordsAffected
}

func GetAllMealPlans() []DataModels.DB_MealPlan {
	client := Singletons.GetClientOptions()
	collection := client.Database("WebServiceDB").Collection("MealPlans")
	filter := bson.M{}

	var results []DataModels.DB_MealPlan

	cur, _ := collection.Find(context.TODO(), filter)

	for cur.Next(context.TODO()) {
		var mealPlan DataModels.DB_MealPlan
		err := cur.Decode(&mealPlan)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, mealPlan)
	}
	return results
}

func GetAllMealPlansByMarketId(marketId primitive.ObjectID) []DataModels.DB_MealPlan {
	client := Singletons.GetClientOptions()
	collection := client.Database("WebServiceDB").Collection("MealPlan_x_Market")
	
	filter := bson.M{"MarketId": marketId}

	var results []DataModels.DB_MealPlan

	cur, _ := collection.Find(context.TODO(), filter)

	for cur.Next(context.TODO()) {
		var mealPlan_x_market DataModels.DB_MealPlan_x_Market
		err := cur.Decode(&mealPlan_x_market)
		if err != nil {
			log.Fatal(err)
		}
		var mealPlan = GetMealPlanById(mealPlan_x_market.MealPlanId)
		results = append(results, mealPlan)
	}
	return results
}


func DeleteMealPlanById(id primitive.ObjectID) []DataModels.NumRecordsAffected {
	client := Singletons.GetClientOptions()
	mealPlansCollection := client.Database("WebServiceDB").Collection("MealPlans")
	associationCollectoin := client.Database("WebServiceDB").Collection("MealPlan_x_Market")
	userCollectoin := client.Database("WebServiceDB").Collection("Users")

	//Delete MealPlan_x_Market associations	
	associationFilter := bson.M{"MealPlanId": id}
	associationDeleteResult, _ := associationCollectoin.DeleteMany(context.TODO(), associationFilter)	
	numMealPlanMarketAssociationsDeleted := DataModels.NumRecordsAffected{
		NumRecordsAffected: associationDeleteResult.DeletedCount,
		TableName: "MealPlan_x_Market",
		Action: "Delete",
	}

	//Update MealPlanId of subscribed users to null
	nullObjectIdDoc, _ := primitive.ObjectIDFromHex("000000000000000000000000")
	userFilter := bson.D{{"MealPlanId", id}}
	update := bson.D{
		{"$set", bson.D{
			{"MealPlanId", nullObjectIdDoc},
		}},
	}
	userUpdateResult, _ := userCollectoin.UpdateMany(context.TODO(), userFilter, update)
	numUsersUpdated := DataModels.NumRecordsAffected{
		NumRecordsAffected: userUpdateResult.ModifiedCount,
		TableName: "Users",
		Action: "Updated",
	}

	//Delete The MealPlan
	filter := bson.M{"_id": id}	

	deleteResult, _ := mealPlansCollection.DeleteOne(context.TODO(), filter) 

	numMealPlanRecordsAffected := DataModels.NumRecordsAffected{
		NumRecordsAffected: deleteResult.DeletedCount,
		TableName: "MealPlans",
		Action: "Delete",
	}
	
	var results []DataModels.NumRecordsAffected
	results = append(results, numMealPlanMarketAssociationsDeleted)
	results = append(results, numUsersUpdated)
	results = append(results, numMealPlanRecordsAffected)
	return results
}