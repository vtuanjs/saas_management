package dto

import (
	"encoding/json"

	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/domain/entity"
	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/infrastructures/postgres/sqlc"
	pkgstring "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_string"
)

type UserDto struct{}

func NewUserDto() *UserDto {
	return &UserDto{}
}

func (m *UserDto) ModelToEntity(model *sqlc.User) *entity.User {
	if model == nil {
		return nil
	}

	var avatar *entity.Attachment
	if model.Avatar != nil {
		var attachment entity.Attachment
		if err := json.Unmarshal(model.Avatar, &attachment); err == nil {
			avatar = &attachment
		}
	}

	return &entity.User{
		ID:                   model.ID,
		Email:                model.Email,
		Phone:                model.Phone,
		FirstName:            model.FirstName,
		LastName:             model.LastName,
		Name:                 *model.Name,
		Avatar:               avatar,
		LastLogin:            model.LastLogin,
		Ref:                  pkgstring.DerefString(model.Ref),
		IsLocked:             model.IsLocked,
		IsActivated:          model.IsActivated,
		IsAdmin:              model.IsAdmin,
		IsChangePassRequired: model.IsChangePassRequired,
		CreatedAt:            model.CreatedAt,
		UpdatedAt:            model.UpdatedAt,
		CreatedByID:          *model.CreatedByID,
		UpdatedByID:          *model.UpdatedByID,
		DeletedAt:            model.DeletedAt,
		Version:              model.Version,
	}
}
func (m *UserDto) EntityToModel(entity *entity.User) *sqlc.User {
	if entity == nil {
		return nil
	}

	var avatar []byte
	if entity.Avatar != nil {
		if data, err := json.Marshal(entity.Avatar); err == nil {
			avatar = data
		}
	}

	return &sqlc.User{
		ID:                   entity.ID,
		Email:                entity.Email,
		Phone:                entity.Phone,
		FirstName:            entity.FirstName,
		LastName:             entity.LastName,
		Avatar:               avatar,
		LastLogin:            entity.LastLogin,
		Ref:                  pkgstring.PointerString(entity.Ref),
		IsLocked:             entity.IsLocked,
		IsActivated:          entity.IsActivated,
		IsAdmin:              entity.IsAdmin,
		IsChangePassRequired: entity.IsChangePassRequired,
		CreatedAt:            entity.CreatedAt,
		UpdatedAt:            entity.UpdatedAt,
		CreatedByID:          &entity.CreatedByID,
		UpdatedByID:          &entity.UpdatedByID,
		DeletedAt:            entity.DeletedAt,
		Version:              entity.Version,
	}
}

func (m *UserDto) EntityToResponse(entity *entity.User) {}
