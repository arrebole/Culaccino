package main

import (
	"net/http"

	"github.com/arrebole/culaccino/controller"
)

func main() {
	http.HandleFunc("/", controller.FileServer())
	http.HandleFunc("/api/papers", controller.Papers)
	http.HandleFunc("/api/papers/", controller.Paper)
	http.ListenAndServe("127.0.0.1:3000", nil)
}
