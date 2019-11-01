package DataModels

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DB_Market struct {
	Id		primitive.ObjectID 	`bson:"_id, omitempty"`
	Name	string				
}
