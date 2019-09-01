package sql

import "github.com/arrebole/culaccino/service/module"

func (p *client) UserExist(username string, password string) bool {
	var user = &module.User{}
	p.db.Where("username = ?", username).First(&user)

	if user.Password != hash(hash(password)) {
		return false
	}
	return true
}
