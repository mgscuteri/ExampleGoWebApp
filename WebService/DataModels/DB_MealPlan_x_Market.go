package DataModels

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DB_MealPlan_x_Market struct {	
	Id			primitive.ObjectID	`bson:"_id, omitempty"`
	MealPlanId	primitive.ObjectID	`bson:"MealPlanId, omitempty"`
	MarketId	primitive.ObjectID	`bson:"MarketId, omitempty"`	
}

