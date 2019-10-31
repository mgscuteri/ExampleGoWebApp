package DataModels

import (
	"time"
)

type Semester struct {
	SemesterId int
	StartDate  time.Time
	EndDate    time.Time
	Name       string
}
