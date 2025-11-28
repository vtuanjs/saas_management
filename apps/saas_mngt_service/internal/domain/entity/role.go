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

func (r Role) ValidateName() bool {
	for _, char := range r.Name {
		if char == ' ' {
			return false
		}
	}
	return true
}

type RolePermission struct {
	ID           string
	UserID       string
	RoleID       string
	PermissionID string
	OrgID        string
	Resources    []string
	CreatedAt    time.Time
	CreatedByID  string
}
