package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	basePath = "http://localhost:8080"
)

// Logger manages the security by validating the token
func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL
		localTime := time.Now()
		method := ctx.Request.Method

		ctx.Next()

		var size int
		if ctx.Writer != nil {
			size = ctx.Writer.Size()
		}

		fmt.Printf("TIME: %v\n"+
			"PATH: %s%s\n"+
			"METHOD: %s\n"+
			"SIZE: %d",
			localTime, basePath, path, method, size)
	}
}
