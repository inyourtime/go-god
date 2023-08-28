package service

import (
	"gopher/src/coreplugins"
	"gopher/src/logs"
	"gopher/src/model"
	"gopher/src/repository"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUsers() ([]model.UserResponse, error)
	GetUser(id uint) (*model.UserResponse, error)
	NewUser(request model.NewUserRequest) (*model.UserResponse, error)
	Login(request model.LoginRequest) (*model.LoginResponse, error)
	UpdateUser(id uint, request model.UpdateUserRequest) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return userService{userRepo: userRepo}
}

func (s userService) Login(request model.LoginRequest) (*model.LoginResponse, error) {
	user, err := s.userRepo.GetByEmail(request.Email)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	if user == nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	ok := coreplugins.CheckPasswordHash(request.Password, user.Password)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	accessClaims := jwt.MapClaims{
		"id":       user.ID,
		"email":    user.Email,
		"name":     user.Name,
		"surname":  user.Surname,
		"nickname": user.Nickname,
		"exp":      coreplugins.AccessTokenExpireTime(),
	}
	refreshClaims := jwt.MapClaims{}
	for key, value := range accessClaims {
		refreshClaims[key] = value
	}

	refreshClaims["exp"] = coreplugins.RefreshTokenExpireTime()
	accessToken, err := coreplugins.Token(accessClaims, coreplugins.Config.JwtSecret)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	refreshToken, err := coreplugins.Token(refreshClaims, coreplugins.Config.JwtSecret)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	response := model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &response, nil
}

func (s userService) NewUser(request model.NewUserRequest) (*model.UserResponse, error) {
	userExist, err := s.userRepo.GetByEmail(request.Email)
	if userExist != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "email already exist")
	}

	hash, err := coreplugins.HashPassword(request.Password)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	user := model.User{
		Email:    request.Email,
		Password: hash,
		Name:     request.Name,
		Surname:  request.Surname,
		Nickname: request.Nickname,
		Age:      request.Age,
		Gender:   request.Gender,
	}

	newUser, err := s.userRepo.Create(user)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	response := model.UserResponse{
		ID:        newUser.ID,
		Email:     newUser.Email,
		Name:      newUser.Name,
		Surname:   newUser.Surname,
		Nickname:  newUser.Nickname,
		Age:       newUser.Age,
		Gender:    newUser.Gender,
		UpdatedAt: newUser.UpdatedAt,
	}
	return &response, nil
}

func (s userService) GetUsers() ([]model.UserResponse, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	userResponses := []model.UserResponse{}
	for _, user := range users {
		userResponse := model.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			Name:      user.Name,
			Surname:   user.Surname,
			Nickname:  user.Nickname,
			Age:       user.Age,
			Gender:    user.Gender,
			UpdatedAt: user.UpdatedAt,
		}
		userResponses = append(userResponses, userResponse)
	}
	return userResponses, nil
}

func (s userService) GetUser(id uint) (*model.UserResponse, error) {
	user, err := s.userRepo.GetById(id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	if user == nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	userResponse := model.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Surname:   user.Surname,
		Nickname:  user.Nickname,
		Age:       user.Age,
		Gender:    user.Gender,
		UpdatedAt: user.UpdatedAt,
	}
	return &userResponse, nil
}

// ไว้มาแก้เรื่อง update ทีหลัง
func (s userService) UpdateUser(id uint, request model.UpdateUserRequest) error {
	user := model.User{}
	if !reflect.ValueOf(request.Password).IsZero() {
		hash, err := coreplugins.HashPassword(request.Password)
		if err != nil {
			logs.Error(err)
			return err
		}
		user.Password = hash
	}
	if !reflect.ValueOf(request.Name).IsZero() {
		user.Name = request.Name
	}
	if !reflect.ValueOf(request.Surname).IsZero() {
		user.Surname = request.Surname
	}
	if request.Nickname != nil {
		user.Nickname = request.Nickname
	}
	if request.Age != nil {
		user.Age = request.Age
	}
	if !reflect.ValueOf(request.Gender).IsZero() {
		user.Gender = request.Gender
	}

	err := s.userRepo.Update(id, user)
	if err != nil {
		return err
	}

	return nil
}
