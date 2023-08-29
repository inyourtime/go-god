package handler

import (
	"gopher/src/errs"
	"gopher/src/model"
	"gopher/src/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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

	err = h.postService.NewPost(GetUserID(c), request)
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

	err = h.postService.NewLike(GetUserID(c), request)
	if err != nil {
		return errs.FiberError(c, err)
	}

	return c.JSON("success")
}

func (h postHandler) Comment(c *fiber.Ctx) error {
	req := model.CommentRequest{}
	if err := c.BodyParser(&req); err != nil {
		return errs.FiberError(c, fiber.ErrUnprocessableEntity)
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return errs.FiberError(c, fiber.ErrBadRequest)
	}

	err := h.postService.NewComment(GetUserID(c), req)
	if err != nil {
		return errs.FiberError(c, err)
	}
	return c.JSON("success")
}
