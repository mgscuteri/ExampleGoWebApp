package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"ExampleGoWebApp/WebService/Controllers"
	"ExampleGoWebApp/WebService/Singletons"
)

func main() {

	// Bind Routes to Controller Methods
	router := mux.NewRouter()
	router.HandleFunc("/GetMealPlanById/{id}", Controllers.GetMealPlanById).Methods("GET")
	router.HandleFunc("/CreateMealPlan", Controllers.CreateMealPlan).Methods("POST")
	router.HandleFunc("/UpdateMealPlan", Controllers.UpdateMealPlan).Methods("POST")

	// Serve up the bundled react app
	router.PathPrefix("/Site").Handler(http.StripPrefix("/Site", http.FileServer(http.Dir("Web/build"))))
	router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("Web/build/static"))))

	log.Fatal(http.ListenAndServe(":8080", router))

	// Close the MongoDB connection when the server exits
	var client = Singletons.GetClientOptions()

	var err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
