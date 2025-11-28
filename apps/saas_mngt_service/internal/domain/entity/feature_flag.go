package entity

type FeatureFlag struct {
	ID          string
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
