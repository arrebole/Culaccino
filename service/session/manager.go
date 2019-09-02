package session

import (
	"errors"
	"time"

	"github.com/arrebole/culaccino/service/share"
)

// Token ...
type Token string

// Store Session Manager interface
type Store interface {
	Get(token string) (Session, error)
	Set(*UserInfo) Token
}

// Manager ...
type Manager struct {
	storage map[string]*Session
	maxAge  time.Duration
}

// Get 从session仓库中获取用户信息
func (p *Manager) Get(token string) (Session, error) {
	user := p.storage[token]
	if user == nil {
		return Session{}, errors.New("no this session")
	}

	if user.Expired(p.maxAge) {
		delete(p.storage, token)
		return Session{}, errors.New("session Expired")
	}
	return *user, nil
}

// Set 从store从存放session，返回user_session
func (p *Manager) Set(info *UserInfo) Token {
	sessionID := share.RandString()
	info.Token = sessionID

	p.storage[sessionID] = &Session{
		UserInfo: *info,
		CreateAt: time.Now(),
	}
	return Token(sessionID)
}
