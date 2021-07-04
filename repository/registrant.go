package repository

import (
	"github.com/forderation/hospital-information-system/db"
	"gorm.io/gorm"
)

func GetRegistrantsWithRelation(dbc *gorm.DB, filter *db.Registrant) (registrants []db.Registrant, err error) {
	err = dbc.Preload("User").Preload("DoctorAppointment").Where(&filter).Find(&registrants).Error
	return
}

func GetRegistrantsById(dbc *gorm.DB, id []uint) (registrants []db.Registrant, err error) {
	err = dbc.Preload("User").Preload("DoctorAppointment").Find(&registrants, id).Error
	return
}

func CountRegistrantsByDoctorId(dbc *gorm.DB, doctorIds []uint) (registrants []db.Registrant, count int64,  err error) {
	err = dbc.Where("doctor_appointment_id IN ?", doctorIds).Find(&registrants).Count(&count).Error
	return
}

func CreateRegistrant(dbc *gorm.DB, userId uint, doctorAppointmentId uint) (registrant db.Registrant, err error) {
	registrant.DoctorAppointmentID = doctorAppointmentId
	registrant.UserID = userId
	err = dbc.Create(&registrant).Error
	if err != nil {
		return
	}
	err = dbc.Preload("User").Preload("DoctorAppointment").Find(&registrant).Error
	return
}

func UpdateRegistrant(dbc *gorm.DB, registrant db.Registrant) (err error) {
	err = dbc.Model(&registrant).Updates(&registrant).Error
	return
}
