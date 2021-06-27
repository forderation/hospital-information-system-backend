package repository

import (
	"crypto/md5"
	"fmt"
	"github.com/forderation/hospital-information-system/db"
	"gorm.io/gorm"
	"io"
)

func CreateUser(dbc *gorm.DB, user *db.User) (err error) {
	h := md5.New()
	_, err = io.WriteString(h, user.Password)
	if err != nil {
		return
	}
	user.Password = fmt.Sprintf("%x", h.Sum(nil))
	err = dbc.Create(user).Error
	if err != nil {
		return
	}
	return
}

func CheckUserExisting(dbc *gorm.DB, username string, email string) (isExist bool, err error) {
	var users []db.User
	err = dbc.Where("username = ? or email = ?", username, email).Find(&users).Error
	if err != nil {
		return false, err
	}
	if len(users) > 0 {
		return true, nil
	}
	return false, nil
}

func GetUser(dbc *gorm.DB, filter db.User) (users []db.User, err error) {
	err = dbc.Where(&filter).Find(&users).Error
	return
}

func GetUsersById(dbc *gorm.DB, ids []uint) (users []db.User, err error) {
	err = dbc.Preload("Registrants.DoctorAppointment").Preload("Registrants.User").Find(&users, ids).Error
	return
}

func DeleteUser(dbc *gorm.DB, filter *db.User) (err error) {
	err = dbc.Delete(filter).Error
	if err != nil {
		return
	}
	return
}
