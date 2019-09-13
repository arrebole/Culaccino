package sql

import(
	"github.com/arrebole/culaccino/service/model"
)

// Interface sql对外接口
type Interface interface{
	SetStorage(arg *model.Storage) error
	GetStorage(query string) *model.Storage
	ExistsStorage(query string) bool

	SetRepo(arg *model.Repo) error
	GetRepo(query string) *model.Repo
	GetRepos(arg ...string) []model.Repo
	DelRepo(arg string) error
	ExistsRepo(query string) bool

	SetChapter(arg *model.Chapter) error
	GetChapter(query string) *model.Chapter
	GetChapters(arg ...string) []model.Chapter
	DelChapter(arg string) error
	ExistsChapter(query string) bool

	Explore(page int64, perPage int64) []string
	AddExploreItem(RepoName string)
	DelExploreItem(RepoName string)
}

