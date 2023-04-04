# git push

ref
[https://git-scm.com/docs/git-push](https://git-scm.com/docs/git-push)
[https://git-scm.com/book/en/v2/Git-Branching-Remote-Branches](https://git-scm.com/book/en/v2/Git-Branching-Remote-Branches)

## Digest
syntax
```
git push [options] [repository_alias] [branch]
```
updates remote refs using local refs

- `git push`

  默认推送 origin/{current_branch}

- `git push origin`

  默认推送 origin/{current_branch}，等价于 `git push origin HEAD`

- `git push origin master`

  推送到 origin master 分支

- `git push origin HEAD:master`

  不需要考虑本地当前使用的分支是否匹配 remote repositoy，直接推送到 origin master 分支

## Optional args

- `--all`

  push all branches

- `-n | --dry-run`

  不实际上传数据到 remote repository

- `-d | --delete `

  删除 remote 指定分支

  ```
  λ ~/gitlab/.git/ master git push -d origin iss53 
  To https://github.com/dhay3/gitlab.git
   - [deleted]         iss53
  ```

- `--tags`

  push 本地所有的 tags 到 remote repository

- `-f | --force`

  如果 remote repository 对应 branch 中的内容不是当前需要推送的内容的历史内容就会被拒绝。可以使用这个参数 表示强制推送，有风险

- `-v | --verbose`

