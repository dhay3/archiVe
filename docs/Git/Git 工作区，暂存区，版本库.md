# Git 工作区，暂存区，版本库

参考：

https://www.runoob.com/git/git-workspace-index-repo.html

https://git-scm.com/book/zh/v2/%E8%B5%B7%E6%AD%A5-Git-%E6%98%AF%E4%BB%80%E4%B9%88%EF%BC%9F

> ==Git只会管理使用了Git命令的文件==
>
> Git有三种状态:
>
> 1. committed(已提交)，已经提交到本地数据库
> 2. modified(已修改)，文件已修改，但是还没保存到本地数据库
> 3. staged(已暂存)，对一个已修改的文件标记，在下次提交的快照中

![Snipaste_2020-12-05_11-20-35](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/Snipaste_2020-12-05_11-20-35.68iq8zdcro40.png)

- working directory 工作区

  即本地磁盘上的目录(==如果删除了文件只是不显示，但是还是存储在磁盘上==)

- staging area / index 暂存区

  是一个文件，保存了下次要提交的文件列表信息(快照)

- repository 仓库目录(版本库)

  `.git`文件，保存所有的版本信息的指针索引，用于数据的前进和回退

![Snipaste_2020-12-05_10-49-32](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/Snipaste_2020-12-05_10-49-32.5g6xqy7mx0s0.png)

==上图的objects表示Git的对象库(本地数据库与)实际位于 `.git/objects` 目录下，里面包含了创建的各种对象及内容==

HEAD是一个指向版本的指针，通过HEAD指向master的不同分支控制，版本的前进与后退

对工作区文件修改或新增执行`git add`，暂存区目录树被更新。同时工作区修改或新增的文件内容被写入到对象库中的一个新对象中，而该对象的ID被记录在暂存区中的文件索引中。

当执行`git commit`时，暂存区的目录树写到对象库中，master指向的分支会做相应的更新。即 master 指向的目录树就是提交时暂存区的目录树。

当执行`git reset HEAD`时，暂存区的目录树会被重写(HEAD指针回退一个版本)，被 master 分支指向的目录树所替换，==但是工作区不受影响==。

当执行`git rm --cached <file>`时，会直接从暂存区删除文件，工作区不做变化

当执行`git checkout `替换分支时，==会用暂存区指定的文件替换工作区的文件，这个命令也是极具危险性的，因为不但会清除工作区中未提交的改动，也会清除暂存区中未提交的改动。==。

