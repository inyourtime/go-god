package repository

import (
	"gopher/src/model"

	"gorm.io/gorm"
)

type PostRepository interface {
	GetAll() ([]model.Post, error)
	Create(post model.Post) error
	Like(like model.Like) error
	Comment(comment model.Comment) error
	Reply(reply model.Reply) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return postRepository{db: db}
}

func (r postRepository) GetAll() ([]model.Post, error) {
	posts := []model.Post{}
	err := r.db.Preload("User",
	).Preload("Comments",
	).Preload("Comments.User",
	).Preload("Likes",
	).Preload("Likes.User").Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r postRepository) Create(post model.Post) error {
	err := r.db.Create(&post).Error
	if err != nil {
		return err
	}
	return nil
}

func (r postRepository) Like(like model.Like) error {
	alreadyLike := []model.Like{}
	err := r.db.Where(&like).Limit(1).Find(&alreadyLike).Error
	if err != nil {
		return err
	}

	if len(alreadyLike) == 0 {
		err := r.db.Create(&like).Error
		if err != nil {
			return err
		}
	} else {
		err := r.db.Unscoped().Delete(&alreadyLike).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (r postRepository) Comment(comment model.Comment) error {
	err := r.db.Create(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

func (r postRepository) Reply(reply model.Reply) error {
	return nil
}
