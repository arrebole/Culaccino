package model

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

// Issues 评论
type Issues struct {
	Parents  string `json:"parents"`
	Seat     int    `json:"Seat"`
	Contents string `json:"contents"`
}
