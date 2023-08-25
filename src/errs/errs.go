package errs

import "github.com/gofiber/fiber/v2"

func FiberError(c *fiber.Ctx, err *fiber.Error) error {
	return c.Status(err.Code).JSON(err)
}
