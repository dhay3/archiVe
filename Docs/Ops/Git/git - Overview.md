# git - Overview

## Digest

Git 是一个提供 Version Control (VC) 的工具。例如大名鼎鼎的 SVN, Perforce, Bazaar 都是 VC 工具

## Version Control

什么是 Version Control 呢？中文也被称为版本控制，顾名思义为文件提供版本控制。主要有 3 种模型

1. local version control systems - 基于本地的版本控制系统
2. centralized version control systems - 中心化的版本控制系统
3. distributed version control systems - 分布式的版本控制系统

### Local Version Control Systems

![Local version control diagram](https://git-scm.com/book/en/v2/images/local.png)

### Centrallized Version Control Systems

![Centralized version control diagram](https://git-scm.com/book/en/v2/images/centralized.png)

### Distributed Version Control Systems

![Distributed version control diagram](https://git-scm.com/book/en/v2/images/distributed.png)

## Git vs VCs

Git 和其他的 VCs 的区别在哪？为什么需要使用 Git？答案是速度以及效率。

一般的 VC 工具都是基于文件的 delta-base，而 Git 是基于 snapshots (快照) 的

如果有 3 个文件 A,B,C 作为 version 1，后面修改了 AC 文件中的内容，然后提交作为 version 2。version 2 会保存对比 version 1 增量内容

![Storing data as changes to a base version of each file](https://git-scm.com/book/en/v2/images/deltas.png)

但是 Git 不一样，Git 会将数据像文件系统快照一样管理

![Git stores data as snapshots of the project over time](https://git-scm.com/book/en/v2/images/snapshots.png)

## Main sections

Git 中会用到类似树结构来管理数据，每个子叶对应一个文件，都有自己的状态。分别是

- Untracked (实际不包含，为了方便记忆)

  the file have not been added

  不在 staging area 或者 repository 中的文件，即从未 `git add` 过的文件

- Staged

  means that you have marked a modified file in its current version to go into your next commit snapshot.

  文件已经暂存到 staging area，即 `git add` 后的文件

- Modified

  means that you have changed the file but have not committed it to your database yet.

  简单的理解就是文件已经 tracked，即 `git add` 或者 `git commit`后，修改了文件内容，且未 `git commit` 的文件

- Commited

  means that the data is safely stored in your local database.

  文件已经从 staging area 转移到 repository 中，即 `git commit` 后的文件

几种状态引出 3 块逻辑上的区域

![Index 1](https://git-scm.com/images/about/index1@2x.png)



### Working Directory/worktree

中文也叫做 工作区。是从 `.git` 目录 ( repository )中提取出来的文件对应某个版本，注意 working directory 并不等价于你用 `ls` 命令看到目录。假如一个文件没有被加入到 staging area 即 untracked 的状态，那么这个文件就不在 working directory 中 (因为文件并没有通过 `git commit` 的方式被加入版本库中)，`git reset --hard` 也就不会对该文件产生影响

### Staging Area/index

中文也叫做 暂存区。实际存储在 `.git` 中。保存了已经 `git add` 但是暂未 `git commit` 的文件

### Repository

中文也叫做 版本库。主要用于存储所有 commit 的记录以及对应的指针索引，通过索引可以回滚到对应 commit working directory 中的内容



**references**

1. https://git-scm.com/book/en/v2/Getting-Started-About-Version-Control

2. https://git-scm.com/book/en/v2/Getting-Started-What-is-Git%3F

3. https://git-scm.com/docs/git
4. https://git-scm.com/about/staging-area

