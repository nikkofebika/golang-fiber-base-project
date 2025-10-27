package controllers

import (
	"fmt"
	"golang-fiber-base-project/app/exceptions"
	"golang-fiber-base-project/app/helpers"
	"golang-fiber-base-project/app/http/requests"
	"golang-fiber-base-project/app/services"
	"golang-fiber-base-project/app/validators"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService services.AuthServiceInterface
	validator   *validators.Validator
}

func NewAuthController(authService services.AuthServiceInterface, validator *validators.Validator) *AuthController {
	return &AuthController{authService, validator}
}

func (controller *AuthController) Login(ctx *fiber.Ctx) error {
	var request requests.LoginRequest

	if err := ctx.BodyParser(&request); err != nil {
		return exceptions.NewBadRequestException()
	}

	if err := controller.validator.Struct(request); err != nil {
		formatted := controller.validator.FormatValidationErrors(err, request)
		return exceptions.NewValidationException(formatted)
	}

	jwtToken, err := controller.authService.Login(ctx.Context(), request.Email, request.Password)
	if err != nil {
		return exceptions.NewBadRequestException(err.Error())
	}

	return helpers.NewResponse(ctx, jwtToken)
}

func (controller *AuthController) Register(ctx *fiber.Ctx) error {
	request, err := validators.ValidateBody[requests.RegisterRequest](controller.validator, ctx)
	fmt.Println("request", request)
	if err != nil {
		return err
	}

	err = controller.authService.Register(ctx.Context(), request)
	if err != nil {
		return err
	}

	return helpers.NewResponseMessage(ctx, 201, "Registered Successfully")
}
