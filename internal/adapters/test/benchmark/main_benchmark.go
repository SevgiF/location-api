package benchmark

import (
	"github.com/SevgiF/location-api/internal/adapters/handler"
	"github.com/SevgiF/location-api/internal/adapters/repository"
	"github.com/SevgiF/location-api/internal/adapters/test"
	"github.com/SevgiF/location-api/internal/core/location/service"
	"github.com/SevgiF/location-api/pkg/database/mysql"
	"github.com/SevgiF/location-api/pkg/validator"
	"gorm.io/gorm"
)

var locationHandler *handler.LocationHandler

var DB = InitBenchmark()

func InitBenchmark() *gorm.DB {
	validator.Validate = validator.InitValidator()
	db := mysql.NewGormDBManager().DB

	test.SeedingForTest(db)

	initLocationModule(db)
	return db
}

func initLocationModule(db *gorm.DB) {
	repo := repository.NewLocationAdapter(db)
	locationService := service.NewLocationService(repo)
	locationHandler = handler.NewLocationHandler(locationService)
}
