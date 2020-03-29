package model

import (
	"encoding/json"
)

// Paper 一篇文章
type Paper struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Type     string `json:"type"`
	Cover    string `json:"cover"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
	Summary  string `json:"summary"`
	Content  string `json:"content"`
}

// ToMap ...
func (p *Paper) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":        p.ID,
		"title":     p.Title,
		"type":      p.Type,
		"cover":     p.Cover,
		"create_at": p.CreateAt,
		"update_at": p.UpdateAt,
		"summary":   p.Summary,
		"content":   p.Content,
	}
}

// PaperBuilder ...
func PaperBuilder(data map[string]string, excludes []string) *Paper {
	var result = &Paper{}

	// 移除排除数据
	for _, v := range excludes {
		data[v] = ""
	}

	// 序列化为结构体
	jsonStr, _ := json.Marshal(data)
	json.Unmarshal(jsonStr, result)
	return result
}

// itos 接口类型数组转string数组
func itos(data []interface{}) []string {
	var result []string
	for _, v := range data {
		switch v.(type) {
		case string:
			result = append(result, v.(string))
		}
	}
	return result
}
