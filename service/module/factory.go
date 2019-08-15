package module

func makeCode(code string) *Signal {
	return &Signal{
		Code: code,
	}
}

// Count 返回查询数量请求
func Count(c int) *Signal {
	result := makeCode("success")
	result.Count = c
	return result
}

// Power 权限
func Power() *Signal {
	result := makeCode("success")
	result.Power = "admin"
	return result
}

// LoginOk 返回登陆成功
func LoginOk() *Signal {
	result := makeCode("success")
	result.Token = NewToken()
	return result
}

// RepData 返回查询数据
func RepData(data ...interface{}) *Signal {
	result := makeCode("success")

	for _, value := range data {
		switch value.(type) {
		case *Archive:
			result.Main = value.(*Archive)
		case []Archive:
			result.Dir = value.([]Archive)
		case int:
			result.Count = value.(int)
		default:
			panic("data error\n")
		}
	}
	return result
}

// Success 返回成功json
func Success() *Signal {
	return makeCode("success")
}

// Fail 返回失败json
func Fail() *Signal {
	return makeCode("fail")
}

// ToArchive 将用户提交的article 转化成数据库的article格式
func ToArchive(data *PostArchive) *Archive {
	var result = &Archive{
		PostArchive: *data,
		Views:       0,
	}
	return result
}
