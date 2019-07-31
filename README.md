# simple blog
开发中

### This use

+ gin
+ gorm + sqlite3
+ vue+ scss + typescript

### API

| method | url                     | 功能          |
| ------ | -----------------       | -----------  |
| GET    | /api/table/:page        | 获取索引      |
| GET    | /api/contents/:id       | 获取详细内容  |
| GET    | /api/login              | 登录         |
| GET    | /api/admin/table/all    | 获取所有索引  |
| POST   | /api/admin/add/text     | 添加文章     |
| PUT    | /api/admin/update/:id   | 更新文章     |
| DELETE | /api/admin/delete/:id   | 删除文章     |

