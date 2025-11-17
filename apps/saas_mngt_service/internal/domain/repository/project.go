package repository

import "github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/domain/entity"

type ProjectRepository interface {
	FindByID(id string) (*entity.Project, error)
	Save(project *entity.Project) (*entity.Project, error)
}
