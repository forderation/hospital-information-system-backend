package service

import (
	"github.com/forderation/hospital-information-system/util"
	"github.com/gin-gonic/gin"
)

func IsAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := util.ExtractPayload(c)
		if err != nil {
			UnauthorizedResponse(err, c)
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}