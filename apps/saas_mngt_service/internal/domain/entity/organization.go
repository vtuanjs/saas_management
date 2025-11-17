package entity

import "time"

type Organization struct {
	ID          string
	Code        string
	Name        string
	Logo        *Attachment
	IsLocked    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedByID string
	UpdatedByID string
	DeletedAt   *time.Time
	Version     int32
}
