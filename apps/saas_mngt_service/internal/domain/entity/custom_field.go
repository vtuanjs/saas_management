package entity

import "time"

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
)

type CustomField struct {
	OrgID       string
	Name        string
	Description string
	Required    bool
	Type        CustomFieldType
	EntityType  CustomFieldEntityType
	Options     []string

	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedByID string
	UpdatedByID string
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
	CustomFieldID string
	EntityID      string
	Value         string

	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedByID string
	UpdatedByID string
}
