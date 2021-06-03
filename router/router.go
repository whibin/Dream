package router

import (
	"Dream/controllers"
	"Dream/middlewares"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	// 在使用路由组之前用中间件
	Router.Use(middlewares.Cors())

	dream := Router.Group("/dream")
	{

		dream.GET("/my/:uid", controllers.SelectOwnDream)
		dream.GET("/count/type/:uid/", controllers.CountByDreamType)
		dream.GET("/count/time/:uid", controllers.CountByTime)
		dream.DELETE("/del/:uid/:id", controllers.Delete)
		dream.PUT("/update", controllers.Update)
		dream.GET("/all/:uid", controllers.CountDreamsByUser)
		dream.GET("/match/:uid/:id", controllers.DreamMatch)

		dream.GET("/openid", controllers.GetOpenId)
		dream.POST("/user/register", controllers.Register)
		dream.PUT("/user/update", controllers.UpdateUser)
		dream.GET("/user", controllers.GetUser)
		dream.GET("/user/likes/:uid", controllers.GetReceivedLikes)

		dream.GET("/explain", controllers.ExplainDream)
		dream.POST("/save", controllers.Save)
		dream.POST("/file/draw", controllers.DrawUpload)
		dream.POST("/file/sound", controllers.SoundUpload)

		dream.GET("/like/check/:userId/:dreamId", controllers.HasLike)
		dream.GET("/like/:userId/:dreamId", controllers.Like)
		dream.GET("/unlike/:userId/:dreamId", controllers.Unlike)
		dream.GET("/like/amount/:dreamId", controllers.GetLikeAmount)
		dream.GET("/like/dream", controllers.GetDreamByLike)
		dream.GET("/bytime", controllers.GetDreamByTime)
		dream.GET("/bytype/:type", controllers.GetDreamByType)
		dream.POST("/com/add", controllers.AddComment)
		dream.DELETE("/com/del/:id", controllers.DeleteComment)
		dream.GET("/com/:id", controllers.GetCommentsByDream)

	}

}
