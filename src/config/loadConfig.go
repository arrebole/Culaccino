package config

import (
	"io/ioutil"
	"path/filepath"
)

// 从文件中读取配置文件
func loadConfig() []byte {
	content, err := ioutil.ReadFile(configPath())
	if err != nil {
		panic("can not load config file :" + configPath())
	}
	return content
}

// 返回配置文件路径
func configPath() string {
	currentPath, _ := filepath.Abs("./")
	return filepath.Join(currentPath, "./config/config.json")
}
