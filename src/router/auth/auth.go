package authRoute

import (
	"gopher/src/database"
	"gopher/src/handler"
	"gopher/src/middlewere"
	"gopher/src/repository"
	"gopher/src/service"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoute(router fiber.Router) {
	auth := router.Group("/auth")

	// instance other repository
	userRepo := repository.NewUserRepositoryDB(database.SqlDB)

	// instance other service
	userService := service.NewUserService(userRepo)

	// instance other handler
	authHandler := handler.NewAuthHandler(userService)

	auth.Post("/login", authHandler.Login)
	auth.Post("/signup", authHandler.SignUp)
	auth.Get("/test", middlewere.Authenticate(), authHandler.Test)
}
