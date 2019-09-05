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
	Get(token string) (User, error)
	Set(*User) User
}

// Manager ...
type Manager struct {
	storage map[string]*Session
	maxAge  time.Duration
}

// Get 从session仓库中获取用户信息
func (p *Manager) Get(token string) (User, error) {
	aSession := p.storage[token]
	if aSession == nil {
		return User{}, errors.New("no this session")
	}

	if aSession.Expired(p.maxAge) {
		delete(p.storage, token)
		return User{}, errors.New("session Expired")
	}
	return aSession.User, nil
}

// Set 从store从存放session，返回user_session
func (p *Manager) Set(user *User) User {
	sessionID := share.RandString()

	user.Token = sessionID
	var newSession = &Session{
		User:     *user,
		CreateAt: time.Now(),
	}
	p.storage[sessionID] = newSession
	return newSession.User
}
