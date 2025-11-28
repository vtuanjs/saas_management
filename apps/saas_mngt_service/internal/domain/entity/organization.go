package entity

import "time"

type Organization struct {
	ID          string
	Code        string
	Name        string
	Logo        *Attachment
	IsLocked    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedByID string
	UpdatedByID string
	DeletedAt   *time.Time
	Version     int32
}

type OrganizationMembership struct {
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

type OrganizationFeatureFlag struct {
	ID            string
	OrgID         string
	FeatureFlagID string
	Enable        bool
}

type OrganizationFeatureConfig struct {
	ID               string
	OrgFeatureFlagID string
	FeatureConfigID  string
	OrgID            string
	Value            string
}
