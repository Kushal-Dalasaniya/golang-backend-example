package services

import (
	"github.com/Kushal-Dalasaniya/golang-backend/entity"
	"github.com/Kushal-Dalasaniya/golang-backend/repositories"
)

type UserService interface {
	CreateUser(user *entity.User) error
	GetUsers() ([]entity.User, error)
	GetUserByID(id uint) (entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id uint) error
}

type UserServiceImpl struct {
	Repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{Repo: repo}
}

func (s *UserServiceImpl) CreateUser(user *entity.User) error {
	return s.Repo.Create(user)
}

func (s *UserServiceImpl) GetUsers() ([]entity.User, error) {
	return s.Repo.GetAll()
}

func (s *UserServiceImpl) GetUserByID(id uint) (entity.User, error) {
	return s.Repo.GetByID(id)
}

func (s *UserServiceImpl) UpdateUser(user *entity.User) error {
	return s.Repo.Update(user)
}

func (s *UserServiceImpl) DeleteUser(id uint) error {
	return s.Repo.Delete(id)
}
