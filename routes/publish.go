package routes

import (
	"TikTokLite/common"
	"TikTokLite/controller"

	"github.com/gin-gonic/gin"
)

func PublishRoutes(r *gin.RouterGroup) {
	publish := r.Group("publish")
	{
		publish.POST("/action/", common.AuthWithOutMiddleware(), controller.PublishAction)
		publish.GET("/list/", common.AuthWithOutMiddleware(), controller.GetPublishList)
	}
}
