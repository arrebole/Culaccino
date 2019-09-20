package session

import (
	"errors"
	"time"

	"github.com/arrebole/culaccino/service/share"
)

// Store Session Manager interface
type Store interface {
	Get(sessionID string) (Body, error)
	Set(*Body) Body
}

// Manager ...
type Manager struct {
	SessionIDMap map[string]*Session
	SecretMap    map[string]*Session
	maxAge       time.Duration
}

// Get 从session仓库中获取用户信息
func (p *Manager) Get(sessionID string) (Body, error) {
	item, ok := p.SessionIDMap[sessionID]
	if !ok {
		return Body{}, errors.New("this session is undefinde")
	}

	if item.Expired(p.maxAge) {
		// 同时删除 SessionIDMap与SecretMap表中的内容
		delete(p.SessionIDMap, item.SessionID)
		delete(p.SecretMap, item.Secret)
		return Body{}, errors.New("session Expired")
	}

	return item.Body, nil
}

// Set 从store从存放session，返回user_session
func (p *Manager) Set(body *Body) Body {

	// 验证是否已登录
	item, ok := p.SecretMap[body.Secret]

	if !ok {
		body.SessionID = share.RandString()
		item = &Session{
			Body:     *body,
			CreateAt: time.Now(),
		}
		p.SessionIDMap[body.SessionID] = item
		p.SecretMap[body.Secret] = item
	}
	
	return item.Body
}
