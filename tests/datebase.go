package main

import (
	"fmt"

	"github.com/arrebole/culaccino/service/sql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	var client = sql.New()
	fmt.Printf("%s", client.Query(1))
}
