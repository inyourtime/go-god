package router

import (
	authRoute "gopher/src/router/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// setup other routes
	authRoute.SetupAuthRoute(api)
}