package routes

import (
	"golang-fiber-base-project/app/http/controllers"
	"golang-fiber-base-project/app/http/middlewares"
	"golang-fiber-base-project/config"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app            *fiber.App
	config         *config.AppConfig
	authController *controllers.AuthController
	userController *controllers.UserController
}

func NewRouter(
	app *fiber.App,
	config *config.AppConfig,
	authController *controllers.AuthController,
	userController *controllers.UserController,
) *Router {
	return &Router{
		app:            app,
		config:         config,
		authController: authController,
		userController: userController,
	}
}

// SetupRoutes configures all the routes in the application
func (r *Router) SetupRoutes() {
	// Auth routes
	r.setupAuthRoutes()

	// User routes
	r.setupUserRoutes()
}

// setupAuthRoutes configures all the authentication routes
func (r *Router) setupAuthRoutes() {
	authRoute := r.app.Group("auth")
	authRoute.Post("login", r.authController.Login)
	authRoute.Post("register", r.authController.Register)
}

// setupUserRoutes configures all the user management routes
func (r *Router) setupUserRoutes() {
	userRoute := r.app.Group("users", middlewares.AuthMiddleware(r.config))
	userRoute.Get("/", r.userController.Index)
	userRoute.Get(":id", r.userController.Show)
	userRoute.Post("/", r.userController.Create)
	userRoute.Patch(":id", r.userController.Update)
	userRoute.Delete(":id", r.userController.Delete)
}
