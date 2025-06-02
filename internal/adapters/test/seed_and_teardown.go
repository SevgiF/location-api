package test

import (
	"github.com/SevgiF/location-api/internal/core/location/domain"
	"gorm.io/gorm"
	"time"
)

func SeedingForTest(db *gorm.DB) {
	locations := []domain.Location{
		{
			Name:        "Test Location 1",
			Latitude:    40.712776,
			Longitude:   -74.005974,
			MarkerColor: "ff0000",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Test Location 2",
			Latitude:    34.052235,
			Longitude:   -118.243683,
			MarkerColor: "00ff00",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	if err := db.Create(&locations).Error; err != nil {
		panic(err)
	}
}

func TeardownTest(db *gorm.DB) {
	// Tüm test verilerini silmek için "Test Location" ismini içeren kayıtları kaldırıyoruz
	if err := db.Where("name LIKE ?", "Test Location%").Delete(&domain.Location{}).Error; err != nil {
		panic(err)
	}
}
