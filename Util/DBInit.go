package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"ExampleGoWebApp/WebService/Singletons"
	"ExampleGoWebApp/WebService/DataModels"
	"ExampleGoWebApp/WebService/Repositories"
)

func main() {

	var client = Singletons.GetClientOptions()

	//Create Markets
	drexel := DataModels.DB_Market {Name: "Drexel"}
	uPenn := DataModels.DB_Market {Name: "UPenn"}
	temple := DataModels.DB_Market {Name: "Temple"}
	westChester := DataModels.DB_Market {Name: "West Chester"}
	pennState := DataModels.DB_Market {Name: "Penn State"}
	rutgers := DataModels.DB_Market {Name: "Rutgers"}

	drexelMarketId := Repositories.CreateMarket(drexel)
	uPennMarketId := Repositories.CreateMarket(uPenn)
	templeMarketId := Repositories.CreateMarket(temple)
	westChesterMarketId := Repositories.CreateMarket(westChester)
	pennStateMarketId := Repositories.CreateMarket(pennState)	
	rutgersMarketId := Repositories.CreateMarket(rutgers)

	//Create Semesters	
	DrexelSem1 := DataModels.DB_Semester{
		StartDate: time.Date(2019, time.September, 1, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2019, time.December, 1, 0, 0, 0, 0, time.UTC),
		Name: "Drexel Semester 1",
	}
	DrexelSem2 := DataModels.DB_Semester{
		StartDate: time.Date(2019, time.December, 2, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2020, time.March, 1, 0, 0, 0, 0, time.UTC),
		Name: "Drexel Semester 2",
	}
	DrexelSem3 := DataModels.DB_Semester{
		StartDate: time.Date(2019, time.March, 2, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2020, time.July, 1, 0, 0, 0, 0, time.UTC),
		Name: "Drexel Semester 3",
	}
	DrexelSem4 := DataModels.DB_Semester{
		StartDate: time.Date(2019, time.July, 2, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2020, time.August, 30, 0, 0, 0, 0, time.UTC),
		Name: "Drexel Semester 4",
	}
	uPennSem1 := DataModels.DB_Semester{
		StartDate: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2020, time.May, 1, 0, 0, 0, 0, time.UTC),
		Name: "UPenn Semester 1",
	}
	uPennSem2 := DataModels.DB_Semester{
		StartDate: time.Date(2019, time.May, 2, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2020, time.September, 1, 0, 0, 0, 0, time.UTC),
		Name: "UPenn Semester 2",
	}
	uPennSem3 := DataModels.DB_Semester{
		StartDate: time.Date(2019, time.September, 2, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2020, time.December, 31, 0, 0, 0, 0, time.UTC),
		Name: "UPenn Semester 3",
	}
	genericSem1 := DataModels.DB_Semester{
		StartDate: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2020, time.June, 30, 0, 0, 0, 0, time.UTC),
		Name: "Semester 1",
	}
	genericSem2 := DataModels.DB_Semester{
		StartDate: time.Date(2019, time.July, 1, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2020, time.December, 31, 0, 0, 0, 0, time.UTC),
		Name: "Semester 2",
	}
	drexelSemester1Id := Repositories.CreateSemester(DrexelSem1)
	drexelSemester2Id := Repositories.CreateSemester(DrexelSem2)
	drexelSemester3Id := Repositories.CreateSemester(DrexelSem3)
	drexelSemester4Id := Repositories.CreateSemester(DrexelSem4)
	uPennSem1Id := Repositories.CreateSemester(uPennSem1)
	uPennSem2Id := Repositories.CreateSemester(uPennSem2)
	uPennSem3Id := Repositories.CreateSemester(uPennSem3)
	genericSem1Id := Repositories.CreateSemester(genericSem1)
	genericSem2Id := Repositories.CreateSemester(genericSem2)

	//Associate Markets with Semesters
	Repositories.CreateMarketSemesterAssociation(drexelMarketId.Id, drexelSemester1Id.Id)
	Repositories.CreateMarketSemesterAssociation(drexelMarketId.Id, drexelSemester2Id.Id)
	Repositories.CreateMarketSemesterAssociation(drexelMarketId.Id, drexelSemester3Id.Id)
	Repositories.CreateMarketSemesterAssociation(drexelMarketId.Id, drexelSemester4Id.Id)
	Repositories.CreateMarketSemesterAssociation(uPennMarketId.Id, uPennSem1Id.Id)
	Repositories.CreateMarketSemesterAssociation(uPennMarketId.Id, uPennSem2Id.Id)
	Repositories.CreateMarketSemesterAssociation(uPennMarketId.Id, uPennSem3Id.Id)
	Repositories.CreateMarketSemesterAssociation(templeMarketId.Id, genericSem1Id.Id)
	Repositories.CreateMarketSemesterAssociation(templeMarketId.Id, genericSem2Id.Id)
	Repositories.CreateMarketSemesterAssociation(westChesterMarketId.Id, genericSem1Id.Id)
	Repositories.CreateMarketSemesterAssociation(westChesterMarketId.Id, genericSem2Id.Id)
	Repositories.CreateMarketSemesterAssociation(pennStateMarketId.Id, genericSem1Id.Id)
	Repositories.CreateMarketSemesterAssociation(pennStateMarketId.Id, genericSem2Id.Id)
	Repositories.CreateMarketSemesterAssociation(rutgersMarketId.Id, genericSem1Id.Id)
	Repositories.CreateMarketSemesterAssociation(rutgersMarketId.Id, genericSem2Id.Id)

	//Create Meal Plans
	DrexelMealPlan1 := DataModels.DB_MealPlan{
		Name: "Drexel Meal Plan 1",
		WeeklyCost: 109.99,
	}
	DrexelMealPlan2 := DataModels.DB_MealPlan{
		Name: "Drexel Meal Plan 2",
		WeeklyCost: 152.99,
	}
	UPennMealPlan := DataModels.DB_MealPlan{
		Name: "UPenn Meal Plan",
		WeeklyCost: 149.00,
	}
	MultiMarketMealPlan1 := DataModels.DB_MealPlan{
		Name: "Best Meal Plan",
		WeeklyCost: 99.99,
	}
	MultiMarketMealPlan2 := DataModels.DB_MealPlan{
		Name: "Premmium Meal Plan",
		WeeklyCost: 199.99,
	}
	DrexelMealPlan1Id := Repositories.CreateMealPlan(DrexelMealPlan1)
	DrexelMealPlan2Id := Repositories.CreateMealPlan(DrexelMealPlan2)
	UPennMealPlanId := Repositories.CreateMealPlan(UPennMealPlan)
	MultiMarketMealPlan1Id := Repositories.CreateMealPlan(MultiMarketMealPlan1)
	MultiMarketMealPlan2Id := Repositories.CreateMealPlan(MultiMarketMealPlan2)

	//Associate Meal Plans with Markets
	Repositories.CreateMealPlanMarketAssociation(DrexelMealPlan1Id.Id, drexelMarketId.Id)
	Repositories.CreateMealPlanMarketAssociation(DrexelMealPlan2Id.Id, drexelMarketId.Id)
	Repositories.CreateMealPlanMarketAssociation(UPennMealPlanId.Id, uPennMarketId.Id)
	Repositories.CreateMealPlanMarketAssociation(MultiMarketMealPlan1Id.Id, templeMarketId.Id)
	Repositories.CreateMealPlanMarketAssociation(MultiMarketMealPlan2Id.Id, templeMarketId.Id)
	Repositories.CreateMealPlanMarketAssociation(MultiMarketMealPlan1Id.Id, westChesterMarketId.Id)
	Repositories.CreateMealPlanMarketAssociation(MultiMarketMealPlan2Id.Id, pennStateMarketId.Id)
	Repositories.CreateMealPlanMarketAssociation(MultiMarketMealPlan1Id.Id, rutgersMarketId.Id)
	Repositories.CreateMealPlanMarketAssociation(MultiMarketMealPlan2Id.Id, rutgersMarketId.Id)
	//Add Users
	user1 := DataModels.DB_User{
		MealPlanId: DrexelMealPlan1Id.Id,
		MarketId: drexelMarketId.Id,
		Name: "SuperDrexelUser1",
	}
	user2 := DataModels.DB_User{
		MealPlanId: DrexelMealPlan2Id.Id,
		MarketId: drexelMarketId.Id,
		Name: "DrexelPerson2",
	}
	user3 := DataModels.DB_User{
		MealPlanId: UPennMealPlanId.Id,
		MarketId: uPennMarketId.Id,
		Name: "UPennMegaUser1",
	}
	user4 := DataModels.DB_User{
		MealPlanId: MultiMarketMealPlan1Id.Id,
		MarketId: templeMarketId.Id,
		Name: "TempleUser1",
	}
	user5 := DataModels.DB_User{
		MealPlanId: MultiMarketMealPlan2Id.Id,
		MarketId: westChesterMarketId.Id,
		Name: "WestChesterUser1",
	}
	user6 := DataModels.DB_User{
		MealPlanId: MultiMarketMealPlan1Id.Id,
		MarketId: pennStateMarketId.Id,
		Name: "PennStateUser1",
	}
	user7 := DataModels.DB_User{
		MealPlanId: MultiMarketMealPlan2Id.Id,
		MarketId: rutgersMarketId.Id,
		Name: "RutgersUser1",
	}
	Repositories.CreateUser(user1)
	Repositories.CreateUser(user2)
	Repositories.CreateUser(user3)
	Repositories.CreateUser(user4)
	Repositories.CreateUser(user5)
	Repositories.CreateUser(user6)
	Repositories.CreateUser(user7)

	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
