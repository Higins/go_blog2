package UserRepository

import (
	"github.com/Higins/go_blog2/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Login(user domain.User) (domain.User, error) {
	if user.Username == "admin" && user.Password == "admin" {
		return domain.User{
			Username: user.Username,
		}, nil
	}
	return user, nil
}
