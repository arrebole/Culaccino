package module

import (
	"math/rand"
	"time"
)

var local *Session

func init() {
	local = &Session{
		CreateAt: time.Now(),
		MaxAge:   time.Minute * 40,
	}
	local.ChangeToken()
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

// ChangeToken ...
func (p *Session) ChangeToken() *Session {
	rand.Seed(time.Now().Unix())
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randRune := make([]rune, 16)
	for i := range randRune {
		randRune[i] = letters[rand.Intn(len(letters))]
	}
	p.Token = string(randRune)
	return p
}

// NewSession 根据是否超时创建 token
func NewSession() *Session {
	if local.IsOverTime() {
		local.ChangeToken()
	}
	return local
}
