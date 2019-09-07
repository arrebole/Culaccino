package main

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/sql"
)

func main() {

	if err := TestSet(&module.Storage{Name: "root", Password: "world"}); err != nil {
		panic(err.Error())
	}

	if err := TestSet(&module.Repo{Name: "first", Parents: "root"}); err != nil {
		panic(err.Error())
	}

	if err := TestSet(&module.Chapter{Name: "index", Parents: "first"}); err != nil {
		panic(err.Error())
	}
}

// TestSet ...
func TestSet(arg interface{}) error {
	return sql.Set(arg)
}
