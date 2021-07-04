package db

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	Age         uint           `json:"age"`
	Email       string         `gorm:"unique" json:"email"`
	Username    string         `gorm:"unique" json:"username"`
	Password    string         `json:"password"`
	IsAdmin     bool           `json:"is_admin"`
	Registrants []Registrant   `json:"registrants"`
}

type Registrant struct {
	ID                  uint              `gorm:"primarykey" json:"id"`
	CreatedAt           time.Time         `json:"created_at"`
	UpdatedAt           time.Time         `json:"updated_at"`
	DeletedAt           gorm.DeletedAt    `gorm:"index" json:"deleted_at"`
	UserID              uint              `json:"user_id"`
	DoctorAppointmentID uint              `json:"doctor_appointment_id"`
	IsCanceled          bool              `json:"is_canceled"`
	User                User              `gorm:"foreignKey:UserID" json:"user"`
	DoctorAppointment   DoctorAppointment `gorm:"foreignKey:DoctorAppointmentID" json:"doctor_appointment"`
}

type DoctorAppointment struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	DoctorName    string         `json:"doctor_name"`
	Description   string         `json:"description"`
	MaxRegistrant uint           `json:"max_registrant"`
	Registrants   []Registrant   `json:"registrants"`
}
