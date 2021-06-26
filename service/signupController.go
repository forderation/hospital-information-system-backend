package service

import (
	"github.com/forderation/hospital-information-system/db"
	"github.com/forderation/hospital-information-system/repository"
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	
}

func Signup(c *gin.Context) {
	var request RegisterRequest
	messageOk := "Register user successfully"
	err := c.BindJSON(&request)
	if err != nil {
		InvalidRequestResponse("", c)
	}
	user := db.User{}
	err = repository.CreateUser(DB, &user)
	if err != nil {
		InternalServerErrorResponse(c)
	}
	StandardResponse(Response{
		Message: messageOk,
		Data:    request,
	}, c)
}
