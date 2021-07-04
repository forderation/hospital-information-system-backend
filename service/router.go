package service

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func InitRoute(db *gorm.DB) *gin.Engine {
	DB = db
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
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
				doctor.GET("/:id", GetDetailDoctorAppointments)
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
