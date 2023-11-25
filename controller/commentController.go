package controller

import (
	"TikTokLite/log"
	"TikTokLite/response"
	"TikTokLite/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 发布评论

func CommentAction(ctx *gin.Context) {
	var err error

	tokenUids, _ := ctx.Get("UserId") //用于获取当前登录的用户id
	tokenUid := tokenUids.(int64)     //并且将其转化为int64类型

	video_id := ctx.Query("video_id")
	comment_text := ctx.Query("comment_text")
	actionType := ctx.Query("action_type")
	comment_id := ctx.Query("comment_id")
	//ToDo：如何将comment_id string类型转化为int64类型
	commentId, err := strconv.ParseInt(comment_id, 10, 64) //用于对其进行转码操作.

	if err != nil {
		log.Errorf("commentId error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	log.Info(commentId) //此时表示成功将commentId进行转码
	videoId, err := strconv.ParseInt(video_id, 10, 64)
	if err != nil {
		log.Errorf("videoId error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}

	commentResponse, err := service.CommentAction(commentId, videoId, tokenUid, comment_text, actionType)
	if err != nil {
		log.Errorf("comment error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", commentResponse)
}

// 获取评论列表

func GetCommentList(ctx *gin.Context) {
	var err error
	video_id := ctx.Query("video_id")
	/* token := ctx.Query("token")
	_, err = util.VerifyToken(token)
	if err != nil {
		log.Errorf("token error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	} */
	videoId, err := strconv.ParseInt(video_id, 10, 64)
	if err != nil {
		log.Errorf("videoId error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	listResponse, err := service.CommentList(videoId)
	if err != nil {
		log.Infof("list error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", listResponse)
}
