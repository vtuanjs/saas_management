package dto

import (
	"encoding/json"

	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/domain/entity"
	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/infrastructures/postgres/sqlc"
	pkgerr "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_err"
)

type RefreshTokenDto struct{}

func NewRefreshTokenDto() *RefreshTokenDto {
	return &RefreshTokenDto{}
}

func (m *RefreshTokenDto) ModelToEntity(model *sqlc.RefreshToken) (*entity.RefreshToken, error) {
	if model == nil {
		return nil, pkgerr.NewValidationError("ModelToEntity", "model is nil", nil)
	}

	var device *entity.Device
	if model.Device != nil {
		var d entity.Device
		if err := json.Unmarshal(model.Device, &d); err != nil {
			return nil, err
		}
		device = &d
	}

	return &entity.RefreshToken{
		ID:        model.ID,
		TokenID:   model.TokenID,
		ExpiresAt: model.ExpiresAt,
		Device:    device,
		IP:        model.Ip,
		UserID:    model.UserID,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		Version:   model.Version,
	}, nil
}
func (m *RefreshTokenDto) EntityToModel(entity *entity.RefreshToken) (*sqlc.RefreshToken, error) {
	if entity == nil {
		return nil, pkgerr.NewValidationError("EntityToModel", "entity is nil", nil)
	}

	var device []byte
	if entity.Device != nil {
		data, err := json.Marshal(entity.Device)
		if err != nil {
			return nil, err
		}
		device = data
	}

	return &sqlc.RefreshToken{
		ID:        entity.ID,
		TokenID:   entity.TokenID,
		ExpiresAt: entity.ExpiresAt,
		Device:    device,
		Ip:        entity.IP,
		UserID:    entity.UserID,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		Version:   entity.Version,
	}, nil
}
