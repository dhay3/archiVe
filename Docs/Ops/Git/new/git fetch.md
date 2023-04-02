ref
[https://git-scm.com/docs/git-fetch](https://git-scm.com/docs/git-fetch)
[https://git-scm.com/book/en/v2/Git-Basics-Working-with-Remotes](https://git-scm.com/book/en/v2/Git-Basics-Working-with-Remotes)
## Digest
syntax
```
git fetch [options] [remote_alias]
```
从 remote repository 拉取 branches 和 tags 到本地，如果没有指定 remote_alias 默认使用 origin，不会对本地的 branches 造成影响
## Optional args

- `-j | --jobs <n>`

fetch 时指定需要使用的子进程数 

- `--dry-run`

不实际 fetch，只测试

- `-n | --no-tags`

fetch 时不 fetch tags

- `-v | --verbose`

fetch 时显示详情

- `--depth=<depth>`

和 `git clone` 的 `--depth` 含义相同
