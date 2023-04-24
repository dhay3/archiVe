# git fetch

## Digest
syntax
```
git fetch [options] [remote_alias]
```
从 remote repository 拉取 remote references 以及 tags 到本地，如果没有指定 remote_alias 默认使用 origin，不会对本地的 working directory 造成影响

```
(base) cpl in /tmp/test on main λ git init
(base) cpl in /tmp/test on main λ ls
(base) cpl in /tmp/test on main λ git remote add origin https://github.com/github/gitignore.git
(base) cpl in /tmp/test on main λ git fetch
remote: Enumerating objects: 9739, done.
remote: Total 9739 (delta 0), reused 0 (delta 0), pack-reused 9739
Receiving objects: 100% (9739/9739), 2.31 MiB | 3.50 MiB/s, done.
Resolving deltas: 100% (5296/5296), done.
From https://github.com/github/gitignore
 * [new branch]      add-metadata            -> origin/add-metadata
 * [new branch]      annotating-visualstudio -> origin/annotating-visualstudio
 * [new branch]      ghfw                    -> origin/ghfw
 * [new branch]      main                    -> origin/main
 * [new branch]      old-ghfw                -> origin/old-ghfw
 * [new branch]      reduce-noise            -> origin/reduce-noise
 * [new branch]      rmw-universe            -> origin/rmw-universe
 * [new branch]      rmw-universe-2022       -> origin/rmw-universe-2022
 (base) cpl in /tmp/test on main λ ls
```

## Optional args

- `-j | --jobs <n>`

  fetch 时指定需要使用的子进程数 

- `--dry-run`

  不实际 fetch，只测试

- `-n | --no-tags`

  fetch 时不 fetch tags

- `-v | --verbose`

  fetch 时显示详情

- `--depth=<depth>`

  和 `git clone` 的 `--depth` 含义相同
  
  

**references**

1. [https://git-scm.com/docs/git-fetch](https://git-scm.com/docs/git-fetch)
2. https://git-scm.com/book/en/v2/Git-Basics-Working-with-Remotes](https://git-scm.com/book/en/v2/Git-Basics-Working-with-Remotes)
