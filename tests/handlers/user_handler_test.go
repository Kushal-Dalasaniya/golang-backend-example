package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Kushal-Dalasaniya/golang-backend/entity"
	"github.com/Kushal-Dalasaniya/golang-backend/handlers"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) GetUsers() ([]entity.User, error) {
	args := m.Called()
	return args.Get(0).([]entity.User), args.Error(1)
}

func (m *MockUserService) GetUserByID(id uint) (entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserService) UpdateUser(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) DeleteUser(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateUserHandler(t *testing.T) {
	mockService := new(MockUserService)
	h := handlers.NewUserHandler(mockService)

	user := &entity.User{Name: "John"}
	mockService.On("CreateUser", user).Return(nil)

	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	h.CreateUser(res, req)
	assert.Equal(t, http.StatusCreated, res.Code)
	mockService.AssertExpectations(t)
}

func TestGetUsersHandler(t *testing.T) {
	mockService := new(MockUserService)
	h := handlers.NewUserHandler(mockService)

	expected := []entity.User{{Name: "Alice"}}
	mockService.On("GetUsers").Return(expected, nil)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	res := httptest.NewRecorder()

	h.GetUsers(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	mockService.AssertExpectations(t)
}

func TestGetUserByIDHandler(t *testing.T) {
	mockService := new(MockUserService)
	h := handlers.NewUserHandler(mockService)

	expected := entity.User{Name: "John", Email: "john@me.com"}
	mockService.On("GetUserByID", uint(1)).Return(expected, nil)

	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	res := httptest.NewRecorder()

	h.GetUserByID(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateUserHandler(t *testing.T) {
	mockService := new(MockUserService)
	h := handlers.NewUserHandler(mockService)

	user := &entity.User{Name: "Updated", Email: "updated@me.com"}
	// mockService.On("UpdateUser", user).Return(nil)
	mockService.On("UpdateUser", mock.MatchedBy(func(u *entity.User) bool {
		return u.ID == 1 && u.Name == "Updated" && u.Email == "updated@me.com"
	})).Return(nil)

	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	res := httptest.NewRecorder()

	h.UpdateUser(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteUserHandler(t *testing.T) {
	mockService := new(MockUserService)
	h := handlers.NewUserHandler(mockService)

	mockService.On("DeleteUser", uint(1)).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	res := httptest.NewRecorder()

	h.DeleteUser(res, req)
	assert.Equal(t, http.StatusNoContent, res.Code)
	mockService.AssertExpectations(t)
}

func TestCreateUserHandler_Error(t *testing.T) {
	mockService := new(MockUserService)
	h := handlers.NewUserHandler(mockService)

	user := &entity.User{Name: "Invalid"}
	err := errors.New("DB failure")
	mockService.On("CreateUser", user).Return(err)

	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	h.CreateUser(res, req)
	assert.Equal(t, http.StatusInternalServerError, res.Code)
	mockService.AssertExpectations(t)
}
