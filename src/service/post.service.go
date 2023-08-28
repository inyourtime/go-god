package service

import (
	"gopher/src/logs"
	"gopher/src/model"
	"gopher/src/repository"
)

type PostService interface {
	GetPosts() ([]model.PostResponse, error)
	NewPost(userID uint, request model.NewPostRequest) error
	NewLike(userID uint, request model.LikeRequest) error
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
			Name: post.User.Name + " " + post.User.Surname,
		}
		likesBy := []model.LikeBy{}
		for _, like := range post.Likes {
			likeBy := model.LikeBy{
				UserID: like.UserID,
				Name: like.User.Name + " " + like.User.Surname,
				CreatedAt: like.CreatedAt,
			}
			likesBy = append(likesBy, likeBy)
		}
		
		response := model.PostResponse{
			ID: post.ID,
			CreatedBy: createdBy,
			Content: post.Content,
			LikesBy: likesBy,
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
	if request.PostID != nil {
		like.PostID = request.PostID
	} else if request.CommentID != nil {
		like.CommentID = request.CommentID
	} else if request.ReplyID != nil {
		like.ReplyID = request.ReplyID
	}

	err := s.postRepo.Like(like)
	if err != nil {
		logs.Error(err)
		return err
	}

	return nil
}