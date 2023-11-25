package controller

import (
	"TikTokLite/log"
	"TikTokLite/response"
	"TikTokLite/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"strconv"
)

//视频发布

func PublishAction(ctx *gin.Context) {
	// publishResponse := &message.DouyinPublishActionResponse{}
	//userId, _ := ctx.Get("UserId")
	//token := ctx.PostForm("token")
	//userId, err := common.VerifyToken(token)
	//title := ctx.PostForm("title")

	data, err := ctx.FormFile("data")
	if err != nil {
		logrus.Info(err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	filename := filepath.Base(data.Filename) //返回data的路径的最后一个参数，即为文件名
	//finalName := fmt.Sprintf("%s_%s", util.RandomString(), filename) //根据文件名生成一个新的文件名
	saveFile := "C:\\Users\\0\\Downloads\\" + filename //获取到视频路径
	//saveFile := filepath.Join(videoPath, filename) //将其进行一个合并成为文件的最终路径

	log.Info("saveFile:", saveFile)

	if err := ctx.SaveUploadedFile(data, saveFile); err != nil { //这段代码是在用来将将要上传的文件data保存在指定路径saveFile上.
		logrus.Info(err)
	}
	logrus.Infof(filename)
	logrus.Info(saveFile)
	publish, err := service.PublishVideo(saveFile, filename)
	/*publish, err := service.PublishVideo(userId, saveFile)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}*/
	logrus.Infof("publish:%v", publish)
	response.Success(ctx, "success", publish)

}

// 获取视频列表

func GetPublishList(ctx *gin.Context) {
	tokenUserId, _ := ctx.Get("UserId")
	id := ctx.Query("user_id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
	}
	list, err := service.PublishList(tokenUserId.(int64), userId)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", list)
}
