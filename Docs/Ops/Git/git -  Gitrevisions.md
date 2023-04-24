# git - Gitrevisions

https://git-scm.com/docs/gitrevisions

## Digest

git 很多 subcommand 都会使用 revision，例如  `git log`, `git reset`, `git restore` 以及 `git show` 

来索引相对记录

```
(base) cpl in /tmp/test on main ● λ git reflog
19098b5 (HEAD -> main) HEAD@{0}: reset: moving to 19098b5
6768e6a HEAD@{1}: commit: e
a55d45f HEAD@{2}: commit: d
19098b5 (HEAD -> main) HEAD@{3}: commit: c
6210fd9 HEAD@{4}: commit: b
1a97b4f HEAD@{5}: commit (initial): a
```

当前指向 c 对应的 commit hash 为 `19098b5`

## Specifying Revisions

revisions ( rev 或者被称为 object name ) 一般不仅仅是 names 或者是 commit object，也可以是多种多样的，例如

- `<sha1>` e.g. *dae86e1950b1277e545cee180551750029cfe735*, *dae86e*

  对应 commit object hash 值

- `<refname>` e.g. *master*, *heads/master*, *refs/heads/master*

  分支名 (分支指针)

- `@`

  等价于 HEAD

- `<rev>:<path>`  e.g. *HEAD:README*, *master:./README*

  表示 rev 中对应 path 的一部分，如果需要表示相对目录需要使用 `./<path>` 方式

- `<rev>^[<n>]` eg.  `HEAD^`, `v1.5.1^0` 

  a suffix `^` to revision parameter means the first parent of that commit object. `^<n>` means the `<n>`th parent

  ```
  (base) cpl in /tmp/test on main ● λ git reset HEAD^  
  (base) cpl in /tmp/test on main ● λ git reflog
  6210fd9 (HEAD -> main) HEAD@{0}: reset: moving to HEAD^
  19098b5 HEAD@{1}: reset: moving to 19098b5
  6768e6a HEAD@{2}: commit: e
  a55d45f HEAD@{3}: commit: d
  19098b5 HEAD@{4}: commit: c
  6210fd9 (HEAD -> main) HEAD@{5}: commit: b
  1a97b4f HEAD@{6}: commit (initial): a
  ```

  `HEAD^` 等价于 `HEAD^1`

- `<rev>~[<n>]` eg. `HEAD~`, `master~3`

  a suffix `~` to a revision parameter means the first parent of that commit object. `~<n>` to  a revision parameter means the commit object that is the `<n>`th generation ancestor 

  ```
  (base) cpl in /tmp/test on main ● λ git reset HEAD~
  (base) cpl in /tmp/test on main ● λ git reflog
  1a97b4f (HEAD -> main) HEAD@{0}: reset: moving to HEAD~
  6210fd9 HEAD@{1}: reset: moving to HEAD^
  19098b5 HEAD@{2}: reset: moving to 19098b5
  6768e6a HEAD@{3}: commit: e
  a55d45f HEAD@{4}: commit: d
  19098b5 HEAD@{5}: commit: c
  6210fd9 HEAD@{6}: commit: b
  1a97b4f (HEAD -> main) HEAD@{7}: commit (initial): a
  ```

## Cheat Sheet

```
G   H   I   J
 \ /         \ /
  D   E   F
   \   |     / \
    \  |  /     |
     \ | /      |
      B      C
      \      /
        \  /
         A
```

如当前分支如上图，B 和 C 都是 A 的 parents， commit 顺序按照图表从右到左提交

可以转换为如下公式

```
A =      = A^0
B = A^   = A^1     = A~1
C =      = A^2
D = A^^  = A^1^1   = A~2
E = B^2  = A^^2
F = B^3  = A^^3
G = A^^^ = A^1^1^1 = A~3
H = D^2  = B^^2    = A^^^2  = A~2^2
I = F^   = B^3^    = A^^3^
J = F^2  = B^3^2   = A^^3^2
```

