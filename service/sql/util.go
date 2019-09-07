package sql

import (
	"encoding/json"
	"os"

	"github.com/arrebole/culaccino/service/share"
)

// 返回两个数中的最大值
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// 获取环境变量中的默认密码
func defaultPassWord() string {
	passwd := os.Getenv("PASSWORD")
	if passwd == "" {
		passwd = "root@culaccino"
	}
	return PassWord(passwd)
}

// PassWord 两次hash计算产生密码
func PassWord(passwd string) string {
	return share.HashHexStr(share.HashHexStr(passwd))
}

// adapter 将结构体转化为 map
func adapter(arg interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	j, _ := json.Marshal(arg)
	json.Unmarshal(j, &result)
	return result
}
