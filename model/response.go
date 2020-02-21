package model

import "encoding/json"

// Response ...
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Build ...
func (p *Response) Build(code int, msg string, data interface{}) []byte {
	p.Code, p.Message, p.Data = code, msg, data
	result, _ := json.Marshal(p)
	return result
}
