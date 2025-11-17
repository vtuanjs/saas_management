package repository

import "github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/domain/entity"

type UserRepository interface {
	FindByID(id string) (*entity.User, error)
	Save(user *entity.User) (*entity.User, error)
}
