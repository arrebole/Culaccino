package model

import (
	"fmt"
)

// Repo 仓库
type Repo struct {
	Storage string `json:"storage"`
	Name    string `json:"name"`
	Area    string `json:"area"`
	Star    int    `json:"star"`
	Hate    int    `json:"hate"`
	Tag     string `json:"tag"`
	Cover   string `json:"cover"`
	Summary string `json:"summary"`
	Weights int    `json:"weigths"`
	Views   int    `json:"views"`
}

// Symbol 指向父名
func (p *Repo) Symbol() string {
	return fmt.Sprintf("%s:%s", p.Storage, p.Name)
}

// Parents 指向父名 dev:a
func (p *Repo) Parents() string {
	return p.Storage
}

// NewRepo 唯一标识
func NewRepo(arg ...string) *Repo {
	if len(arg) != 2 {
		panic("new Repo arg need 2")
	}
	return &Repo{
		Storage: arg[0],
		Name:    arg[1],
	}
}
