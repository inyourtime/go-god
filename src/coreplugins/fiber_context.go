package coreplugins

import (
	"github.com/gofiber/fiber/v2"
)

var Ctx *fiber.Ctx

func SetContext(ctx *fiber.Ctx) {
	Ctx = ctx
}
