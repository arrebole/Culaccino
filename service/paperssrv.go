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
		if data, err := client.HGetAll(key).Result(); err == nil {
			result = append(result, model.PaperBuilder(data, []string{"content"}))
		}
	}
	return result
}
