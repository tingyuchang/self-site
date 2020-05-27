package apiv1

import (
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context){
	c.JSON(200, gin.H{
		"msg": "OK",
	})
}