package model

// Sucess 单纯的成功
func Sucess() IResponse {
	return Response{
		StatusCode: 200,
		Body: map[string]string{
			"message": "sucess",
		},
	}
}

// CreateSucesss 成功创建时返回
func CreateSucesss(data *Paper) IResponse {
	return Response{
		StatusCode: 201,
		Body:       data,
	}
}

// OnePaper 获取单个paper
func OnePaper(data *Paper) IResponse {
	return Response{
		StatusCode: 200,
		Body:       data,
	}
}

// Papers 获取多个paper
func Papers(data []*Paper) IResponse {
	return Response{
		StatusCode: 200,
		Body:       data,
	}
}
