package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	User     User
	UserID   uint
	Content  string
	Likes    []Like
	Comments []Comment
}

type Like struct {
	gorm.Model
	User      User
	UserID    uint
	PostID    *uint
	CommentID *uint
	ReplyID   *uint
}

type Comment struct {
	gorm.Model
	User    User
	UserID  uint
	PostID  uint
	Content string
	Likes   []Like
	Replies []Reply
}

type Reply struct {
	gorm.Model
	User      User
	UserID    uint
	CommentID uint
	Content   string
	Likes     []Like
}

type NewPostRequest struct {
	Content string `json:"content"`
}

type PostResponse struct {
	ID        uint      `json:"id"`
	CreatedBy CreatedBy `json:"created_by"`
	Content   string    `json:"content"`
	LikesBy   []LikeBy  `json:"likes_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreatedBy struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
}

type LikeBy struct {
	UserID    uint      `json:"user_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type LikeRequest struct {
	PostID    *uint `json:"post_id" validate:"required_without_all=CommentID ReplyID"`
	CommentID *uint `json:"comment_id" validate:"required_without_all=PostID ReplyID"`
	ReplyID   *uint `json:"reply_id" validate:"required_without_all=PostID CommentID"`
}
