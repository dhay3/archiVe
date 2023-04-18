# hugo - Directory Structure

当使用 `hugo new site example` 时，会创建如下的目录结构
```
example/
├── archetypes/
│   └── default.md
├── assets/
├── content/
├── data/
├── layouts/
├── public/
├── static/
├── themes/
└── config.toml
```
有主要几个目录

- archetypes

存储站点文章使用的 template，当使用 `hugo new` 创建文章时会自动使用该目录中的 template

- content

存储站点的文章

- config

存储站点配置

- data

存储生成站点时使用的数据，类似一个小型的数据库

- layouts

存储站点使用的自定义 template

- public

存储当使用 `hugo` 时会生成的静态页面

- static

存储站点的静态内容，例如图片, CSS, Javascript。在文章或者 template 中可以通过 `/` 来直接调用 static 目录下内容

- themes

存储站点使用的主题 template



**references**

1. [https://gohugo.io/getting-started/directory-structure/](https://gohugo.io/getting-started/directory-structure/)
