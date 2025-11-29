package entity

import "time"

type Organization struct {
	Code     string
	Name     string
	Logo     *Attachment
	IsLocked bool

	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedByID string
	UpdatedByID string
}

type OrganizationMembership struct {
	UserID string
	OrgID  string

	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedByID string
	UpdatedByID string
}

type OrganizationFeatureFlag struct {
	OrgID         string
	FeatureFlagID string
	Enable        bool

	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedByID string
	UpdatedByID string
}

type OrganizationFeatureConfig struct {
	OrgID            string
	OrgFeatureFlagID string
	FeatureConfigID  string
	Value            string

	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedByID string
	UpdatedByID string
}
