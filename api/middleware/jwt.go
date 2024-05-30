package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/GrandOichii/todoapp-golang/api/config"
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

func CreateJwtMiddleware(config *config.Configuration, userService services.UserService) *JwtMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte(config.AuthSecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,

		SendCookie:     true,
		SecureCookie:   false, // ! non HTTPS dev environments
		CookieHTTPOnly: true,  // JS can't modify
		CookieDomain:   "localhost:" + config.Port,
		CookieName:     "token",                  // default jwt
		CookieSameSite: http.SameSiteDefaultMode, //SameSiteDefaultMode, SameSiteLaxMode, SameSiteStrictMode, SameSiteNoneMode

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

			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			if strings.HasPrefix(c.FullPath(), "/api") {
				c.JSON(code, gin.H{
					"code":    code,
					"message": message,
				})
				return
			}
			c.HTML(200, "login", nil)
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			if strings.HasPrefix(c.FullPath(), "/api") {
				c.JSON(http.StatusOK, gin.H{
					"code":   code,
					"token":  token,
					"expire": expire.Format(time.RFC3339),
				})
				return
			}
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: token",
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
