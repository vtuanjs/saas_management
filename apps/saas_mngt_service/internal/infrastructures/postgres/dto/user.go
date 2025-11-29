package dto

import (
	"encoding/json"

	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/domain/entity"
	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/infrastructures/postgres/sqlc"
	pkgerr "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_err"
	pkgstring "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_string"
)

type UserDto struct{}

func NewUserDto() *UserDto {
	return &UserDto{}
}

func (m *UserDto) ModelToEntity(model *sqlc.User) (*entity.User, error) {
	if model == nil {
		return nil, pkgerr.NewValidationError("ModelToEntity", "model is nil", nil)
	}

	var avatar *entity.Attachment
	if model.Avatar != nil {
		var attachment entity.Attachment
		if err := json.Unmarshal(model.Avatar, &attachment); err != nil {
			return nil, err
		}
		avatar = &attachment
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
	}, nil
}
func (m *UserDto) EntityToModel(entity *entity.User) (*sqlc.User, error) {
	if entity == nil {
		return nil, pkgerr.NewValidationError("EntityToModel", "entity is nil", nil)
	}

	var avatar []byte
	if entity.Avatar != nil {
		data, err := json.Marshal(entity.Avatar)
		if err != nil {
			return nil, err
		}
		avatar = data
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
	}, nil
}

func (m *UserDto) EntityToResponse(entity *entity.User) {}
