package middleware

import (
	"log"
	"time"

	dto "github.com/GrandOichii/todoapp-golang/api/dto/user"
	"github.com/GrandOichii/todoapp-golang/api/models"
	services "github.com/GrandOichii/todoapp-golang/api/services/user"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

const (
	IDKey string = "id"
)

type JwtMiddleware struct {
	// Middleware

	// AuthMiddleware
	Middle *jwt.GinJWTMiddleware
}

func GetSecretKey() []byte {
	// TODO add actual secret key
	return []byte("secret key")
}

func CreateJwtMiddleware(userService services.UserService) *JwtMiddleware {

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        GetSecretKey(),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,

		IdentityKey: IDKey,

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*dto.GetUser); ok {
				return jwt.MapClaims{
					IDKey: v.Id,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{
				Id: claims[IDKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals dto.PostUser
			if err := c.BindJSON(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			result, err := userService.Login(&loginVals)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return result, nil

		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// if v, ok := data.(*models.User); ok && v.UserName == "admin" {
			// 	return true
			// }

			// TODO figure out what this is for

			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	err = authMiddleware.MiddlewareInit()

	if err != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + err.Error())
	}

	result := &JwtMiddleware{
		Middle: authMiddleware,
	}
	return result
}

func (jm *JwtMiddleware) GetMiddlewareFunc() gin.HandlerFunc {
	return jm.Middle.MiddlewareFunc()
}
