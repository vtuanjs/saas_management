package entity

import "time"

type RefreshToken struct {
	ID        string
	TokenID   string
	ExpiresAt *time.Time
	Device    *Device
	IP        string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int32
}
