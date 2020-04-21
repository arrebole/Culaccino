package controller

import (
	"net/http"
	"os"
	"path/filepath"
)

// ServerFile ...
type ServerFile struct {
	root string
}

func (p ServerFile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if p.fileExist(filepath.Join(p.root, r.RequestURI)) {
		http.FileServer(http.Dir(p.root)).ServeHTTP(w, r)
		return
	}
	http.ServeFile(w, r, filepath.Join(p.root, "index.html"))
}

// fileExist 判断文件是否存在
func (p ServerFile) fileExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// CreateServerFiles ...
func CreateServerFiles(path string) http.Handler {
	var staticRoot, _ = filepath.Abs(path)
	return ServerFile{root: staticRoot}
}
