ref
[https://cli.github.com/manual/gh_search](https://cli.github.com/manual/gh_search)
默认所有 github 上所有的内容，不仅限于当前的用户
## Repos
syntax
```
gh search repos [<query>] [flags]
```
按照参数查找包含 query 关键字的 repos，可以和 Github search syntax 一起使用
[https://docs.github.com/en/search-github/searching-on-github/searching-for-repositories](https://docs.github.com/en/search-github/searching-on-github/searching-for-repositories)
```
λ ~/ gh search repos in:name followers:\>5000 gitlab
```
注意特殊符号需要在 shell 中转译
### Optional args

- `--followers <number>`
- `--forks <number>`
- `--stars <number>`
- `--language <string>`
- `--match <strings>`

从 repo 的 `{name|description|readme}` 中过滤

- `--sort <string>`

按照指定的内容排序，`{forks|help-wanted-issues|stars|updated}`

- `--order <string>`

指定输出结果的排序方式，`{asc|desc}`

- `--onwer`

按照 repo 所有者过滤

- `-w | --web`

过滤后使用 browser 打开
### Examples
```
#可以使用比较符号但是需要注意转译
λ ~/ gh search repos linux --stars \>10000 --forks \>10000 --match name
#等价,同样需要注意转译
λ ~/ gh search repos stars:\>10000 forks:\>10000 in:name linux
```
## Commits
syntax
```
gh search commits [<query>] [flags]
```
按照参数查找包含 query 关键字的 commits 记录，可以和 Github search syntax 一起使用
[https://docs.github.com/search-github/searching-on-github/searching-commits](https://docs.github.com/search-github/searching-on-github/searching-commits)
### Optional args

- `--author <string>`
- `--author-email <string>`
- `--owner <string>`
- `--repo <strings>`
- `--hash <strings>`

按照 hash 过滤

- `-w | --browser`
### Exmaples
```
λ ~/ gh search commits init --repo torvalds/linux 
#等价
λ ~/ gh search commits init repo:torvalds/linux 
```
## Issues
syntax
```
gh search issues [<query>] [flags]
```
按照参数查找包含 query 关键字的 issues，可以和 Github search syntax 一起使用
[https://docs.github.com/search-github/searching-on-github/searching-issues-and-pull-requests](https://docs.github.com/search-github/searching-on-github/searching-issues-and-pull-requests)
### Optional args

- `--owner <string>`
- `--repo <string>`
- `--state <string>`

按照 issues 状态过滤，可以是 `{open|closed}`

- `--closed <date>`

按照 issues closed 日期过滤

- `--created <date>`

按照 issues created 日期过滤

- `--update <date>`

按照 last update 日期过滤

- `--sort <string>`

按照 `{comments|created|updated|reactions...}`字段排序，reactions 表示表情具体对应规则参考文档
### Examples
```
λ ~/ gh search issues --repo spring-projects/spring-boot --sort comments --state open
```
## Prs
syntax
```
gh search prs [<query>] [flags]
```
按照参数查找包含 query 关键字的 pull requests，可以和 Github search syntax 一起使用
[https://docs.github.com/search-github/searching-on-github/searching-issues-and-pull-requests](https://docs.github.com/search-github/searching-on-github/searching-issues-and-pull-requests)
### Optional args

- `--author <string>`
- `--owner <string>`
- `--repo <string>`
- `--closed <date>`
- `--created <date>`
- `--review <strings>`

按照 review 的状态过滤，`{none|required|approved|changes_requested}`

- `--draft`

过滤草稿状态的

- `--merged`

过滤已经 merged 的 pull requests

- `--sort <string>`

按照 `{comments|created|updated|reactions...}`字段排序，reactions 表示表情具体对应规则参考文档
### Exmaples
```
λ ~/ gh search prs --repo spring-projects/spring-boot --review approved --sort comments
```
