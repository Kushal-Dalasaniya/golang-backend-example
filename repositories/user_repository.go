package repositories

import (
	"github.com/Kushal-Dalasaniya/golang-backend/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
	GetAll() ([]entity.User, error)
	GetByID(id uint) (entity.User, error)
	Update(user *entity.User) error
	Delete(id uint) error
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) Create(user *entity.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepositoryImpl) GetAll() ([]entity.User, error) {
	var users []entity.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepositoryImpl) GetByID(id uint) (entity.User, error) {
	var user entity.User
	err := r.DB.First(&user, id).Error
	return user, err
}

func (r *UserRepositoryImpl) Update(user *entity.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepositoryImpl) Delete(id uint) error {
	return r.DB.Delete(&entity.User{}, id).Error
}
