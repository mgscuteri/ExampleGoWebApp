package Controllers

import (
	"encoding/json"
	"net/http"

	"ExampleGoWebApp/WebService/DataModels"
	"ExampleGoWebApp/WebService/Repositories"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	result := Repositories.GetAllUsers()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
	return
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	jsonDecoder := json.NewDecoder(r.Body)
	var user DataModels.DB_User
	err := jsonDecoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createdUserResult := Repositories.CreateUser(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUserResult)
	return
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	jsonDecoder := json.NewDecoder(r.Body)
	var user DataModels.DB_User
	err := jsonDecoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedUserResult := Repositories.UpdateUser(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedUserResult)
	return
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	userDocId, _ := primitive.ObjectIDFromHex(userId)

	deleteResult := Repositories.DeleteUserById(userDocId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deleteResult)
	return
}