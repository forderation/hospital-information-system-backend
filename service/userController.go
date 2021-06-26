package service

import (
	"github.com/forderation/hospital-information-system/db"
	"github.com/forderation/hospital-information-system/repository"
	"github.com/forderation/hospital-information-system/util"
	"github.com/gin-gonic/gin"
)

type RequestListRegistrant struct {
	UserId uint `json:"user_id"`
}

func GetListRegistrant(c *gin.Context) {
	var request RequestListRegistrant
	messageOk := "Successfully get registrants"
	err := util.BindAndValidateRequest(&request, c)
	if err != nil {
		InvalidRequestResponse(err.Error(), c)
		return
	}
	var registrants []db.Registrant
	if request.UserId != 0 {
		registrants, err = repository.GetRegistrantsWithRelation(DB, &db.Registrant{
			UserID: request.UserId,
		})
	} else {
		registrants, err = repository.GetRegistrantsWithRelation(DB, nil)
	}
	if err != nil {
		InternalServerErrorResponse(err, c)
		return
	}
	StandardResponse(Response{
		Message: messageOk,
		Data:    registrants,
	}, c)
	return
}

type RequestAddRegistrant struct {
	UserId              uint `json:"user_id" validate:"required"`
	DoctorAppointmentId uint `json:"doctor_appointment_id" validate:"required"`
}

func AddRegistrant(c *gin.Context) {
	var request RequestAddRegistrant
	messageOK := "Successfully create registrant"
	err := util.BindAndValidateRequest(&request, c)
	if err != nil {
		InvalidRequestResponse(err.Error(), c)
		return
	}
	users, err := repository.GetUsersById(DB, []uint{request.UserId})
	if err != nil {
		InternalServerErrorResponse(err, c)
		return
	}
	if len(users) < 1 {
		NotFoundResponse("User id not found", c)
		return
	}
	doctors, err := repository.GetDoctorsById(DB, []uint{request.DoctorAppointmentId})
	if err != nil {
		InternalServerErrorResponse(err, c)
		return
	}
	if len(doctors) < 1 {
		NotFoundResponse("Doctor appointment id not found", c)
		return
	}
	registrant, err := repository.CreateRegistrant(DB, request.UserId, request.DoctorAppointmentId)
	if err != nil {
		InternalServerErrorResponse(err, c)
		return
	}
	StandardResponse(Response{
		Message: messageOK,
		Data:    registrant,
	}, c)
	return
}
