package sql

import (
	"strings"
	"errors"
	"fmt"
	"github.com/arrebole/culaccino/service/model"
	
	"github.com/go-redis/redis"
)

func (p *SQL) SetChapter(arg *model.Chapter) error {
	p.SetChapterMap(arg.Symbol())
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

// GetChapterMap ...
func (p *SQL) GetChapterMap(arg string) []string {
	return p.ChapterDB.ZRange(arg, 0, -1).Val()
}

// SetChapterMap ...
// src  -> storageName `root:dev`
// dist -> repo.Symbol() `root:dev:index`
func (p *SQL) SetChapterMap(arg string) error {
	split := strings.Split(arg, ":")
	if len(split) != 3 {
		return errors.New("arg need xxx:xxx:xxx")
	}

	parents := fmt.Sprintf("%s:%s", split[0], split[1])

	return p.MapChapterDB.ZAdd(parents, redis.Z{
		Score:  0,
		Member: arg,
	}).Err()
}

// DelChapterMap ...
// arg: chapter.Symbol() `root:dev:xxx`
func (p *SQL) DelChapterMap(arg string) error {
	split := strings.Split(arg, ":")
	if len(split) != 3 {
		return errors.New("arg need xxx:xxx")
	}
	parents := fmt.Sprintf("%s:%s", split[0], split[1])
	return p.MapChapterDB.ZRem(parents, arg).Err()
}

// DeleteChapter 删除章节
func (p *SQL) DelChapter(arg string) error {
	p.DelChapterMap(arg)
	p.ChapterDB.Del(arg)
	return nil
}

// ExistsChapter 判断chapter是否存在
func (p *SQL) ExistsChapter(query string) bool {
	if p.ChapterDB.Exists(query).Val() > 0 {
		return true
	}
	return false
}
