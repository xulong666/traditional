package models

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Model struct {
	ID uint `grom:"primarykey autoIncrement"`
}
type User struct {
	Model
	Name     string `gorm:"unique, not null " json:"username" form:"username"`
	Password string `gorm:"not null" json:"password" form:"password"`
}

//用户注册
func Register(c *gin.Context) {
	var reg User
	if err := c.Bind(&reg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Printf("user: %#v", reg)
	if reg.Name == "" || reg.Password == "" {
		fmt.Printf("name and password can not be null")
	} else {
		if result := db.Create(&reg); result.Error != nil {
			fmt.Printf("%s insert erro:%v", reg.Name, result.Error)
		} else {
			c.JSON(http.StatusOK, gin.H{"register": "success"})
			c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1/user/login")
		}
	}
}

//用户登录
func Login(c *gin.Context) {
	fmt.Println("hello login")
	var login User
	var loginSql User
	c.Bind(&login)
	fmt.Println("username:", login.Name, "password", login.Password)
	if name := login.Name; name == "" {
		c.JSON(http.StatusOK, gin.H{"code":400, "error": "user is null"})
		fmt.Println("name is erro")
	} else {
		if result := db.Where("name=?", name).First(&loginSql); result.RowsAffected == 0 {
			c.JSON(http.StatusOK, gin.H{"code": 400, "error": "user is not exist"})
		} else {
			if login.Password != loginSql.Password {
				c.JSON(http.StatusOK, gin.H{"code": 400, "error": "password is invalid"})
				fmt.Println("Login: password is erro")
			} else {
				ReleaseTocken(c, name)
			}
		}
	}

}

//判断用户是否在数据库中, 通过ID
func UserIsexist(id uint) string {
	var user User
	db.Table("users").Where("ID = ?", id).First(&user)
	return user.Name

}

//通过string 判断用户是否存在
func UserIsexistByString(name string) string {
	var user User
	db.Table("users").Where("name = ?", name).First(&user)
	return user.Name

}

//获取用户信息
func Getuserinfo(c *gin.Context) {
	name, err := c.Get("user")
	if name != "" {
		c.JSON(http.StatusOK, gin.H{"code": 200, "username": name})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": 401, "info": "该用户不存在"})
	}
	fmt.Println("name: ", name, "error: ", err)

}

//通过name获取id
func GetId(name string) (ID uint) {
	var user User
	db.Table("users").Where("name=?", name).First(&user)
	return user.ID
}

//获取自己的文章列表
func MyArticle(c *gin.Context) {
	name, _ := c.Get("user")

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": "401", "error": "用户权限不足"})
	}

	//判断用户是否存在
	if UserIsexistByString(name.(string)) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "error": "用户不存在"})
		return
	}

	//获取改用户的文章id
	var indexs []struct {
		ID    uint
		Title string
	}
	db.Table("articles").Select("id", "title").Where("author = ?", name).Scan(&indexs)
	c.JSON(http.StatusBadRequest, gin.H{"code": 200, "article index": indexs})
}
