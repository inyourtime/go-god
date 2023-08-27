package handler

import (
	"gopher/src/errs"
	"gopher/src/model"
	"gopher/src/service"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{userService: userService}
}

func (h userHandler) GetList(c *fiber.Ctx) error {
	users, err := h.userService.GetUsers()
	if err != nil {
		return errs.FiberError(c, err)
	}
	return c.JSON(users)
}

func (h userHandler) GetDetail(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errs.FiberError(c, fiber.ErrBadRequest)
	}

	if id < 0 {
		return errs.FiberError(c, fiber.ErrBadRequest)
	}

	user, err := h.userService.GetUser(uint(id))
	if err != nil {
		return errs.FiberError(c, err)
	}
	return c.JSON(user)
}

func (h userHandler) UpdateDetail(c *fiber.Ctx) error {
	request := model.UpdateUserRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return errs.FiberError(c, fiber.ErrUnprocessableEntity)
	}
	return c.JSON(request)
}
