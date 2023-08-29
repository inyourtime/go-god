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
	post := router.Group("/post")

	// instance other repository
	postRepo := repository.NewPostRepository(database.SqlDB)

	// instance other service
	postService := service.NewPostService(postRepo)

	// instance other handler
	postHandler := handler.NewPostHandler(postService)

	post.Post("/", middlewere.Authenticate(), postHandler.CreateNewPost)
	post.Get("/", middlewere.Authenticate(), postHandler.GetList)
	post.Patch("/like", middlewere.Authenticate(), postHandler.Like)
	post.Post("/comment", middlewere.Authenticate(), postHandler.Comment)
	post.Post("/comment/reply", middlewere.Authenticate(), postHandler.Reply)
}
