package sql

import "github.com/arrebole/culaccino/service/module"

// Commit 修改
func Commit(arg interface{}) bool {
	switch arg.(type) {

	case *module.Storage:
		return commitStorage(arg.(*module.Storage))

	case *module.Repo:
		return commitRepo(arg.(*module.Repo))

	case *module.Chapter:
		return commitChapter(arg.(*module.Chapter))

	default:
		return false
	}
}

func commitStorage(arg *module.Storage) bool {
	return false
}

func commitRepo(arg *module.Repo) bool {
	return false
}

func commitChapter(arg *module.Chapter) bool {
	return false
}
