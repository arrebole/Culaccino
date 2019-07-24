package module

// Success 返回成功json
func Success() *Signal {
	return &Signal{
		Message: "success",
	}
}

// Fail 返回失败json
func Fail() *Signal {
	return &Signal{
		Message: "fail",
	}
}

// ToArticle 将用户提交的article 转化成数据库的article格式
func ToArticle(data *PostArticle) *Article {
	var result = &Article{
		PostArticle: *data,
		Views:       0,
	}
	return result
}
