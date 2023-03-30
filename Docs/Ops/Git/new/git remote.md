# git remote

ref

https://git-scm.com/docs/git-remote

## Digest

syntax

```
git remote [options] | [[subcommand] [options]]
```

管理 tracked repositories ( remote repositories )

## Subcommands

### add

用于添加 remote repository 例如

```
git remote add origin https://github.com/github/gitignore.git
```

就会添加 `https://github.com/github/gitignore.git` remote repository。本地以 origin 做为 remote repository alias

- `-f`

  添加 remote repository 的同时，做`git fetch origin` 

- `--tags`

  添加 remote repository 的同时，做 `git fetch origin`

### rename

用于对 remote repository alias 重命名 例如

```
root@v2:/home/ubuntu/test# git remote rename origin or
root@v2:/home/ubuntu/test# git remote -v
or      https://github.com/github/gitignore.git (fetch)
or      https://github.com/github/gitignore.git (push)
```

### remove/rm

用于 remove remote repository 例如

```
git remote rm origin
```

### set-url

用于更新设置 remote repository url

```
git remote set-url origin https://github.com/github/linux.git 
```

### show

显示 remote repository 信息

```
(base) cpl in /tmp/gitignore on test ● λ git remote show origin
* remote origin
  Fetch URL: https://github.com/github/gitignore.git
  Push  URL: https://github.com/github/gitignore.git
  HEAD branch: main
  Remote branches:
    add-metadata            tracked
    annotating-visualstudio tracked
    ghfw                    tracked
    main                    tracked
    old-ghfw                tracked
    reduce-noise            tracked
    rmw-universe            tracked
    rmw-universe-2022       tracked
  Local branch configured for 'git pull':
    main merges with remote main
  Local ref configured for 'git push':
    main pushes to main (up to date)
```

表示当前 remote repository 有 `add-metadata`, `annotating-visualstudio`, `ghfw`, `main`, `old-ghfw`, `reduce-noise`, `rmw-universe`, `rmw-universe-2022`; 如果运行 `git pull` 默认 merge 的分支是 remote repository 的 `main` 分支，如果运行 `git push` 默认 push 的分支是本地的 `main` 分支 

### update

从 remote repository 同步更新到本地 

## Optional args

- `-v | --verbose`

  verbose 输出 remote 信息

  ```
  root@v2:/home/ubuntu/test# git remote -v
  or      https://github.com/github/gitignore.git (fetch)
  or      https://github.com/github/gitignore.git (push)
  ```

  