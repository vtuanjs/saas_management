package entity

import "time"

type OrgUser struct {
	ID          string
	UserID      string
	OrgID       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedByID string
	UpdatedByID string
	DeletedAt   *time.Time
	Version     int32
}
