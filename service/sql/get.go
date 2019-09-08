package sql

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/arrebole/culaccino/service/module"
)

// Get 查询多个
func Get(out interface{}, query ...string) error {
	switch out.(type) {

	case *module.Storage:
		return getStorage(out, query...)

	case *module.Repo:
		return getRepo(out, query...)

	case *module.Chapter:
		return getChapter(out, query...)

	default:
		return errors.New("")
	}
}

func getStorage(arg interface{}, query ...string) error {
	reflectSet(arg, StorageDB.HGetAll(query[0]).Val())
	return nil
}

func getRepo(arg interface{}, query ...string) error {
	reflectSet(arg, RepoDB.HGetAll(query[0]+"/"+query[1]).Val())
	return errors.New("")
}

func getChapter(arg interface{}, query ...string) error {
	reflectSet(arg, ChapterDB.HGetAll(query[0]+"/"+query[1]+"/"+query[2]).Val())
	return errors.New("")
}

func reflectSet(arg interface{}, data map[string]string) {
	v, t := reflect.ValueOf(arg), reflect.TypeOf(arg)
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
