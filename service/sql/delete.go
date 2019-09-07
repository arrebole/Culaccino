package sql

import (
	"errors"

	"github.com/arrebole/culaccino/service/module"
)

// Delete 删除
func Delete(arg interface{}) error {
	switch arg.(type) {

	case *module.Repo:
		return deleteRepo(arg.(*module.Repo))

	case *module.Chapter:
		return deleteChapter(arg.(*module.Chapter))

	default:
		return errors.New("")
	}
}

func deleteRepo(arg *module.Repo) error {
	return errors.New("")
}

func deleteChapter(arg *module.Chapter) error {
	return errors.New("")
}
