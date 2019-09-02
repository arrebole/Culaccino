package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Cofig 导出的配置文件
var (
	DBName     string
	RootDir    string
	UploadDir  string
	ListenPort string
)

func init() {
	var _config *Config
	json.Unmarshal(loadConfig(), _config)
	export(_config)
	mkDir(UploadDir)
}

func export(config *Config) {
	DBName = config.DBName
	RootDir = config.RootDir
	UploadDir = config.UploadDir
	ListenPort = config.ListenPort
}

// 初始化上传文件夹
func mkDir(path string) {
	uploadPath, err := filepath.Abs(path)
	if err != nil {
		os.Mkdir(uploadPath, os.ModePerm)
	}
}

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
