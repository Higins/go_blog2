package UserRepository

import (
	"fmt"
	"log"

	"github.com/Higins/go_blog2/domain"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Login(user domain.User) (domain.User, error) {
	var domainUser domain.User
	u.db.Where("username = ?", user.Username).First(&domainUser)

	err := bcrypt.CompareHashAndPassword([]byte(domainUser.Password), []byte(user.Password))
	if err != nil {
		log.Printf("ERROR:Login: %s", err.Error())
		return domain.User{}, err
	}

	return domainUser, nil
}

func (u *UserRepository) Registration(user domain.User) (domain.User, error) {
	var err error

	if user.ID > 0 {
		err = u.db.Create(&user).Error
	} else {
		err = u.db.Save(&user).Error
	}
	if err != nil {
		fmt.Println(err)
		return domain.User{}, err
	}
	return user, nil
}
