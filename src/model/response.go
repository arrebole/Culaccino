package model

// FullArticles 用于返回给用户的数据格式
type FullArticles struct {
	Articles   []FullArticle `json:"articles"`
	Pagination *Pagination   `json:"pagination"`
}

// FullArticle 用于返回给用户的数据格式
type FullArticle struct {
	Article
	Sections []Link `json:"sections"`
}

// Link ...
type Link struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Pagination 分页
type Pagination struct {
	TotalSize  int `json:"total_size"`  // 数据总量
	RemainSize int `json:"remain_size"` // 剩余总量
}

// Error ...
type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
