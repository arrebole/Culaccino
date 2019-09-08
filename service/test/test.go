package main

import (
	"fmt"

	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/sql"
)

func main() {
	TestSet()
	TestGet()
}

// TestGet ...
func TestGet() {
	var storage = &module.Storage{}
	sql.Get(storage, "arch")
	fmt.Println(storage)

	var repo = &module.Repo{}
	sql.Get(repo, "arch", "dev")
	fmt.Println(repo)

	var chapter = &module.Chapter{}
	sql.Get(chapter, "arch", "dev", "index")
	fmt.Println(chapter)
}

// TestSet ...
func TestSet() {
	if err := sql.Set(&module.Storage{Name: "arch", Password: "123"}); err != nil {
		panic(err.Error())
	}

	if err := sql.Set(module.NewRepo("arch", "dev")); err != nil {
		panic(err.Error())
	}

	if err := sql.Set(module.NewChapter("arch", "dev", "index")); err != nil {
		panic(err.Error())
	}
}
