references

1. [https://gohugo.io/content-management/organization/](https://gohugo.io/content-management/organization/)
2. [https://gohugo.io/templates/section-templates/](https://gohugo.io/templates/section-templates/)
3. [https://gohugo.io/content-management/urls/](https://gohugo.io/content-management/urls/)

hugo 通常会按照 `content` 目录下的阶级生成对应的 URL，例如 `content/posts/post-1.md` 会生成 `https://example.org/posts/post-1/`
## Example
假设 `content` 目录下的内容如下
```
.
└── content
    └── about
    |   └── index.md  // <- https://example.com/about/
    ├── posts
    |   ├── firstpost.md   // <- https://example.com/posts/firstpost/
    |   ├── happy
    |   |   └── ness.md  // <- https://example.com/posts/happy/ness/
    |   └── secondpost.md  // <- https://example.com/posts/secondpost/
    └── quote
        ├── first.md       // <- https://example.com/quote/first/
        └── second.md      // <- https://example.com/quote/second/
```
如果没有配置任何其他关联的参数，想要访问对应的内容，URL 映射关系如上
## Path breakdown
假设在配置文件配置了 `baseURL='http://example.com'`，那么在 hugo 中的相关的路径解析如下
### List pages
原文档中并没有 list pages 对应的内容，这里为了方便理解才添加
list pages 可以将其想象成`content` 下的所有目录( `content`自己本身也是 )，会使用 list templates 来渲染
```
.         url
.       ⊢--^-⊣
.        path 
.       ⊢--^-⊣
.        filepath
.       ⊢-^-⊣
content/posts/
```
那么对应的解析如下
```
                     url ("/posts/")
                    ⊢-^-⊣
       baseurl      section ("posts")
⊢--------^---------⊣⊢-^-⊣
        permalink
⊢----------^-------------⊣
https://example.com/posts/
```
section 对应 `content` 下的目录
### Single pages
`content` 目录下的文章，会被 single template 渲染
例如，有一个 `my-first-hugo-post.md` 文件
```
                   path ("posts/my-first-hugo-post.md")
.       ⊢-----------^------------⊣
.      section        slug
.       ⊢-^-⊣⊢--------^----------⊣
content/posts/my-first-hugo-post.md
```
hugo 将会将其使用 single template 渲染，对应的 URL 映射如下
```
                               url ("/posts/my-first-hugo-post/")
                   ⊢------------^----------⊣
       baseurl     section     slug
⊢--------^--------⊣⊢-^--⊣⊢-------^---------⊣
                 permalink
⊢--------------------^---------------------⊣
https://example.com/posts/my-first-hugo-post/index.html
```
实际我们可以直接使用 `https://example.com/posts/my-first-hugo-post`
### Index pages
`_index.md` 是一个特殊的文件，list template 可以获取 `_index.md` 中的内容，为对应渲染的目录添加 `_index.md` 中的 front matter 或者其他内容。`content` 下的所有目录都可以有一个 `_index.md` ( 包含 `content` 目录本身 )
例如 `content`下有如下内容
```
.         url
.       ⊢--^-⊣
.        path    slug
.       ⊢--^-⊣⊢---^---⊣
.           filepath
.       ⊢------^------⊣
content/posts/_index.md
```
那么对应的路径解析如下
```
                     url ("/posts/")
                    ⊢-^-⊣
       baseurl      section ("posts")
⊢--------^---------⊣⊢-^-⊣
        permalink
⊢----------^-------------⊣
https://example.com/posts/index.html
```
实际我们可以通过 `https://example.com/posts` 获取到
## Front matter
我们可以通过 front matter 来修改对应 path breakdown 对应的部分
### slug
例如 `content/posts/post-1.md` 中的 front matter 如下
```
---
slug: my-first-post
title: My First Post
---
```
如果没有设置 slug，URL 是 `exmaple.com/posts/post-1`。设置如上 slug 后，URL 为 `exmaple.com/posts/my-first-post`
### url
例如 `content/posts/post-1.md` 中的 front matter 如下
```
---
title: My First Article
url: /articles/my-first-article
---
```
如果没有设置 url，URL 是 `exmaple.com/posts/post-1`。设置如上 url 后，URL 为 `exmaple.com/articles/my-first-article`

如果同时设置了 url 和 slug，url 优先级更高
