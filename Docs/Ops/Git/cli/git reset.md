git add assets/# git reset

ref
[https://git-scm.com/docs/git-reset](https://git-scm.com/docs/git-reset)

## Digest
reset 时需要注意自己所处的分支

syntax

```
git reset [options] [commit]
```
用于移动 HEAD 指针到对应的 commit object，或者是回退 staging area 和 working directory 中的内容

也可以使用下面的命令来回退 staging area 中的内容

```
git reset [options] [file_path]
```
即 `git add file_path` 的反向操作，等价于 `git restore --staged file_path`

例如

```
root@v2:/home/ubuntu/test# git add file5
root@v2:/home/ubuntu/test# git st
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        new file:   file5

Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        deleted:    file4

root@v2:/home/ubuntu/test# git reset file5
Unstaged changes after reset:
D       file4
root@v2:/home/ubuntu/test# git st
On branch master
Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        deleted:    file4

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        file5

no changes added to commit (use "git add" and/or "git commit -a")
```
## Optional args

- `--soft <commit>`

  回退到指定 commit ，但是不修改 index file 和 working direcotry 的状态。和 `git commit --amend` 类似，状态会回到 "changes to be committed"

- `--mixed <commit>`

  mixed 是 reset 默认使用的模式，将  HEAD reset 到指定 commit object 时，还会 reset staging area 但是不会 reset working directory

- `--hard <commit>`

  回退到指定 commit ，修改 staging area 和 working tree 的状态

## mixed

mixed 是 reset 默认使用的模式，将  HEAD reset 到指定 commit object 时，还会 reset staging area 但是不会 reset working directory

```
(base) cpl in /tmp/test on master ● λ ls
 a   b   c   d
(base) cpl in /tmp/test on master λ git lg1
* 8d455c9 - (3 seconds ago) d - 4liceh  (HEAD -> master)
* 2adb676 - (34 seconds ago) c - 4liceh 
* ed02c53 - (39 seconds ago) b - 4liceh 
* bd537a2 - (46 seconds ago) a - 4liceh %
(base) cpl in /tmp/test on master λ git reset ed02c53
(base) cpl in /tmp/test on master ● λ git lg1
* ed02c53 - (54 seconds ago) b - 4liceh  (HEAD -> master)
* bd537a2 - (61 seconds ago) a - 4liceh %
(base) cpl in /tmp/test on master ● λ ls
 a   b   c   d
```

现在在 staging area 中添加一个 e，但是不做 commit

```
(base) cpl in /tmp/test on master λ git add e
(base) cpl in /tmp/test on master ● λ git st
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        new file:   e
(base) cpl in /tmp/test on master ● λ ls
 a   b   c   d   e
```

然后直接做 reset

```
(base) cpl in /tmp/test on master ● λ ls
 a   b   c   d   e
(base) cpl in /tmp/test on master ● λ git reset ed02c53  
(base) cpl in /tmp/test on master ● λ git lg1          
* ed02c53 - (21 minutes ago) b - 4liceh  (HEAD -> master)
* bd537a2 - (22 minutes ago) a - 4liceh
(base) cpl in /tmp/test on master ● λ git st
On branch master
Untracked files:
  (use "git add <file>..." to include in what will be committed)
        c
        d
        e
```

这里可以明显发现原本 staging area 中的 e 变成 untrack 的状态了，但是 working directory 中的内容没有改变

## soft

将  HEAD reset 到指定 commit object 时，但是不会 reset staging area（如果之前的这台是 added 就是 added） 和 working directory

初始状态如下

```
(base) cpl in /tmp/test on master ● λ git lg1
* 8d455c9 - (23 minutes ago) d - 4liceh  (HEAD -> master)
* 2adb676 - (23 minutes ago) c - 4liceh 
* ed02c53 - (23 minutes ago) b - 4liceh 
* bd537a2 - (24 minutes ago) a - 4liceh
(base) cpl in /tmp/test on master ● λ git st
On branch master
nothing to commit, working tree clean
(base) cpl in /tmp/test on master ● λ ls
 a   b   c   d
(base) cpl in /tmp/test on master λ echo e > e
(base) cpl in /tmp/test on master λ git add e
(base) cpl in /tmp/test on master ● λ git st
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        new file:   e
```

执行 reset

```
(base) cpl in /tmp/test on master λ git reset --soft ed02c53
(base) cpl in /tmp/test on master ● λ git lg1
* ed02c53 - (24 minutes ago) b - 4liceh  (HEAD -> master)
* bd537a2 - (24 minutes ago) a - 4liceh
(base) cpl in /tmp/test on master ● λ git st
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        new file:   c
        new file:   d
        new file:   e
(base) cpl in /tmp/test on master ● λ ls
 a   b   c   d
```

## Hard

将  HEAD reset 到指定 commit object 时，会同时 reset staging area 和 working directory

初始状态如下，文件 e 加入到 staging area，但是文件 f 没有

```
(base) cpl in /tmp/test on master ● λ git lg1
* 8d455c9 - (23 minutes ago) d - 4liceh  (HEAD -> master)
* 2adb676 - (23 minutes ago) c - 4liceh 
* ed02c53 - (23 minutes ago) b - 4liceh 
* bd537a2 - (24 minutes ago) a - 4liceh
(base) cpl in /tmp/test on master λ git st
On branch master
nothing to commit, working tree clean
(base) cpl in /tmp/test on master λ ls  
 a   b   c   d
(base) cpl in /tmp/test on master λ echo e > e
(base) cpl in /tmp/test on master λ echo f > f
(base) cpl in /tmp/test on master λ git add e
(base) cpl in /tmp/test on master ● λ git st
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        new file:   e

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        f
```

执行 reset

```
(base) cpl in /tmp/test on master ● λ git reset --hard ed02c53
HEAD is now at ed02c53 b
(base) cpl in /tmp/test on master λ git lg1
* ed02c53 - (41 minutes ago) b - 4liceh  (HEAD -> master)
* bd537a2 - (41 minutes ago) a - 4liceh %
(base) cpl in /tmp/test on master ● λ git st
On branch master
Untracked files:
  (use "git add <file>..." to include in what will be committed)
        f
nothing added to commit but untracked files present (use "git add" to track)
nothing to commit, working tree clean
(base) cpl in /tmp/test on master ● λ ls
 a   b   f
```

可以看到的一点是 working directory 改变了，但是不修改 unstaged 的文件，因为该文件并没有通过 `git commit` 加入到版本库中，也就不在 working directory 中

还有一种特殊的情况，假设当前的状态如下

```
(base) cpl in /tmp/test on main λ ls                      
 a   b   e   f
(base) cpl in /tmp/test on main λ git sw -
Switched to branch 'topic'
(base) cpl in /tmp/test on topic λ ls
 a   b   c   d
(base) cpl in /tmp/test on topic λ git lg1
* 82e3551 - (3 minutes ago) f - cyberPelican  (main)
* 1e8cd71 - (3 minutes ago) e - cyberPelican 
| * b822af5 - (3 minutes ago) d - cyberPelican  (HEAD -> topic)
| * 47bc052 - (3 minutes ago) c - cyberPelican 
|/  
* 3c6c505 - (3 minutes ago) b - cyberPelican 
* 0b6ec6c - (4 minutes ago) a - cyberPelican
```

现在在不切换 branch 的情况下，直接使用如下命令

```
(base) cpl in /tmp/test on topic λ git reset --hard 82e3551
```

那么当前 branch 的 working directory 同样也会发生变化

```
(base) cpl in /tmp/test on topic λ ls
 a   b   e   f
```

