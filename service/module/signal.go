package module

// Signal 返回json约定
type Signal struct {
	Code  string    `json:"code"`
	Token string    `json:"token"`
	Main  *Article  `json:"articles"`
	Dir   []Article `json:"dir"`
}
