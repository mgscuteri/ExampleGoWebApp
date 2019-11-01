package Controllers

import (
	"encoding/json"
	"net/http"
	"fmt"

	"ExampleGoWebApp/WebService/DataModels"
	"ExampleGoWebApp/WebService/Repositories"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMealPlanById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mealPlanId := vars["id"]
	mealPlanDocId, _ := primitive.ObjectIDFromHex(mealPlanId)
	var mealPlan DataModels.DB_MealPlan = Repositories.GetMealPlanById(mealPlanDocId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mealPlan)
	return
}

func CreateMealPlan(w http.ResponseWriter, r *http.Request) {
	jsonDecoder := json.NewDecoder(r.Body)
	var mealPlan DataModels.DB_MealPlan
	vars := mux.Vars(r)
	marketId := vars["marketId"]
	marketIdDoc, _ := primitive.ObjectIDFromHex(marketId)
	err := jsonDecoder.Decode(&mealPlan)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Create the meal plan
	createdMealPlanResult := Repositories.CreateMealPlan(mealPlan)

	//Associate it with a market
	createdMealPlanMarketAssociationResult := Repositories.CreateMealPlanMarketAssociation(createdMealPlanResult.Id, marketIdDoc)
	fmt.Println(createdMealPlanMarketAssociationResult.Id.String())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdMealPlanResult)
	return
}

func UpdateMealPlan(w http.ResponseWriter, r *http.Request) {
	jsonDecoder := json.NewDecoder(r.Body)
	var mealPlan DataModels.DB_MealPlan
	err := jsonDecoder.Decode(&mealPlan)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedMealPlanResult := Repositories.UpdateMealPlan(mealPlan)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedMealPlanResult)
	return
}

func GetAllMealPlans(w http.ResponseWriter, r *http.Request) {
	var mealPlans []DataModels.DB_MealPlan = Repositories.GetAllMealPlans()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mealPlans)
	return
}

func GetAllMealPlansByMarketId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	marketId := vars["id"]
	marketDocId, _ := primitive.ObjectIDFromHex(marketId)
	var mealPlans []DataModels.DB_MealPlan = Repositories.GetAllMealPlansByMarketId(marketDocId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mealPlans)
	return
}

func DeleteMealPlanById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mealPlanId := vars["id"]
	mealPlanDocId, _ := primitive.ObjectIDFromHex(mealPlanId) 
	deleteResult := Repositories.DeleteMealPlanById(mealPlanDocId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deleteResult)
	return
}