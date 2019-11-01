package Controllers

import (
	"encoding/json"
	"net/http"
	"fmt"

	"ExampleGoWebApp/WebService/Repositories"
	"ExampleGoWebApp/WebService/DataModels"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateSemester(w http.ResponseWriter, r *http.Request) {
	jsonDecoder := json.NewDecoder(r.Body)
	var semester DataModels.DB_Semester
	vars := mux.Vars(r)
	marketId := vars["marketId"]
	marketIdDoc, _ := primitive.ObjectIDFromHex(marketId)
	err := jsonDecoder.Decode(&semester)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Create the semester
	createdSemesterResult := Repositories.CreateSemester(semester)

	//Associate it with a market
	createdMarketSemesterAssociationResult := Repositories.CreateMarketSemesterAssociation(marketIdDoc, createdSemesterResult.Id)
	fmt.Println(createdMarketSemesterAssociationResult.Id.String())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdSemesterResult)
	return
}


func GetAllSemestersByMarketId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	marketId := vars["id"]
	marketDocId, _ := primitive.ObjectIDFromHex(marketId)
	result := Repositories.GetAllSemestersByMarketId(marketDocId)	

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
	return
}

func UpdateSemester(w http.ResponseWriter, r *http.Request) {
	jsonDecoder := json.NewDecoder(r.Body)
	var semester DataModels.DB_Semester
	err := jsonDecoder.Decode(&semester)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedSemesterresult := Repositories.UpdateSemester(semester)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedSemesterresult)
	return
}

