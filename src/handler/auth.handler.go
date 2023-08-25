package handler

import (
	"gopher/src/errs"
	"gopher/src/model"
	"gopher/src/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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
		return errs.FiberError(c, fiber.ErrUnprocessableEntity)
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		return errs.FiberError(c, fiber.ErrBadRequest)
	}

	response, err := h.userService.Login(request)
	if err != nil {
		if appErr, ok := err.(*fiber.Error); ok {
			return errs.FiberError(c, appErr)
		}
		return errs.FiberError(c, fiber.ErrInternalServerError)
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
		if appErr, ok := err.(*fiber.Error); ok {
			return errs.FiberError(c, appErr)
		}
		return errs.FiberError(c, fiber.ErrInternalServerError)
	}
	return c.JSON(response)
}

func (h authHandler) Test(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	surname := claims["surname"].(string)
	return c.JSON(fiber.Map{
		"hello": name + surname,
	})
}
