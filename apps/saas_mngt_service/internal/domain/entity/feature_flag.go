package entity

type FeatureFlag struct {
	ID          string
	OrgID       string
	Name        string
	Description string
	Enable      bool
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
	ID            string
	OrgID         string
	FeatureFlagID string
	Name          string
	Type          FeatureFlagConfigType
}
