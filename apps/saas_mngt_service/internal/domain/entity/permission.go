package entity

import "time"

type Permission struct {
	Name        string
	Description string
	System      bool
	OrgID       string

	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedByID string
	UpdatedByID string
}

func (p Permission) ValidateName() bool {
	for _, char := range p.Name {
		if char == ' ' {
			return false
		}
	}
	return true
}
