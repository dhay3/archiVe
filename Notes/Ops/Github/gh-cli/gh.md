# gh

ref
[https://cli.github.com/](https://cli.github.com/)
[https://cli.github.com/manual/](https://cli.github.com/manual/)

## Digest
gh 是 github 推出的一款面向 terminal 的工具
## Installation
[https://github.com/cli/cli#installation](https://github.com/cli/cli#installation)
## Configuration
在使用 gh 前还需要配置
```
λ master gh auth login
? What account do you want to log into? GitHub.com
? You're already logged into github.com. Do you want to re-authenticate? Yes
? What is your preferred protocol for Git operations? HTTPS
? How would you like to authenticate GitHub CLI? Paste an authentication token
Tip: you can generate a Personal Access Token here https://github.com/settings/tokens
The minimum required scopes are 'repo', 'read:org', 'workflow'.
? Paste your authentication token: 
```
创建管理 authentication token 具体位置在 settings -> developer settings -> personal access tokens，这里推荐使用 Tokens(classic) 的方式来创建 token
token 必须含有 `repo`，`read:org`，`workflow` 这几个权限。如果还需要删除的权限可以按照错误提示另外加
## Completion
命令补全参考
[https://cli.github.com/manual/gh_completion](https://cli.github.com/manual/gh_completion)
## Subcommands
gh 和 git 一样由几个 subcommands 组成，大部分命令都提供了 interactively mode 会自动 prompt

- `help`

和 `git help` 一样，用于查看帮助信息

- `status`

- `alias`

和 `git config --alias` 类似用于创建别名

- `api`
- `auth`

为 git 和 gh 授权使用 github

- `browser`

使用 browser 执行操作

- `config`

和 `git config` 一样用于配置 gh 的行为

- `extensions`

用于扩展 gh

- `gist`

操作 github gist

- `gpg-key`

操作管理 gpg

- `issue`

操作管理 github issues

- `pr`

操作管理 pull request

- `release`

操作管理 release

- `repo`

操作管理 repo

