package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/iamsushank/jwt-go/utils/token"
	"net/http"
)

func JwtAuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
