package domain

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
}
type UserApi struct {
	ID       int
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRepository interface {
	Login(user User) (User, error)
	Registration(user User) (User, error)
}
type UserUsecase interface {
	Login(user UserApi) User
	Registration(user User) error
}

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Login(user User) (User, error) {
	args := m.Called(user)
	return args.Get(0).(User), args.Error(1)
}

func (m *MockUserRepository) Registration(user User) (User, error) {
	args := m.Called(user)
	return args.Get(0).(User), args.Error(1)
}
