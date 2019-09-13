package main

import (
	"fmt"

	"github.com/arrebole/culaccino/service/model"
	"github.com/arrebole/culaccino/service/sql"
)

func main() {
	TestSet()
	TestGet()
	TestExists()
	TestDel()
}

// TestGet ...
func TestGet() {
	var storage = &model.Storage{}
	sql.Get(storage, "arch")
	fmt.Println(storage)

	var repo = &model.Repo{}
	sql.Get(repo, "arch", "dev")
	fmt.Println(repo)

	var chapter = &model.Chapter{}
	sql.Get(chapter, "arch", "dev", "index")
	fmt.Println(chapter)
}

// TestSet ...
func TestSet() {
	if err := sql.Set(&model.Storage{Name: "root", Password: "123"}); err != nil {
		panic(err.Error())
	}

	if err := sql.Set(model.NewRepo("arch", "dev")); err != nil {
		panic(err.Error())
	}
	if err := sql.Set(model.NewRepo("arch", "master")); err != nil {
		panic(err.Error())
	}

	if err := sql.Set(model.NewChapter("arch", "dev", "index")); err != nil {
		panic(err.Error())
	}
}

// TestExists ...
func TestExists() {
	fmt.Println(sql.Exists("root"))
	fmt.Println(sql.Exists("arch", "dev"))
	fmt.Println(sql.Exists("arch", "master"))
	fmt.Println(sql.Exists("arch", "dev", "index"))
}

func TestDel() {
	sql.Delete(model.NewRepo("arch", "dev"))
}
