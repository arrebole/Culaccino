package main

import (
	"net/http"

	"github.com/arrebole/culaccino/service/config"
	"github.com/arrebole/culaccino/service/controllers"
)

func main() {
	http.HandleFunc("/api/directory", controllers.DirectoryCtl(config.DirectoryfilePath))
	http.HandleFunc("/api/source", controllers.SourceCtl(config.SourceRoot))
	http.HandleFunc("/", controllers.FileServer(config.StaticRoot))
	http.ListenAndServe(":3000", nil)
}
