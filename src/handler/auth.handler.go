package handler

import (
	"gopher/src/errs"
	"gopher/src/model"
	"gopher/src/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	userService service.UserService
}

func NewAuthHandler(userService service.UserService) authHandler {
	return authHandler{
		userService: userService,
	}
}

func (h authHandler) Login(c *fiber.Ctx) error {
	request := model.LoginRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return errs.FiberError(c, fiber.ErrUnprocessableEntity)
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		return errs.FiberError(c, err)
	}

	response, err := h.userService.Login(request)
	if err != nil {
		return errs.FiberError(c, err)
	}
	return c.JSON(response)
}

func (h authHandler) SignUp(c *fiber.Ctx) error {
	request := model.NewUserRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return errs.FiberError(c, fiber.ErrUnprocessableEntity)
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		return errs.FiberError(c, fiber.ErrBadRequest)
	}

	response, err := h.userService.NewUser(request)
	if err != nil {
		return errs.FiberError(c, err)
	}
	return c.JSON(response)
}

func (h authHandler) Test(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"id":    GetUserID(c),
		"hello": GetFullname(c),
	})
}
