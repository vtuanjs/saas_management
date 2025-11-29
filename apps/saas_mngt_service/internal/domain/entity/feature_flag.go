package entity

import "time"

type FeatureFlag struct {
	OrgID       string
	Name        string
	Description string
	Enable      bool

	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedByID string
	UpdatedByID string
}

func (f FeatureFlag) ValidateName() bool {
	for _, char := range f.Name {
		if char == ' ' {
			return false
		}
	}
	return true
}

type FeatureFlagConfigType int

const (
	FeatureFlagConfigTypeUnknown FeatureFlagConfigType = iota
	FeatureFlagConfigTypeBoolean
	FeatureFlagConfigTypeString
	FeatureFlagConfigTypeNumber
)

type FeatureConfig struct {
	OrgID         string
	FeatureFlagID string
	Name          string
	Type          FeatureFlagConfigType

	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedByID string
	UpdatedByID string
}

func (f FeatureConfig) ValidateName() bool {
	for _, char := range f.Name {
		if char == ' ' {
			return false
		}
	}
	return true
}
