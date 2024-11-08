# git reflog

ref

https://git-scm.com/docs/git-reflog

## Digest

syntax

```
git reflog [subcommd] [options]
```

显示 reference logs，即 HEAD 指针指向的位置，如果没有指定 subcommand 默认表示 show

## reflog vs log

```
(base) cpl in /tmp/test on main ● ● performing a merge λ git log --all --format=oneline
85048ef70ea81dce9262b85886e45cc4e93053fa (topic) topic b
94715fe71a692892a2c896ec2a5fd59df42f0d09 (HEAD -> main) master b
038f625efdbe6d3f65ed9cc278d46eb8508ca0f1 d
6e5f3ac2c14c7c2c72ec36f7a9f91a8629597003 c
fa5e06f8631cbfba3bfa780e27934318131bba7f b
b38d959335ff69563c764aea7156952dcff002bc a

(base) cpl in /tmp/test on main ● ● performing a merge λ git reflog
94715fe (HEAD -> main) HEAD@{0}: checkout: moving from topic to main
85048ef (topic) HEAD@{1}: commit: topic b
038f625 HEAD@{2}: checkout: moving from main to topic
94715fe (HEAD -> main) HEAD@{3}: commit: master b
fa5e06f HEAD@{4}: checkout: moving from topic to main
038f625 HEAD@{5}: commit: d
6e5f3ac HEAD@{6}: checkout: moving from main to topic
fa5e06f HEAD@{7}: checkout: moving from topic to main
6e5f3ac HEAD@{8}: commit: c
fa5e06f HEAD@{9}: checkout: moving from main to topic
fa5e06f HEAD@{10}: checkout: moving from topic to main
fa5e06f HEAD@{11}: checkout: moving from main to topic
fa5e06f HEAD@{12}: commit: b
b38d959 HEAD@{13}: commit (initial): a
```

可以明显观察到 reflog 的记录比 log 多，因为记录的是 HEAD 移动的日志，所以也会包含 checkout

如果现在使用 reset 回退到 a 对应的 commit 记录，reflog 还会有记录吗

```
(base) cpl in /tmp/test on main λ git reflog
b38d959 (HEAD -> main) HEAD@{0}: reset: moving to b38d959335ff69563c764aea7156952dcff002bc
94715fe HEAD@{1}: checkout: moving from topic to main
85048ef (topic) HEAD@{2}: commit: topic b
038f625 HEAD@{3}: checkout: moving from main to topic
94715fe HEAD@{4}: commit: master b
fa5e06f HEAD@{5}: checkout: moving from topic to main
038f625 HEAD@{6}: commit: d
6e5f3ac HEAD@{7}: checkout: moving from main to topic
fa5e06f HEAD@{8}: checkout: moving from topic to main
6e5f3ac HEAD@{9}: commit: c
fa5e06f HEAD@{10}: checkout: moving from main to topic
fa5e06f HEAD@{11}: checkout: moving from topic to main
fa5e06f HEAD@{12}: checkout: moving from main to topic
fa5e06f HEAD@{13}: commit: b
b38d959 (HEAD -> main) HEAD@{14}: commit (initial): a
```

显然还是有的，但是也可以观察到新增了一条记录指向了 b38d959

```
(base) cpl in /tmp/test on main ● ● performing a merge λ git reset --hard b38d959335ff69563c764aea7156952dcff002bc
HEAD is now at b38d959 a
                                                                                                                                                         
(base) cpl in /tmp/test on main λ git log  --format=oneline 
b38d959335ff69563c764aea7156952dcff002bc (HEAD -> main) a
```

对比 `git log` 可以发现已经没有对应的 commit 的记录了，如果我们需要回退之前怎么办呢？

可以先利用 reflog 查询 commit 对应 HEAD  的记录

```
(base) cpl in /tmp/test on main λ git reset --hard 94715fe
HEAD is now at 94715fe master b
                                                                                                                                                         
(base) cpl in /tmp/test on main λ git log --format=oneline
94715fe71a692892a2c896ec2a5fd59df42f0d09 (HEAD -> main) master b
fa5e06f8631cbfba3bfa780e27934318131bba7f b
b38d959335ff69563c764aea7156952dcff002bc a
```

