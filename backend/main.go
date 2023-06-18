package main

import (
	"letters/pkg/config"
	"letters/pkg/controllers"
	"letters/pkg/middlewares"
	"letters/pkg/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func ApiRoutes(api fiber.Router, jwt func(*fiber.Ctx) error) {
	// Letter
	api.Get("/letters", controllers.GetLetters)
	api.Get("/letter/:id", controllers.GetLetter)
	api.Post("/letter", jwt, controllers.PostLetter)

	// User
	api.Get("/user", jwt, controllers.GetUser)
	api.Get("/user/:id", controllers.GetUserId)
	api.Post("/register", controllers.PostRegisterUser)

	// Auth
	api.Post("/login", controllers.PostLogin)
}

func main() {
	config.Init()

	repository.InitConnection()

	app := fiber.New()

	// Modules
	app.Use(cors.New())

	jwt := middlewares.NewAuthMiddleware(config.Secret)

	// Routes
	app.Static("/", "./public")

	api := app.Group("/api")
	ApiRoutes(api, jwt)

	app.Listen(":" + config.Port)
}
