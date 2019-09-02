package sql

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/session"
)

func (p *client) UserBaseInfo(username string, password string) *session.UserInfo {
	var user = &module.User{}
	p.db.Where("username = ?", username).First(&user)

	if user.Password != PassWord(password) {
		return nil
	}
	return &session.UserInfo{
		UID:        user.ID,
		Uname:      user.Username,
		Permission: user.Permission,
	}
}

// func (p *client) UserAddArchiveIDs(uid uint, archiveID uint) {
// 	user := &module.User{}
// 	p.db.First(&user, uid)
// }
