package service

import "C"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func StandardResponse(response Response, c *gin.Context) {
	response.Code = "00"
	c.JSON(http.StatusOK, response)
}

func InternalServerErrorResponse(err error, c *gin.Context) {
	resp := Response{
		Code:    "INTERNAL-SERVER-REQUEST",
		Message: "Oops there is something wrong in our server, please try again",
		Data:    err.Error(),
	}
	c.JSON(http.StatusOK, resp)
}

func UnauthorizedResponse(err error, c *gin.Context) {
	resp := Response{
		Code:    "REQUEST-UNAUTHORIZED",
		Message: "Request is unauthorized please login first",
		Data:    err.Error(),
	}
	c.JSON(http.StatusOK, resp)
}

func NotAllowedResponse(notAllowedMessage string, c *gin.Context) {
	resp := Response{
		Code:    "REQUEST-NOT-ALLOWED",
		Message: notAllowedMessage,
		Data:    nil,
	}
	c.JSON(http.StatusOK, resp)
}

func InvalidRequestResponse(invalidMessage string, c *gin.Context) {
	resp := Response{
		Code:    "INVALID-REQUEST",
		Message: fmt.Sprintf("Request is invalid: %s", invalidMessage),
		Data:    nil,
	}
	c.JSON(http.StatusOK, resp)
}

func NotFoundResponse(notfoundMessage string, c *gin.Context) {
	resp := Response{
		Code:    "NOT-FOUND-REQUEST",
		Message: notfoundMessage,
		Data:    nil,
	}
	c.JSON(http.StatusOK, resp)
}
