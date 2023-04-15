# gh pr

ref

https://cli.github.com/manual/gh_pr

用于管理 pull requests

## list

syntax

```
ph pr list [flags]
```

显示当前 repository 的 pull requests

可以使用 github pull request 查询语法

https://docs.github.com/en/search-github/searching-on-github/searching-issues-and-pull-requests

### Optional args

- `-A | --author <string>`

  按照 author 过滤

- `-l | --label <strings>`

  按照 label 过滤

- `-L | --limit <int>`

  显示最多的条目

- `-s | --state <string>`

  按照 pull requests 的状态过滤，`{open|closed|all}`

- `-d | --draft`

  查询 draft 状态的 pull requests

- `-S | --search <query>`

  如果需要查询 pull requests，就必须使用该参数，可以使用 github pull requests 查询语法

- `-R | --repo <[HOST/]OWNER/REPO>`

- `-w | --web`

### Examples

```
gh pr list -R cli/cli
```

## view

syntax

```
gh pr view [<number> | <url> | <branch>] [flags]
```

查看 pull request 详情

### Optional args

- `-c | --comments`

  显示 pull requests comments

- `-w | --web`

## status