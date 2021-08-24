package models

import (
	"fmt"
	"net/http"
	_ "path"
	_ "path/filepath"
	 "strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Model
	Author string `gorm:"not null, foreignKey"`
	Title  string `gorm:"not null" json:"title" form:"title"`
	Abstract string `gorm:"not null" json:"content" form:"content"`
	CreateAt time.Time
	Update int64 `grom:"autoCreateTime"`
}

type Index struct{
	Model
	Author string
	Title string
	Abstract string
}

//获取所有文章
func GetAllArticle(c *gin.Context){
	index,_ := strconv.Atoi(c.DefaultQuery("id", ""))
  var num Index
	//获取总文章数量
	db.Table("articles").Last(&num)
	allnum := num.ID
	fmt.Printf("num: %v", num)
	if allnum == 0{
		fmt.Println("获取文章总数失败")
		return
	}
	//分页 每页10条记录
	var articles []Index
	db.Table("articles").Limit(10).Offset(index * 10).Find(&articles)
	c.JSON(http.StatusOK, gin.H{"code":200, "articles":articles, "allnum":allnum})

}


//上传文章
func ArticleUpload(c *gin.Context){
	var article Article
	c.Bind(&article)
	// 获取文章数据等待存入bolt， mysql中只存如文摘就好[100字]
	content := article.Abstract
	fmt.Println("len(content): ", len(content))
	if len(content)  > 100 {
		abstract := content[0:100]
		article.Abstract = abstract
	}

	fmt.Println("content:", content)
	
  article.CreateAt = time.Now()
	//获取用户名
  name, _ := c.Get("user")
	article.Author = name.(string)
	db.Create(&article)
	c.JSON(http.StatusOK, gin.H{"code":200, "info": "article is upload success"})
	// 将content存入bolt中， 首先获取该文章id
	var index struct{
		ID int
	}
	db.Table("articles").Select("id").Where(&article).Find(&index)
	if index.ID == 0{
		c.JSON(http.StatusOK, gin.H{"code":400, "error":"文章上传失败"})
		return
	}

	if err := ArticleSet([]byte(content), []byte(string(index.ID))); err != nil{
		fmt.Println("文章落库失败: ", err)
		c.JSON(http.StatusOK, gin.H{"code":400, "error":"文章上传失败"})
		return
	}else{
		c.JSON(http.StatusOK, gin.H{"code":200, "error":"文章上传成功"})
	}

}


// func ArticleUpload(c *gin.Context) {
// 	//获取文件名， 并对后缀名进行判断
// 	file, _ := c.FormFile("file")
// 	var suffix string
// 	if suffix = path.Ext(file.Filename); suffix != ".txt" && suffix != ".md" {
// 		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "error": "The suffix of file is error"})
// 		return
// 	}

// 	//将文章索引加入数据库
// 	// if name, _ := c.Get("user"); name != " " {
// 	// 	c.JSON(http.StatusBadRequest, gin.H{"code": "400", "error": "用户权限不足"})
// 	// } else {
// 	// 	AddArticle(name.(string), file.Filename)
// 	// }
// 	index := AddArticle("weufh", file.Filename)
// 	index2 := strconv.Itoa(int(index)) + suffix
// 	path := filepath.Join("./art", index2)

// 	c.SaveUploadedFile(file, path)
// 	// if err := c.SaveUploadedFile(file, path); err != nil{
// 	// 	c.JSON(http.StatusBadRequest, c.JSON("code":"400", "error":"save error"))
// 	// }
// }

//通过用户名获取文章列表
func GetarticleOfUser(c *gin.Context) {
	name := c.Param("user")
	fmt.Println("name:", name)
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": "401", "error": "用户权限不足"})
	}

	//判断用户是否存在
	if UserIsexistByString(name) == "" {
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

//文章索引加入数据库, 返回文件唯一id

func AddArticle(author, title string) uint {
	db.Create(&Article{
		Author: author,
		Title:  title,
	})

	var article Article
	db.Where("title = ?", title).Last(&article)
	return article.ID

}

//-----------------------------------------------------获取某一篇文章--------------------------------------
// func Getarticle(c *gin.Context){
// 	articleid := c.Param("articleid")

// }
