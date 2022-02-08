package UserUsecase

import (
	"testing"

	"github.com/Higins/go_blog2/domain"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	loginUser := domain.User{
		Username: "admin",
		Password: "admin",
	}
	testLogin := new(domain.MockUserRepository)
	testLogin.On("Login", loginUser).Return(domain.User{}, nil)
	user := NewUserUsecase(&domain.MockUserRepository{})
	user.Login(domain.UserApi{Username: "admin", Password: "admin"})

	assert := assert.New(t)
	assert.Equal(domain.UserApi{Username: "admin", Password: "admin"}, user, "ok")

}
func TestRegistration(t *testing.T) {
	regUser := domain.User{
		Username: "admin",
		Password: "admin",
	}
	testRegistration := new(domain.MockUserRepository)
	testRegistration.On("Registration", regUser).Return(domain.User{}, nil)

	user := NewUserUsecase(&domain.MockUserRepository{})
	user.Registration(domain.User{Username: "admin", Password: "admin"})

	assert := assert.New(t)
	assert.Equal(domain.User{Username: "admin", Password: "admin"}, user, "ok")
}
