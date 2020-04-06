package main

import (
	"net/http"

	"github.com/arrebole/culaccino/src/controller"
)

func main() {
	http.HandleFunc("/", controller.FileServer())
	http.HandleFunc("/api/papers", controller.Articles)
	http.HandleFunc("/api/papers/", controller.Articles)
	http.ListenAndServe("127.0.0.1:3000", nil)
}
