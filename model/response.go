package model

import "encoding/json"

// IResponse ...
type IResponse interface {
	Code() int
	Data() []byte
}

// Response 实现
type Response struct {
	StatusCode int
	Body       interface{}
}

// Code ...
func (p Response) Code() int {
	return p.StatusCode
}

// Data ...
func (p Response) Data() []byte {
	data, err := json.Marshal(p.Body)
	if err != nil {
		return nil
	}
	return data
}
