package main

import (
	"fmt"
	"golang-fiber-base-project/app/exceptions"
	"golang-fiber-base-project/app/http/controllers"
	"golang-fiber-base-project/app/http/middlewares"
	"golang-fiber-base-project/app/repositories"
	"golang-fiber-base-project/app/services"
	"golang-fiber-base-project/app/validators"
	"golang-fiber-base-project/config"
	"golang-fiber-base-project/routes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.LoadAppConfig()

	config.InitLogger(&cfg)

	db, err := config.NewDB(&cfg)
	if err != nil {
		logrus.Error("Cannot connect to database: ", err)
	}

	validator := validators.NewValidator(db)

	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		Prefork:      cfg.AppPrefork,
		ErrorHandler: exceptions.ErrorHandler,
	})

	app.Use(recover.New())
	app.Use(middlewares.RequestLogMiddleware())

	// Initialize repositories
	userRepository := repositories.NewUserRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService(userRepository, &cfg)

	// Initialize controllers
	userController := controllers.NewUserController(userService, validator)
	authController := controllers.NewAuthController(authService, validator)

	// Setup routes
	router := routes.NewRouter(app, &cfg, authController, userController)
	router.SetupRoutes()

	if err := app.Listen(":" + cfg.AppPort); err != nil {
		panic(err)
	}

	fmt.Println("APP RUNNING ON PORT ", cfg.AppPort)
}
