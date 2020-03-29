package transfer

import (
	"net/http"

	"github.com/arrebole/culaccino/model"
)

// Send 发送
func Send(w http.ResponseWriter, resp model.IResponse) {
	w.WriteHeader(resp.Code())
	w.Write(resp.Data())
}
