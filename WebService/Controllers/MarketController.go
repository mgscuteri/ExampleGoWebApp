package Controllers

import (
	"encoding/json"
	"net/http"

	"ExampleGoWebApp/WebService/DataModels"
	"ExampleGoWebApp/WebService/Repositories"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllMarkets(w http.ResponseWriter, r *http.Request) {
	result := Repositories.GetAllMarkets()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
	return
}

func CreateMarket(w http.ResponseWriter, r *http.Request) {
	jsonDecoder := json.NewDecoder(r.Body)
	var market DataModels.DB_Market
	err := jsonDecoder.Decode(&market)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createdMarketResult := Repositories.CreateMarket(market)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdMarketResult)
	return
}

func UpdateMarket(w http.ResponseWriter, r *http.Request) {
	jsonDecoder := json.NewDecoder(r.Body)
	var market DataModels.DB_Market
	err := jsonDecoder.Decode(&market)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedMarketResult := Repositories.UpdateMarket(market)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedMarketResult)
	return
}

func DeleteMarketById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	marketId := vars["id"]
	marketDocId, _ := primitive.ObjectIDFromHex(marketId)

	deleteResult := Repositories.DeleteMarketById(marketDocId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deleteResult)
	return
}