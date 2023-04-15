# hugo - Archetypes
refernece

1. [https://gohugo.io/content-management/archetypes/](https://gohugo.io/content-management/archetypes/)

Archetypes 是 content 的 template，即当使用 `hugo new` 创建站点的文章时，会自动使用该目录中对应的 template
template 中包含预设的 front matter ( 和 content 中的 front matter 不同，template 中的 front matter 可以使用 hugo variables/functions )以及其他的一些内容。如果当前 `archetypes` 中没有对应的 template，且配置了 template ，默认也会有查看对应 themes 中的 `archetypes`目录
例如，在配置文件中配置了 `mytheme` , 并且使用了如下命令
```
hugo new posts/mypost.md
```
会在 `content` 中创建一个 `posts` 目录，以及目录的 `post1.md`，那么就会按照如下顺序使用 archetypes

1. `archetypes/posts.md`
2. `archetypes/default.md`
3. `themes/mytheme/archetypes/posts.md`
4. `themes/mytheme/archetypes/default.md`
## Create a New Archetype Template
例如在 `archetypes` 中创建一个 `newsletter.md`，内容如下
```
---
title: "{{ replace .Name "-" " " | title }}"
date: {{ .Date }}
draft: true
---

**Insert Lead paragraph here.**

## New Cool Posts

{{ range first 10 ( where .Site.RegularPages "Type" "cool" ) }}
* {{ .Title }}
{{ end }}
```
使用如下命令创建 `newsletter/the-last-cool.stuff.md`	
```
hugo new newsletter/the-last-cool.stuff.md
```
## Directory base archetypes
在 `0.49` 版本后，hugo 还可以使用 direcotory archetypes。当创建站点的目录时，会自动将 direcotory archetypes 中的内容加入到对应目录
例如：archetypes 中的目录如下
```
archetypes
├── default.md
└── post-bundle
    ├── bio.md
    └── index.md
```
使用如下命令
```
hugo new --kind post-bundle posts/my-post
```
会在 `content` 目录下创建 `post/my-post` 目录，同时会创建 `archetypes/post-bundle` 目录下的内容
```
content
└── posts
    └── my-post
        ├── bio.md
        └── index.md
```
