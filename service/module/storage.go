package module

import "time"

// Storage 主存储单元
type Storage struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Permission int    `json:"permission"`
	Status     int    `json:"status"`
	Avatar     string `json:"avatar"`
	Popular    int    `json:"popular"`
	Options    string `json:"options"`
}

// Repo 仓库
type Repo struct {
	Name    string `json:"name"`
	Parents string `json:"parents"`
	Area    string `json:"area"`
	Star    int    `json:"star"`
	Hate    int    `json:"hate"`
	Tag     string `json:"tag"`
	Cover   string `json:"cover"`
	Summary string `json:"summary"`
	Weights int    `json:"weigths"`
	Views   int    `json:"views"`
}

// Chapter Repo的一个章节
type Chapter struct {
	Index    int       `json:"index"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	Parents  string    `json:"parents"`
	Contents string    `json:"contents"`
}

// Issues 评论
type Issues struct {
	Parents  string `json:"parents"`
	Seat     int    `json:"Seat"`
	Contents string `json:"contents"`
}
