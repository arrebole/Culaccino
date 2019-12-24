package controllers

import (
	"github.com/arrebole/culaccino/service/config"
	"github.com/arrebole/culaccino/service/model"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// SourceCtl ...
func SourceCtl(w http.ResponseWriter, req *http.Request) {
	symbol := req.URL.Query().Get("symbol")
	if symbol == "" {
		w.WriteHeader(400)
		return
	}

	// 读取文件内容
	file, err := os.Open(filepath.Join(config.SourceRootPath, symbol+".md"))
	if err != nil {
		w.WriteHeader(400)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.Write(model.RespDirector(&model.SourceBuilder{}, data))
}
