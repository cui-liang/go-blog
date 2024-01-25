package utils

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

// Config 是配置结构体
type Config struct {
	Title string
	MySQL database
}

type database struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

// ParseConfig 用于解析 TOML 文件
func LoadConfig() (Config, error) {
	var config Config
	_, err := toml.DecodeFile("conf/config.toml", &config)
	if err != nil {
		fmt.Println("Error decoding TOML:", err)
		return Config{}, err
	}
	return config, nil
}
