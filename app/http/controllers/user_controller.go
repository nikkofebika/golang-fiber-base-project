package controllers

import (
	"golang-fiber-base-project/app/http/requests"
	"golang-fiber-base-project/app/http/resources"
	"golang-fiber-base-project/app/services"
	"golang-fiber-base-project/app/validators"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type userController struct {
	service services.UserServiceInterface
}

func NewUserController(service services.UserServiceInterface) *userController {
	return &userController{service}
}

func (controller *userController) Index(c *fiber.Ctx) error {
	users, err := controller.service.FindAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return resources.ToResponsePagination(c, users, &resources.Meta{})
	// return c.JSON(users)
}

func (controller *userController) Show(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user, err := controller.service.FindByID(uint(id))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return resources.ToResponse(c, user)
}

func (controller *userController) Create(c *fiber.Ctx) error {
	request, err := validators.ValidateBody[requests.UserCreateRequest](c)
	if err != nil {
		return resources.ToResponse(c, err)
	}

	if err := controller.service.Create(request); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return resources.ToResponseCreated(c)
}
