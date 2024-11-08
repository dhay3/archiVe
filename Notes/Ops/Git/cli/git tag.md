# git tag

ref

https://git-scm.com/docs/git-tag

https://git-scm.com/book/en/v2/Git-Basics-Tagging

## Digest

syntax

```
git tag [options] [tagname]
```

用于管理 tag，默认做 list 显示所有的 tag

```
(base) cpl in /tmp/test on main λ git tag -a v1 -m "version one"
(base) cpl in /tmp/test on main λ git tag -a v2 -m "version two"
(base) cpl in /tmp/test on main λ git tag
v1
v2
```

## Optional args

- `-a | -annotate`

  创建 annotated tag

- `-m | --message msg`

   和 `-a` 一起使用，做 tag 记录

- `-s | --sign`

  创建 tag 的时候使用 GPG

-  `-d | --delete`

  删除 tag

