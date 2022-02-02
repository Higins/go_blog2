package UserUsecase

import (
	"log"
	"time"

	"github.com/Higins/go_blog2/domain"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(user domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepository: user,
	}
}

func (u *userUsecase) authMiddleware(user domain.UserApi) (userDomain domain.User, err error) {
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
			err = u.Login(user)

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
func (u *userUsecase) Login(userApi domain.UserApi) error {
	_, err := u.authMiddleware(userApi)
	if err != nil {
		log.Fatal("Login error")
	}
	return nil
}
