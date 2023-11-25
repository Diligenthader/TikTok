package controller

import (
	"TikTokLite/log"
	"TikTokLite/response"
	"TikTokLite/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 用户登录
func UserLogin(ctx *gin.Context) {
	var err error
	userName := ctx.Query("username")
	password := ctx.Query("password")
	if len(userName) > 32 || len(password) > 32 { //最长32位字符
		response.Fail(ctx, "username or password invalid", nil)
		return
	}
	loginResponse, err := service.UserLogin(userName, password)
	if err != nil {
		log.Infof("login error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", loginResponse)
}

func UserRegister(ctx *gin.Context) {
	var err error
	userName := ctx.Query("username")
	password := ctx.Query("password")
	if len(userName) > 32 || len(password) > 32 { //最长32位字符
		response.Fail(ctx, "username or password invalid", nil)
		return
	}
	registResponse, err := service.UserRegister(userName, password)
	if err != nil {
		log.Infof("registe error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", registResponse)

}

//获取用户信息

func GetUserInfo(ctx *gin.Context) {

	var err error
	userId := ctx.Query("user_id") //获取Param请求的的user_id
	uids, _ := ctx.Get("UserId")   //这个表示为获取token中的UserId，这是可以从用户发出的请求中的token解析得到

	uid := uids.(int64)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}

	if strconv.FormatInt(uid, 10) != userId { //则表示为登录的用户和发送查询用户的id不同，则无法查询
		response.Fail(ctx, "token error", nil)
		return
	}
	userinfo, err := service.UserInfo(uid)
	if err != nil {
		log.Infof("get userinfo  error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", userinfo)

}
