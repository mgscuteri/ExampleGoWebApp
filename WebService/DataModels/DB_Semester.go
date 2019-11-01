package DataModels

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DB_Semester struct {
	Id			primitive.ObjectID 	`bson:"_id, omitempty"`
	StartDate  	time.Time			`json:"StartDate"`
	EndDate    	time.Time			`json:"EndDate"`
	Name       	string				`json:"Name"`
}
