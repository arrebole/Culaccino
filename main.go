package main

import (
	"net/http"

	"github.com/arrebole/culaccino/controller"
)

func main() {
	http.HandleFunc("/api/new", controller.New())
	http.HandleFunc("/api/update", controller.Update())
	http.HandleFunc("/api/delete", controller.Del())
	http.HandleFunc("/api/paper", controller.Get())
	http.HandleFunc("/api/explore", controller.Explore())
	http.HandleFunc("/", controller.FileServer())
	http.ListenAndServe("127.0.0.1:3000", nil)
}
