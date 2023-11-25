package common

import (
	"TikTokLite/log"
	"TikTokLite/response"
	"errors"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
)

var (
	Secret = []byte("TikTok")
	// TokenExpireDuration = time.Hour * 2 过期时间
)

type JWTClaims struct {
	UserId   int64  `json:"user_id"` //ToDo:表示为将UseId进行序列化时，其字段将会变为"user_id"
	Username string `json:"user_name"`
	jwt.RegisteredClaims
}

//生成token

func GenToken(userid int64, userName string) (string, error) {
	claims := JWTClaims{
		UserId:   userid,
		Username: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "Xui", //表示为发行者
			//ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),可用于设定token过期时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //ToDo:这个函数表示创建一个新的JWT对象，其中第一个参数表示为签名的方法，另一个是包含JWT负载的Claims对象负载是JWT的有效负载部分。它包含了实际的数据，例如用户 ID、用户名等。
	signedToken, err := token.SignedString([]byte("TikTok"))   //对token进行一个签名
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

//解析token

func ParsenToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

//验证token

func VerifyToken(tokenString string) (int64, error) {

	log.Debugf("tokenString:%v", tokenString)

	if tokenString == "" {
		return int64(0), nil
	}
	claims, err := ParsenToken(tokenString)
	if err != nil {
		return int64(0), err
	}
	return claims.UserId, nil
}

//=============================gin的中间件，就是一个函数，返回gin 的HandlerFunc======================================================

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		tokenString := c.GetHeader("token") //从Header的请求头中获取数据.

		userId, err := VerifyToken(tokenString)
		if err != nil || userId == int64(0) {
			response.Fail(c, "auth error", nil)
			c.Abort()
		}

		c.Set("UserId", userId)
		c.Next()
	}
}

// 部分接口不需要用户登录也可访问，如feed，publish list，favList，follow/follower list

func AuthWithOutMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		tokenString := c.GetHeader("token")

		userId, err := VerifyToken(tokenString)
		if err != nil {
			response.Fail(c, "auth error", nil)
			c.Abort()
		}

		c.Set("UserId", userId)
		//c *gin.Context是一个指向Gin上下文的指针123。Gin上下文是一个结构体，它包含了HTTP请求的所有信息，如请求头、请求体、查询参数等123。
		//c.Set("UserId", userId)是一个方法调用，它将键值对(“UserId”, userId)存储到Gin上下文中123。这个方法可以用来在请求的生命周期内存储和传递数据123
		//ToDo: *gin.Context这个用于设置上下文的标识，这个标识是从token解析来的.
		c.Next()
	}
}
