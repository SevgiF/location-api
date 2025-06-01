package outbound

import (
	"context"
	"github.com/SevgiF/location-api/internal/core/location/domain"
)

type LocationOutboundPort interface {
	Create(ctx context.Context, location *domain.Location) error
	List(ctx context.Context) ([]domain.Location, error)
	DetailById(ctx context.Context, id uint) (*domain.Location, error)
	Update(ctx context.Context, location *domain.Location, id uint) error
}
