package entity

import "time"

type User struct {
	ID                   string
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
	CreatedAt            time.Time
	UpdatedAt            time.Time
	CreatedByID          string
	UpdatedByID          string
	DeletedAt            *time.Time
	Version              int32
}
