package config

// Config 配置文件对象,用于从json序列化
type Config struct {
	ListenPort string `json:"port"`
	DBAddr     string `json:"db_addr"`
	RootDir    string `json:"root_dir"`
	UploadDir  string `json:"upload_dir"`
}
