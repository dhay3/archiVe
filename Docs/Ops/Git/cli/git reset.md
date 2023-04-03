ref
[https://git-scm.com/docs/git-reset](https://git-scm.com/docs/git-reset)
## Digest
syntax
```
git reset [options] [commit]
```
用于回退 commit，即将当前分支的指针指向指定的 commit
或者是回退 staging area 中指定文件，即 `git add file_path` 的反向操作，等价于 `git restore --staged file_path`
```
git reset [options] [file_path]
```
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

回退到指定 commit ，修改 staging area 但是不修改 working directory 的状态，`git reset` 默认使用该模

- `--hard <commit>`

回退到指定 commit ，修改 staging area 和 working tree 的状态
## Examples
回撤 `git add` 操作
```
λ ~/lab/ master* git add a
                                                                                                                                     
λ ~/lab/ master* git st
On branch master

No commits yet

Changes to be committed:
  (use "git rm --cached <file>..." to unstage)

        new file:   a

                                                                                                                                     
λ ~/lab/ master* git reset
                                                                                                                                     
λ ~/lab/ master* git st
On branch master

No commits yet

Untracked files:
  (use "git add <file>..." to include in what will be committed)

        a

nothing added to commit but untracked files present (use "git add" to track)
```
假设现在需要回退到 a 对应的 commit 记录（回撤 `git commit` 操作），同时不修改已经创建并 commit 的文件
```
root@v2:/home/ubuntu/test# git lg2
*   f624c02 - Mon, 3 Apr 2023 17:44:36 +0800 (32 seconds ago) (HEAD -> master)
|\            Merge branch 'topic' - hl4ce
| * c0ebb2e - Mon, 3 Apr 2023 17:42:34 +0800 (3 minutes ago) (topic)
| |           d - hl4ce
* | f0ad174 - Mon, 3 Apr 2023 17:42:01 +0800 (3 minutes ago)
|/            c - hl4ce
* 0c6f9e0 - Mon, 3 Apr 2023 16:59:26 +0800 (46 minutes ago)
|           b - hl4ce
* b4da5d7 - Mon, 3 Apr 2023 16:58:46 +0800 (46 minutes ago)
            a - hl4ce

```
可以使用 `--soft` 参数
```
root@v2:/home/ubuntu/test# git reset --soft b4da5d7
root@v2:/home/ubuntu/test# ls
a  b  c  d
root@v2:/home/ubuntu/test# git status
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        new file:   b
        new file:   c
        new file:   d
```
如果需要回退到 a 对应的 commit 记录，同时 working directory 也回撤到同记录时的内容，可以使用 `--hard`
```
root@v2:/home/ubuntu/test# git reset --hard b4da5d7
HEAD is now at b4da5d7 a
root@v2:/home/ubuntu/test# ls
a
```
