package sql

import (
	"github.com/arrebole/culaccino/service/model"
)

// SetChapter ...
func (p *SQL) SetChapter(arg *model.Chapter) error {
	return p.ChapterDB.HMSet(arg.Symbol(), adapter(arg)).Err()
}

// GetChapter ...
func (p *SQL) GetChapter(query string) *model.Chapter {
	var result = &model.Chapter{}
	reflectMapToStruct(result, p.ChapterDB.HGetAll(query).Val())
	return result
}

// GetChapters ...
func (p *SQL) GetChapters(arg ...string) []model.Chapter {
	var result = make([]model.Chapter, len(arg))
	cmds := getHashPipeline(p.ChapterDB, arg...)
	for i := 0; i < len(arg); i++ {
		reflectMapToStruct(&result[i], cmds[i].Val())
	}
	return result
}

// DelChapter 删除章节
func (p *SQL) DelChapter(arg string) error {
	return p.ChapterDB.Del(arg).Err()
}

// ExistsChapter 判断chapter是否存在
func (p *SQL) ExistsChapter(query string) bool {
	if p.ChapterDB.Exists(query).Val() > 0 {
		return true
	}
	return false
}
