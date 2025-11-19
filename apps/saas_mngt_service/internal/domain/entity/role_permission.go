package entity

import "time"

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
