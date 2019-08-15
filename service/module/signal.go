package module

// Signal 返回json约定
type Signal struct {
	Code  string    `json:"code"`
	Power string    `json:"power"`
	Token string    `json:"token"`
	Main  *Archive  `json:"archive"`
	Count int       `json:"count"`
	Dir   []Archive `json:"dir"`
}
