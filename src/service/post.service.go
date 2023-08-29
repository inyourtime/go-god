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
	NewReply(userID uint, req model.ReplyRequest) error
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
		createdBy := model.CreateCreatedBy(post)
		likesBy := formatLike(post.Likes)
		comments := formatComment(post.Comments)

		response := model.PostResponse{
			ID:        post.ID,
			CreatedBy: createdBy,
			Content:   post.Content,
			LikesBy:   likesBy,
			Comments:  comments,
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

func (s postService) NewReply(userID uint, req model.ReplyRequest) error {
	reply := model.Reply{}
	reply.UserID = userID
	reply.CommentID = req.CommentID
	reply.Content = req.Content

	if err := s.postRepo.Reply(reply); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func formatLike(likes []model.Like) []model.LikeBy {
	likesBy := []model.LikeBy{}
	for _, like := range likes {
		likeBy := model.LikeBy{
			UserID:    like.UserID,
			Name:      like.User.Name + " " + like.User.Surname,
			CreatedAt: like.CreatedAt,
		}
		likesBy = append(likesBy, likeBy)
	}
	return likesBy
}

func formatReply(replies []model.Reply) []model.ReplyResponse {
	repliesRes := []model.ReplyResponse{}
	for _, reply := range replies {
		createdBy := model.CreateCreatedBy(reply)
		likesBy := formatLike(reply.Likes)
		replyRes := model.ReplyResponse{
			ID:        reply.ID,
			CreatedBy: createdBy,
			CommentID: reply.CommentID,
			Content:   reply.Content,
			LikesBy:   likesBy,
			CreatedAt: reply.CreatedAt,
			UpdatedAt: reply.UpdatedAt,
		}
		repliesRes = append(repliesRes, replyRes)
	}
	return repliesRes
}

func formatComment(comments []model.Comment) []model.CommentResponse {
	commentsRes := []model.CommentResponse{}
	for _, comment := range comments {
		createdBy := model.CreateCreatedBy(comment)
		replies := formatReply(comment.Replies)
		likesBy := formatLike(comment.Likes)
		res := model.CommentResponse{
			ID:        comment.ID,
			CreatedBy: createdBy,
			PostID:    comment.PostID,
			Content:   comment.Content,
			Replies:   replies,
			LikesBy:   likesBy,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		}
		commentsRes = append(commentsRes, res)
	}
	return commentsRes
}
