package entity

import "time"

type RefreshToken struct {
	TokenID   string
	ExpiresAt *time.Time
	Device    *Device
	IP        string
	UserID    string

	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}
