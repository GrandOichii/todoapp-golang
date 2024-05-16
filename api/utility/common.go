package utility

import (
	"errors"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func MapSlice[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func Extract(key string, c *gin.Context) (string, error) {
	claims := jwt.ExtractClaims(c)
	result := claims[key]
	if result == nil {
		return "", errors.New("key " + key + " doesn't exist")
	}

	return result.(string), nil
}
