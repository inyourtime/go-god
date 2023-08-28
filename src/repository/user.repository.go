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
	Update(id uint, user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{db: db}
}

func (r userRepository) GetByEmail(email string) (*model.User, error) {
	user := []model.User{}
	err := r.db.Where("Email=?", email).Limit(1).Find(&user).Error
	if err != nil {
		return nil, err
	}
	if len(user) == 0 {
		return nil, nil
	}
	return &user[0], nil
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
	user := []model.User{}
	err := r.db.Where(id).Limit(1).Find(&user).Error
	if err != nil {
		return nil, err
	}
	if len(user) == 0 {
		return nil, nil
	}
	return &user[0], nil
}

func (r userRepository) Update(id uint, user model.User) error {
	err := r.db.Model(&model.User{}).Where(id).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}
