package entity

import "time"

type Permission struct {
	ID          string
	Code        string
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
