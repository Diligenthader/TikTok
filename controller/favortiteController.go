package controller

import (
	"TikTokLite/log"
	"TikTokLite/response"
	"TikTokLite/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FavActionParams struct {
	// 暂时没 user_id ，因为客户端出于安全考虑没给出
	/*	Token      string `form:"token" binding:"required"`*/
	VideoId    int64 `form:"video_id" binding:"required"`              //ToDo:在这个变量的定义过程中，这两个变量都是在一个表单中提交的数据，需要进行验证，同时binding:"required" 则表示在这个表单中，这个字段是必填的，否则就会验证失败
	ActionType int8  `form:"action_type" binding:"required,oneof=1 2"` //binding:"required,oneof=1 2" 标签告诉Gin框架这个参数是必需的，而且它的值必须是1或2，否则Gin会返回一个错误。
	/*
		form:"video_id" 标签告诉Gin框架这个字段对应请求中的 video_id 参数。binding:"required" 标签告诉Gin框架这个参数是必需的，如果请求中没有这个参数，Gin会返回一个错误。
	*/
}

type FavListParams struct {
	/*	Token  string `form:"token" binding:"required"`*/
	UserId int64 `form:"user_id" binding:"required"`
}

// 点赞视频

func FavoriteAction(ctx *gin.Context) {
	var favInfo FavActionParams
	err := ctx.ShouldBindQuery(&favInfo) //ToDo: 这一行代码调用了ShouldBindQuery函数，将HTTP请求的查询参数绑定到favInfo变量上
	/*
		首先声明了一个名为favInfo的变量，这是一个已经定义好的结构体变量
		err := ctx.ShouldBindQuery(&favInfo)
		这行代码尝试将请求的查询参数绑定到 favInfo 变量。ctx.ShouldBindQuery 是Gin框架提供的一个方法，它会解析HTTP请求的查询字符串，并尝试将结果填充到提供的结构体中。在这个例子中，它尝试将查询参数填充到 favInfo 中。
		ToDo:ctx是一个指向gin.Context类型的指针。gin.Context是Gin框架中的一个结构体类型，用于封装HTTP请求和响应的信息，以及提供一些方法，用于获取请求和响应的信息、设置响应头、设置响应状态码等操作1。
	*/
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	tokenUids, _ := ctx.Get("UserId")
	//和前面的操作一样，根据传入的token解析用户身份
	tokenUid := tokenUids.(int64)

	if err != nil {
		log.Errorf("token error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	err = service.FavoriteAction(tokenUid, favInfo.VideoId, favInfo.ActionType) //传入了三个参数，其中包含有用户的Uid，发送的请求包括的点赞的视频Id，以及视频的类型，表示为点赞的视频与被删除的视频

	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", nil)
}

// 获取点赞列表

func GetFavoriteList(ctx *gin.Context) {

	UserId := ctx.Query("user_id") //根据用户id查询点赞的视频集合列表
	tokenUids, _ := ctx.Get("UserId")
	tokenUid := tokenUids.(int64)
	uid, err := strconv.ParseInt(UserId, 10, 64)
	if err != nil {
		log.Errorf("userid error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	favList, err := service.FavoriteList(tokenUid, uid)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", favList)
}
