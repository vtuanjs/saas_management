package entity

import "time"

type Role struct {
	ID          string
	Name        string
	Description string
	System      bool
	OrgID       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedByID string
	UpdatedByID string
	DeletedAt   *time.Time
	Version     int32
}
