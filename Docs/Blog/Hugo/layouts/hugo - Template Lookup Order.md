# hugo - Template Lookup Order

当 hugo 渲染页面时，会按照特定的顺序选择 layouts 中的 template 作为渲染的模板。具体查看官方文档
## Regular pages
对应 `content` 目录下的文章
例如 `content/posts` 中的文章就会按照如下的顺序选择 template

1. layouts/posts/single.html.html
2. layouts/posts/single.html
3. layouts/_default/single.html.html
4. layouts/_default/single.html
## Home page
对应 homepage，会按照如下的顺序选择 template

1. layouts/index.html.html
2. layouts/home.html.html
3. layouts/list.html.html
4. layouts/index.html
5. layouts/home.html
6. layouts/list.html
7. layouts/_default/index.html.html
8. layouts/_default/home.html.html
9. layouts/_default/list.html.html
10. layouts/_default/index.html
11. layouts/_default/home.html
12. layouts/_default/list.html
## Section pages
对应 `content`目录下目录
例如 `content/posts` section，会按照如下的顺序选择 template

1. layouts/posts/posts.html.html
2. layouts/posts/section.html.html
3. layouts/posts/list.html.html
4. layouts/posts/posts.html
5. layouts/posts/section.html
6. layouts/posts/list.html
7. layouts/section/posts.html.html
8. layouts/section/section.html.html
9. layouts/section/list.html.html
10. layouts/section/posts.html
11. layouts/section/section.html
12. layouts/section/list.html
13. layouts/_default/posts.html.html
14. layouts/_default/section.html.html
15. layouts/_default/list.html.html
16. layouts/_default/posts.html
17. layouts/_default/section.html
18. layouts/_default/list.html



**references**

1. [https://gohugo.io/templates/lookup-order/](https://gohugo.io/templates/lookup-order/)
