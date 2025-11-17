package entity

import "time"

type Project struct {
	ID          string
	Name        string
	Code        string
	ParentID    string
	Logo        *Attachment
	OrgID       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedByID string
	UpdatedByID string
	DeletedAt   *time.Time
	Version     int32
}
