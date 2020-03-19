package service

import "github.com/arrebole/culaccino/model"

// Table ...
func (p PaperSrv) Table(limit int64) []*model.Paper {
	keys := client.Keys("*").Val()
	return batch(keys)
}

// batch 使用key数组，批量获取paper
func batch(keys []string) []*model.Paper {
	var result = []*model.Paper{}
	for _, key := range keys {
		if data, err := client.HGetAll(key).Result(); err == nil {
			result = append(result, model.PaperBuilder(data, []string{"content"}))
		}
	}
	return result
}
