package controllers

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/arrebole/culaccino/service/config"
	"github.com/arrebole/culaccino/service/model"
)

// DirectoryCtl ...
func DirectoryCtl(w http.ResponseWriter, req *http.Request) {
	// 读取文件内容
	file, err := os.Open(config.DirectoryfilePath)
	if err != nil {
		panic("丢失目录文件" + config.DirectoryfilePath)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic("读取文件失败" + config.DirectoryfilePath)
	}

	w.Write(model.RespDirector(&model.DirectoryBuilder{}, data))
}
