package router

import (
	authRoute "gopher/src/router/auth"
	postRoute "gopher/src/router/post"
	userRoute "gopher/src/router/user"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// setup other routes
	authRoute.SetupAuthRoute(api)
	userRoute.SetupUserRoute(api)
	postRoute.SetupPostRoute(api)

	// api.Post("/upload", func(c *fiber.Ctx) error {
	// 	file, err := c.FormFile("file")
	// 	if err != nil {
	// 		logs.Error(err)
	// 	}

	// 	path, err := s3.NewS3Handler().UploadFile(file)
	// 	if err != nil {
	// 		logs.Error(err)
	// 	}
	// 	fmt.Println(*path)

	// 	return c.JSON("g")
	// })
}
