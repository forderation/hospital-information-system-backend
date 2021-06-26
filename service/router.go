package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Response struct {
	Code    string
	Message string
	Data    interface{}
}

var DB *gorm.DB

func InitRoute(db *gorm.DB) *gin.Engine {
	DB = db
	router := gin.Default()
	api := router.Group("/api")
	auth := api.Group("/auth")
	auth.POST("/login", Login)
	return router
}
