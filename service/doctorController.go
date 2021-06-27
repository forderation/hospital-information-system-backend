package service

import (
	"github.com/forderation/hospital-information-system/db"
	"github.com/forderation/hospital-information-system/repository"
	"github.com/forderation/hospital-information-system/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AddDoctorRequest struct {
	DoctorName    string `json:"doctor_name" validate:"required"`
	Description   string `json:"description"`
	MaxRegistrant uint   `json:"max_registrant" validate:"required,gt=0"`
}

func AddDoctorAppointment(c *gin.Context) {
	var request AddDoctorRequest
	messageOk := "Add doctor appointment are successfully"
	err := util.BindAndValidateRequest(&request, c)
	if err != nil {
		InvalidRequestResponse(err.Error(), c)
		return
	}
	doctor := db.DoctorAppointment{
		DoctorName:    request.DoctorName,
		Description:   request.Description,
		MaxRegistrant: request.MaxRegistrant,
	}
	err = repository.CreateDoctorAppointment(DB, &doctor)
	if err != nil {
		InternalServerErrorResponse(err, c)
		return
	}
	StandardResponse(Response{
		Message: messageOk,
		Data:    doctor,
	}, c)
	return
}

func GetDoctorAppointments(c *gin.Context) {
	messageOk := "Get doctors successfully"
	doctors, err := repository.GetDoctor(DB, &db.DoctorAppointment{})
	if len(doctors) < 1 {
		NotFoundResponse("Doctor appointment is empty", c)
		return
	}
	if err != nil {
		InternalServerErrorResponse(err, c)
		return
	}
	StandardResponse(Response{
		Message: messageOk,
		Data:    doctors,
	}, c)
	return
}

type EditDoctorRequest struct {
	DoctorName    string `json:"doctor_name" validate:"required"`
	Description   string `json:"description"`
	MaxRegistrant uint   `json:"max_registrant" validate:"required,gt=0"`
}

func EditDoctorAppointment(c *gin.Context) {
	messageOk := "Update doctor successfully"
	var request EditDoctorRequest
	err := util.BindAndValidateRequest(&request, c)
	if err != nil {
		InvalidRequestResponse(err.Error(), c)
		return
	}
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	doctors, err := repository.GetDoctorsById(DB, []uint{uint(id)})
	if len(doctors) < 1 {
		NotFoundResponse("Doctor appointment is not found", c)
		return
	}
	if err != nil {
		InternalServerErrorResponse(err, c)
		return
	}
	doctor := doctors[0]
	doctor.DoctorName = request.DoctorName
	doctor.MaxRegistrant = request.MaxRegistrant
	doctor.Description = request.Description
	err = repository.UpdateDoctor(DB, doctor)
	if err != nil {
		InternalServerErrorResponse(err, c)
		return
	}
	StandardResponse(Response{
		Message: messageOk,
		Data:    doctor,
	}, c)
	return
}

func DeleteDoctorAppointment(c *gin.Context) {
	messageOk := "Delete doctor successfully"
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	doctors, err := repository.GetDoctorsById(DB, []uint{uint(id)})
	if len(doctors) < 1 {
		NotFoundResponse("Doctor appointment is not found", c)
		return
	}
	if err != nil {
		InternalServerErrorResponse(err, c)
		return
	}
	doctor := doctors[0]
	err = repository.DeleteDoctor(DB, doctor)
	if err != nil {
		InternalServerErrorResponse(err, c)
		return
	}
	StandardResponse(Response{
		Message: messageOk,
		Data:    doctor,
	}, c)
	return
}
