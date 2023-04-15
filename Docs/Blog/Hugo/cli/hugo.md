refernce

1. [https://gohugo.io/getting-started/usage/](https://gohugo.io/getting-started/usage/)
2. [https://gohugo.io/commands/hugo/](https://gohugo.io/commands/hugo/)
3. [https://discourse.gohugo.io/t/whats-the-difference-between-hugo-and-hugo-server/4742](https://discourse.gohugo.io/t/whats-the-difference-between-hugo-and-hugo-server/4742)
## Digest
syntax
```
hugo [flags]
```
`hugo` 用于编译生成站点的静态内容，存储在 `public` 目录下
## Optional args

- `-b | --baseURL string`

指定站点的 baseURL

- `-D | --buildDrafts`

drafts 也会被编译

- `-E | --buildExpired`

expired 的内容也会被编译

- `-c | --contentDir`

存储站点文章的目录

- `--config string`

指定使用的配置文件路径

- `--configDir string`，默认 `config.yaml | json | toml`

指定使用的配置文件目录，默认 `config`

- `-d | --destination string`

指定生成静态文件的目录，默认使用 `public` 目录

- `--enableGitInfo`

编译时还会生成 GIt 相关的信息

- `--noChmod`

修改文件后，不同步文件的权限

- `--noTimes`

修改文件后，不同步文件的时间

- `--minify`

以 minify 的格式输出对应的 HTML, XML 等

- `--templateMetrics`

生成静态内容时，显示和 template 相关的 metrics，包括耗时

- `-t | --theme strings`

指定使用的 theme 的路径

- `--debug`
- `--log`

记录日志

- `--logFile string`

存储日志的地址，默认使用 `--log`

- `--verbose`
- `-w | --watch`

监控模式，显示文件的变更
## hugo vs hugo server
`hugo` 用于生成站点的静态内容，存储在 `public` 目录下。而 `hugo server` 并不会直接生成静态内容到 `public` 目录下，而是会直接运行一个 webserve
所以我们可以使用 `hugo` 来生成静态内容，然后可以使用类似 nginx 的 webserver 来做转发 ，如果使用 `hugo server` 我们可以省去大量的配置
