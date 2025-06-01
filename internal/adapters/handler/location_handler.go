package handler

import (
	"context"
	"github.com/SevgiF/location-api/internal/core/location/domain"
	"github.com/SevgiF/location-api/internal/core/location/ports/inbound"
	"github.com/SevgiF/location-api/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

type LocationHandler struct {
	service inbound.LocationInboundPort
}

func NewLocationHandler(service inbound.LocationInboundPort) *LocationHandler {
	return &LocationHandler{
		service: service,
	}
}

func (h *LocationHandler) AddLocation(c *fiber.Ctx) error {
	var location domain.Location
	err := c.BodyParser(&location)
	if err != nil {
		return err
	}

	err = validator.Validate.Validator.Struct(location)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	err = h.service.AddLocation(ctx, &location)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	response := domain.Response{
		Message: "Location added",
	}

	return c.JSON(response)
}

func (h *LocationHandler) ListLocation(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	locations, err := h.service.ListLocation(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	response := domain.Response{
		Message: "Location list",
		Data:    locations,
	}

	return c.JSON(response)
}

func (h *LocationHandler) DetailLocation(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	location, err := h.service.DetailLocation(ctx, uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	response := domain.Response{
		Message: "Location detail",
		Data:    location,
	}

	return c.JSON(response)
}

func (h *LocationHandler) UpdateLocation(c *fiber.Ctx) error {
	var location domain.Location
	err := c.BodyParser(&location)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = validator.Validate.Validator.Struct(location)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	err = h.service.UpdateLocation(ctx, &location, uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	response := domain.Response{
		Message: "Location updated",
		Data:    location,
	}

	return c.JSON(response)
}

func (h *LocationHandler) RotateLocation(c *fiber.Ctx) error {
	var target domain.TargetLocation
	err := c.BodyParser(&target)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = validator.Validate.Validator.Struct(target)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	locations, err := h.service.RouteLocation(ctx, &target)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	response := domain.Response{
		Message: "Location rotated",
		Data:    locations,
	}

	return c.JSON(response)
}
