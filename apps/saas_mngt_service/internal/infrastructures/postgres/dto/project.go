package dto

import (
	"encoding/json"

	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/domain/entity"
	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/infrastructures/postgres/sqlc"
	pkgerr "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_err"
	pkgstring "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_string"
)

type ProjectDto struct{}

func NewProjectDto() *ProjectDto {
	return &ProjectDto{}
}

func (m *ProjectDto) ModelToEntity(model *sqlc.Project) (*entity.Project, error) {
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

	return &entity.Project{
		ID:          model.ID,
		Name:        model.Name,
		Code:        model.Code,
		ParentID:    pkgstring.DerefString(model.ParentID),
		Logo:        logo,
		OrgID:       model.OrgID,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		CreatedByID: pkgstring.DerefString(model.CreatedByID),
		UpdatedByID: pkgstring.DerefString(model.UpdatedByID),
		DeletedAt:   model.DeletedAt,
		Version:     model.Version,
	}, nil
}
func (m *ProjectDto) EntityToModel(entity *entity.Project) (*sqlc.Project, error) {
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

	return &sqlc.Project{
		ID:          entity.ID,
		Name:        entity.Name,
		Code:        entity.Code,
		ParentID:    pkgstring.PointerString(entity.ParentID),
		Logo:        logo,
		OrgID:       entity.OrgID,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		CreatedByID: pkgstring.PointerString(entity.CreatedByID),
		UpdatedByID: pkgstring.PointerString(entity.UpdatedByID),
		DeletedAt:   entity.DeletedAt,
		Version:     entity.Version,
	}, nil
}
