package entity

import "time"

type UserAvatar struct {
	ID  string
	URL string
}

type User struct {
	ID                   string
	Email                string
	Phone                string
	FirstName            string
	LastName             string
	Name                 string
	Avatar               UserAvatar
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
