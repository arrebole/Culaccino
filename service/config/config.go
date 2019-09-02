package config

// Config 配置文件对象,用于从json序列化
type Config struct {
	ListenPort string `json:"port"`
	DBName     string `json:"dbname"`
	RootDir    string `json:"root_dir"`
	UploadDir  string `json:"upload_dir"`
}
