package entity

type CustomFieldType int

const (
	CustomFieldUnknown CustomFieldType = iota
	CustomFieldTypeText
	CustomFieldTypeNumber
	CustomFieldTypeSelect
	CustomFieldTypeCheckbox
)

type CustomFieldEntityType int

const (
	CustomFieldEntityTypeUnknown CustomFieldEntityType = iota
	CustomFieldEntityTypeUser
	CustomFieldEntityTypeOrganization
	CustomFieldEntityTypeOrgFeatureFlag
)

type CustomField struct {
	ID          string
	OrgID       string
	Name        string
	Description string
	Required    bool
	Type        CustomFieldType
	EntityType  CustomFieldEntityType
	Options     []string
}

func (c CustomField) ValidateName() bool {
	for _, char := range c.Name {
		if char == ' ' {
			return false
		}
	}
	return true
}

type CustomFieldValue struct {
	ID            string
	CustomFieldID string
	EntityID      string
	Value         string
}
