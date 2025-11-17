package dto

import (
	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/domain/entity"
	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/infrastructures/postgres/sqlc"
	pkgerr "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_err"
	pkgstring "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_string"
)

type RoleDto struct{}

func NewRoleDto() *RoleDto {
	return &RoleDto{}
}

func (m *RoleDto) ModelToEntity(model *sqlc.Role) (*entity.Role, error) {
	if model == nil {
		return nil, pkgerr.NewValidationError("ModelToEntity", "model is nil", nil)
	}

	return &entity.Role{
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
func (m *RoleDto) EntityToModel(entity *entity.Role) (*sqlc.Role, error) {
	if entity == nil {
		return nil, pkgerr.NewValidationError("EntityToModel", "entity is nil", nil)
	}

	return &sqlc.Role{
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
