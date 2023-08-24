package handler

import (
	"fmt"
	"gopher/src/model"
	"gopher/src/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	userService service.UserService
}

func NewAuthHandler(userService service.UserService) authHandler {
	return authHandler{userService: userService}
}

func (h authHandler) Login(c *fiber.Ctx) error {
	request := model.LoginRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		return fiber.ErrBadRequest
	}
	
	fmt.Println(request)
	return c.SendString("Hello, Boat")
}

func (h authHandler) SignUp(c *fiber.Ctx) error {
	return nil
}
