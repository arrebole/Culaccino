package main

import (
	"net/http"

	"github.com/arrebole/culaccino/src/controller"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	// auth
	router.HandlerFunc(http.MethodPost, "/api/auth/token", controller.CreateToken)
	router.HandlerFunc(http.MethodGet, "/api/auth/token/check", controller.CheckToken)

	// articles
	router.HandlerFunc(http.MethodGet, "/api/articles", controller.ListArticles)
	router.HandlerFunc(http.MethodPost, "/api/articles", controller.CreateArticle)

	// article
	router.HandlerFunc(http.MethodGet, "/api/articles/:article", controller.GetArticle)
	router.HandlerFunc(http.MethodPatch, "/api/articles/:article", controller.UpdateArticle)
	router.HandlerFunc(http.MethodDelete, "/api/articles/:article", controller.DeleteArticle)

	// Serve static files
	router.NotFound = controller.CreateServerFiles("./theme/build")

	http.ListenAndServe("0.0.0.0:80", router)
}
