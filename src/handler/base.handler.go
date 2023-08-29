package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// type baseHandler struct {
// 	c        *fiber.Ctx
// 	userID   uint
// 	fullname string
// 	nickname string
// }

// func NewBaseHandler(c *fiber.Ctx) baseHandler {
// 	user := c.Locals("user").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	name := claims["name"].(string)
// 	surname := claims["surname"].(string)
// 	nickname := claims["nickname"].(string)
// 	id := uint(claims["id"].(float64))
// 	return baseHandler{
// 		c:        c,
// 		userID:   id,
// 		fullname: name + " " + surname,
// 		nickname: nickname,
// 	}
// }

func GetUserID(c *fiber.Ctx) uint {
	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	return uint(claims["id"].(float64))
}

func GetFullname(c *fiber.Ctx) string {
	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	return claims["name"].(string) + " " + claims["surname"].(string)
}
