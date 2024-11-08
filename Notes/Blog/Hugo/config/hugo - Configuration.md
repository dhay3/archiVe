# hugo - Confiuration


也被称为 site configuration
除了使用单一的配置文件 `config.toml`，还可以使用 `configDir`( 默认 `config/` ) 来存储所有的配置文件。在 `configDir` 中的每一个配置文件，可以通过文件名来标识不同模块。例如 `languages.toml` 配置就对应语言的模块
配置文件同时支持 3 种格式，`yaml`，`toml` 以及 `json`
假设当前目录如下

```
├── config
│   ├── _default
│   │   ├── config.toml
│   │   ├── languages.toml
│   │   ├── menus.en.toml
│   │   ├── menus.zh.toml
│   │   └── params.toml
│   ├── production
│   │   ├── config.toml
│   │   └── params.toml
│   └── staging
│       ├── config.toml
│       └── params.toml
```
当使用 `hugo --enviroment staging` 时，会将 `config/_default` 已 `config/staging` 种的内容合并。可以将 `config/_default` 理解成类似 springboot 中的 application.yaml 即全局变量，对所有的环境生效。而 `config/production` 和 `config/staging` 为 production.yaml 和 staging.yaml，只对特定环境的生效
假设你需要在 production 和 staging 环境中使用 Google Analytics，当时不想在 development 环境中使用 Google Analytics。那么你需要

1. 无需在 `_default/config.toml` 中设置 Google Analytics 的配置，分别在 `config/production` 和 `config/staging` 中设置 Google Analytics 的配置
2. 如果需要使用 Google Analytics，那么只需要在启动 `hugo server` 时使用 `--environment production | staging` 来指定使用的环境，会同时加载 `config/_default` 和 `config/production | staging` 中的配置

这里还需要注意的一点是 `hugo server` 默认使用的环境的是 development，`hugo` 默认使用的是 production
## Common Directives

> 可以通过 `.Site` 的方式来获取对应 directives 的值

有如下的 directives

- baseURL

站点使用的 baseURL，例如 `https://bep.is`

- buildDrafts

构建的时候会包含 Drafts，默认 false

- buildExpired

构建的时候会包含已经过期的内容，默认 false

- contentDir

存储站点文章的目录，默认 content

- copyright

footer 部分显示的 copyright，默认为空

- defaultContentLanguage

默认显示的语言信息，默认 en

- disableLiveReload

关闭热部署功能，默认 false

- enableEmoji

开启 emoji 功能，默认 false

- enableGitInfo

开启 `.GitInfo`, 文章的内容和 git commit 的时间相同，默认 false

- enableInlineShortcodes

开启 inline shortcode 功能，默认 false

- enableRobotsTxT

允许生成 robots.txt 文件，默认 false

- googleAnalytics

google analytics tacking ID

- hasCJKLanuage

自动检测文档中是否包含，Chinese/Japanese/Korean，如果文档中有对应内容，需要开启该功能。默认 false

- noChmod

文件修改权限时不同步，默认 false

- noTImes

文件修改时不同步，默认 false

- paginnate

每页的文章数，默认 10

- publishDir

存储生成静态页面的目录

- theme

站点使用的主题

- themesDir

储存站点主题的目录

- timeZone

设置使用的 timeZone

- title

设置站点的名字

- permanlnks
## Custom Directives



**reference**

1. [https://gohugo.io/getting-started/configuration/](https://gohugo.io/getting-started/configuration/)
