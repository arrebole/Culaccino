package sql

import "github.com/arrebole/culaccino/service/module"

// Create 创建
func Create(arg interface{}) bool {
	switch arg.(type) {

	case *module.Storage:
		return createStorage(arg.(*module.Storage))

	case *module.Repo:
		return createRepo(arg.(*module.Repo))

	case *module.Chapter:
		return createChapter(arg.(*module.Chapter))

	default:
		return false
	}
}

func createStorage(arg *module.Storage) bool {
	return false
}

func createRepo(arg *module.Repo) bool {
	return false
}

func createChapter(arg *module.Chapter) bool {
	return false
}
