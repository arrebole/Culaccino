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
	Get(token string) (Body, error)
	Set(*Body) Body
}

// Manager ...
type Manager struct {
	storage map[string]*Session
	maxAge  time.Duration
}

// Get 从session仓库中获取用户信息
func (p *Manager) Get(token string) (Body, error) {
	item := p.storage[token]
	if item == nil {
		return Body{}, errors.New("no this session")
	}

	if item.Expired(p.maxAge) {
		delete(p.storage, token)
		return Body{}, errors.New("session Expired")
	}
	return item.Body, nil
}

// Set 从store从存放session，返回user_session
func (p *Manager) Set(body *Body) Body {
	sessionID := share.RandString()

	body.SessionID = sessionID
	item := &Session{
		Body:     *body,
		CreateAt: time.Now(),
	}
	p.storage[sessionID] = item
	return item.Body
}
