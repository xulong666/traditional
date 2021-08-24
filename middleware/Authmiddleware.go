package middleware

import (
	_ "fmt"
	"net/http"
	"traditional/models"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusServiceUnavailable, gin.H{"code": 401, "msg": "权限不足"})
		c.Abort()
		return
	}
	token, claims, err := models.ParseTocken(tokenString)
	if err != nil || !token.Valid {
		c.JSON(http.StatusServiceUnavailable, gin.H{"code": 401, "msg": "权限不足"})
		c.Abort()
		return
	}

	if username := models.UserIsexist(claims.UserId); username == "" {
		c.JSON(http.StatusServiceUnavailable, gin.H{"code": 401, "msg": "权限不足"})
		c.Abort()
	} else {
		//如果用户存在， 见用户写入上下文
		c.Set("user", username)
		c.Next()
	}

}
