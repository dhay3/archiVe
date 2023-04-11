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

可以使用 github issues 搜索的语法

https://docs.github.com/en/search-github/searching-on-github/searching-issues-and-pull-requests

### Optional args

- `-A`

  按照 author 过滤

- `-l | --label <strings>`

  按照 label 过滤

- `-L | --limit <int>`

  显示最多的条目

- `-S | --search <query>`

  使用 issues 搜索语法

- `-R | --repo <[HOST/]OWNER/REPO>`

### Examples

```
(base) cpl in ~/hugo λ gh issue list -R cli/cli -S pager 

Showing 4 of 4 issues in cli/cli that match your search

#1727  PAGER + Cygwin = partial output                                             bug, windows, p3, needs-investigation, help wanted  about 1 month ago
#6463  When upgrading an extension, print information about changes since the ...  enhancement                                         about 5 months ago
#6179  Prioritize PR check name/context over URI                                   enhancement, help wanted                            about 5 months ago
#1980  Better document per-host configuration                                      enhancement, core                                   about 1 year ago
```

