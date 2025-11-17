package repository

import "github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/domain/entity"

type RefreshTokenRepository interface {
	FindByID(id string) (*entity.RefreshToken, error)
	Save(refreshToken *entity.RefreshToken) (*entity.RefreshToken, error)
}
