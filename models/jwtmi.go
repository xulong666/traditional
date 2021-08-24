package models

import (
	"fmt"
	"net/http"
	_ "strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

var jwtkey = []byte("hello world")

//颁发token
func ReleaseTocken(c *gin.Context) {
	username := c.PostForm("username")
	ID := GetId(username) //获取数据库中的id， 赋值给claims.UserId
	expirationTime := time.Now().Add(2 * time.Hour)
	claims := &Claims{
		UserId: ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",  // 签名颁发者
			Subject:   "user token", //签名主题
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenString, err := token.SignedString(jwtkey); err != nil {
		fmt.Println("setting token error:", err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"login":"success",
			"token:": tokenString,
		})
	}

}

//解析token
func Getting(c *gin.Context) Claims {
	var cl = Claims{}
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusServiceUnavailable, gin.H{"code": 401, "msg": "权限不足"})
		c.Abort()
		return cl
	}
	token, claims, err := ParseTocken(tokenString)
	fmt.Println("*********************************************************")
	fmt.Println("token:", token, "  claims: ", claims, "  error: ", err)
	fmt.Println("*********************************************************")
	if err != nil || !token.Valid {
		c.JSON(http.StatusServiceUnavailable, gin.H{"code": 401, "msg": "权限不足"})
		c.Abort()
		return cl
	}
	return *claims

}

//获取tocken内容
func ParseTocken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}
