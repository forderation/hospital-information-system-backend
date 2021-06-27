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
				doctor.PATCH("/:id", EditDoctorAppointment)
				doctor.DELETE("/:id", DeleteDoctorAppointment)
			}
			registrant := res.Group("/registrant")
			{
				registrant.GET("/", GetListRegistrant)
				registrant.POST("/apply", AddRegistrant)
				registrant.POST("/cancel", CancelRegistrant)
			}
			user := res.Group("/user")
			{
				user.GET("/:id", GetUser)
			}
		}
	}
	return router
}
