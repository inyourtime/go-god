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
	ID        uint              `json:"id"`
	CreatedBy CreatedBy         `json:"created_by"`
	Content   string            `json:"content"`
	LikesBy   []LikeBy          `json:"likes_by"`
	Comments  []CommentResponse `json:"comments"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
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
	LikeTo    string `json:"like_to" validate:"required,oneof=post comment reply"`
	PostID    *uint  `json:"post_id" validate:"required_without_all=CommentID ReplyID"`
	CommentID *uint  `json:"comment_id" validate:"required_without_all=PostID ReplyID"`
	ReplyID   *uint  `json:"reply_id" validate:"required_without_all=PostID CommentID"`
}

type CommentRequest struct {
	PostID  uint   `json:"post_id" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type CommentResponse struct {
	ID        uint            `json:"id"`
	CreatedBy CreatedBy       `json:"created_by"`
	PostID    uint            `json:"post_id"`
	Content   string          `json:"content"`
	Replies   []ReplyResponse `json:"replies"`
	LikesBy   []LikeBy        `json:"likes_by"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type ReplyRequest struct {
	CommentID uint   `json:"comment_id" validate:"required"`
	Content   string `json:"content" validate:"required"`
}

type ReplyResponse struct {
	ID        uint      `json:"id"`
	CreatedBy CreatedBy `json:"created_by"`
	CommentID uint      `json:"comment_id"`
	Content   string    `json:"content"`
	LikesBy   []LikeBy  `json:"likes_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Creator interface {
	GetUserID() uint
	GetName() string
	GetSurname() string
}

func CreateCreatedBy(c Creator) CreatedBy {
	createdBy := CreatedBy{
		UserID: c.GetUserID(),
		Name:   c.GetName() + " " + c.GetSurname(),
	}
	return createdBy
}

func (p Post) GetUserID() uint {
	return p.UserID
}

func (p Post) GetName() string {
	return p.User.Name
}

func (p Post) GetSurname() string {
	return p.User.Surname
}

func (c Comment) GetUserID() uint {
	return c.UserID
}

func (c Comment) GetName() string {
	return c.User.Name
}

func (c Comment) GetSurname() string {
	return c.User.Surname
}

func (r Reply) GetUserID() uint {
	return r.UserID
}

func (r Reply) GetName() string {
	return r.User.Name
}

func (r Reply) GetSurname() string {
	return r.User.Surname
}
