package main

import (
	"net/http"

	"github.com/arrebole/culaccino/src/controller"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	// articles
	router.HandlerFunc(http.MethodGet, "/api/articles", controller.ListArticles)
	router.HandlerFunc(http.MethodPost, "/api/articles", controller.CreateArticle)

	// article
	router.HandlerFunc(http.MethodGet, "/api/articles/:article", controller.GetArticle)
	router.HandlerFunc(http.MethodPatch, "/api/articles/:article", controller.UpdateArticle)

	router.HandlerFunc(http.MethodPost, "/api/articles/:article", controller.CreateSection)

	// section
	router.HandlerFunc(http.MethodGet, "/api/articles/:article/:section", controller.GetSection)
	router.HandlerFunc(http.MethodPatch, "/api/articles/:article/:section", controller.UpdateSection)

	// Serve static files
	router.NotFound = http.FileServer(http.Dir("theme/build"))

	http.ListenAndServe("127.0.0.1:3000", router)
}
