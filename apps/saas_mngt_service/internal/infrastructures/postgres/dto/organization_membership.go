package dto

import (
	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/domain/entity"
	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/infrastructures/postgres/sqlc"
	pkgerr "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_err"
	pkgstring "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_string"
)

type OrganizationMembershipDto struct{}

func NewOrganizationMembershipDto() *OrganizationMembershipDto {
	return &OrganizationMembershipDto{}
}

func (m *OrganizationMembershipDto) ModelToEntity(model *sqlc.OrganizationMembership) (*entity.OrganizationMembership, error) {
	if model == nil {
		return nil, pkgerr.NewValidationError("ModelToEntity", "model is nil", nil)
	}

	return &entity.OrganizationMembership{
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
func (m *OrganizationMembershipDto) EntityToModel(entity *entity.OrganizationMembership) (*sqlc.OrganizationMembership, error) {
	if entity == nil {
		return nil, pkgerr.NewValidationError("EntityToModel", "entity is nil", nil)
	}

	return &sqlc.OrganizationMembership{
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
