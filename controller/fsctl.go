package controller

import (
	"net/http"
	"os"
	"path/filepath"
)

// FileServer ...
func FileServer() http.HandlerFunc {
	var StaticRoot, _ = filepath.Abs("themes/build")

	return func(w http.ResponseWriter, r *http.Request) {
		if fileExist(filepath.Join(StaticRoot, r.RequestURI)) {
			http.FileServer(http.Dir(StaticRoot)).ServeHTTP(w, r)
			return
		}
		http.ServeFile(w, r, filepath.Join(StaticRoot, "index.html"))
	}
}

// fileExist 判断文件是否存在
func fileExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
