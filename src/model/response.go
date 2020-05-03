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

// Status ...
type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

// Token 访问权限控制
type Token struct {
	AccessToken string `json:"access_token"`
	// TokenType: bearer
	TokenType string `json:"token_type"`
	ExpiresAt int    `json:"expires_at"`
	// RefreshToken string `json:"refresh_token"`
}
