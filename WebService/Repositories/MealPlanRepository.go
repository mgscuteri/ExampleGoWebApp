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

func GetMealPlanById(id int) DataModels.MealPlan {
	client := Singletons.GetClientOptions()

	collection := client.Database("WebServiceDB").Collection("MealPlans")

	filter := bson.D{{"mealplanid", id}}

	var mealPlanResult DataModels.MealPlan = DataModels.MealPlan{}

	result := collection.FindOne(context.TODO(), filter)
	result.Decode(&mealPlanResult)

	return mealPlanResult
}

func CreateMealPlan(mealPlan DataModels.MealPlan) string {
	client := Singletons.GetClientOptions()

	collection := client.Database("WebServiceDB").Collection("MealPlans")
	res, err := collection.InsertOne(context.TODO(), mealPlan)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted Meal Plan: ", res.InsertedID)

	insertedId := res.InsertedID.(primitive.ObjectID).String()

	return insertedId
}

func UpdateMealPlan(mealPlan DataModels.MealPlan) DataModels.UpdatedMealPlanResult {
	client := Singletons.GetClientOptions()

	collection := client.Database("WebServiceDB").Collection("MealPlans")

	filter := bson.D{{"mealplanid", mealPlan.MealPlanId}}

	update := bson.D{
		{"$set", bson.D{
			{"MealPlanId", mealPlan.MealPlanId},
			{"Name", mealPlan.Name},
			{"WeeklyCost", mealPlan.WeeklyCost},
			{"MarketId", mealPlan.MarketId},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	updatedMealPlanResult := DataModels.UpdatedMealPlanResult{
		NumRecordsUpdated: updateResult.MatchedCount,
	}

	return updatedMealPlanResult
}
