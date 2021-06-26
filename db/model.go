package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName   string
	LastName    string
	Age         uint
	Email       string `gorm:"unique"`
	Username    string `gorm:"unique"`
	Password    string
	IsAdmin     bool
	Registrants []Registrant
}

type Registrant struct {
	gorm.Model
	UserID              uint
	DoctorAppointmentID uint
	IsCanceled          bool
	User User `gorm:"foreignKey:UserID"`
	DoctorAppointment DoctorAppointment `gorm:"foreignKey:DoctorAppointmentID"`
}

type DoctorAppointment struct {
	gorm.Model
	DoctorName    string
	Description   string
	MaxRegistrant uint
	Registrants   []Registrant
}