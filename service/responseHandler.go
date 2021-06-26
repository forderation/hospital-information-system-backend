package service

import "C"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Message string
	Data    interface{}
}

func StandardResponse(response Response, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Code":    "00",
		"Message": response.Message,
		"Data":    response.Data,
	})
}

func InternalServerErrorResponse(err error, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"Code":    "INTERNAL-SERVER-REQUEST",
		"Message": "Oops there is something wrong in our server, please try again",
		"Data":    err.Error(),
	})
}

func UnauthorizedResponse(err error, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"Code":    "REQUEST-UNAUTHORIZED",
		"Message": "Request is unauthorized please login first",
		"Data":    err.Error(),
	})
}

func InvalidRequestResponse(invalidMessage string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"Code":    "INVALID-REQUEST",
		"Message": fmt.Sprintf("Request is invalid: %s", invalidMessage),
		"Data":    nil,
	})
}

func NotFoundResponse(notfoundMessage string, c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"Code":    "NOT-FOUND-REQUEST",
		"Message": notfoundMessage,
		"Data":    nil,
	})
}
