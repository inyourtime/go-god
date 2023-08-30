package userRoute

import (
	"gopher/src/database"
	"gopher/src/handler"
	"gopher/src/middlewere"
	"gopher/src/repository"
	"gopher/src/service"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoute(router fiber.Router) {
	user := router.Group("/user")

	// instance other repository
	userRepo := repository.NewUserRepository(database.SqlDB)

	// instance other service
	userService := service.NewUserService(userRepo)

	// instance other handler
	userHandler := handler.NewUserHandler(userService)

	user.Get("/", middlewere.Authenticate(), userHandler.GetList)
	user.Get("/:id/detail", middlewere.Authenticate(), userHandler.GetDetail)
	user.Put("/:id/detail", middlewere.Authenticate(), userHandler.UpdateDetail)
	user.Post("/profile/image", middlewere.Authenticate(), userHandler.UpdateProfile)
}
