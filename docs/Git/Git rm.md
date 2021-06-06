# Git rm

## 概述

git rm 用于删除工作区(会从当前目录消失)和索引树中的文件。==只有commit之后才会生效==

pattern：`git rm [opitons] <files>`

## 参数

- 默认，==使用`git reset --heard <hash>`==可以还原

  ```
  82341@bash MINGW64 /f/test (master)
  $ git rm test
  rm 'test'
  
  82341@bash MINGW64 /f/test (master)
  $ git status
  On branch master
  Changes to be committed:
    (use "git restore --staged <file>..." to unstage)
          deleted:    test
  
  
  82341@bash MINGW64 /f/test (master)
  $ git commit -m"t"
  [master 95ed2d4] t
   1 file changed, 1 deletion(-)
   delete mode 100644 test
  
  82341@bash MINGW64 /f/test (master)
  $ git status
  On branch master
  nothing to commit, working tree clean
  
  ```

- `--cached`

  只从索引树中删除，不会对工作区产生影响。使用该参数可以起到`.gitignore`的作用

  ```
  82341@bash MINGW64 /d/workspace_for_idea/谷粒学院/guli_parent/common/common_util/src/main/java/com/chz/utils (master)
  $ git rm --cached AliyunUtils.java
  rm 'guli_parent/common/common_util/src/main/java/com/chz/utils/AliyunUtils.java'
  
  82341@bash MINGW64 /d/workspace_for_idea/谷粒学院/guli_parent/common/common_util/src/main/java/com/chz/utils (master)
  $ git status
  On branch master
  Your branch is behind 'origin/master' by 1 commit, and can be fast-forwarded.
    (use "git pull" to update your local branch)
  
  Changes to be committed:
    (use "git restore --staged <file>..." to unstage)
          deleted:    AliyunUtils.java
  
  Untracked files:
    (use "git add <file>..." to include in what will be committed)
          AliyunUtils.java
  
  ```

- `-r`

  递归

