package sql

import (
	"encoding/json"
	"os"
	"reflect"
	"strconv"

	"github.com/arrebole/culaccino/service/share"
	"github.com/go-redis/redis"
)

// max 返回两个数中的最大值
func max(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// min 返回俩个数的最小值
func min(a int64, b int64) int64 {
	if a < b {
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

// reflectMapToStruct 将map赋值给结构体
func reflectMapToStruct(out interface{}, data map[string]string) {
	v, t := reflect.ValueOf(out), reflect.TypeOf(out)
	for i := 0; i < t.Elem().NumField(); i++ {
		// 获取结构体json化后的字段名
		tag := t.Elem().Field(i).Tag.Get("json")

		switch v.Elem().Field(i).Kind() {
		case reflect.Int:
			n, _ := strconv.ParseInt(data[tag], 10, 64)
			v.Elem().Field(i).SetInt(n)
		case reflect.String:
			v.Elem().Field(i).SetString(data[tag])
		}

	}
}

// GetPipeline use hash
// 用于批量读取redis hash类型数据
func getHashPipeline(db *redis.Client, arg ...string) []*redis.StringStringMapCmd {
	cmds := make([]*redis.StringStringMapCmd, len(arg))
	db.Pipelined(func(pipe redis.Pipeliner) error {
		for i := 0; i < len(arg); i++ {
			cmds[i] = pipe.HGetAll(arg[i])
		}
		pipe.Exec()
		return nil
	})
	return cmds
}


// 增加映射数据库
func addmap(db *redis.Client, src string, dist string) error {
	return db.ZAdd(src, redis.Z{
		Score:  0,
		Member: dist,
	}).Err()
}