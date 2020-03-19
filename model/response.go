package model

import "encoding/json"

// IResponse 抽象接口
type IResponse interface {
	ToBytes() []byte
}

// Response ...
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ToBytes 将数据转换为[]byte
func (p Response) ToBytes() []byte {
	result, _ := json.Marshal(p)
	return result
}

// CreateResponse Responsego的工厂方法
func CreateResponse(code int, msg string, data interface{}) IResponse {
	return Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}
