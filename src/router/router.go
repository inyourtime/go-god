package router

import (
	authRoute "gopher/src/router/auth"
	userRoute "gopher/src/router/user"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// setup other routes
	authRoute.SetupAuthRoute(api)
	userRoute.SetupUserRoute(api)
}