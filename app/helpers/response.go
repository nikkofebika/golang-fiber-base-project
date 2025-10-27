package helpers

import (
	"github.com/gofiber/fiber/v2"
)

type Response[T any] struct {
	Data    *T                  `json:"data,omitempty"`
	Meta    *Meta               `json:"meta,omitempty"`
	Message string              `json:"message,omitempty"`
	Errors  map[string][]string `json:"errors,omitempty"`
}

type Meta struct {
	CurrentPage int `json:"current_page"`
	From        int `json:"from"`
	PerPage     int `json:"per_page"`
	To          int `json:"to"`
	Total       int `json:"total"`
	TotalPages  int `json:"total_pages"`
}

func NewResponse[T any](c *fiber.Ctx, data T) error {
	return c.Status(fiber.StatusOK).JSON(Response[T]{
		Data: &data,
	})
}

func NewResponsePagination[T any](c *fiber.Ctx, data []T, meta *Meta) error {
	return c.JSON(Response[[]T]{
		Data: &data,
		Meta: meta,
	})
}

func NewResponseMessage(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(Response[string]{
		Message: message,
	})
}

func NewResponseErrors(c *fiber.Ctx, status int, errs map[string][]string) error {
	// If err implements the error interface, use its Error() message.
	// var errMsg any = err

	// if e, ok := err.(error); ok {
	// 	errMsg = e.Error()
	// }

	// // Log original error for debugging.
	// fmt.Println("NewResponseError NewResponseError", err)

	return c.Status(status).JSON(Response[any]{
		Errors: errs,
	})
}

func NewResponseCreated(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(Response[string]{
		Message: "Data created successfully",
	})
}

func NewResponseUpdated(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(Response[string]{
		Message: "Data updated successfully",
	})
}

func NewResponseDeleted(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(Response[string]{
		Message: "Data deleted successfully",
	})
}
