package handler

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
var MySecret = []byte("rfiders") //密钥
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请携带token访问",
			})
			c.Abort()
			return
		}
		// 按空格分割

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := ParseToken(authHeader)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "token失效",
			})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		c.Set("id",mc.Id)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
func GenToken(username string)(string,error){
	c := MyClaims{username,jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
		Issuer: "devicesManager",
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,c)
	return token.SignedString(MySecret)
}
func DoorMiddleWare() func(c *gin.Context){
	return func(c *gin.Context) {
		door := c.GetHeader("door")
		if door != "william"{
			c.JSON(200,gin.H{
				"code":500,
				"msg":"无权访问",
			})
			return
		}
		c.Next()
	}
}
func ParseToken(tokenString string)(*MyClaims,error){
	token,err := jwt.ParseWithClaims(tokenString,&MyClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return MySecret,nil
		},
	)
	if err!=nil{
		return nil,err
	}
	if claims,ok := token.Claims.(*MyClaims);ok && token.Valid{
		return claims,nil
	}
	return nil,errors.New("invalid token")
}