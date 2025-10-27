package controllers

import (
	"fmt"
	"golang-fiber-base-project/app/exceptions"
	"golang-fiber-base-project/app/helpers"
	"golang-fiber-base-project/app/http/requests"
	"golang-fiber-base-project/app/services"
	"golang-fiber-base-project/app/validators"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service   services.UserServiceInterface
	validator *validators.Validator
}

func NewUserController(service services.UserServiceInterface, validator *validators.Validator) *UserController {
	return &UserController{service, validator}
}

func (controller *UserController) Index(ctx *fiber.Ctx) error {
	users, err := controller.service.FindAll(ctx.Context())
	if err != nil {
		return err
	}

	return helpers.NewResponsePagination(ctx, users, &helpers.Meta{})
}

func (controller *UserController) Show(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	user, err := controller.service.FindByID(ctx.Context(), uint(id))
	if err != nil {
		return err
	}

	return helpers.NewResponse(ctx, user)
}

func (controller *UserController) Create(ctx *fiber.Ctx) error {
	request, err := validators.ValidateBody[requests.UserCreateRequest](controller.validator, ctx)
	if err != nil {
		return exceptions.NewBadRequestException()
	}

	if err := controller.service.Create(ctx.Context(), request); err != nil {
		return err
	}

	return helpers.NewResponseCreated(ctx)
}

func (controller *UserController) Update(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	fmt.Println("idku", id)

	request, err := validators.ValidateBody[requests.UserUpdateRequest](controller.validator, ctx)
	if err != nil {
		return exceptions.NewBadRequestException()
	}

	if err := controller.service.Update(ctx.Context(), uint(id), request); err != nil {
		return fiber.ErrInternalServerError
	}

	return helpers.NewResponseUpdated(ctx)
}

func (controller *UserController) Delete(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	if err := controller.service.Delete(ctx.Context(), uint(id)); err != nil {
		return fiber.ErrInternalServerError
	}

	return helpers.NewResponseDeleted(ctx)

}
