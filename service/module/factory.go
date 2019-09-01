package module

// ResponseSuccess 返回成功json
func ResponseSuccess() *Response {
	return (&Response{}).SetSuccess()
}

// ResponseFail 返回失败json
func ResponseFail() *Response {
	return (&Response{}).SetFail()
}

// ResponseData 返回查询数据
func ResponseData(data ...interface{}) *Response {
	result := (&Response{}).SetSuccess()

	for _, value := range data {
		switch value.(type) {
		case *Archive:
			result.Archive = value.(*Archive)
		case []Archive:
			result.Dir = value.([]Archive)
		case *Count:
			result.Count = value.(*Count)
		case *Session:
			result.Token = value.(*Session).Token
		case Islogin:
			result.Islogin = value.(Islogin)
		case *File:
			f := value.(*File)
			result.FileInFo = &FileInfo{
				Hash: f.CheckHash(),
				URL:  f.URL(),
			}
		default:
			panic("data error\n")
		}
	}
	return result
}

// ToArchive 将用户提交的article 转化成数据库的article格式
func ToArchive(data *PostArchive) *Archive {
	var result = &Archive{
		PostArchive: *data,
		Views:       0,
	}
	return result
}
