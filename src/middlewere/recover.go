package middlewere

import (
	"fmt"
	"gopher/src/logs"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Recover() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				// Log the error
				t := fmt.Sprintf("Panic recovered: %v", r)
				log.Println(t)

				// Respond with a 500 status code to the client
				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"code":  fiber.StatusInternalServerError,
					"error": "Internal Server Error",
				})

				logs.Error(r)
			}
		}()
		// Next is called to execute the actual route handler
		return c.Next()
	}
}
