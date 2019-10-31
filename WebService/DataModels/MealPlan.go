package DataModels

type MealPlan struct {
	MealPlanId int
	Name       string
	WeeklyCost int
	MarketId   int
}

type CreatedMealPlanResult struct {
	MealPlanId string
}

type UpdatedMealPlanResult struct {
	NumRecordsUpdated int64
}
