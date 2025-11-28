package dto

import (
	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/domain/entity"
	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/infrastructures/postgres/sqlc"
	pkgerr "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_err"
	pkgstring "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_string"
)

type PermissionDto struct{}

func NewPermissionDto() *PermissionDto {
	return &PermissionDto{}
}

func (m *PermissionDto) ModelToEntity(model *sqlc.Permission) (*entity.Permission, error) {
	if model == nil {
		return nil, pkgerr.NewValidationError("ModelToEntity", "model is nil", nil)
	}

	return &entity.Permission{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		System:      model.System,
		OrgID:       model.OrgID,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		CreatedByID: pkgstring.DerefString(model.CreatedByID),
		UpdatedByID: pkgstring.DerefString(model.UpdatedByID),
		DeletedAt:   model.DeletedAt,
		Version:     model.Version,
	}, nil
}
func (m *PermissionDto) EntityToModel(entity *entity.Permission) (*sqlc.Permission, error) {
	if entity == nil {
		return nil, pkgerr.NewValidationError("EntityToModel", "entity is nil", nil)
	}

	return &sqlc.Permission{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		System:      entity.System,
		OrgID:       entity.OrgID,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		CreatedByID: pkgstring.PointerString(entity.CreatedByID),
		UpdatedByID: pkgstring.PointerString(entity.UpdatedByID),
		DeletedAt:   entity.DeletedAt,
		Version:     entity.Version,
	}, nil
}
