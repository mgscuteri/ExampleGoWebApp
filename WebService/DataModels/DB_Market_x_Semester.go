package DataModels

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DB_Market_x_Semester struct {	
	Id        	primitive.ObjectID 	`bson:"_id, omitempty"`
	MarketId 	primitive.ObjectID 	`bson:"MarketId, omitempty"`
	SemesterId 	primitive.ObjectID 	`bson:"SemesterId, omitempty"`	
}

