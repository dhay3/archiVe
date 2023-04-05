# git

ref

https://git-scm.com/book/en/v2/Getting-Started-About-Version-Control

https://git-scm.com/book/en/v2/Getting-Started-What-is-Git%3F

https://git-scm.com/docs/git

## VC

在谈 Git 前需要了解一下 VC。所以，什么是 Version Control 呢？ VC 中文也被称为版本控制，顾名思义为文件提供版本控制。主要有 3 种模型

- Local Version Control Systems

  ![Local version control diagram](https://git-scm.com/book/en/v2/images/local.png)

- Centrallized Version Control Systems

  ![Centralized version control diagram](https://git-scm.com/book/en/v2/images/centralized.png)

- Distributed Version Control Systems

  ![Distributed version control diagram](https://git-scm.com/book/en/v2/images/distributed.png)

## Git

Git 是一个提供 Version Control (VC) 功能的工具，例如大名鼎鼎的 SVN, Perforce, Bazaar 都是 VC 工具。所以，为什么需要使用 Git 呢？答案是效率以及速度。

一般的 VC 工具都是基于文件的 delta-base

如果有 3 个文件 A,B,C 作为 version 1，后面修改了 AC 文件中的内容，然后提交作为 version 2。version 2 会保存对比 version 1 增量内容

![Storing data as changes to a base version of each file](https://git-scm.com/book/en/v2/images/deltas.png)

但是 Git 不一样，Git 会将数据像文件系统快照一样管理

![Git stores data as snapshots of the project over time](https://git-scm.com/book/en/v2/images/snapshots.png)

### Tree states

在 Git 中会使用类似树结构来管理数据，每个子叶都有各自的状态

分别是一下 3 种状态

- Modified

  means that you have changed the file but have not committed it to your database yet.

  简单的理解就是文件已经 tracked，但是没有 commit，即已经 `git add` 且有一次 `git commit` 后，未 `git add` 的文件

- Staged

  means that you have marked a modified file in its current version to go into your next commit snapshot.

  文件以及暂存到 staging area，即 `git add` 后的文件

- commited

  means that the data is safely stored in your local database.

  文件以及从 staging area 转移到 xxxx，即 `git commit` 后的文件

3 种状态引出有 3 块逻辑上区域

![Working tree, staging area, and Git directory](https://git-scm.com/book/en/v2/images/areas.png)

- Working Directory 工作区

  是从 `.git` 目录中提取出来的文件对应某个版本，注意 working directory 并不等价于你用 `ls` 命令看到目录。假如一个文件没有被假如到 staging area 即 unstaged 的状态，那么这个文件就不在 working directory 中 (因为文件并没有通过 `git commit` 的方式被加入版本库中)，`git reset --hard` 也就不会对该文件产生影响

- Staging Area 暂存区

  也被称为 index，实际存储在 `.git` 中。保存了已经 `git add` 但是暂未 `git commit` 的文件

- Git directory

  也被称为 Repository，用于存储项目的 metadata。包含所有版本信息的指针索引( 版本库 )

![Snipaste_2020-12-05_10-49-32](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/Snipaste_2020-12-05_10-49-32.5g6xqy7mx0s0.png)

## Optional args

- `-v | --version`

  查看版本信息，等价于 `git version`

- `-h | --help`

  查看帮助信息，等价于 `git help`

- `-p | --paginate`

  pipe all output into `less`

  将输出内容通过管道符到 less

- `-P | --no-pager`

  do not pipe Git output inot a pager