package DataModels

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DB_MealPlan struct {	
	Id         primitive.ObjectID `bson:"_id, omitempty"`
	Name       string             `json:"Name"`
	WeeklyCost float32            `json:"WeeklyCost"`	
}
