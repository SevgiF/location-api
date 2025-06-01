package repository

import (
	"context"
	"errors"
	"github.com/SevgiF/location-api/internal/core/location/domain"
	"gorm.io/gorm"
)

type LocationAdapter struct {
	db *gorm.DB
}

// NewLocationAdapter creates a LocationAdapter with a given database connection.
func NewLocationAdapter(db *gorm.DB) *LocationAdapter {
	return &LocationAdapter{
		db: db,
	}
}

func (r *LocationAdapter) Create(ctx context.Context, location *domain.Location) error {
	result := r.db.WithContext(ctx).Create(location)

	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("failed to create location: " + result.Error.Error())
	}
	return nil
}

func (r *LocationAdapter) List(ctx context.Context) ([]domain.Location, error) {
	var locations []domain.Location
	err := r.db.WithContext(ctx).Find(&locations).Error
	if err != nil {
		return nil, err
	}
	return locations, nil
}

func (r *LocationAdapter) DetailById(ctx context.Context, id uint) (*domain.Location, error) {
	var location *domain.Location
	err := r.db.WithContext(ctx).First(&location, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return location, nil
}

func (r *LocationAdapter) Update(ctx context.Context, location *domain.Location, id uint) error {
	err := r.db.WithContext(ctx).Model(&domain.Location{}).Where("id = ?", id).Updates(location).Error
	if err != nil {
		return err
	}
	return nil
}
