package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitRoute(db *gorm.DB) *gin.Engine {
	DB = db
	router := gin.Default()
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", Login)
			auth.POST("/signup", Signup)
		}
		res := api.Group("/res")
		res.Use(IsAuthentication())
		{
			doctor := res.Group("/doctor")
			{
				doctor.GET("/", GetDoctorAppointments)
				doctor.POST("/", AddDoctorAppointment)
			}
			registrant := res.Group("/registrant")
			{
				registrant.GET("/", GetListRegistrant)
				registrant.POST("/", AddRegistrant)
			}
		}
	}
	return router
}
