package DataModels

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreatedRecordResult struct {
	Id primitive.ObjectID
}

type NumRecordsAffected struct {
	NumRecordsAffected int64
	TableName string
	Action string
}
