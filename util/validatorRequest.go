package util

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strings"
)

func BindAndValidateRequest(request interface{}, c *gin.Context) (err error) {
	err = c.BindJSON(&request)
	if err != nil {
		return errors.New("failed to binding request")
	}
	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		var errorMessage []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage = append(errorMessage, fmt.Sprintf("%s", err.Field()))
		}
		return errors.New(strings.Join(errorMessage, ","))
	}
	return nil
}
