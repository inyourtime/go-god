package repository

import (
	"gopher/src/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetById(uint) (*model.User, error)
	Create(model.User) (*model.User, error)
	GetByEmail(string) (*model.User, error)
}

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) UserRepository {
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) GetByEmail(email string) (*model.User, error) {
	user := model.User{}
	err := r.db.Where(model.User{Email: email}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepositoryDB) Create(user model.User) (*model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepositoryDB) GetAll() ([]model.User, error) {
	users := []model.User{}
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r userRepositoryDB) GetById(id uint) (*model.User, error) {
	user := model.User{}
	err := r.db.Where(id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
