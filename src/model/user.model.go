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
	Email    string
	Password string
	Name     string
	Surname  string
	Nickname string
	Age      int
	Gender   genderType `gorm:"type:gender_type"`
}

type UserResponse struct {
	ID        uint       `json:"id"`
	Email     string     `json:"email"`
	Name      string     `json:"name"`
	Surname   string     `json:"surname"`
	Nickname  string     `json:"nickname"`
	Age       int        `json:"age"`
	Gender    genderType `json:"gender"`
	UpdatedAt time.Time  `json:"updated_at"`
}
