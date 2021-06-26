package service

import (
	"errors"
	"fmt"
	"github.com/forderation/hospital-information-system/db"
	"github.com/forderation/hospital-information-system/repository"
	"github.com/forderation/hospital-information-system/util"
	"github.com/gin-gonic/gin"
)

type SignupRequest struct {
	Firstname string `json:"firstname" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email" validate:"email,required"`
	Password  string `json:"password" validate:"required,min=6"`
	Age       uint   `json:"age" validate:"required,gt=0"`
}

func Signup(c *gin.Context) {
	var request SignupRequest
	messageOk := "Register user successfully"
	err := util.BindAndValidateRequest(&request, c)
	if err != nil {
		InvalidRequestResponse(err.Error(), c)
		return
	}
	userExist, err := repository.CheckUserExisting(DB, request.Username, request.Email)
	if err != nil {
		InternalServerErrorResponse(err, c)
		return
	}
	if userExist {
		InvalidRequestResponse(errors.New(fmt.Sprintf(
			"user with email %s or username %s already existing", request.Email, request.Username),
		).Error(), c)
		return
	}
	user := db.User{
		FirstName: request.Firstname,
		LastName:  request.Lastname,
		Password:  request.Password,
		Username:  request.Username,
		Email:     request.Email,
		Age:       request.Age,
	}
	err = repository.CreateUser(DB, &user)
	if err != nil {
		InternalServerErrorResponse(err, c)
		return
	}
	StandardResponse(Response{
		Message: messageOk,
		Data:    request,
	}, c)
}
