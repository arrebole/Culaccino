package controllers

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/arrebole/culaccino/service/model"
)

// DirectoryCtl ...
func DirectoryCtl(directoryfilePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// 读取文件内容
		file, err := os.Open(directoryfilePath)
		if err != nil {
			panic("丢失目录文件" + directoryfilePath)
		}
		defer file.Close()

		data, err := ioutil.ReadAll(file)
		if err != nil {
			panic("读取文件失败" + directoryfilePath)
		}

		w.Write(model.RespDirector(&model.DirectoryBuilder{}, data))
	}
}
