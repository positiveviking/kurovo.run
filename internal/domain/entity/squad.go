package entity

import "time"

type SquadAccount struct {
	UserID string

	SquadID string
	Token   string
	Expire  time.Time
}
