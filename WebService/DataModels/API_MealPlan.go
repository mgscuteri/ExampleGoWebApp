package DataModels

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type API_MealPlan struct {	
	Id         primitive.ObjectID `bson:"_id, omitempty"`
	Name       string             `json:"Name"`
	WeeklyCost int                `json:"WeeklyCost"`
	MarketId   int                `json:"MarketId"`
}
