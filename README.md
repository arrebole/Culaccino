# simple blog

version: 0.2.0 (开发中)

### This use

+ gin
+ gorm + sqlite3
+ vue + scss + typescript + github-markdown-css
+ github markdown Render api

```bash
chmod +x ./build.sh && ./build.sh
chmod +x ./start.h && ./start.h
```

### API

| method | url                     | 功能          |
| ------ | -----------------       | -----------  |
| GET    | /api/count              | 获取文章总数   |
| GET    | /api/table/:page        | 获取索引      |
| GET    | /api/contents/:id       | 获取详细内容  |
| GET    | /api/login              | 登录         |
| GET    | /api/admin/table/all    | 获取所有索引  |
| POST   | /api/admin/add/text     | 添加文章     |
| PUT    | /api/admin/update/:id   | 更新文章     |
| DELETE | /api/admin/delete/:id   | 删除文章     |

