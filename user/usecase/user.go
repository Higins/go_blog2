package UserUsecase

import (
	"log"
	"time"

	"github.com/Higins/go_blog2/domain"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type userUsecase struct {
	userRepository domain.userRepository
}
func NewUserUsecase(user domain.UserRepository) domain.UserRepository {
	return &userRepository{
		userRepository: user,
	}
}

func (u *userUsecase) AuthMiddleware(user domain.User) (user domain.User,error) {

	 jwt.New(&jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
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
			user,err = u.Login(user)
			if err != nil {
				return user
			}

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

}
