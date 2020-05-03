package repos

import "github.com/arrebole/culaccino/src/model"

// UsersRepository ...
type UsersRepository struct{}

// Check ...
func (p UsersRepository) Check(username, password string) bool {
	var count int
	db.Model(&model.User{}).Where("username = ? AND password = ?", username, password).Count(&count)
	return count > 0
}
