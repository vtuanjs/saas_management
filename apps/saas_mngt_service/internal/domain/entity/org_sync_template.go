package entity

import "time"

type TemplateType int

const (
	TemplateTypePermission TemplateType = iota + 1
	TemplateTypeRole
)

type TemplateDataPermission struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TemplateDataPermissions = []TemplateDataPermission

type TemplateDataRole struct {
	Code        string   `json:"code"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
}

type OrgSyncTemplate struct {
	ID           string
	TemplateName string
	TemplateType TemplateType
	Active       bool
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	Version      int32

	TemplateDataPermissions *TemplateDataPermissions
	TemplateDataRole        *TemplateDataRole
}
