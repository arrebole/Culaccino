package session

import "time"

// Body ...
type Body struct {
	Secret     string `json:"secret"`
	SessionID  string `json:"cookie"`
	Permission int    `json:"permission"`
}

// Session 储存会话
type Session struct {
	Body

	CreateAt time.Time
	Options  string
}

// Expired 判断是否过期
func (p *Session) Expired(maxAge time.Duration) bool {
	return (time.Now().Sub(p.CreateAt) > maxAge)
}
