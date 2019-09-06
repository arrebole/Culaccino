package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 导出的配置文件
var (
	DBAddr     string
	RootDir    string
	UploadDir  string
	ListenPort string
)

func init() {
	var config = &Config{}
	json.Unmarshal(loadConfig(), config)
	export(config)
	mkDir(UploadDir)
}

func export(config *Config) {

	DBAddr, _ = filepath.Abs(config.DBAddr)
	RootDir, _ = filepath.Abs(config.RootDir)
	UploadDir, _ = filepath.Abs(config.UploadDir)
	ListenPort = config.ListenPort
}

// 初始化上传文件夹
func mkDir(path string) {
	uploadPath, _ := filepath.Abs(path)
	if err := os.Mkdir(uploadPath, os.ModePerm); err != nil && !os.IsExist(err) {
		panic(err.Error())
	}
}

// 从文件中读取配置文件
func loadConfig() []byte {
	content, err := ioutil.ReadFile(configPath())
	if err != nil {
		panic(err.Error())
	}
	return content
}

// 返回配置文件路径
func configPath() string {
	currentPath, _ := filepath.Abs("./")
	return filepath.Join(currentPath, "config/config.json")
}
