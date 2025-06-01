package main

import (
	"github.com/SevgiF/location-api/internal/core/location"
	"github.com/SevgiF/location-api/pkg/database/mysql"
	"github.com/SevgiF/location-api/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"strings"
)

func main() {
	//validation
	validator.Validate = validator.InitValidator()
	//mysql db connection
	db := mysql.NewGormDBManager().DB

	// Fiber settings
	fiberConfig := fiber.Config{
		AppName:       "Location API",
		CaseSensitive: true,
		StrictRouting: true,
		BodyLimit:     1024 * 1024 * 10,
	}

	// Create fiber app
	app := fiber.New(fiberConfig)

	// CORS settings
	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOriginsFunc: nil,
		AllowOrigins:     "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodOptions,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: false,
	}))

	// Version
	v1 := app.Group("/v1")

	// DOMAINS
	admin := v1.Group("/location")

	location.SetupLocation(admin, db)

	err := app.Listen(":80")
	if err != nil {
		panic(err)
	}
}
