package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	MySQL MySQLConfig
	Redis RedisConfig
}

type MySQLConfig struct {
	Host           string
	Port           string
	User           string
	Password       string
	Charset        string
	Database       string
	MaxIdleConn    uint
	SetMaxOpenConn uint
}

type RedisConfig struct {
	IP   string
	Port int
	Auth string
}

func GetConfig(key string) string {
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	_ = viper.ReadInConfig()
	value := viper.GetString(key)
	if value == "" {
		viper.Reset()
		viper.SetConfigType("json")
		viper.AddConfigPath(".")
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("[Warring]: Fatal error config file: %w \n", err)
		}
		value = viper.GetString(key)
	}
	return value
}

func GetConfigAndDefualt(key string, defualt string) string {
	value := GetConfig(key)
	if value == "" {
		return defualt
	}
	return value
}

func GetConfigByMap(key string) map[string]any {
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	_ = viper.ReadInConfig()
	value := viper.GetStringMap(key)
	if value == nil {
		viper.Reset()
		viper.SetConfigType("json")
		viper.AddConfigPath(".")
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("[Warring]: Fatal error config file: %w \n", err)
		}
		value = viper.GetStringMap(key)
	}
	return value
}
