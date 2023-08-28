package handler

import (
	"gopher/src/errs"
	"gopher/src/model"
	"gopher/src/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type postHandler struct {
	postService service.PostService
}

func NewPostHandler(postService service.PostService) postHandler {
	return postHandler{postService: postService}
}

func (h postHandler) GetList(c *fiber.Ctx) error {
	posts, err := h.postService.GetPosts()
	if err != nil {
		return errs.FiberError(c, err)
	}
	// fmt.Println(posts)
	return c.JSON(posts)
}

func (h postHandler) CreateNewPost(c *fiber.Ctx) error {
	request := model.NewPostRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return errs.FiberError(c, fiber.ErrUnprocessableEntity)
	}

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := uint(claims["id"].(float64))

	err = h.postService.NewPost(id, request)
	if err != nil {
		return errs.FiberError(c, err)
	}
	return c.JSON("sd")
}

func (h postHandler) Like(c *fiber.Ctx) error {
	request := model.LikeRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return errs.FiberError(c, fiber.ErrUnprocessableEntity)
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		return errs.FiberError(c, fiber.ErrBadRequest)
	}

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := uint(claims["id"].(float64))

	err = h.postService.NewLike(id, request)
	if err != nil {
		return errs.FiberError(c, err)
	}

	return c.JSON("good")
}
