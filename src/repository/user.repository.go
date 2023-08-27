package repository

import (
	"gopher/src/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetById(id uint) (*model.User, error)
	Create(user model.User) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Update(id uint, user model.User) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{db: db}
}

func (r userRepository) GetByEmail(email string) (*model.User, error) {
	user := model.User{}
	err := r.db.Where(model.User{Email: email}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepository) Create(user model.User) (*model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepository) GetAll() ([]model.User, error) {
	users := []model.User{}
	err := r.db.Order("created_at desc").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r userRepository) GetById(id uint) (*model.User, error) {
	user := model.User{}
	err := r.db.Where(id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepository) Update(id uint, user model.User) (*model.User, error) {
	return nil, nil
}
