package domain

import (
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
}
type UserUsecase interface {
	Login(user UserApi) error
}
