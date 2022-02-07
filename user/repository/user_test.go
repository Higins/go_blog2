package UserRepository

import (
	"testing"

	"github.com/Higins/go_blog2/domain"
)

func TestLogin(t *testing.T) {
	loginUser := domain.User{
		Username: "admin",
		Password: "admin",
	}
	testLogin := new(domain.MockUserRepository)
	testLogin.On("Login", loginUser).Return(domain.User{}, nil)
}
func TestRegistration(t *testing.T) {
	regUser := domain.User{
		Username: "admin",
		Password: "admin",
	}
	testRegistration := new(domain.MockUserRepository)
	testRegistration.On("Registration", regUser).Return(domain.User{}, nil)
}
