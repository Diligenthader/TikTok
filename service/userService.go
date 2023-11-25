package service

import (
	"TikTokLite/common"
	message "TikTokLite/proto/pkg" //为这个包起了一个别名
	"TikTokLite/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// 用户注册

func UserRegister(userName, password string) (*message.DouyinUserRegisterResponse, error) {
	err := repository.UserNameIsExist(userName)
	if err != nil {
		return nil, err
	}
	info, err := repository.InsertUser(userName, password)
	if err != nil {
		return nil, err
	}
	token, err := common.GenToken(info.Id, userName)
	if err != nil {
		return nil, err
	}
	registResponse := &message.DouyinUserRegisterResponse{ //这个proto序列化返回结果
		UserId: info.Id,
		Token:  token,
	}
	return registResponse, nil
}

// 用户登录

func UserLogin(userName, password string) (*message.DouyinUserLoginResponse, error) {
	info, err := repository.GetUserInfo(userName)
	if err != nil {
		return nil, err
	}
	//验证密码是否正确
	err = bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(password)) //将数据库中的密码与用户带验证的密码进行比对
	if err != nil {
		return nil, errors.New("password error")
	}
	token, err := common.GenToken(info.Id, userName)
	if err != nil {
		return nil, err
	}
	loginResponse := &message.DouyinUserLoginResponse{
		UserId: info.Id,
		Token:  token,
	}
	return loginResponse, nil
}

//获取登录用户的信息

func UserInfo(userID int64) (*message.DouyinUserResponse, error) {
	info, err := repository.GetUserInfo(userID)
	if err != nil {
		return nil, err
	}
	user := messageUserInfo(info)
	return &message.DouyinUserResponse{User: user}, nil
}

func messageUserInfo(info repository.User) *message.User {
	return &message.User{
		Id:              info.Id,
		Name:            info.Name,
		FollowCount:     info.Follow,
		FollowerCount:   info.Follower,
		IsFollow:        false,
		Avatar:          info.Avatar,
		BackgroundImage: info.BackgroundImage,
		Signature:       info.Signature,
		TotalFavorited:  info.TotalFav,
		FavoriteCount:   info.FavCount,
	}
}
