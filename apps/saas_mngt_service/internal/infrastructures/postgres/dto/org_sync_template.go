package dto

import (
	"encoding/json"

	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/domain/entity"
	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/infrastructures/postgres/sqlc"
	pkgerr "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_err"
)

type OrgSyncTemplateDto struct{}

func NewOrgSyncTemplateDto() *OrgSyncTemplateDto {
	return &OrgSyncTemplateDto{}
}

func (m *OrgSyncTemplateDto) ModelToEntity(model *sqlc.OrgSyncTemplate) (*entity.OrgSyncTemplate, error) {
	if model == nil {
		return nil, pkgerr.NewValidationError("ModelToEntity", "model is nil", nil)
	}

	var templateType entity.TemplateType
	switch model.TemplateType {
	case sqlc.TemplateTypePERMISSION:
		templateType = entity.TemplateTypePermission
	case sqlc.TemplateTypeROLE:
		templateType = entity.TemplateTypeRole
	}

	var templateDataPermissions *entity.TemplateDataPermissions
	var templateDataRole *entity.TemplateDataRole
	var err error
	if model.TemplateData != nil {
		switch model.TemplateType {
		case sqlc.TemplateTypePERMISSION:
			err = json.Unmarshal(model.TemplateData, templateDataPermissions)
		case sqlc.TemplateTypeROLE:
			err = json.Unmarshal(model.TemplateData, templateDataRole)
		}
	}

	if err != nil {
		return nil, err
	}

	return &entity.OrgSyncTemplate{
		ID:           model.ID,
		TemplateName: model.TemplateName,
		TemplateType: templateType,
		Active:       model.Active,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		Version:      model.Version,

		TemplateDataPermissions: templateDataPermissions,
		TemplateDataRole:        templateDataRole,
	}, nil
}
func (m *OrgSyncTemplateDto) EntityToModel(e *entity.OrgSyncTemplate) (*sqlc.OrgSyncTemplate, error) {
	if e == nil {
		return nil, pkgerr.NewValidationError("EntityToModel", "entity is nil", nil)
	}

	var templateType sqlc.TemplateType
	switch e.TemplateType {
	case entity.TemplateTypePermission:
		templateType = sqlc.TemplateTypePERMISSION
	case entity.TemplateTypeRole:
		templateType = sqlc.TemplateTypeROLE
	}

	var templateData []byte
	var err error
	switch e.TemplateType {
	case entity.TemplateTypeRole:
		templateData, err = json.Marshal(e.TemplateDataRole)
	case entity.TemplateTypePermission:
		templateData, err = json.Marshal(e.TemplateDataPermissions)
	}

	if err != nil {
		return nil, err
	}

	return &sqlc.OrgSyncTemplate{
		ID:           e.ID,
		TemplateName: e.TemplateName,
		TemplateType: templateType,
		TemplateData: templateData,
		Active:       e.Active,
		CreatedAt:    e.CreatedAt,
		UpdatedAt:    e.UpdatedAt,
		Version:      e.Version,
	}, nil
}
