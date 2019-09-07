package sql

import(
	"errors"
)
// Exist 查询多个
func Exist(query ...string) error {
	switch len(query){
	case 1:
		return existStorage(query...)
	case 2:
		return existRepo(query...)
	case 3:
		return existChapter(query...)
	default:
		return errors.New("")
	}
}

func existStorage(query ...string) error {
	return errors.New("")
}

func existRepo(query ...string) error {
	return errors.New("")
}
func existChapter(query ...string) error {
	return errors.New("")
}