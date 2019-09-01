package module

import (
	"time"
)

var local *Session

func init() {
	local = &Session{
		CreateAt: time.Now(),
		MaxAge:   time.Hour * 3,
	}
	local.changeToken()
}

// Session 储存会话
type Session struct {
	Token    string
	CreateAt time.Time
	MaxAge   time.Duration
}

// IsOverTime 判断是否过期
func (p *Session) IsOverTime() bool {
	return (time.Now().Sub(p.CreateAt) > p.MaxAge)
}

func (p *Session) refreshTime() {
	p.CreateAt = time.Now()
}

func (p *Session) changeToken() {
	p.Token = randString(16)
}

// NewSession 根据是否超时创建 token
func NewSession() *Session {
	if local.IsOverTime() {
		local.changeToken()
		local.refreshTime()
	}
	return local
}
