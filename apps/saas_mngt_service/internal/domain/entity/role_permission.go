package entity

import "time"

type RolePermission struct {
	ID           string
	UserID       string
	RoleID       string
	PermissionID string
	ProjectID    string
	OrgID        string
	CreatedAt    time.Time
	CreatedByID  string
}
