# gh browse

ref

https://cli.github.com/manual/gh_browse

syntax

```
gh browse [<number> | <path> | <commit-SHA>] [flags]
```

gh 主要使用 github 提供的 api 接口实现的，所以当然也可以用 browser 来查看指定的 repository

默认打开当前目录对应的 repository

## Optional args

- `-b | --branch <string>`

  打开对应 repository 的 branch 页面

- `-c | --commit <string>`

  查看指定 commit 的信息，非常有用

- `-r | --release`

  打开对应 repository 的 release 页面

- `-R | --repo <[HOST/]OWNER/REPO>`

  使用指定的 repository

- `-s | --settings`

  打开对应 repository 的 settings 页面

