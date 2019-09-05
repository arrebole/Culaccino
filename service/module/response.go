package module

// Response 返回的json类型
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    `json:"data"`
}

// Data ...
type Data struct {
	User     UserInfo  `json:"user"`
	Archive  *Archive  `json:"archive"`
	Count    *Count    `json:"count"`
	Dir      []Archive `json:"dir"`
	FileInFo *FileInfo `json:"file"`
}

// UserInfo ...
type UserInfo struct {
	UID    uint   `json:"uid"`
	Domain string `json:"domain"`
	Token  string `json:"token"`
}

// Count ...
type Count struct {
	Total  int `json:"total"`
	Remain int `json:"remain"`
}

// FileInfo ...
type FileInfo struct {
	Hash string `json:"hash"`
	URL  string `json:"url"`
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
