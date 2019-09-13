package sql

import (
	"github.com/arrebole/culaccino/service/model"
)

// SetStorage 设置storage
func (p *SQL) SetStorage(arg *model.Storage) error {
	return p.StorageDB.HMSet(arg.Name, adapter(arg)).Err()
}

// GetStorage 通过名称获取storage的信息
func (p *SQL) GetStorage(query string) *model.Storage {
	var result = &model.Storage{}
	reflectMapToStruct(result, p.StorageDB.HGetAll(query).Val())
	return result
}

// ExistsStorage 判断storage是否存在
func (p *SQL) ExistsStorage(query string) bool {
	if p.StorageDB.Exists(query).Val() > 0{
		return true
	}
	return false
}

