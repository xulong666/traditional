package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

//读取配置
func (c *Config) Initconfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("config")
		viper.SetConfigName("conf")
	}
	viper.SetConfigType("yaml")

	return viper.ReadInConfig()

}

// 监控配置改动
func (c *Config) WatchConfig(change chan int) {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("配置已经被改变: %s", e.Name)
		change <- 1
	})
}
