package router

import (
	"github.com/gin-gonic/gin"

	"Dream/controllers"
	"Dream/middlewares"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	// 在使用路由组之前用中间件
	Router.Use(middlewares.Cors())

	dream := Router.Group("/dream")
	{
		dream.GET("/my/:uid", controllers.SelectOwnDream)
		dream.POST("/save", controllers.Save)
		dream.POST("/file/draw", controllers.DrawUpload)
		dream.POST("/file/sound", controllers.SoundUpload)
	}

}
