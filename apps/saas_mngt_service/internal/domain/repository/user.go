package repository

import (
	"context"

	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/domain/entity"
)

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*entity.User, error)
	Save(ctx context.Context, user *entity.User) (*entity.User, error)
}
