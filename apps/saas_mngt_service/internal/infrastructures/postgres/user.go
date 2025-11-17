package postgres

import (
	"context"

	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/domain/entity"
	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/domain/repository"
	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/infrastructures/postgres/dto"
	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/infrastructures/postgres/sqlc"
)

type UserRepository struct {
	query sqlc.Querier
}

func NewUserRepository(querier sqlc.Querier) repository.UserRepository {
	return &UserRepository{
		query: querier,
	}
}

// FindByID implements repository.UserRepository.
func (u *UserRepository) FindByID(ctx context.Context, id string) (*entity.User, error) {
	result, err := u.query.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	dto := dto.NewUserDto()
	return dto.ModelToEntity(result), nil
}

// Save implements repository.UserRepository.
func (u *UserRepository) Save(ctx context.Context, user *entity.User) (*entity.User, error) {
	dto := dto.NewUserDto()
	in := dto.EntityToModel(user)

	result, err := u.query.SaveUser(ctx, sqlc.SaveUserParams{
		ID:                   in.ID,
		Email:                in.Email,
		Phone:                in.Phone,
		FirstName:            in.FirstName,
		LastName:             in.LastName,
		Avatar:               in.Avatar,
		LastLogin:            in.LastLogin,
		Ref:                  in.Ref,
		IsLocked:             in.IsLocked,
		IsActivated:          in.IsActivated,
		IsChangePassRequired: in.IsChangePassRequired,
		CreatedAt:            in.CreatedAt,
		UpdatedAt:            in.UpdatedAt,
		CreatedByID:          in.CreatedByID,
		UpdatedByID:          in.UpdatedByID,
		DeletedAt:            in.DeletedAt,
	})
	if err != nil {
		return nil, err
	}

	return dto.ModelToEntity(result), nil
}
