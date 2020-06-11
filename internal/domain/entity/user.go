package entity

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	UserName  string
	IsBot     bool

	BlockReason null.String
	Created     time.Time
	Modified    time.Time
}
