package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) {
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("../../config/")
	config.AddConfigPath("config/")
	config.AddConfigPath("/etc/xiuexcel/")
	config.AddConfigPath("C:\\")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatalf("error on parseing configuration file, error: %v", err)
	}
}

// GetConfig 返回配置信息
func GetConfig() *viper.Viper {
	return config
}
