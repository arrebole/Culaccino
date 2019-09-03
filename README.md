# blog

> version: 0.8.0 (开发中)
预览: https://gkdark.xyz

TODO:
+ 评论功能
+ 多用户注册管理

### This use

+ gin + gorm + sqlite3
+ vue + scss + typescript
+ github-markdown-css + marked + highlight.js

```bash
docker build -t arrebole/culaccino .
docker run  -p 8080:80 -d arrebole/culaccino
```

### API

| method | url                     | 功能          |
| ------ | -----------------       | -----------  |
| GET    | /api/archive/dashboard/:page        | 获取索引      |
| GET    | /api/archives/details/:id       | 获取详细内容  |
| GET    | /api/session/login              | 登录         |
| GET    | /api/archive/all        | 获取所有索引  |
| POST   | /api/archive/create          | 添加文章     |
| PUT    | /api/archive/update/:id   | 更新文章     |
| DELETE | /api/archive/delete/:id   | 删除文章     |

