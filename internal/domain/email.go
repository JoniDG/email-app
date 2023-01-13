package domain

import "time"

type Body struct {
	Name        string
	TimeStamp   time.Time
	Temperature int8
}
