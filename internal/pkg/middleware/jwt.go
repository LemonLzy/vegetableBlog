package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lemonlzy/vegetableBlog/internal/pkg/error"
	"github.com/lemonlzy/vegetableBlog/internal/pkg/resp"
	"strings"
	"time"
)

const (
	TokenExpireDuration = time.Hour * 2
	Issuer              = "v-blog"
)

var mySecret = []byte("show me the code.")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们需要额外记录username字段用于区分用户，所以自定义结构体
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userID int64, username string) (aToken, rToken string, err error) {
	// access token
	c := jwt.MapClaims{
		"UserID":    userID,
		"Username":  username,
		"Issuer":    Issuer,
		"ExpiresAt": time.Now().Add(TokenExpireDuration).Unix(),
	}
	// 使用指定的签名方法创建签名对象(注意使用HS512，不要使用ES512)
	aToken, _ = jwt.NewWithClaims(jwt.SigningMethodHS512, c).SignedString(mySecret)

	// refresh token
	rToken, _ = jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Millisecond * 30).Unix(),
		Issuer:    Issuer,
	}).SignedString(mySecret)
	// 使用指定的secret签名并获得完整编码后的字符串token
	return
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

// RefreshToken 刷新AccessToken
func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	_, err = jwt.Parse(rToken, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return
	}

	// 从旧access token中解析出claims数据
	var mc = new(MyClaims)
	_, err = jwt.ParseWithClaims(aToken, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	v, _ := err.(*jwt.ValidationError)

	// 当access token是过期错误，并且refresh token没有过期时就创建一个新的access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(mc.UserID, mc.Username)
	}

	return
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// Authorization: Bearer xxxxx.xxx.xxxx
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			resp.ResponseErrorWithMsg(c, errCode.ServerInvalidToken, "")
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			resp.ResponseErrorWithMsg(c, errCode.ServerInvalidToken, "")
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := ParseToken(parts[1])
		if err != nil {
			resp.ResponseErrorWithMsg(c, errCode.ServerInvalidToken, "")
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set("user_id", mc.UserID)
		c.Next() // 后续的处理请求的函数中 可以通过c.Get(ContextUserIDKey) 来获取当前请求的用户信息
	}
}
