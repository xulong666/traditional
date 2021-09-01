package models

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// --------------------------------------------------------mysql-----------------------------------------------------------------
var db *gorm.DB

//初始化数据库
func InitialMysql(user, password, database string) {
	var err error
	dsn := user + ":" + password + "@tcp(127.0.0.1:3306)/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "root:123@tcp(127.0.0.1:3306)/yang?charset=utf8mb4&parseTime=True&loc=Local"
	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err == nil {
		fmt.Print("mysql connect success")
	} else {
		panic("mysql connect erro:")
	}
	db.AutoMigrate(&User{}, &Article{})

}

func SetConn(maxid, maxopen int) (err error) {
	if sqlDB, err := db.DB(); err != nil {
		fmt.Println("get sqldb err:", err)
		return err
	} else {
		sqlDB.SetMaxIdleConns(maxid)
		sqlDB.SetMaxOpenConns(maxopen)
	}

	return nil
}

//关闭数据库
func Close() {
	//注意： gromv2不在支持db.close(), 而是需要转为对应的sqldb在close， 其他类型的数据库也一样
	if sqldb, err := db.DB(); err != nil {
		fmt.Println("sqldb close err:", err)
	} else {
		sqldb.Close()
	}

}

//-----------------------------------------------------------bolt部分-----------------------------------------------------

var boltdb *bolt.DB
var bucketArt = []byte("arti") //存储文章的分组
//初始化bolt
func InitBolt(sleepNum int) *bolt.DB {
	//bolt一定传入绝对路径
	var err error
	boltdb, err = bolt.Open("D:/go_workspace/src/traditional/art/article.db", 0600, &bolt.Options{Timeout:time.Duration(sleepNum) * time.Second})
	if err != nil{
		fmt.Println("InitBolt:", err)
	}
	return boltdb
}

//关闭bolt
func BlotClose(){
	boltdb.Close()
}

//写入文章
func ArticleSet(artid int, arti string) error{
	if err := boltdb.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucketArt)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(string(artid)), []byte(arti))
		if err != nil {
			return err
		}
		return nil
	}); err != nil{
		return err
	}else{
		return nil
	}
}

//获取文章
func ArticleGet(artid int, content *string) {
	 boltdb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketArt)
		v := b.Get([]byte(string(artid)))
    *content = string(v)
		return nil
	})
}

