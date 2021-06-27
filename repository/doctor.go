package repository

import (
	"github.com/forderation/hospital-information-system/db"
	"gorm.io/gorm"
)

func CreateDoctorAppointment(dbc *gorm.DB, da *db.DoctorAppointment) (err error) {
	err = dbc.Create(da).Error
	if err != nil {
		return
	}
	return
}

func GetDoctor(dbc *gorm.DB, doctorGet *db.DoctorAppointment) (doctors []db.DoctorAppointment, err error) {
	err = dbc.Where(&doctorGet).Find(&doctors).Error
	return
}

func GetDoctorsById(dbc *gorm.DB, ids []uint) (doctors []db.DoctorAppointment, err error) {
	err = dbc.Find(&doctors, ids).Error
	return
}

func UpdateDoctor(dbc *gorm.DB, da db.DoctorAppointment) (err error) {
	err = dbc.Model(&da).Updates(da).Error
	return
}

func DeleteDoctor(dbc *gorm.DB, da db.DoctorAppointment) (err error) {
	err = dbc.Delete(&da).Error
	return
}
