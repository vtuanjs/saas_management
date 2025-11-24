//go:generate sh -c "mockgen -source=$GOFILE -destination=mock/$(basename $GOFILE .go)_test.go -package=mock"
package repository

import (
	"context"

	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/domain/entity"
)

type OrgUserRepository interface {
	FindByID(ctx context.Context, id string) (*entity.OrgUser, error)
	Save(ctx context.Context, orgUser *entity.OrgUser) (*entity.OrgUser, error)
}
