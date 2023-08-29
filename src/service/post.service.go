package service

import (
	"gopher/src/logs"
	"gopher/src/model"
	"gopher/src/repository"

	"github.com/gofiber/fiber/v2"
)

type PostService interface {
	GetPosts() ([]model.PostResponse, error)
	NewPost(userID uint, request model.NewPostRequest) error
	NewLike(userID uint, request model.LikeRequest) error
	NewComment(userID uint, req model.CommentRequest) error
}

type postService struct {
	postRepo repository.PostRepository
}

func NewPostService(postRepo repository.PostRepository) PostService {
	return postService{postRepo: postRepo}
}

func (s postService) GetPosts() ([]model.PostResponse, error) {
	posts, err := s.postRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	responses := []model.PostResponse{}

	for _, post := range posts {
		createdBy := model.CreatedBy{
			UserID: post.UserID,
			Name:   post.User.Name + " " + post.User.Surname,
		}
		likesBy := []model.LikeBy{}
		for _, like := range post.Likes {
			likeBy := model.LikeBy{
				UserID:    like.UserID,
				Name:      like.User.Name + " " + like.User.Surname,
				CreatedAt: like.CreatedAt,
			}
			likesBy = append(likesBy, likeBy)
		}
		coms := []model.CommentResponse{}
		for _, com := range post.Comments {
			createdBy := model.CreatedBy{
				UserID: com.UserID,
				Name:   com.User.Name + " " + com.User.Surname,
			}
			res := model.CommentResponse{
				ID:        com.ID,
				CreatedBy: createdBy,
				PostID:    com.PostID,
				Content:   com.Content,
				CreatedAt: com.CreatedAt,
				UpdatedAt: com.UpdatedAt,
			}
			coms = append(coms, res)
		}

		response := model.PostResponse{
			ID:        post.ID,
			CreatedBy: createdBy,
			Content:   post.Content,
			LikesBy:   likesBy,
			Comments:  coms,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		}
		responses = append(responses, response)
	}
	return responses, nil
}

func (s postService) NewPost(userID uint, request model.NewPostRequest) error {
	post := model.Post{}
	post.UserID = userID
	post.Content = request.Content

	err := s.postRepo.Create(post)
	if err != nil {
		logs.Error(err)
		return err
	}

	return nil
}

func (s postService) NewLike(userID uint, request model.LikeRequest) error {
	like := model.Like{}
	like.UserID = userID

	switch likeTo := request.LikeTo; likeTo {
	case "post":
		if request.PostID == nil {
			return fiber.NewError(fiber.StatusBadRequest, "post_id is required")
		}
		like.PostID = request.PostID
	case "comment":
		if request.CommentID == nil {
			return fiber.NewError(fiber.StatusBadRequest, "comment_id is required")
		}
		like.CommentID = request.CommentID
	case "reply":
		if request.ReplyID == nil {
			return fiber.NewError(fiber.StatusBadRequest, "reply_id is required")
		}
		like.ReplyID = request.ReplyID
	default:
		return fiber.NewError(fiber.StatusBadRequest, "Bad request")
	}

	err := s.postRepo.Like(like)
	if err != nil {
		logs.Error(err)
		return err
	}

	return nil
}

func (s postService) NewComment(userID uint, req model.CommentRequest) error {
	comment := model.Comment{}
	comment.UserID = userID
	comment.Content = req.Content
	comment.PostID = req.PostID

	if err := s.postRepo.Comment(comment); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
