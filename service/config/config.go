package config

// Config 配置文件对象,用于从json序列化
type Config struct {
	DBName     string `json:"dbName"`
	RootDir    string `json:"root_dir"`
	UploadDir  string `json:"upload_dir"`
	ListenPort string `json:"listen_port"`
}
