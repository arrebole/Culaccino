package main

import (
	"net/http"

	"github.com/arrebole/culaccino/service/controllers"
)

func main() {
	http.HandleFunc("/api/directory", controllers.DirectoryCtl)
	http.HandleFunc("/api/source", controllers.SourceCtl)
	http.Handle("/", http.FileServer(http.Dir("./themes/build")))
	http.ListenAndServe(":3000", nil)
}
