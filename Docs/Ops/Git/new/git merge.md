ref
[https://git-scm.com/docs/git-merge](https://git-scm.com/docs/git-merge)
[https://git-scm.com/book/en/v2/Git-Branching-Basic-Branching-and-Merging](https://git-scm.com/book/en/v2/Git-Branching-Basic-Branching-and-Merging)
[https://www.atlassian.com/git/tutorials/using-branches/git-merge#:~:text=Fast%20Forward%20Merge,to%20the%20target%20branch%20tip.](https://www.atlassian.com/git/tutorials/using-branches/git-merge#:~:text=Fast%20Forward%20Merge,to%20the%20target%20branch%20tip.)
[https://stackoverflow.com/questions/29673869/what-is-git-fast-forwarding](https://stackoverflow.com/questions/29673869/what-is-git-fast-forwarding)
## Digest
syntax
```
git merge 
```
主要用于 branch 之间的 merge 操作
假设分别有 2 个 branch，当前处在 master branch
```
λ ~/test/ topic* git log --all --graph
* commit 987c6c2c27c9c7ca133ea8dcb7b95f6313521c50 (master)
| Author: Your Name <you@example.com>
| Date:   Fri Mar 31 18:16:22 2023 +0800
| 
|     f
| 
* commit b308e152cdea31c578d3bae159ad59994b093c13
| Author: Your Name <you@example.com>
| Date:   Fri Mar 31 18:16:16 2023 +0800
| 
|     e
|   
| * commit f63c52a5d2d542b8e2c245bebbb19e82bd6d9e02 (HEAD -> topic)
| | Author: Your Name <you@example.com>
| | Date:   Fri Mar 31 18:15:59 2023 +0800
| | 
| |     d
| | 
| * commit 4a965cfbbd9c1c374fa7422dd8e5b3d7a8b6f947
|/  Author: Your Name <you@example.com>
|   Date:   Fri Mar 31 18:15:50 2023 +0800
|   
|       c
| 
* commit c8dfd8c9e0258b563fcf8a159666e707fd0fa5b0
| Author: Your Name <you@example.com>
| Date:   Fri Mar 31 18:15:14 2023 +0800
| 
|     b
| 
* commit 7dfde0561ccf51333e93f58d5f6bde73572f608b
  Author: Your Name <you@example.com>
  Date:   Fri Mar 31 18:14:52 2023 +0800
  
      a
```
现在需要将 topic branch merge into master
```
λ ~/test/ topic* git checkout master
Switched to branch 'master'
                                                                                                                                                                                                                                     
λ ~/test/ master* git merge topic
Merge made by the 'recursive' strategy.
 c | 1 +
 d | 1 +
 2 files changed, 2 insertions(+)
 create mode 100644 c
 create mode 100644 d
λ ~/test/ master* ls
a  b  c  d  e  f  g
```
那么 merge 后的状态如下，topic 和 master merge 实际会做一次 commit 生成新的 b0deaac15cd4be605dcb06227efc98401c8fa931 记录
```
λ ~/test/ master* git log --all --graph
*   commit b0deaac15cd4be605dcb06227efc98401c8fa931 (HEAD -> master)
|\  Merge: 987c6c2 f63c52a
| | Author: Your Name <you@example.com>
| | Date:   Mon Apr 3 14:29:21 2023 +0800
| | 
| |     Merge branch 'topic'
| | 
| * commit f63c52a5d2d542b8e2c245bebbb19e82bd6d9e02 (topic)
| | Author: Your Name <you@example.com>
| | Date:   Fri Mar 31 18:15:59 2023 +0800
| | 
| |     d
| | 
| * commit 4a965cfbbd9c1c374fa7422dd8e5b3d7a8b6f947
| | Author: Your Name <you@example.com>
| | Date:   Fri Mar 31 18:15:50 2023 +0800
| | 
| |     c
| | 
* | commit 987c6c2c27c9c7ca133ea8dcb7b95f6313521c50
| | Author: Your Name <you@example.com>
| | Date:   Fri Mar 31 18:16:22 2023 +0800
| | 
| |     f
| | 
* | commit b308e152cdea31c578d3bae159ad59994b093c13
|/  Author: Your Name <you@example.com>
|   Date:   Fri Mar 31 18:16:16 2023 +0800
|   
|       e
| 
* commit c8dfd8c9e0258b563fcf8a159666e707fd0fa5b0
| Author: Your Name <you@example.com>
| Date:   Fri Mar 31 18:15:14 2023 +0800
| 
|     b
| 
* commit 7dfde0561ccf51333e93f58d5f6bde73572f608b
  Author: Your Name <you@example.com>
  Date:   Fri Mar 31 18:14:52 2023 +0800
  
      a
```
## Optional args

- `--commit | --no-commit`

merge 的时候是否做 commit

- `--stat | --summary`

merge 完成后是否显示 dffstat 

- `-v | --verbose`
- `-S | --gpg-sign=keyid`

## Pre-Merge Checks
不管是执行 `git merge` 还是 `git pull` 都会校验 local uncommitted changes 是否和 `git merge` 或者 `git pull` 内容出现重叠覆盖。如果出现重叠覆盖 git 不会做任何操作
如果当前本地的 HEAD 指针和 remote HEAD 指针一致，就会提示 "Already up to date"
## Fast-Forward Merge
当前需要 merge 的节点是当前节点的 ancestor 时，git 就会选择 Fast-Forward Merge。一般出现在 `git pull` 对 tracking 的分支做更新时出现
简单的理解就是当前节点和需要合并的节点在一条线上时，git 会采用 Fast-Forward Merge，将 HEAD 指针指向更新后的节点
但是还有一种特殊情况也会采用 Fast-Forward Merge
假设现在有两个分支分别提交了一次
```
root@v2:/home/ubuntu/test# git add a
root@v2:/home/ubuntu/test# git commit -m "a"
[master (root-commit) 9462119] a
 1 file changed, 1 insertion(+)
 create mode 100644 a
root@v2:/home/ubuntu/test# git checkout -b topic
Switched to a new branch 'topic'
root@v2:/home/ubuntu/test# git add b
root@v2:/home/ubuntu/test# git commit -m "b"
[topic dbcdead] b
 1 file changed, 1 insertion(+)
 create mode 100644 b
```
日志如下
```
root@v2:/home/ubuntu/test# git lg1
* dbcdead - (53 seconds ago) b - hl4ce (HEAD -> topic)
* 9462119 - (2 minutes ago) a - hl4ce (master)
```
现在需要 merge master 和 topic，这时 git 就会使用 fast-forward merge
```
root@v2:/home/ubuntu/test# git checkout master
Switched to branch 'master'
root@v2:/home/ubuntu/test# git merge topic
Updating 9462119..dbcdead
Fast-forward
 b | 1 +
 1 file changed, 1 insertion(+)
 create mode 100644 b
```
这里也是只是将分支的指针移动了
```
root@v2:/home/ubuntu/test# git lg1
* dbcdead - (19 minutes ago) b - hl4ce (HEAD -> master, topic)
* 9462119 - (19 minutes ago) a - hl4ce
```
但是目录内容以及改变了
```
root@v2:/home/ubuntu/test# ls
a  b
```
## True Merge
当前 HEAD 指针指向的分支 commit 和需要合并的分支 commit，不在一条线上，即出现 diverg。会采用 True Merge
例如想要将 master 和 topic 分支对应的 commit 合并
```
root@v2:/home/ubuntu/test# git lg1
* c0ebb2e - (84 seconds ago) d - hl4ce (HEAD -> topic)
| * f0ad174 - (2 minutes ago) c - hl4ce (master)
|/  
* 0c6f9e0 - (45 minutes ago) b - hl4ce
* b4da5d7 - (45 minutes ago) a - hl4ce
root@v2:/home/ubuntu/test# 
```
合并分支
```
root@v2:/home/ubuntu/test# git checkout master
Switched to branch 'master'
root@v2:/home/ubuntu/test# git merge topic
Merge made by the 'recursive' strategy.
 d | 1 +
 1 file changed, 1 insertion(+)
 create mode 100644 d
```
输出结果如下
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
可以看到分支向前走了一个 commit 记录
## How Conflicts Are Presented
当使用 `git merge` 或者 `git pull` 时，如果需要合并的部分有重叠 git 就会提示需要人工合并或者直接放弃合并的操作
例如

