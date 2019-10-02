package sql

import (
	"github.com/arrebole/culaccino/service/model"
)

// Interface sql对外接口
type Interface interface {
	SetStorage(*model.Storage) error
	GetStorage(string) *model.Storage
	ExistsStorage(string) bool

	SetRepo(*model.Repo) error
	GetRepo(string) *model.Repo
	GetRepos(...string) []model.Repo
	DelRepo(string) error
	ExistsRepo(string) bool

	SetChapter(*model.Chapter) error
	GetChapter(string) *model.Chapter
	GetChapters(...string) []model.Chapter
	DelChapter(string) error
	ExistsChapter(string) bool

	Explore(page int64, perPage int64) []string
	AddExploreItem(string)
	DelExploreItem(string)
}
