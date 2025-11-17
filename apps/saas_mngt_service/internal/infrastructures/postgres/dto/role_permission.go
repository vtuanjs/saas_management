package dto

import (
	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/domain/entity"
	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/infrastructures/postgres/sqlc"
	pkgerr "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_err"
	pkgstring "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_string"
)

type RolePermissionDto struct{}

func NewRolePermissionDto() *RolePermissionDto {
	return &RolePermissionDto{}
}

func (m *RolePermissionDto) ModelToEntity(model *sqlc.RolePermission) (*entity.RolePermission, error) {
	if model == nil {
		return nil, pkgerr.NewValidationError("ModelToEntity", "model is nil", nil)
	}

	return &entity.RolePermission{
		ID:           model.ID,
		UserID:       model.UserID,
		RoleID:       model.RoleID,
		PermissionID: model.PermissionID,
		ProjectID:    model.ProjectID,
		OrgID:        model.OrgID,
		CreatedAt:    model.CreatedAt,
		CreatedByID:  pkgstring.DerefString(model.CreatedByID),
	}, nil
}
func (m *RolePermissionDto) EntityToModel(entity *entity.RolePermission) (*sqlc.RolePermission, error) {
	if entity == nil {
		return nil, pkgerr.NewValidationError("EntityToModel", "entity is nil", nil)
	}

	return &sqlc.RolePermission{
		ID:           entity.ID,
		UserID:       entity.UserID,
		RoleID:       entity.RoleID,
		PermissionID: entity.PermissionID,
		ProjectID:    entity.ProjectID,
		OrgID:        entity.OrgID,
		CreatedAt:    entity.CreatedAt,
		CreatedByID:  pkgstring.PointerString(entity.CreatedByID),
	}, nil
}
