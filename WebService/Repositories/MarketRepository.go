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

func CreateMarket(market DataModels.DB_Market) DataModels.CreatedRecordResult {
	client := Singletons.GetClientOptions()

	collection := client.Database("WebServiceDB").Collection("Markets")
	res, err := collection.InsertOne(context.TODO(), bson.D{
		{"Name", market.Name},		
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted: ", res.InsertedID)

	insertedId := res.InsertedID.(primitive.ObjectID)
	createdMarketResult := DataModels.CreatedRecordResult {
		Id: insertedId,
	}

	return createdMarketResult
}

func UpdateMarket(market DataModels.DB_Market) DataModels.NumRecordsAffected {
	client := Singletons.GetClientOptions()

	collection := client.Database("WebServiceDB").Collection("Markets")

	filter := bson.D{{"_id", market.Id}}

	update := bson.D{
		{"$set", bson.D{
			{"_id", market.Id},
			{"Name", market.Name},			
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	numRecordsAffected := DataModels.NumRecordsAffected{
		NumRecordsAffected: updateResult.ModifiedCount,
	}
	return numRecordsAffected
}

func GetAllMarkets() []DataModels.DB_Market {
	client := Singletons.GetClientOptions()
	collection := client.Database("WebServiceDB").Collection("Markets")
	filter := bson.M{}

	var results []DataModels.DB_Market

	cur, _ := collection.Find(context.TODO(), filter)

	for cur.Next(context.TODO()) {
		var market DataModels.DB_Market
		err := cur.Decode(&market)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, market)
	}	
	return results
}

func DeleteMarketById(id primitive.ObjectID) []DataModels.NumRecordsAffected {
	client := Singletons.GetClientOptions()
	marketCollection := client.Database("WebServiceDB").Collection("Markets")
	userCollectoin := client.Database("WebServiceDB").Collection("Users")

	marketFilter := bson.M{"_id": id}	

	//Update MarketId of users associated with the market to nil
	nullObjectIdDoc, _ := primitive.ObjectIDFromHex("000000000000000000000000")
	userFilter := bson.D{{"MarketId", id}}
	update := bson.D{
		{"$set", bson.D{
			{"MarketId", nullObjectIdDoc},
		}},
	}
	userUpdateResult, _ := userCollectoin.UpdateMany(context.TODO(), userFilter, update)
	numUsersUpdated := DataModels.NumRecordsAffected{
		NumRecordsAffected: userUpdateResult.ModifiedCount,
		TableName: "Users",
		Action: "Updated",
	}	

	//Delete the market
	deleteResult, _ := marketCollection.DeleteOne(context.TODO(), marketFilter) 

	numMarketsAffected := DataModels.NumRecordsAffected{
		NumRecordsAffected: deleteResult.DeletedCount,
		TableName: "Markets",
		Action: "Delete",
	}
	
	var results []DataModels.NumRecordsAffected
	results = append(results, numUsersUpdated)
	results = append(results, numMarketsAffected)
	return results
}

func CreateMealPlanMarketAssociation(mealPlanId primitive.ObjectID, marketId primitive.ObjectID) DataModels.CreatedRecordResult {
	client := Singletons.GetClientOptions()

	collection := client.Database("WebServiceDB").Collection("MealPlan_x_Market")
	res, err := collection.InsertOne(context.TODO(), bson.D{
		{"MealPlanId", mealPlanId},
		{"MarketId", marketId},
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted: ", res.InsertedID)

	insertedId := res.InsertedID.(primitive.ObjectID)
	createdAssociation := DataModels.CreatedRecordResult {
		Id: insertedId,
	}

	return createdAssociation
}