package config

import (
	"encoding/json"
)

// Cofig 导出的配置文件
var Cofig Config

// Config 配置文件对象,用于从json序列化
type Config struct {
	SQL    SQL     `json:"sql"`
	Server Server  `json:"server"`
	Admin  []Admin `json:"admin"`
}

// Admin 管理员账户
type Admin struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

// SQL 连接的数据库
type SQL struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbName"`
}

// Server 服务器
type Server struct {
	ListenPort string `json:"listenPort"`
}

func init() {
	json.Unmarshal(loadConfig(), &Cofig)
}
