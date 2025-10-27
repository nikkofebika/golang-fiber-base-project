package exceptions

import (
	"golang-fiber-base-project/app/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case *ValidationException:
		return ctx.Status(e.StatusCode).JSON(helpers.Response[any]{
			Message: e.Message,
			Errors:  e.Errors,
		})
	case *BaseException:
		return ctx.Status(e.StatusCode).JSON(helpers.Response[any]{
			Message: e.Message,
		})
	case *fiber.Error:
		return ctx.Status(e.Code).JSON(helpers.Response[any]{
			Message: e.Error(),
		})
	default:
		code := fiber.StatusInternalServerError
		return ctx.Status(code).JSON(helpers.Response[any]{
			Message: utils.StatusMessage(code),
		})
	}
}
