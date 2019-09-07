package sql

import (
	"errors"

	"github.com/arrebole/culaccino/service/module"
)

// Get 查询多个
func Get(out interface{}, query ...string) error {
	switch out.(type) {

	case *module.Storage:
		return getStorage(out.(*module.Storage), query...)

	case *module.Repo:
		return getRepo(out.(*module.Repo), query...)

	case *module.Chapter:
		return getChapter(out.(*module.Chapter), query...)

	default:
		return errors.New("")
	}
}

func getStorage(arg *module.Storage, query ...string) error {
	return errors.New("")
}

func getRepo(arg *module.Repo, query ...string) error {
	return errors.New("")
}
func getChapter(arg *module.Chapter, query ...string) error {
	return errors.New("")
}
