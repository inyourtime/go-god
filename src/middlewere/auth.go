package middlewere

import (
	"gopher/src/coreplugins"
	"gopher/src/errs"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Authenticate() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(coreplugins.Config.JwtSecret)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return errs.FiberError(c, fiber.ErrUnauthorized)
		},
	})
}
