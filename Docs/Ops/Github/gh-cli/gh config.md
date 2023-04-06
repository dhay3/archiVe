ref
[https://cli.github.com/manual/gh_config](https://cli.github.com/manual/gh_config)
和 git 一样，gh 也有一个工具用于配置，目前允许的参数有

- git_protocol: the protocol to use for git clone and push operations (default: "https")
- editor: the text editor program to use for authoring text
- prompt: toggle interactive prompting in the terminal (default: "enabled")
- pager: the terminal pager program to send standard output to
- http_unix_socket: the path to a Unix socket through which to make an HTTP connection
- browser: the web browser to use for opening URLs
## List
用于查看当前的配置
```
λ ~/ gh config list
git_protocol=https
editor=
prompt=enabled
pager=
http_unix_socket=
browser=
```
## Set
用于设置指定配置
syntax
```
gh config set <key> <value> [flags]
```
例如
```
$ gh config set editor vim
$ gh config set editor "code --wait"
$ gh config set git_protocol ssh --host github.com
$ gh config set prompt disabled
```
## Get
用于获取值
```
$ gh config get git_protocol
https
```
