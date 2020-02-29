package service

import "github.com/arrebole/culaccino/model"

// Table ...
func (p PaperSrv) Table(limit int64) []*model.Paper {
	keys := client.Keys("*").Val()
	return pipeline(keys)
}

func pipeline(keys []string) []*model.Paper {
	var result = []*model.Paper{}
	for _, key := range keys {
		paper := &model.Paper{}
		data := client.HMGet(key, "title", "type", "summary", "create_at", "update_at", "cover").Val()
		result = append(result, paper.BuildFromSlice(data))
	}
	return result
}
