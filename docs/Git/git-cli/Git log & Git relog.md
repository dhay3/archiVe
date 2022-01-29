# Git log & Git relog

参考：

https://blog.csdn.net/Qidi_Huang/article/details/53839591

file:///D:/git/Git/mingw64/share/doc/git-doc/git-log.html

## Git log

参考：https://www.cnblogs.com/wancy86/p/5848028.html

用来查看git commit 的日志（不能查看已被删除的commit记录）。默认使用`--no-decorate`参数

```
82341@bash MINGW64 /d/asset/note (master)
$ git log
commit 48d6272e498e2120be8893d3970aa653c97c20fd (HEAD -> master)
Author: cyberPelican <hostlockdown@gmail.com>
Date:   Wed Dec 2 11:31:21 2020 +0800

    init content
```

### 常用参数

- `--pretty=<format>`

  ```
  82341@bash MINGW64 /d/asset/note (master)
  $ git log --pretty=oneline
  48d6272e498e2120be8893d3970aa653c97c20fd (HEAD -> master) init content
  ```

- `--graph`

  以拓扑图的形式显示，分支的出现和合并

  ```
  *   d5e9fc2 (HEAD -> master) Merge branch 'change_site'
  |\  
  | * 7774248 (change_site) changed the runoob.php
  * | c68142b 修改代码
  |/  
  * c1501a2 removed test.txt、add runoob.php
  * 3e92c19 add test.txt
  * 3b58100 第一次版本提交
  ```

## Git relog

与`git log`不同的是，该命令会==记录commit和reset的日志信息（所以会有多个HEAD指针）==，用于恢复数据

```
82341@bash MINGW64 /d/asset/note (master)
$ git reflog
48d6272 (HEAD -> master) HEAD@{0}: reset: moving to 48d6272
1adf54d HEAD@{1}: commit: init
a5a488f HEAD@{2}: commit: init
302663d HEAD@{3}: commit: init
3e7bd39 (origin/master) HEAD@{4}: commit: content
40dce09 HEAD@{5}: commit: nginx init
48d6272 (HEAD -> master) HEAD@{6}: checkout: moving from Linux to master
48d6272 (HEAD -> master) HEAD@{7}: checkout: moving from master to Linux
48d6272 (HEAD -> master) HEAD@{8}: commit (initial): init content

```



