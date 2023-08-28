package postRoute

import (
	"gopher/src/database"
	"gopher/src/handler"
	"gopher/src/middlewere"
	"gopher/src/repository"
	"gopher/src/service"

	"github.com/gofiber/fiber/v2"
)

func SetupPostRoute(router fiber.Router) {
	auth := router.Group("/post")

	// instance other repository
	postRepo := repository.NewPostRepository(database.SqlDB)

	// instance other service
	postService := service.NewPostService(postRepo)

	// instance other handler
	postHandler := handler.NewPostHandler(postService)

	auth.Post("/", middlewere.Authenticate(), postHandler.CreateNewPost)
	auth.Get("/", middlewere.Authenticate(), postHandler.GetList)
	auth.Patch("/like", middlewere.Authenticate(), postHandler.Like)
}
