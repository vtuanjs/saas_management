package repository

import "github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/domain/entity"

type OrganizationRepository interface {
	FindByID(id string) (*entity.Organization, error)
	Save(organization *entity.Organization) (*entity.Organization, error)
}
