package jwt

import (
	"fmt"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"	
)

func JWT() gin.HandlerFunc {
	return func (c *gin.Context) {
		// check token is valid or not
		fmt.Println(c.FullPath())
		c.Next()
	}
}