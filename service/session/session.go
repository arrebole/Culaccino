package session

import "time"

// User ...
type User struct {
	UID        uint   `json:"uid"`
	Uname      string `json:"domain"`
	Token      string `json:"token"`
	Permission int    `json:"permission"`
}

// Session 储存会话
type Session struct {
	User

	CreateAt time.Time
	Options  string
}

// Expired 判断是否过期
func (p *Session) Expired(maxAge time.Duration) bool {
	return (time.Now().Sub(p.CreateAt) > maxAge)
}
