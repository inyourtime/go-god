package handler

import (
	"fmt"
	"gopher/src/model"
	"gopher/src/service"
	"strings"

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

	fmt.Println(c.Hostname())
	fmt.Println(c.GetReqHeaders()["Authorization"])
	fmt.Println(c.OriginalURL())

	// x := map[string]string{}
	// c.QueryParser(&x)
	fmt.Println(c.Queries())
	bd := map[string]interface{}{}
	c.BodyParser(&bd)
	fmt.Println(bd)
	i := c.GetReqHeaders()["Authorization"]
	if strings.HasPrefix(i, "Bearer ") {
		token := strings.TrimPrefix(i, "Bearer ")
		fmt.Println(token)
	} else {
		fmt.Println("Invalid input format")
	}

	user, _ := h.userService.GetUsers()
	_ = user
	return c.SendString("Hello, Boat")
}

func (h authHandler) SignUp(c *fiber.Ctx) error {
	return nil
}
