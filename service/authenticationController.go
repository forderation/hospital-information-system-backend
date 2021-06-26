package service

import (
	"crypto/md5"
	"fmt"
	"github.com/forderation/hospital-information-system/db"
	"github.com/forderation/hospital-information-system/repository"
	"github.com/forderation/hospital-information-system/util"
	"github.com/gin-gonic/gin"
	"io"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func Login(c *gin.Context) {
	var request LoginRequest
	messageOk := "Login successfully"
	err := util.BindAndValidateRequest(&request, c)
	if err != nil {
		InvalidRequestResponse(err.Error(), c)
		return
	}
	h := md5.New()
	_, err = io.WriteString(h, request.Password)
	if err != nil {
		InternalServerErrorResponse(err, c)
		return
	}
	users, err := repository.GetUser(DB, db.User{
		Username: request.Username,
		Password: fmt.Sprintf("%x", h.Sum(nil)),
	})
	if err != nil {
		InternalServerErrorResponse(err, c)
		return
	}
	if len(users) < 1 {
		NotFoundResponse("Failed to login! user not found", c)
		return
	}
	user := users[0]
	token, err := util.CreateToken(user.ID)
	if err != nil {
		InternalServerErrorResponse(err, c)
		return
	}
	StandardResponse(Response{
		Data: map[string]interface{}{
			"Token": token,
		},
		Message: messageOk,
	}, c)
	return
}
