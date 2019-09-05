package sql

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/session"
)

var (
	dashboardKeys      = "id, title, author, tag, area, cover, summary, views, updated_at, created_at"
	touristDetailskeys = "id,author_id, title, author, tag, area, cover, summary,contents, views, updated_at, created_at"
)

// TODO 分页处理
// GetRepos通过作者id查找
func (p *client) GetRepos(domain string) ([]module.Archive, *module.Count) {
	var result []module.Archive

	p.db.Select(touristDetailskeys).Where("author = ?", domain).Find(&result)
	return result, &module.Count{
		Total:  len(result),
		Remain: 0,
	}
}

//  Dashboard 查询目录
func (p *client) Dashboard(page int, per int) ([]module.Archive, *module.Count) {
	var dir []module.Archive
	p.db.Limit(per).Offset(page * per).Order("id desc").Select(dashboardKeys).Find(&dir)
	return dir, p.remain(page, per)
}

// ArchiveCount 查询文章总量和剩余
func (p *client) CountRepos() *module.Count {
	return p.remain(0, 0)
}

// remain 中用于计算剩余的数目 per 每页显示数量
func (p *client) remain(curr int, per int) *module.Count {
	var total int
	p.db.Model(&module.Archive{}).Count(&total)
	return &module.Count{
		Total:  total,
		Remain: max(total-(curr+1)*per, 0),
	}
}

// NewRepo 添加文章
func (p *client) NewRepo(Archive *module.PostArchive, session *session.User) {
	aArchive := module.ToArchive(Archive)
	aArchive.Author = session.Uname
	aArchive.AuthorID = session.UID
	p.db.Create(aArchive)
}

// CommitRepo 修改文章,只更新修改的字段
func (p *client) CommitRepo(domain string, symbol string, repo *module.Archive) bool {
	p.db.Model(&module.Archive{}).Where("author = ? AND title= ?", domain, symbol).Updates(repo)
	return true
}

// 增加文章浏览量
func (p *client) increaseAccess(Archive *module.Archive) {
	if Archive.Title != "" {
		Archive.Views++
		p.db.Model(Archive).UpdateColumn("views", Archive.Views)
	}
}

// GetRepo ...
func (p *client) GetRepo(domain string, symbol string) *module.Archive {
	var result = &module.Archive{}
	p.db.Where("author = ? AND title = ?", domain, symbol).Select(touristDetailskeys).First(result)
	return result
}

// DelRepo ...
func (p *client) DelRepo(domain string, symbol string) bool {
	p.db.Where("author = ? AND title = ?", domain, symbol).Delete(&module.Archive{})
	return true
}
