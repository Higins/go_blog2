package UserUsecase

import (
	"log"
	"time"

	"github.com/Higins/go_blog2/domain"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(user domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepository: user,
	}
}

func (uc *userUsecase) authMiddleware(user domain.UserApi) (userDomain domain.User, err error) {
	var identityKey = "id"
	jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*domain.User); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {

			if err := c.ShouldBind(&user); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			uc.Login(user)

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*domain.User); ok && v.Username == "admin" {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return userDomain, nil
}
func (uc *userUsecase) Login(userApi domain.UserApi) (userDomain domain.User) {
	userDomain, err := uc.authMiddleware(userApi)
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return userDomain
}
func (uc *userUsecase) Registration(user domain.User) error {
	var userDomain domain.User
	hashedPass, err := encryptPassword(user.Password)
	if err != nil {
		log.Printf("reg error: %s", err.Error())
	}
	userDomain.Password = hashedPass
	userDomain.Username = user.Username
	_, err = uc.userRepository.Login(userDomain)
	if err != nil {
		return err
	}
	return nil
}
func encryptPassword(password string) (string, error) {
	bytePass := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	if err != nil {
		log.Printf("ERROR:EncryptPassword: %s", err.Error())
	}
	return string(hashedPassword), err
}
