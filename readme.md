
### *About*

This project is meant to act as the foundation for a fully self contained website. It consists of a webserver written in Go that serves up static content, and acts as a REST server.  The static content that it serves is a React.js application. Below, are instructions on how to get the server up and running. Below that, is documentation on all of the servers currently supported REST endpoints.  The react app that this application serves can be reached by navigating to http://localhost:8080/site.  This website does not yet expose much of the REST capability.  It exists primarily to demonstrate static content being served on the same port as the webserver. (Thus avoiding any CORS security issues) 

Included in this project in an Advanced Rest Client (ARC) project export, which when imported, provides an easy way to test most of the REST functionality. 


### *Quick Start Instructions*

1) Install the Core Dependencies
	A) go : https://golang.org/doc/install
	B) MongoDB community edition (Run MongoDB on default port of 27017)
	C) npm https://www.npmjs.com/get-npm (This is optional. Only needed if you want to make changes to the ReactApp)
2) Clone this repository to your Go working directory (src folder)
3) Install the Go dependancies
	A) Official MongoDB driver (execute the shell script below)
		- go get go.mongodb.org/mongo-driver
	B) gorrilla/mux (execute the shell script below)
		- go get -u github.com/gorilla/mux
4) Build the React App (Optional. Only necesary if you've changed the react app)
	A) navigate a terminal to /WebService/Web and run the shell script below
		- npm run-script build
4) Populate the database with sample data
	1) Build and run /Util/DBInit.go
5) Run the Web/API Server
	1) Build and run /WebService/Routing.go
6) Navigate to http://localhost:8080/site to see the example react app, or replace "/site" with any of the endpoints below to test the REST API functionality


### *REST API DOCUMENTATION* 
(For verbose REST documentation and testing, import REST_Definition.arc into Advanced Rest Client)
#### *Market Routes*
##### "/Markets/CreateMarket" ("POST")
	- Creates a new Market (AKA College)
	- Responds with the newly created ObjectId (string)
##### "/Markets/UpdateMarket" ("POST")
	- Updates the meatadata of a market
	- Responds with the table name, and number of records affected (1 or 0)
##### "/Markets/DeleteMarketById/{marketId}" ("DELETE")
	- Deletes the specified market, and removes any USER/Market Association
	- Responds with the number of records affected (1 or 0)
##### "/Markets/GetAllMarkets" ("GET")
	- Returns a list of all markets

#### *Meal Plan Routes*
##### "/MealPlans/CreateMealPlan/{marketId}" ("POST")
	- Creates a new meal plan, and associates it with the specified marketId 
	- Responds with the newly created ObjectId (string)
##### "/MealPlans/UpdateMealPlan" ("POST")
	- Updates the metadata of an existing meal plan. 
	- Responds with the table name, and number of records affected (1 or 0)
##### "/MealPlans/GetMealPlanById/{mealPlanId}" ("GET")
	- Gets a single meal plan using the specified mealPlanId
##### "/MealPlans/DeleteMealPlanById/{mealPlanId}" ("DELETE")
	- Removes all associations between markets, and the specified mealPlanId
	- Removes all associations between users, and the specified mealPlanId
	- Deletes the specified Meal Paln 
	- Responds with:
		- The number of market associations removed (typically 1)
		- The number of user associations removed
		- The number of meal plans removed (1 or 0)	
##### "/MealPlans/GetAllMealPlans" ("GET")
	- Returns a list of all meal plans	
##### "/MealPlans/GetAllMealPlansByMarketId/{mealPlanId}" ("GET")
	- Returns a list of all meal plans available in for the given market id (typically 1 or 0)
#### *Semester Routes*
##### "/Semesters/GetAllSemestersByMarketId/{id}" ("GET")
	- Returns a list of all semesters (name, startDate, endDate) for the given market
##### "/Semesters/UpdateSemester" ("POST")
	- Updates an existing semester 
	- Responds with the number of recofds affected (1 or 0)
##### "/Semesters/CreateSemester/{marketId}"
	- Creates a new semeste	r, and associats it with the specified market
	- Responds with the newly created ObjectId (string)
#### *User Routes*
##### "/Users/CreateUser" ("POST")
	- Creates a new user, and associates the user with the specified Market, and MealPlan
	- Responds with the newly created ObjectId (string)
##### "/Users/UpdateUser" ("POST")
	- Updates the metadata of a user
##### "/Users/DeleteUserById/{id}" ("DELETE")
	- Deletes a user
##### "/Users/GetAllUsers" ("GET")
	- Returns a list of all users
