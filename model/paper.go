package model

// Paper 一篇文章
type Paper struct {
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
		"title":     p.Title,
		"type":      p.Type,
		"cover":     p.Cover,
		"create_at": p.CreateAt,
		"update_at": p.UpdateAt,
		"summary":   p.Summary,
		"content":   p.Content,
	}
}

// BuildFromMap ...
func (p *Paper) BuildFromMap(data map[string]string) *Paper {
	p.Title = data["title"]
	p.Type = data["type"]
	p.CreateAt = data["create_at"]
	p.UpdateAt = data["update_at"]
	p.Summary = data["summary"]
	p.Content = data["content"]
	p.Cover = data["cover"]
	return p
}

// BuildFromSlice ...
func (p *Paper) BuildFromSlice(data []interface{}) *Paper {
	var tData = itos(data)
	p.Title = tData[0]
	p.Type = tData[1]
	p.Summary = tData[2]
	p.CreateAt = tData[3]
	p.UpdateAt = tData[4]
	p.Cover = tData[5]
	return p
}

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
