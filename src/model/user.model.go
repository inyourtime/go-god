package model

import (
	"time"

	"gorm.io/gorm"
)

type genderType string

const (
	Male        genderType = "male"
	Female      genderType = "female"
	Unspecified genderType = "unspecified"
)

type User struct {
	gorm.Model
	Email        string `gorm:"unique;uniqueIndex"`
	Password     string
	Name         string
	Surname      string
	Nickname     *string
	Age          *int
	Gender       genderType `gorm:"type:gender_type"`
	ProfileImage *string
}

type UserResponse struct {
	ID         uint       `json:"id"`
	Email      string     `json:"email"`
	Name       string     `json:"name"`
	Surname    string     `json:"surname"`
	Nickname   *string    `json:"nickname"`
	Age        *int       `json:"age"`
	Gender     genderType `json:"gender"`
	ProfileUrl *string    `json:"profile_url"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type NewUserRequest struct {
	Email    string     `json:"email" validate:"required,email"`
	Password string     `json:"password" validate:"required"`
	Name     string     `json:"name" validate:"required"`
	Surname  string     `json:"surname" validate:"required"`
	Nickname *string    `json:"nickname"`
	Age      *int       `json:"age"`
	Gender   genderType `json:"gender" validate:"required,oneof=male female unspecified"`
}

type UpdateUserRequest struct {
	Password string     `json:"password,omitempty"`
	Name     string     `json:"name,omitempty"`
	Surname  string     `json:"surname,omitempty"`
	Nickname *string    `json:"nickname,omitempty"`
	Age      *int       `json:"age,omitempty"`
	Gender   genderType `json:"gender" validate:"omitempty,oneof=male female unspecified"`
}
