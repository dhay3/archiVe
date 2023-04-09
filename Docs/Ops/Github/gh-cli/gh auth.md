# gh auth

ref
[https://cli.github.com/manual/gh_auth](https://cli.github.com/manual/gh_auth)
用于 gh 和 git 认证的一系列子命令

## Status
syntax
```
gh auth status [flags]
```
校验当前账户开启权限
```
λ ~/ gh auth status
github.com
  ✓ Logged in to github.com as dhay3 (/root/.config/gh/hosts.yml)
  ✓ Git operations for github.com configured to use https protocol.
  ✓ Token: ghp_************************************
  ✓ Token scopes: read:org, repo, workflow
```
### Optional args

- `-t | --show-token`

显示明文的 auth token
## Login
syntax
```
gh auth login [flags]
```
登录 github，使用该命令会输出 prompt 让用户自己选择可以无需使用 optional args
```
λ ~/ gh auth login
? What account do you want to log into? GitHub.com
? What is your preferred protocol for Git operations? HTTPS
? How would you like to authenticate GitHub CLI? Paste an authentication token
Tip: you can generate a Personal Access Token here https://github.com/settings/tokens
The minimum required scopes are 'repo', 'read:org', 'workflow'.
? Paste your authentication token: ****************************************
- gh config set -h github.com git_protocol https
✓ Configured git protocol
✓ Logged in as dhay3
```
### Exmaples
```
# start interactive setup
$ gh auth login

# authenticate against github.com by reading the token from a file
$ gh auth login --with-token < mytoken.txt

# authenticate with a specific GitHub instance
$ gh auth login --hostname enterprise.internal
```
## Logout
syntax
```
gh auth logout [flags]
```
登出 github
```
λ ~/ gh auth logout
✓ Logged out of github.com account 'dhay3'
```
## Refresh
syntax
```
gh auth refresh [flags]
```
用于扩展权限，默认会打开浏览器
```
$ gh auth refresh --scopes write:org,read:public_key
# => open a browser to add write:org and read:public_key scopes for use with gh api

$ gh auth refresh
# => open a browser to ensure your authentication credentials have the correct minimum scopes
```
### Optional args

- `-s | --scopes <strings>`

需要额外添加的权限
