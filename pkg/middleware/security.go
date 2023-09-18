package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// Authenticate manages the security by validating the token
func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := os.Getenv("TOKEN")
		tokenHeader := ctx.GetHeader("TOKEN")

		if tokenHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"message": "Token not found"})
			return
		}
		if tokenHeader != token {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"message": "Invalid token"})
			return
		}

		ctx.Next()
	}
}
