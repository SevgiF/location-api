package location

import (
	"github.com/SevgiF/location-api/internal/adapters/handler"
	"github.com/SevgiF/location-api/internal/adapters/repository"
	"github.com/SevgiF/location-api/internal/core/location/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupLocation(app fiber.Router, db *gorm.DB) {
	//adapter
	adapter := repository.NewLocationAdapter(db)

	//service
	locationService := service.NewLocationService(adapter)

	//handler
	locationHandler := handler.NewLocationHandler(locationService)

	app.Post("/add", locationHandler.AddLocation)
	app.Get("/list", locationHandler.ListLocation)
	app.Post("/route", locationHandler.RotateLocation)

	app.Get("/:id", locationHandler.DetailLocation)
	app.Put("/:id", locationHandler.UpdateLocation)
}
