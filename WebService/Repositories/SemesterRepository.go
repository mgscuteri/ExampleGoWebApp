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

func CreateSemester(semester DataModels.DB_Semester) DataModels.CreatedRecordResult {
	client := Singletons.GetClientOptions()

	collection := client.Database("WebServiceDB").Collection("Semesters")
	res, err := collection.InsertOne(context.TODO(), bson.D{
		{"StartDate", semester.StartDate},
		{"EndDate", semester.EndDate},
		{"Name", semester.EndDate},		
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted: ", res.InsertedID)

	insertedId := res.InsertedID.(primitive.ObjectID)
	createdSemesterResult := DataModels.CreatedRecordResult {
		Id: insertedId,
	}

	return createdSemesterResult
}

func CreateMarketSemesterAssociation(marketId primitive.ObjectID, semesterId primitive.ObjectID) DataModels.CreatedRecordResult {
	client := Singletons.GetClientOptions()

	collection := client.Database("WebServiceDB").Collection("Market_x_Semester")
	res, err := collection.InsertOne(context.TODO(), bson.D{
		{"MarketId", marketId},
		{"SemesterId", semesterId},
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

func UpdateSemester(semester DataModels.DB_Semester) DataModels.NumRecordsAffected {
	client := Singletons.GetClientOptions()
	collection := client.Database("WebServiceDB").Collection("Semesters")

	filter := bson.D{{"_id", semester.Id}}
	update := bson.D{
		{"$set", bson.D{
			{"_id", semester.Id},
			{"StartDate", semester.StartDate},
			{"EndDate", semester.EndDate},
			{"Name", semester.Name},
		}},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	numRecordsAffected := DataModels.NumRecordsAffected{
		NumRecordsAffected: updateResult.ModifiedCount,
		TableName: "Semesters",
		Action: "Update",
	}
	return numRecordsAffected
}

func DeleteSemesterById(semesterId string) []DataModels.NumRecordsAffected {
	client := Singletons.GetClientOptions()
	collection := client.Database("WebServiceDB").Collection("Semesters")
	associationCollectoin := client.Database("WebServiceDB").Collection("Market_x_Semester")

	//Delete Market_x_Semester associations
	associationFilter := bson.M{"SemesterId": semesterId}
	associationDeleteResult, _ := associationCollectoin.DeleteMany(context.TODO(), associationFilter)	
	numMarketSemesterAssociationsDeleted := DataModels.NumRecordsAffected{
		NumRecordsAffected: associationDeleteResult.DeletedCount,
		TableName: "Market_x_Semester",
		Action: "Delete",
	}

	//Delete Semester
	docID, _ := primitive.ObjectIDFromHex(semesterId)
	filter := bson.M{"_id": docID}	

	deleteResult, _ := collection.DeleteOne(context.TODO(), filter) 

	numSemestersDeleted := DataModels.NumRecordsAffected{
		NumRecordsAffected: deleteResult.DeletedCount,
		TableName: "Semester",
		Action: "Delete",
	}

	var results []DataModels.NumRecordsAffected
	results = append(results, numMarketSemesterAssociationsDeleted)
	results = append(results, numSemestersDeleted)
	return results
}

func GetSemesterById(semesterId primitive.ObjectID) DataModels.DB_Semester {
	client := Singletons.GetClientOptions()
	collection := client.Database("WebServiceDB").Collection("Semesters")

	filter := bson.M{"_id": semesterId}

	var semesterResult DataModels.DB_Semester = DataModels.DB_Semester{}

	result := collection.FindOne(context.TODO(), filter)
	result.Decode(&semesterResult)

	return semesterResult
}

func GetAllSemestersByMarketId(marketId primitive.ObjectID) []DataModels.DB_Semester {
	client := Singletons.GetClientOptions()
	collection := client.Database("WebServiceDB").Collection("Market_x_Semester")
	
	filter := bson.M{"MarketId": marketId}

	var results []DataModels.DB_Semester

	cur, _ := collection.Find(context.TODO(), filter)

	for cur.Next(context.TODO()) {
		var market_x_semester DataModels.DB_Market_x_Semester
		err := cur.Decode(&market_x_semester)
		if err != nil {
			log.Fatal(err)
		}
		var semester = GetSemesterById(market_x_semester.SemesterId)
		results = append(results, semester)
	}
	return results
}