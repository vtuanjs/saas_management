package dto

import (
	"encoding/json"

	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/domain/entity"
	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/infrastructures/postgres/sqlc"
	pkgerr "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_err"
	pkgstring "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_string"
)

type OrganizationDto struct{}

func NewOrganizationDto() *OrganizationDto {
	return &OrganizationDto{}
}

func (m *OrganizationDto) ModelToEntity(model *sqlc.Organization) (*entity.Organization, error) {
	if model == nil {
		return nil, pkgerr.NewValidationError("ModelToEntity", "model is nil", nil)
	}

	var logo *entity.Attachment
	if model.Logo != nil {
		var attachment entity.Attachment
		if err := json.Unmarshal(model.Logo, &attachment); err != nil {
			return nil, err
		}
		logo = &attachment
	}

	return &entity.Organization{
		ID:          model.ID,
		Code:        model.Code,
		Name:        model.Name,
		Logo:        logo,
		IsLocked:    model.IsLocked,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		CreatedByID: pkgstring.DerefString(model.CreatedByID),
		UpdatedByID: pkgstring.DerefString(model.UpdatedByID),
	}, nil
}
func (m *OrganizationDto) EntityToModel(entity *entity.Organization) (*sqlc.Organization, error) {
	if entity == nil {
		return nil, pkgerr.NewValidationError("EntityToModel", "entity is nil", nil)
	}

	var logo []byte
	if entity.Logo != nil {
		data, err := json.Marshal(entity.Logo)
		if err != nil {
			return nil, err
		}
		logo = data
	}

	return &sqlc.Organization{
		ID:          entity.ID,
		Code:        entity.Code,
		Name:        entity.Name,
		Logo:        logo,
		IsLocked:    entity.IsLocked,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		CreatedByID: pkgstring.PointerString(entity.CreatedByID),
		UpdatedByID: pkgstring.PointerString(entity.UpdatedByID),
	}, nil
}
