# CoolBlog API 文档

## /articles
- GET 获取文章列表
```
[
    {
        "id": 2017102701,
        "title": "文章标题",
        "url": "http://coolblog.top/articles/2017102701"
    },
    ...
]
```
- POST 新建一篇文章
```
{
    "title": "文章标题"
}
```
返回值
```
{
    "id": 2017102701,
    "title": "文章标题",
    "url": "http://coolblog.top/articles/2017102701"
}
```

## /articles/ID
- GET 获取单篇文章内容
```
{
    "content": "<h1>这是文章内容</h1>"
}
```
- PUT 修改文章内容
```
{
    "content": "<h1>这是修改后文章内容</h1>" 
}
```