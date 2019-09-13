package config

import(
	"path/filepath"
	"os"
	"github.com/arrebole/culaccino/service/share"
)

// Config 配置文件对象,用于从json序列化
type Config struct {
	ListenPort string `json:"port"`
	DBAddr     string `json:"db_addr"`
	RootDir    string `json:"root_dir"`
	UploadDir  string `json:"upload_dir"`
	PassWord   string `json:"password"`
}

func envConfig() Config{
	var config = &Config{}
	config.setPassWord().setDBAddr().setRootDir().setUploadDir().setListenPort()
	return *config
}

// 获取环境变量中的默认密码
func (p *Config)setPassWord() *Config {
	v := os.Getenv("PASSWORD")
	if v == "" {
		v = "root@culaccino"
	}
	p.PassWord = share.PassWord(v)
	return p
}

// 获取环境变量中的DBAddr
func (p *Config)setDBAddr() *Config {
	v := os.Getenv("DB_ADDR")
	if v == "" {
		v = "localhost:6379"
	}
	p.DBAddr = v
	return p
}

// 获取环境变量中的DBAddr
func (p *Config)setRootDir() *Config {
	v := os.Getenv("DB_ADDR")
	if v == "" {
		v = "./wwwwroot"
	}
	path,_ := filepath.Abs(v)
	p.RootDir = path
	return p
}

// 获取环境变量中的DBAddr
func (p *Config)setUploadDir() *Config {
	v := os.Getenv("UPLOAD_DIR")
	if v == "" {
		v = "./wwwwroot/static"
	}
	path,_ := filepath.Abs(v)
	p.UploadDir = path
	return p
}

// 获取环境变量中的DBAddr
func (p *Config)setListenPort() *Config {
	v := os.Getenv("PORT")
	if v == "" {
		v = ":3000"
	}
	p.ListenPort = v
	return p
}

