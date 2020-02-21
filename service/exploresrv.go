package service

import "github.com/arrebole/culaccino/model"

// Explore ...
func (p PaperSrv) Explore(limit int64) []*model.Paper {
	keys := client.Keys("*").Val()
	return pipeline(keys)
}

func pipeline(keys []string) []*model.Paper {
	var result []*model.Paper
	for _, key := range keys {
		paper := &model.Paper{}
		data := client.HMGet(key, "title", "type", "summary").Val()
		result = append(result, paper.BuildFromSlice(data))
	}
	return result
}
