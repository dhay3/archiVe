# git checkout

ref

https://git-scm.com/docs/git-checkout

## Digest

syntax

```
git checkout [options] [branch]
```

git branch 一般用于切换分支，例如 `git branch master` 会将当前使用的分支替换到 master，即 HEAD 指针移动到 master (最新的 commit)

```
(base) cpl in /tmp/gitignore on main λ git br -vv
* main 4488915 [origin/main] Merge pull request #3567 from aaronfranke/godot
  test 4488915 Merge pull request #3567 from aaronfranke/godot
```

其中的 `*` 就表示当前指向的 branch，`4488915` 表示当前指向 branch 对应 commit 的 hash，`[origin/main]` 表示对应的 remote-tracking branch

## Optional args

- `-b <new-branch>`

  创建新的 branch 并切换到该分支

## HEAD

HEAD 是一个指针，指向当前所处的分支对应最新的 commit

假设当前 HEAD 指向 branch master commit c

```
	HEAD (refers to branch 'master')
        |
        v
a---b---c branch 'master'
```

现在执行了一次 `git commit` 增加了一个 commit d，那个 HEAD 就会指向 commit d

```
        HEAD (refers to branch 'master')
            |
            v
a---b---c---d branch 'master'
```

假设现在有 2 个 branch b1 b2，b2 上有做过一次 commit ,当前 HEAD b1，需要切换到 b2

```
        HEAD (refers to branch 'b1')
            |
 			v
a---b---c---d branch 'b1'
	|
	e---f branch 'b2'
	
```

如果执行 `git checkout b2` 就会将 HEAD 指向 branch b2 f commit

```
a---b---c---d branch 'b1'
	|
	e---f branch 'b2'
        ^
        |
    HEAD (refers to branch 'b2')
```

## Staging area

假设现在新增了一个 test 文件，并执行了 `git add test`

```
(base) cpl in /tmp/gitignore on main λ git st
On branch main
Your branch is up to date with 'origin/main'.

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        test

nothing added to commit but untracked files present (use "git add" to track)
                                                                                                                                                         
(base) cpl in /tmp/gitignore on main ● λ git add test
```

然后`git checkout test`, staging area 的状态是什么样的呢？

```
(base) cpl in /tmp/gitignore on main ● λ git checkout test
A       test
Switched to branch 'test'
                                                                                                                                                         
(base) cpl in /tmp/gitignore on test ● λ git st
On branch test
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        new file:   test

```

staging area 会和切换分支前的状态一样，即表示 staging area 不会随着切换分支而改变，或保留在 “原” staging area，即可将 staging area 理解成 “公用区间”

## Working directory

staginig area 不会随着分支的切换而改变，那么 working directory 呢？

```
root@v2:/home/ubuntu/test# git br
* master
  test
root@v2:/home/ubuntu/test# ls
file1  file2
root@v2:/home/ubuntu/test# git co test
Switched to branch 'test'
root@v2:/home/ubuntu/test# ls
file1
```

working directory 是会随着 branch 的变化而改变的