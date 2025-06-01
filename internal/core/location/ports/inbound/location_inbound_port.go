package inbound

import (
	"context"
	"github.com/SevgiF/location-api/internal/core/location/domain"
)

type LocationInboundPort interface {
	AddLocation(ctx context.Context, location *domain.Location) error
	ListLocation(ctx context.Context) ([]domain.Location, error)
	DetailLocation(ctx context.Context, id uint) (*domain.Location, error)
	UpdateLocation(ctx context.Context, location *domain.Location, id uint) error
	RouteLocation(ctx context.Context, target *domain.TargetLocation) ([]domain.LocationRouting, error)
}
