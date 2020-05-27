package routers 

import (
	"github.com/gin-gonic/gin"
	"self-site/setting"
	"self-site/routers/api/v1"
	"self-site/middleware/jwt"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(setting.Config.Server.RunMode)
	r := gin.New()
	
	// setup moddleware 
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// regist router
	// r type of  gin.Engine, and gin.Engine also has RouterGroup
	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())

	setupUser(apiV1)

	return r 
}

func setupUser(r *gin.RouterGroup) {
	r.GET("/users", apiv1.GetUserInfo)
}