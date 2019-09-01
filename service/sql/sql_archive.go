package sql

import "github.com/arrebole/culaccino/service/module"

// Query 查询
func (p *client) ArchiveQuery(id uint) *module.Archive {
	var result = &module.Archive{}
	p.db.First(result, id)
	p.increaseAccess(result)
	return result
}

// // QueryDir 查询目录
func (p *client) ArchiveQueryDir(page int, per int) ([]module.Archive, *module.Count) {
	var dir []module.Archive
	var searchItems = "id, title, author, target, cover, summary, views, updated_at, created_at"
	p.db.Limit(per).Offset(page * per).Order("id desc").Select(searchItems).Find(&dir)
	return dir, p.remain(page, per)
}

// Count 查询文章总量和剩余
func (p *client) ArchiveCount() *module.Count {
	return p.remain(0, 0)
}

// 在QueryDir中用于计算剩余的数目 per 每页显示数量
func (p *client) remain(curr int, per int) *module.Count {
	var total int
	p.db.Model(&module.Archive{}).Count(&total)
	return &module.Count{
		Total:  total,
		Remain: max(total-(curr+1)*per, 0),
	}
}

// Delete 删除文章
func (p *client) ArchiveDelete(id uint) {
	p.db.Where("id = ?", id).Delete(&module.Archive{})
}

// Add 添加文章
func (p *client) ArchiveCreate(Archive *module.PostArchive) {
	aNewArchive := module.ToArchive(Archive)
	p.db.Create(aNewArchive)
}

// Update修改文章,只更新修改的字段
func (p *client) ArchiveUpdate(id uint, Archive *module.Archive) {
	p.db.Model(&module.Archive{}).Where("id = ?", id).Updates(Archive)
}

// 增加文章浏览量
func (p *client) increaseAccess(Archive *module.Archive) {
	if Archive.Title != "" {
		Archive.Views++
		p.db.Model(Archive).UpdateColumn("views", Archive.Views)
	}
}