package repository

import (
	"context"

	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/domain/entity"
)

type ProjectRepository interface {
	FindByID(ctx context.Context, id string) (*entity.Project, error)
	Save(ctx context.Context, project *entity.Project) (*entity.Project, error)
}
