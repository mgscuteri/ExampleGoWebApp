package Controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"ExampleGoWebApp/WebService/DataModels"
	"ExampleGoWebApp/WebService/Repositories"

	"github.com/gorilla/mux"
)

func GetMealPlanById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mealPlanId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	var mealPlan DataModels.MealPlan = Repositories.GetMealPlanById(mealPlanId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mealPlan)
	return
}

func CreateMealPlan(w http.ResponseWriter, r *http.Request) {
	jsonDecoder := json.NewDecoder(r.Body)
	var mealPlan DataModels.MealPlan
	err := jsonDecoder.Decode(&mealPlan)
	if err != nil {
		panic(err)
	}

	mealPlanId := Repositories.CreateMealPlan(mealPlan)
	createdMealPlanResult := DataModels.CreatedMealPlanResult{
		MealPlanId: mealPlanId,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdMealPlanResult)
	return
}

func UpdateMealPlan(w http.ResponseWriter, r *http.Request) {
	jsonDecoder := json.NewDecoder(r.Body)
	var mealPlan DataModels.MealPlan
	err := jsonDecoder.Decode(&mealPlan)
	if err != nil {
		panic(err)
	}

	updatedMealPlanResult := Repositories.UpdateMealPlan(mealPlan)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedMealPlanResult)
	return
}
