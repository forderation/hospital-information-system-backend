package repository

import (
	"github.com/forderation/hospital-information-system/db"
	"gorm.io/gorm"
)

func GetRegistrantsWithRelation(dbc *gorm.DB, doctorGet *db.Registrant) (doctors []db.DoctorAppointment, err error) {
	err = dbc.Where(&doctorGet).Find(&doctors).Error
	return
}
