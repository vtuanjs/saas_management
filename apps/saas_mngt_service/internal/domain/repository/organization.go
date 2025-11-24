//go:generate sh -c "mockgen -source=$GOFILE -destination=mock/$(basename $GOFILE .go)_test.go -package=mock"
package repository

import (
	"context"

	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/domain/entity"
)

type OrganizationRepository interface {
	FindByID(ctx context.Context, id string) (*entity.Organization, error)
	Save(ctx context.Context, organization *entity.Organization) (*entity.Organization, error)
}
