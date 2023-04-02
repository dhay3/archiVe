# git branch

ref

https://git-scm.com/docs/git-branch

## Digest

syntax

```
git branch [options] [branch]
```

用于查询，在创建或者删除 branches

## Optional args

- `--list`

  显示当前本地所有的 branches，如果 `git branch` 没有使用任何 arguments 默认会使用该参数


- `-r | --remotes`

  显示所有 remote tracking branches

- `-d | --delete <branch>`

  删除 branch，只有 branch 被 merge 后才可以

- `-D <branch>`

  强制删除 branch，等价与 `--delete --force`

- `-m | --move <branch> <newbranch>`

  对 branch 重命名

- `-M <branch> <newbranch>`

  强制删除，等价与 `--move --force`

- `-c | --copy <branch>`

  复制 branch

- `-r | --remote`

  显示 remote-tracking branches，即 remote repository 有的 branches

- `-v | -vv`

  以 verbose 的形式输出

