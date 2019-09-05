package sql

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/session"
)

func (p *client) UserBaseInfo(username string, password string) *session.User {
	var user = &module.User{}
	p.db.Where("username = ?", username).First(&user)

	if user.Password != PassWord(password) {
		return nil
	}
	return &session.User{
		UID:        user.ID,
		Uname:      user.Username,
		Permission: user.Permission,
	}
}
