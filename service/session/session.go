package session

import "time"

// UserInfo ...
type UserInfo struct {
	UID        uint
	Uname      string
	Token      string
	Permission int
}

// Session 储存会话
type Session struct {
	UserInfo

	CreateAt time.Time
	Options  string
}

// Expired 判断是否过期
func (p *Session) Expired(maxAge time.Duration) bool {
	return (time.Now().Sub(p.CreateAt) > maxAge)
}
