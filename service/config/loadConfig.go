package config

import (
	"os"
)

// 导出的配置文件
var (
	DBAddr     string
	RootDir    string
	UploadDir  string
	ListenPort string
	PassWord   string
)

func init() {
	config := envConfig()
	export(config)
	mkDir(config.UploadDir)
}

func export(config Config) {
	DBAddr = config.DBAddr
	ListenPort = config.ListenPort
	RootDir = config.RootDir
	UploadDir = config.UploadDir
}

// 初始化上传文件夹
func mkDir(path string) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil && !os.IsExist(err) {
		panic(err.Error())
	}
}
