package module

// Signal 返回json约定
type Signal struct {
	Code  string    `json:"code"`
	Power string    `json:"power"`
	Token string    `json:"token"`
	Main  *Article  `json:"article"`
	Count int       `json:"count"`
	Dir   []Article `json:"dir"`
}
