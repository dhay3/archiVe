ref
[https://git-scm.com/docs/git-diff](https://git-scm.com/docs/git-diff)
## Digest
syntax
```
git diff [options] [path...]
```
显示 path 有文件  commits 之间的不同，或者是 commit 和 working directory 之间的不同，或者是和 staging area 的不同。如果没有指定 path，默认比较当前目录下所有的文件
例如

- `git diff [options] [path]`

比较 staging area 和 lastest commit 
```
git diff filename
```

- `git diff [options] --no-index <path> <path>`

直接比较 working direcotry 中的文件
```
git diff filename1 filename2
```

- `git diff [options] <commit> [path]`

比较指定 commit 和 working direcotry 
```
git diff HEAD filename
```

- `git diff [options] --cached [commit] [path]`

比较 staging area 和指定的 commit，如果没有指定 commit 默认 HEAD，表示 HEAD 指针指向的记录
```
git diff --cached HEAD filename
```

- `git diff [options] <commit1>..<commit2>`

对比 commit1 和 commit2 之间
```
git diff HEAD..eb0ec746d880aad712b11834d5a37b008ac2a960
```
## Optional args

- `-s | --no-patch`

suppress diff output，一般和 `git show`一起使用

- `-w | --ignore-all-space`

忽略对比空格数
## Diff Output
```
root@v2:/home/ubuntu/test# git diff --no-index file1 file2
diff --git a/file1 b/file2
index 4554b37..628d2a8 100644
--- a/file1
+++ b/file2
@@ -1,5 +1,4 @@
-1
-3
-5
+2
+4
 6
 8
```
例如上述的 `diff --git a/a b/b` 表示对比的文件分别是 file1 file2，别名分别为 a b
`--- a/file1` 表示为需要对比的源文件
`+++ b/file2` 表示对比文件
`@@ -1,5 +1,4 @@` 表示对比的内容；file1 为 1,5 闭区间，file2 为 1,4 闭区间
`-1 -3 -5` 表示在 file1 中存在，但是不在 file2 中存在，file1 需要删除后才能变成 file2
`+2 +4` 表示在不在 file1 中存在，但是在 file2 中存在，file1 需要添加后才能变成 file2
`6 8` 表示在 file1 file2 中共有的内容

