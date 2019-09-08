package module

import (
	"fmt"
	"strings"
)

// Repo 仓库
type Repo struct {
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
	Area    string `json:"area"`
	Star    int    `json:"star"`
	Hate    int    `json:"hate"`
	Tag     string `json:"tag"`
	Cover   string `json:"cover"`
	Summary string `json:"summary"`
	Weights int    `json:"weigths"`
	Views   int    `json:"views"`
}

// Parents 指向父名
func (p *Repo) Parents() string {
	array := strings.Split(p.Symbol, ":")
	return fmt.Sprintf("%s", array[0])
}

// NewRepo 唯一标识
func NewRepo(arg ...string) *Repo {
	if len(arg) != 2 {
		panic("new Repo arg need 2")
	}
	return &Repo{
		Symbol: fmt.Sprintf("%s:%s", arg[0], arg[1]),
		Name:   arg[1],
	}
}
