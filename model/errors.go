package model

// NotFound 未找到
var NotFound = Response{
	StatusCode: 404,
	Body:       map[string]string{"message": "Not Found"},
}

// Conflict 未找到
var Conflict = Response{
	StatusCode: 409,
	Body:       map[string]string{"message": "Existed"},
}

// DBFail ...
var DBFail = Response{
	StatusCode: 500,
	Body:       map[string]string{"message": "Database Error"},
}

// UnsupportedMediaType ...
var UnsupportedMediaType = Response{
	StatusCode: 415,
	Body:       map[string]string{"message": "Only receive json"},
}
