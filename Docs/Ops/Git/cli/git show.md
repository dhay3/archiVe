# git show

ref

https://git-scm.com/docs/git-show

https://stackoverflow.com/questions/610208/how-to-retrieve-a-single-file-from-a-specific-revision-in-git

## Digest

syntax

```
git show [options] [object...]
```

用于查看一个或者多个 object，object 可以是

1. blobs (简单的理解就是对应的文件)

   For plain blobs, it shows the plain contents.

2. trees

   For trees, it shows the names (equivalent to *git ls-tree* with --name-only).

3. tags

   For tags, it shows the tag message and the referenced objects.

4. commits

   For commits it shows the log message and textual diff. It also presents the merge commit in a special format as produced by *git diff-tree --cc*.

## blobs

假设当前状态如下

```
(base) cpl in /tmp/test on master λ echo a1 > a
(base) cpl in /tmp/test on master ● λ git add a
(base) cpl in /tmp/test on master ● λ git commit -m'a1'
[master (root-commit) f808028] a1
 1 file changed, 1 insertion(+)
 create mode 100644 a
(base) cpl in /tmp/test on master λ echo a2 > a
(base) cpl in /tmp/test on master λ git add a
(base) cpl in /tmp/test on master ● λ git commit -m 'a2'
[master 470c32d] a2
 1 file changed, 1 insertion(+), 1 deletion(-)
(base) cpl in /tmp/test on master λ echo a3 > a
(base) cpl in /tmp/test on master λ git add a
(base) cpl in /tmp/test on master ● λ git commit -m'a3'
[master 72faa15] a3
 1 file changed, 1 insertion(+), 1 deletion(-)
```

如果需要比较文件 a 和前一个版本中 a 的内容，可以使用如下命令

```
(base) cpl in /tmp/test on master λ git show a                          
commit 72faa1582f6b99af406ddc6a92b8fa4dec780aeb (HEAD -> master)
Author: 4liceh <4@liceh.com>
Date:   Wed Apr 5 23:36:25 2023 +0800

    a3

diff --git a/a b/a
index c1827f0..d616f73 100644
--- a/a
+++ b/a
@@ -1 +1 @@
-a2
+a3
```

那如果需要查看一个文件的历史内容呢, 可以使用如下格式

```
git show gitrevision:file
```

例如

```
(base) cpl in /tmp/test on master λ git show HEAD^^:a
a1
(base) cpl in /tmp/test on master λ git show 72faa15:a
a3
```