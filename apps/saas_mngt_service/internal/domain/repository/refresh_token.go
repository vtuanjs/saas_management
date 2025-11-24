//go:generate sh -c "mockgen -source=$GOFILE -destination=mock/$(basename $GOFILE .go)_test.go -package=mock"
package repository

import (
	"context"

	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/domain/entity"
)

type RefreshTokenRepository interface {
	FindByID(ctx context.Context, id string) (*entity.RefreshToken, error)
	Save(ctx context.Context, refreshToken *entity.RefreshToken) (*entity.RefreshToken, error)
}
