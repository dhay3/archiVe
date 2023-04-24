# git branch

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
  
  ```
  $ git branch -vv
    iss53     7e424c3 [origin/iss53: ahead 2] Add forgotten brackets
    master    1ae2a45 [origin/master] Deploy index fix
  * serverfix f8674d9 [teamone/server-fix-good: ahead 3, behind 1] This should do it
    testing   5ea463a Try something new
  ```
  
  - `*` 表示 local HEAD reference
  - `serverfix` 表示当前所处的分支，也是 local reference 对应的名字
  - `f8674d9` 表示当前分支 local reference 指向的 last commit hash
  - `team/server-fix-good` 表示当前分支追踪的 remote branch，也是本地 remote references
  - `ahead3,behind 1` 表示本地有 3 次 commit 没有 push 到 remote repository，remote repository 有被其他人 push 过 1 次，但是没有被本地 merge
  - `This should do it` 表示 serverfix branch last commit message

**references**

1. https://git-scm.com/docs/git-branch
2. https://git-scm.com/book/en/v2/Git-Branching-Branches-in-a-Nutshell
