package controllers

import (
	"net/http"
	"os"
	"path/filepath"
)

// fileExist 判断文件是否存在
func fileExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// FileServer ...
func FileServer(root string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if fileExist(filepath.Join(root, req.RequestURI)) {
			http.FileServer(http.Dir(root)).ServeHTTP(w, req)
			return
		}
		http.ServeFile(w, req, filepath.Join(root, "index.html"))
	}
}
