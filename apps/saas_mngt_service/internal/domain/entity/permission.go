package entity

import "time"

type Permission struct {
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

func (p Permission) ValidateName() bool {
	for _, char := range p.Name {
		if char == ' ' {
			return false
		}
	}
	return true
}
