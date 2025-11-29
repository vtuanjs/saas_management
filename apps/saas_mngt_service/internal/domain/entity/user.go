package entity

import "time"

type User struct {
	Email                string
	Phone                string
	FirstName            string
	LastName             string
	Name                 string
	Avatar               *Attachment
	LastLogin            *time.Time
	Ref                  string
	IsLocked             bool
	IsActivated          bool
	IsAdmin              bool
	IsChangePassRequired bool

	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}
