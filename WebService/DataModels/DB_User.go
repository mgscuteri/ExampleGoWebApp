package DataModels

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DB_User struct {
	Id         	primitive.ObjectID	`bson:"_id, omitempty"`
	MealPlanId 	primitive.ObjectID	`bson:"MealPlanId, omitempty"`
	MarketId	primitive.ObjectID	`bson:"MarketId, omitempty"`
	Name       	string				
}
