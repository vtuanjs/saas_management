package domain

import "time"

type UserEntity struct {
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

type UserRepository interface {
	FindByID(id string) (*UserEntity, error)
	FindByEmail(email string) (*UserEntity, error)
	Save(user *UserEntity) (*UserEntity, error)
}
