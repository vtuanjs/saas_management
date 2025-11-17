package dto

import (
	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/domain/entity"
	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/infrastructures/postgres/sqlc"
	pkgerr "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_err"
	pkgstring "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_string"
)

type OrgUserDto struct{}

func NewOrgUserDto() *OrgUserDto {
	return &OrgUserDto{}
}

func (m *OrgUserDto) ModelToEntity(model *sqlc.OrgUser) (*entity.OrgUser, error) {
	if model == nil {
		return nil, pkgerr.NewValidationError("ModelToEntity", "model is nil", nil)
	}

	return &entity.OrgUser{
		ID:          model.ID,
		UserID:      model.UserID,
		OrgID:       model.OrgID,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		CreatedByID: pkgstring.DerefString(model.CreatedByID),
		UpdatedByID: pkgstring.DerefString(model.UpdatedByID),
		DeletedAt:   model.DeletedAt,
		Version:     model.Version,
	}, nil
}
func (m *OrgUserDto) EntityToModel(entity *entity.OrgUser) (*sqlc.OrgUser, error) {
	if entity == nil {
		return nil, pkgerr.NewValidationError("EntityToModel", "entity is nil", nil)
	}

	return &sqlc.OrgUser{
		ID:          entity.ID,
		UserID:      entity.UserID,
		OrgID:       entity.OrgID,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		CreatedByID: pkgstring.PointerString(entity.CreatedByID),
		UpdatedByID: pkgstring.PointerString(entity.UpdatedByID),
		DeletedAt:   entity.DeletedAt,
		Version:     entity.Version,
	}, nil
}
