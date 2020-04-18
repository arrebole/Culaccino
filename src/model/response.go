package model

// Articles 用于返回给用户的数据格式
type Articles struct {
	Articles   []Article   `json:"articles"`
	Pagination *Pagination `json:"pagination"`
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
