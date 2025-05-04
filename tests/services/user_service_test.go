package services_test

import (
	"errors"
	"testing"

	"github.com/Kushal-Dalasaniya/golang-backend/entity"
	"github.com/Kushal-Dalasaniya/golang-backend/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository mocks the UserRepository interface
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetAll() ([]entity.User, error) {
	args := m.Called()
	return args.Get(0).([]entity.User), args.Error(1)
}

func (m *MockUserRepository) GetByID(id uint) (entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserRepository) Update(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestUserService_CreateUser(t *testing.T) {
	repo := new(MockUserRepository)
	service := services.NewUserService(repo)

	user := &entity.User{Name: "John"}
	repo.On("Create", user).Return(nil)

	err := service.CreateUser(user)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestUserService_GetUsers(t *testing.T) {
	repo := new(MockUserRepository)
	service := services.NewUserService(repo)

	expected := []entity.User{{Name: "Alice"}, {Name: "Bob"}}
	repo.On("GetAll").Return(expected, nil)

	users, err := service.GetUsers()
	assert.NoError(t, err)
	assert.Equal(t, expected, users)
	repo.AssertExpectations(t)
}

func TestUserService_GetUserByID(t *testing.T) {
	repo := new(MockUserRepository)
	service := services.NewUserService(repo)

	expected := entity.User{Name: "John",Email: "john@me.com"}
	repo.On("GetByID", uint(1)).Return(expected, nil)

	user, err := service.GetUserByID(1)
	assert.NoError(t, err)
	assert.Equal(t, expected, user)
	repo.AssertExpectations(t)
}

func TestUserService_UpdateUser(t *testing.T) {
	repo := new(MockUserRepository)
	service := services.NewUserService(repo)

	user := &entity.User{Name: "Updated",Email: "updated@me.com"}
	repo.On("Update", user).Return(nil)

	err := service.UpdateUser(user)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestUserService_DeleteUser(t *testing.T) {
	repo := new(MockUserRepository)
	service := services.NewUserService(repo)

	repo.On("Delete", uint(1)).Return(nil)

	err := service.DeleteUser(1)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestUserService_CreateUser_Error(t *testing.T) {
	repo := new(MockUserRepository)
	service := services.NewUserService(repo)

	errExpected := errors.New("create error")
	user := &entity.User{Name: "John"}
	repo.On("Create", user).Return(errExpected)

	err := service.CreateUser(user)
	assert.EqualError(t, err, "create error")
	repo.AssertExpectations(t)
}
