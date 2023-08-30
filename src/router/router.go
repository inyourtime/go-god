package router

import (
	"fmt"
	"gopher/src/logs"
	authRoute "gopher/src/router/auth"
	postRoute "gopher/src/router/post"
	userRoute "gopher/src/router/user"
	s3 "gopher/src/utils"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// setup other routes
	authRoute.SetupAuthRoute(api)
	userRoute.SetupUserRoute(api)
	postRoute.SetupPostRoute(api)

	api.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			logs.Error(err)
		}
		h := file.Header["Content-Type"][0]
		fmt.Println(h)

		s3.NewS3Handler().UploadFile(file)

		return c.JSON("g")
	})
}