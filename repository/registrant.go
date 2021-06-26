package repository

import (
	"github.com/forderation/hospital-information-system/db"
	"gorm.io/gorm"
)

func GetRegistrantsWithRelation(dbc *gorm.DB, filter *db.Registrant) (registrants []db.Registrant, err error) {
	err = dbc.Preload("User").Preload("DoctorAppointment").Where(&filter).Find(&registrants).Error
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
