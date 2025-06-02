package location_helpers

import (
	"github.com/SevgiF/location-api/internal/adapters/test/helpers"
	"github.com/gofiber/fiber/v2"
)

func GetRequestsResponsesForCreateLocation() []helpers.RequestAndResponses {
	return []helpers.RequestAndResponses{
		// Nil request
		{
			Request: nil,
			ResponseBody: &[]helpers.ValidationResponse{
				{Field: "Name", Message: "Field Name is required"},
				{Field: "Latitude", Message: "Field Latitude is required"},
				{Field: "Longitude", Message: "Field Longitude is required"},
				{Field: "MarkerColor", Message: "Field MarkerColor is required"},
			},
			StatusCode: fiber.StatusBadRequest,
		},
		// Empty fields
		{
			Request: &helpers.Location{
				Name:        "",
				Latitude:    0,
				Longitude:   0,
				MarkerColor: "",
			},
			ResponseBody: &[]helpers.ValidationResponse{
				{Field: "Name", Message: "Field Name is required"},
				{Field: "Latitude", Message: "Field Latitude is required"},
				{Field: "Longitude", Message: "Field Longitude is required"},
				{Field: "MarkerColor", Message: "Field MarkerColor is required"},
			},
			StatusCode: fiber.StatusBadRequest,
		},
		// Invalid color format
		{
			Request: &helpers.Location{
				Name:        "Test Location",
				Latitude:    40.7128,
				Longitude:   -74.0060,
				MarkerColor: "zzzzzz",
			},
			ResponseBody: &[]helpers.ValidationResponse{
				{Field: "MarkerColor", Message: "Field MarkerColor is not valid hexadecimal"},
			},
			StatusCode: fiber.StatusBadRequest,
		},
		// Name too short
		{
			Request: &helpers.Location{
				Name:        "A",
				Latitude:    40.7128,
				Longitude:   -74.0060,
				MarkerColor: "00ff00",
			},
			ResponseBody: &[]helpers.ValidationResponse{
				{Field: "Name", Message: "Field Name must be at least 3 characters"},
			},
			StatusCode: fiber.StatusBadRequest,
		},
		// Valid request
		{
			Request: &helpers.Location{
				Name:        "Test Location",
				Latitude:    40.7128,
				Longitude:   -74.0060,
				MarkerColor: "00ff00",
			},
			ResponseBody: helpers.Response{
				Message: "Location created successfully",
				Data:    nil,
			},
			StatusCode: fiber.StatusCreated,
		},
	}
}
