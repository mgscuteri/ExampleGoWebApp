package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"ExampleGoWebApp/WebService/Controllers"
)

func main() {
	router := mux.NewRouter()
	
	//Market Routes
	router.HandleFunc("/Markets/CreateMarket", Controllers.CreateMarket).Methods("POST")
	router.HandleFunc("/Markets/UpdateMarket", Controllers.UpdateMarket).Methods("POST")
	router.HandleFunc("/Markets/DeleteMarketById/{id}", Controllers.DeleteMarketById).Methods("DELETE")
	router.HandleFunc("/Markets/GetAllMarkets", Controllers.GetAllMarkets).Methods("GET")
	
	//Meal Plans Routes
	router.HandleFunc("/MealPlans/CreateMealPlan/{marketId}", Controllers.CreateMealPlan).Methods("POST")
	router.HandleFunc("/MealPlans/UpdateMealPlan", Controllers.UpdateMealPlan).Methods("POST")
	router.HandleFunc("/MealPlans/GetMealPlanById/{id}", Controllers.GetMealPlanById).Methods("GET")
	router.HandleFunc("/MealPlans/DeleteMealPlanById/{id}", Controllers.DeleteMealPlanById).Methods("DELETE")
	router.HandleFunc("/MealPlans/GetAllMealPlans", Controllers.GetAllMealPlans).Methods("GET")	
	router.HandleFunc("/MealPlans/GetAllMealPlansByMarketId/{id}", Controllers.GetAllMealPlansByMarketId).Methods("GET")

	//Semester Routes
	router.HandleFunc("/Semesters/GetAllSemestersByMarketId/{id}", Controllers.GetAllSemestersByMarketId).Methods("GET")
	router.HandleFunc("/Semesters/UpdateSemester", Controllers.UpdateSemester).Methods("POST")
	router.HandleFunc("/Semesters/CreateSemester/{marketId}", Controllers.UpdateSemester).Methods("POST")

	//User Routes
	router.HandleFunc("/Users/CreateUser", Controllers.CreateUser).Methods("POST")
	router.HandleFunc("/Users/UpdateUser", Controllers.UpdateUser).Methods("POST")
	router.HandleFunc("/Users/DeleteUserById/{id}", Controllers.DeleteUserById).Methods("DELETE")
	router.HandleFunc("/Users/GetAllUsers", Controllers.GetAllUsers).Methods("GET")

	// Serve up the bundled react app
	router.PathPrefix("/site").Handler(http.StripPrefix("/site", http.FileServer(http.Dir("Web/build"))))
	router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("Web/build/static"))))

	log.Fatal(http.ListenAndServe(":8080", router))
}
