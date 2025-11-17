package repository

import "github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/domain/entity"

type OrgUserRepository interface {
	FindByID(id string) (*entity.OrgUser, error)
	Save(orgUser *entity.OrgUser) (*entity.OrgUser, error)
}
