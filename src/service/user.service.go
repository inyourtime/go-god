package service

import (
	"gopher/src/coreplugins"
	"gopher/src/model"
	"gopher/src/repository"
	"log"
)

type UserService interface {
	GetUsers() ([]model.UserResponse, error)
	GetUser(uint) (*model.UserResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return userService{userRepo: userRepo}
}

func (s userService) GetUsers() ([]model.UserResponse, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		log.Println(err)
		coreplugins.WebhookSend(coreplugins.NewDiscord(), err.Error())
		return nil, err
	}

	userResponses := []model.UserResponse{}
	for _, user := range users {
		userResponse := model.UserResponse{
			ID: user.ID,
			Email: user.Email,
			Name: user.Name,
			Surname: user.Surname,
			Nickname: user.Nickname,
			Age: user.Age,
			Gender: user.Gender,
			UpdatedAt: user.UpdatedAt,
		}
		userResponses = append(userResponses, userResponse)
	}
	return userResponses, nil
}

func (s userService) GetUser(id uint) (*model.UserResponse, error) {
	user, err := s.userRepo.GetById(id)
	if err != nil {
		log.Println(err)
		coreplugins.WebhookSend(coreplugins.NewDiscord(), err.Error())
		return nil, err
	}
	userResponse := model.UserResponse{
		ID: user.ID,
		Email: user.Email,
		Name: user.Name,
		Surname: user.Surname,
		Nickname: user.Nickname,
		Age: user.Age,
		Gender: user.Gender,
		UpdatedAt: user.UpdatedAt,
	}
	return &userResponse, nil
}