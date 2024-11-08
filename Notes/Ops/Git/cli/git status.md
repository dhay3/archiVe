# git status

ref
[https://git-scm.com/docs/git-status](https://git-scm.com/docs/git-status)

## Digest
syntax
```
git status [options] [pathspec...]
```
显示 staging area 和 working directory 中状态变更或者未追踪的文件，如果指定了 pathspec，查看指定的目录
```
root@v2:/home/ubuntu/test# git status
On branch master

No commits yet

Changes to be committed:
  (use "git rm --cached <file>..." to unstage)
        new file:   a

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        b
```
例如上述，就表示文件 `b` 还没有使用 `git add` 加入到 staging area，文件 `a` 使用了 `git add` 加入到了 staging area，但是还没有使用 `git commit` 加入到 repository
## Optional args

- `-s | --short`

  已 short-format 的格式出现文件的状态。
  其中 `M` 表示 modified，`A` 表示 added，`D` 表示 deleted，`R` 表示 renamed

- `-v | --verbose`

  如果文件处于 modified 状态，还会显示修改的内容

