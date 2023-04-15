# hugo - Front matter
reference

1. [https://gohugo.io/content-management/front-matter/](https://gohugo.io/content-management/front-matter/)

Front matter 是描述 content file 的元数据，可以被 template 通过 go template 的语法调用，支持如下几种格式

- TOML
```
+++
title = "{{ replace .Name "-" " " | title }}"
date = {{ .Date }}
draft = true
+++
```

- YAML
```
---
title: "{{ replace .Name "-" " " | title }}"
date: {{ .Date }}
draft: true
---
```

- JSON

不推荐

- ORG

不推荐
## Predefined variables
只记录常用的 variables

- title

文章的 title

- date

标识文章的时间

- description

标识文章的摘要

- draft

标识文章是否是 draft，只有当使用 `-D | --buildDrafts` 时才会渲染在站点

- expiryDate

标识文章的过期时间，如果超过指定时间，需要使用 `-E | --buldExpired` 时才会渲染在站点

- images

标识文章中的图片使用的目录

- keywords

标识文章的关键字

- 

## User-defined variables
也可以在 front matter 中添加用户自定义的变量，这样就可以在 template 中调用对应的自定义变量
例如，在 `content/post1.md`添加了一个自定义变量 `myvar`
```
myvar: "this var is defined by me"
```
添加的这个变量会被作为 dict 的一对键值对，存储在 `.Params` 中，这样我们就可以通过 `.Params` 在 templates 中的调用对应的自定义变量
例如，在 `layouts/single.html`中有如下一段内容
```
<h1>{{.Params.myvar}}</h1>
```
那么最后在 hugo 的渲染下会变成 `<h1>this var is defined by me</h1>`
