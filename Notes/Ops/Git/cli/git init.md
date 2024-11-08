# git init

ref

https://git-scm.com/docs/git-init

## Digest

git init 用于创建空的 Git repository (`.git`)。如果在一个已有 Git repository 的目录下使用 `git init` 也是安全的，不会 overwrite 原有的内容

## Optional args

- `--bare`

  创建一个 bare repository

- `-b | --initial-branch branch-name`

  指定创建 repository 时 branch 名。如果没有指定 默认使用 `main`。可以通过 `init.defaultBranch` 配置参数来设定