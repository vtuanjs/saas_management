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
}

func (r Role) ValidateName() bool {
	for _, char := range r.Name {
		if char == ' ' {
			return false
		}
	}
	return true
}

type RolePermission struct {
	UserID       string
	RoleID       string
	PermissionID string
	OrgID        string
	Resources    []string

	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedByID string
	UpdatedByID string
}
