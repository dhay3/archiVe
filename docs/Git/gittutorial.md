# gittutorial

git 可以从两种方式获取对应命令的帮助

1. `man git-log`
2. `git help log`

在使用git前需要设置username和useremail

```
           $ git config --global user.name "Your Name Comes Here"
           $ git config --global user.email you@yourdomain.example.com
```

## importing a new project

例如有一个`project.tar.gz`，需要纳入 git revsion controll

1. git init

   初始化working directory，会生成一个`.git`隐藏文件

2. git add .

   告诉git对当前目录生产snapshot，将会被存储在index（暂存区）

3. git commit 

   将index中的snap永久存储到repository（版本库）

## making changes

修该文件，然后将文件上传到index

1. git add file1 file2 file3

   添加新文件

2. git diff --cached

   对比文件，`--cached`只会对比index中文件的修改。如果不带`--cached`会比较working directory以及index 和 repository 中的不同

3. git status

   查看index中更新的文件

4. git commit

   对修改的文件进行提交，可以使用`git commit -a`提交(git add + git commitS)

## viewing project history

1. git log 

   版本序列 

2. git log -p

   修改过内容

3. git log --stat --summary

   修改过的行

## managing branches

1. git branch experimental

   新增一个branch

2. git branch

   查看所有的branch，及当前使用的分区

3. git switch experimental

   切换分区

4. git merge experimental

   将连个分区中不同的内容合并到当前的分区

5. git commit -a

   提交合并的内容

6. git branch -d experimental

   删除分区

## exploring history