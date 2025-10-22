package resources

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Response[T any] struct {
	Data    *T     `json:"data,omitempty"`
	Meta    *Meta  `json:"meta,omitempty"`
	Message string `json:"message,omitempty"`
	Error   any    `json:"error,omitempty"`
}

type Meta struct {
	CurrentPage int `json:"current_page"`
	From        int `json:"from"`
	PerPage     int `json:"per_page"`
	To          int `json:"to"`
	Total       int `json:"total"`
	TotalPages  int `json:"total_pages"`
}

func ToResponse[T any](c *fiber.Ctx, data T) error {
	return c.Status(fiber.StatusOK).JSON(Response[T]{
		Data: &data,
	})
}

func ToResponsePagination[T any](c *fiber.Ctx, data []T, meta *Meta) error {
	return c.JSON(Response[[]T]{
		Data: &data,
		Meta: meta,
	})
}

func ToResponseMessage(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(Response[string]{
		Message: message,
	})
}

func ToResponseError(c *fiber.Ctx, status int, err any) error {
	fmt.Println("ToResponseError ToResponseError", err)
	return c.Status(status).JSON(Response[any]{
		Error: err,
	})
}

func ToResponseCreated(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(Response[string]{
		Message: "Data created successfully",
	})
}

func ToResponseUpdated(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(Response[string]{
		Message: "Data updated successfully",
	})
}

func ToResponseDeleted(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(Response[string]{
		Message: "Data deleted successfully",
	})
}
