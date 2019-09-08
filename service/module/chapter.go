package module

import (
	"fmt"
	"strings"
	"time"
)

// Chapter Repo的一个章节
type Chapter struct {
	Index    int       `json:"index"`
	Name     string    `json:"name"`
	Symbol   string    `json:"symbol"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	Contents string    `json:"contents"`
}

// Parents 指向父名
func (p *Chapter) Parents() string {
	array := strings.Split(p.Symbol, ":")
	return fmt.Sprintf("%s:%s", array[0], array[1])
}

// NewChapter 唯一标识
func NewChapter(arg ...string) *Chapter {
	if len(arg) != 3 {
		panic("new chapter arg need 3")
	}
	return &Chapter{
		Symbol: fmt.Sprintf("%s:%s:%s", arg[0], arg[1], arg[2]),
		Name:   arg[2],
	}
}
