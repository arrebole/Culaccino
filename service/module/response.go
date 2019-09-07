package module

import "github.com/arrebole/culaccino/service/session"

// Response 返回的json类型
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    `json:"data"`
}

// Data ...
type Data struct {
	Session session.Body `json:"session"`
	Storage *Storage     `json:"storage"`
	Repos   []Repo       `json:"repos"`
	Repo    *Repo        `json:"repo"`
	Issues  []Issues     `json:"issues"`
	Count   *Count       `json:"count"`
	File    *FileStatus  `json:"file"`
}

// Count ...
type Count struct {
	Total  int `json:"total"`
	Remain int `json:"remain"`
}

func (p *Response) setCode(code int) *Response {
	p.Code = code
	return p
}

func (p *Response) setMessage(msg string) *Response {
	p.Message = msg
	return p
}

// SetSuccess 将code设置成成功
func (p *Response) SetSuccess() *Response {
	return p.setCode(0).setMessage("success")
}

//SetFail 设置code为失败
func (p *Response) SetFail() *Response {
	return p.setCode(-1).setMessage("fail")
}
