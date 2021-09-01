package main

import (
	"fmt"

	"traditional/config"
	"traditional/models"
	"traditional/routers"

	"github.com/spf13/viper"
	_ "gorm.io/driver/mysql"
)

var con = config.Config{
	"config/conf.yaml",
}

type User struct {
	ID   uint
	Name string
	Age  int
}

func main() {
	//初始化配置文件
	if err := con.Initconfig(); err != nil {
		fmt.Println("read config erro:", err)
		return
	}
	//初始化mysql
	
	mysqlUser := viper.GetString("mysql.user")
	mysqlPassword := viper.GetString("mysql.password")
	mysqlDatabase := viper.GetString("mysql.database")
	fmt.Println(mysqlUser, mysqlPassword, mysqlDatabase)
	models.InitialMysql(mysqlUser, mysqlPassword, mysqlDatabase)
	defer models.Close()

	//初始化bolt
	models.InitBolt(1)
  defer models.BlotClose()
	//开启http服务

	router.SetupRouter()
	router.Run()

}
