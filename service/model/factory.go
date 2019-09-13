package model

import "github.com/arrebole/culaccino/service/session"

// ResponseSuccess 返回成功json
func ResponseSuccess(data ...interface{}) *Response {
	return setResponseData(data...).SetSuccess()
}

// ResponseFail 返回失败json
func ResponseFail(data ...interface{}) *Response {
	return setResponseData(data...).SetFail()
}

// ResponseData 返回查询数据
func setResponseData(data ...interface{}) *Response {
	result := &Response{}

	for _, value := range data {
		switch value.(type) {
		case string:
			result.Message = value.(string)

		case *Storage:
			result.Storage = value.(*Storage)

		case []Repo:
			result.Repos = value.([]Repo)

		case *Repo:
			result.Repo = value.(*Repo)

		case *Count:
			result.Count = value.(*Count)

		case session.Body:
			result.Session = value.(session.Body)

		case *FileStatus:
			result.File = value.(*FileStatus)

		default:
			panic("data error\n")
		}
	}
	return result
}
