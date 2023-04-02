ref
[https://git-scm.com/docs/git-clone](https://git-scm.com/docs/git-clone)
## Digest
syntax
```
git clone [options] <repository> [directory]
```
Clones a repository into a newly created direcotry, creates remote-tracking branches for each branch in the cloned repository ( 可以使用 `git branch --remotes` 来查看 )，and creates and checks out an initial branch that is forked from the cloned repository's currently active branch
git clone 时会从 remote repository ( Git URLs ) 克隆所有的内容到本地的 directory (如果没有指定 directory 时，目录名默认同克隆的 repository ) ，然后创建 init branch ( 一般是 master 或者是 main ) 并  checkout init branch。所以当使用 `git branch -v` 查看只有一个分支时，并不是只克隆一个分支，而是克隆所有的分支。可以使用 `git branch` 来校验
```
root@v2:/home/ubuntu# git clone https://gitee.com/mindspore/graphengine.git
Cloning into 'graphengine'...
remote: Enumerating objects: 42312, done.
remote: Counting objects: 100% (1118/1118), done.
remote: Compressing objects: 100% (600/600), done.
remote: Total 42312 (delta 590), reused 884 (delta 474), pack-reused 41194
Receiving objects: 100% (42312/42312), 14.68 MiB | 25.52 MiB/s, done.
Resolving deltas: 100% (32664/32664), done.
root@v2:/home/ubuntu# cd graphengine/
root@v2:/home/ubuntu/graphengine# git branch -v
* master 5c7fe0ef !2154 upgrade ascend software package 22 mar 23 Merge pull request !2154 from yanghaoran/master
root@v2:/home/ubuntu/graphengine# git branch --remote
  origin/HEAD -> origin/master
  origin/development
  origin/ffts_plus_dev
  origin/master
  origin/r0.1
  origin/r0.2
  origin/r0.3
  origin/r0.5
  origin/r0.6
  origin/r0.7
  origin/r1.0
  origin/r1.0.1
  origin/r1.1
  origin/r1.1.1
  origin/r1.2
  origin/r1.2.0
  origin/r1.3
  origin/r1.3.0
  origin/r1.5
  origin/r1.5.0
  origin/r1.6
  origin/r1.7
  origin/r1.8
  origin/r1.9
  origin/release
  origin/update_with_ascend_newest
```
## Optional args

- `-o | --origin <name>`

Instead of using the remote name `origin` to keep track of the upstream repository
追踪的 remote repository 以 name 命名而不是默认的 origin 

- `--no-tags`

don't clone any tags
不克隆 tags

- `--depth <depth>`

shallow clone with a history truncated to the specified number of commits
浅克隆，克隆时按照 depth 克隆 log 条目( 也不会克隆 tags )，默认会带上 `--single-branch`。如果不想克隆 remote repository 其他分支，同时不想保留多余的 commit log 可以使用这个参数，可以有效的减少下载量
```
git clone --depth 1  https://gitee.com/mindspore/graphengine.git gra
cd gra
git log --oneline
5c7fe0e (grafted, HEAD -> master, origin/master, origin/HEAD) !2154 upgrade ascend software package 22 mar 23 Merge pull request !2154 from yanghaoran/master
```
将 [https://gitee.com/mindspore/graphengine.git](https://gitee.com/mindspore/graphengine.git) 以浅克隆的方式 ( 只有最新一条的 log )到本地的 `gra` 目录

- `-b | --branch <name>`

从 remote repository 克隆所有的 branch，但是 create 和 check out 的分支是 name
```
root@v2:/home/ubuntu/gra# git clone -b release https://gitee.com/mindspore/graphengine.git
Cloning into 'graphengine'...
remote: Enumerating objects: 42312, done.
remote: Counting objects: 100% (1118/1118), done.
remote: Compressing objects: 100% (600/600), done.
remote: Total 42312 (delta 590), reused 884 (delta 474), pack-reused 41194
Receiving objects: 100% (42312/42312), 14.68 MiB | 5.78 MiB/s, done.
Resolving deltas: 100% (32664/32664), done.
root@v2:/home/ubuntu/gra# cd graphengine/
root@v2:/home/ubuntu/gra/graphengine# git br -v
* release ebd097ad !2151 headers for Ascend 7 Mar 23 Merge pull request !2151 from yanghaoran/release
```

- `--single-branch`

只克隆单分支 , 不追踪 remote-tracking branches，可以和 `-b` 一起使用指定分支名
```

root@v2:/home/ubuntu# git clone --single-branch https://gitee.com/mindspore/graphengine.git
Cloning into 'graphengine'...
remote: Enumerating objects: 36719, done.
remote: Counting objects: 100% (7264/7264), done.
remote: Compressing objects: 100% (1076/1076), done.
remote: Total 36719 (delta 6455), reused 6536 (delta 6153), pack-reused 29455
Receiving objects: 100% (36719/36719), 14.19 MiB | 24.72 MiB/s, done.
Resolving deltas: 100% (28459/28459), done.
root@v2:/home/ubuntu# cd graphengine/
root@v2:/home/ubuntu/graphengine# git branch -v
* master 5c7fe0ef !2154 upgrade ascend software package 22 mar 23 Merge pull request !2154 from yanghaoran/master
root@v2:/home/ubuntu/graphengine# git branch --remote
  origin/HEAD -> origin/master
  origin/master
```
## Git URLS
repository 的值可以是 Git URLS，例如

- ssh://[user@]host.xz[:port]/path/to/repo.git/
- git://host.xz[:port]/path/to/repo.git/
- http[s]://host.xz[:port]/path/to/repo.git/
- ftp[s]://host.xz[:port]/path/to/repo.git/
## Branches in clone
git clone 会 clone remote repository 上所有的 branches 吗？
```
root@v2:/home/ubuntu# https://gitee.com/mindspore/graphengine.git gra
bash: https://gitee.com/mindspore/graphengine.git: No such file or directory
root@v2:/home/ubuntu# git clone https://gitee.com/mindspore/graphengine.git gra
Cloning into 'gra'...
remote: Enumerating objects: 42312, done.
remote: Counting objects: 100% (1118/1118), done.
remote: Compressing objects: 100% (600/600), done.
remote: Total 42312 (delta 591), reused 884 (delta 474), pack-reused 41194
Receiving objects: 100% (42312/42312), 14.68 MiB | 5.87 MiB/s, done.
Resolving deltas: 100% (32665/32665), done.
root@v2:/home/ubuntu# cd gra/
154 from yanghaoran/master
```
这里可以看到只克隆了 master 分支，但是 remote repository 的其他 branches 并没有一起克隆下来
```
root@v2:/home/ubuntu/gra# git branch -v
* master 5c7fe0ef !2154 upgrade ascend software package 22 mar 23 Merge pull request !2
```
除了在 remote repository 的平台上分支，还可以使用 `git branch`来查看
```
root@v2:/home/ubuntu/gra# git branch --remote
  origin/HEAD -> origin/master
  origin/development
  origin/ffts_plus_dev
  origin/master
  origin/r0.1
  origin/r0.2
  origin/r0.3
  origin/r0.5
  origin/r0.6
  origin/r0.7
  origin/r1.0
  origin/r1.0.1
  origin/r1.1
  origin/r1.1.1
  origin/r1.2
  origin/r1.2.0
  origin/r1.3
  origin/r1.3.0
  origin/r1.5
  origin/r1.5.0
  origin/r1.6
  origin/r1.7
  origin/r1.8
  origin/r1.9
  origin/release
  origin/update_with_ascend_newest
```
