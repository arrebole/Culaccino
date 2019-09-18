# blog

> version: 0.9.0 (开发中)
预览: https://gkdark.xyz

TODO:
+ 美化编辑页面 ⌛
+ 评论功能 ⌛
+ 多用户注册管理 ⌛
+ sqlite迁移至redis 👍
+ 将文章升级成仓库 👍

### This use

+ gin + go-redis + redis
+ react + scss + typescript
+ github-markdown-css + marked + highlight.js

```bash
docker build -t arrebole/culaccino:latest .
docker run  -p 8080:80 -d arrebole/culaccino
```

### API

| method | url                     | 功能          |
| ------ | -----------------       | -----------  |
| POST   | /api/new                 | 创建内容|
| GET    | /api/export              | 获取内容      |
| PUT    | /api/commit               | 修改内容 |
| DELETE | /api/delete              | 删除内容|

