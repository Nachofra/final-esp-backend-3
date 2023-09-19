package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// response is a struc for responses
type response struct {
	Data interface{} `json:"data"`
}

// errorResponse is a struc for errors
type errorResponse struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Response creates the base response that the Success and Error functions will then use
func Response(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

// Success creates the successful response with its status and object data in an interface
func Success(c *gin.Context, status int, data interface{}) {
	Response(c, status, response{Data: data})
}

// Error creates a new error with the given status code and the message
// formatted according to args and format.
func Error(c *gin.Context, status int, format string, args ...interface{}) {
	err := errorResponse{
		Code:    strings.ReplaceAll(strings.ToLower(http.StatusText(status)), " ", "_"),
		Message: fmt.Sprintf(format, args...),
		Status:  status,
	}

	Response(c, status, err)
}
