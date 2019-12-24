package model

// Response 服务器返回的json
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// DirectoryItem 目录项目
type DirectoryItem struct {
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
}

// RespDirector ...
func RespDirector(builder Builder, data []byte) []byte {
	return builder.Build(data)
}
