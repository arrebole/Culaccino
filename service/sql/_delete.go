package sql

import "github.com/arrebole/culaccino/service/module"

// Delete 删除
func Delete(arg interface{}) bool {
	switch arg.(type) {

	case *module.Repo:
		return deleteRepo(arg.(*module.Repo))

	case *module.Chapter:
		return deleteChapter(arg.(*module.Chapter))

	default:
		return false
	}
}

func deleteRepo(arg *module.Repo) bool {
	return false
}

func deleteChapter(arg *module.Chapter) bool {
	return false
}
