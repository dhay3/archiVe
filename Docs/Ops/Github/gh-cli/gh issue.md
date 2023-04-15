# gh issue

ref

https://cli.github.com/manual/gh_issue

用于管理 issues

## list

syntax

```
gh issue list [flags]
```

显示当前 repository  的 issues

可以使用 github issues 查询语法

https://docs.github.com/en/search-github/searching-on-github/searching-issues-and-pull-requests

### Optional args

- `-A | --author <string>`

  按照 author 过滤

- `-l | --label <strings>`

  按照 label 过滤

- `-L | --limit <int>`

  显示最多的条目

- `-s | --state <string>`

  按照 issue 的状态过滤，`{open|closed|all}`

- `-S | --search <query>`

  如果需要查询 issues，就必须使用该参数，可以使用 github issues 查询语法

- `-R | --repo <[HOST/]OWNER/REPO>`

- `-w | --web`

### Examples

```
gh issue list -R cli/cli -S pager 
gh issue list -R cli/cli -s closed -S pager
gh issue list -R cli/cli -s closed -l bug -S pager
```

## view

syntax

```
gh issue view {<number> | <url>} [flags]
```

用于查看 issue 的详情

### Optional args

- `-c`

  显示 issue 的 comments

- `-w | --web`

  在 browser 中打开当前的 issue

- `-R | --repo <[HOST/]OWNER/REPO>`

### Examples

```
gh issue view -R cli/cli 6658
gh issue view -R cli/cli 6658 --web
```

