package service

import (
	"context"
	"github.com/SevgiF/location-api/internal/core/location/domain"
	"github.com/SevgiF/location-api/internal/core/location/ports/outbound"
	"github.com/SevgiF/location-api/pkg"
	"github.com/SevgiF/location-api/pkg/sort"
)

type LocationService struct {
	repo outbound.LocationOutboundPort
}

func NewLocationService(repo outbound.LocationOutboundPort) *LocationService {
	return &LocationService{
		repo: repo,
	}
}

func (s *LocationService) AddLocation(ctx context.Context, location *domain.Location) error {
	err := s.repo.Create(ctx, location)
	if err != nil {
		return err
	}
	return nil
}

func (s *LocationService) ListLocation(ctx context.Context) ([]domain.Location, error) {
	locations, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return locations, nil
}

func (s *LocationService) DetailLocation(ctx context.Context, id uint) (*domain.Location, error) {
	location, err := s.repo.DetailById(ctx, id)
	if err != nil {
		return nil, err
	}
	return location, nil
}

func (s *LocationService) UpdateLocation(ctx context.Context, location *domain.Location, id uint) error {
	err := s.repo.Update(ctx, location, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *LocationService) RouteLocation(ctx context.Context, target *domain.TargetLocation) ([]domain.LocationRouting, error) {
	//get all locations from db
	locations, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	var routes []domain.LocationRouting

	for _, location := range locations {
		distance := pkg.Haversine(target.Latitude, target.Longitude, location.Latitude, location.Longitude)
		route := domain.LocationRouting{
			Name:     location.Name,
			Distance: distance,
		}
		routes = append(routes, route)
	}

	sortedList := sort.SortingDistanceWithBubbleSort(routes, len(routes))

	return sortedList, nil
}
