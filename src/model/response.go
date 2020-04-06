package model

// Pagination 分页
type Pagination struct {
	Total   int // 数据总量
	Pages   int // 页数总量
	Current int //当前页
}

// HetoasSectionLink ...
type HetoasSectionLink struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// ArticlesResponse 文章列表返回的数据
type ArticlesResponse struct {
	Articles   []ArticleResponse `json:"articles"`
	Pagination Pagination        `json:"pagination"`
}

// ArticleResponse 单个文章返回的数据
type ArticleResponse struct {
	Article      Article             `json:"article"`
	ArticleURL   string              `josn:"url"`
	SectionLinks []HetoasSectionLink `json:"section_links"`
}

// SectionResponse 文章返回数据
type SectionResponse struct {
	Section Section `json:"section"`
}
