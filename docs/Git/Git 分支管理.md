# Git 分支管理

### 查看分支

使用`-v`参数，查看分支的具体合并情况

```
82341@bash MINGW64 /f/test (master)
$ git branch
* master
```

### 创建分支

创建的分支的内容使用当前分支的快照（内容相同）

```
82341@bash MINGW64 /f/test (master)
$ git branch test1

82341@bash MINGW64 /f/test (master)
$ git branch
* master
  test1
```

### 删除分支

```
82341@bash MINGW64 /f/test (master)
$ git branch -d test1
Deleted branch test1 (was 2f0b0b7).
```

### 切换分支

> 使用`git switch`可以达到同样的效果

工作区内容换为切换的分支内容(commit的内容)

```
82341@bash MINGW64 /f/test (master)
$ git checkout test1
Switched to branch 'test1'

82341@bash MINGW64 /f/test (test1)
$ git branch
  master
* test1
```

### 合并分支

将指定分支合并到当前分支(commit的内容)，==合并后分支并不会消失==

```
82341@bash MINGW64 /f/test (master)
$ git merge test1
Updating 9c2ffeb..3b8923c
Fast-forward
 3.txt | 1 +
 1 file changed, 1 insertion(+)
 create mode 100644 3.txt
```

