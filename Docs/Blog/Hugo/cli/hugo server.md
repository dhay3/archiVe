reference

1.  [https://gohugo.io/commands/hugo_server/](https://gohugo.io/commands/hugo_server/)
## Digest
syntax
```
hugo server [flags]
```
## Optional args
只列出常用选项，具体参考 help 信息

- `--bind string`

指定绑定的 interface，默认绑定 127.0.0.1

- `-p | --port int`

指定监听的端口，默认 1313

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
## Exmaples
```
hugo server -p 80 --bind 0.0.0.0
```
