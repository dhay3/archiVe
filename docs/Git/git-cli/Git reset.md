# Git reset

## 概述

`git reset`用于控制本地库的版本，==不会对工作区产生影响==。

**HEAD说明**

- HEAD表示当前版本
- HEAD^ 上一个版本
- HEAD^^ 上上一个版本，依次类推
- HEAD~0 表示当前版本
- HEAD~1 表示上一个版本，依次类推

## 常用参数

### `git reset [<mode>] [<commit>]`

- `--mixed`

  缺省值，不修改工作区，修改分支的HEAD的指针和暂存区

  ```
  82341@bash MINGW64 /f/test (master)
  $ git reflog
  2f0b0b7 (HEAD -> master) HEAD@{0}: reset: moving to 2f0b0b7a5aa3555dccdc045da9e5239888ac3605
  d6bd2fc HEAD@{1}: reset: moving to HEAD
  d6bd2fc HEAD@{2}: reset: moving to d6bd2fc
  2f0b0b7 (HEAD -> master) HEAD@{3}: reset: moving to HEAD^
  d6bd2fc HEAD@{4}: reset: moving to HEAD
  d6bd2fc HEAD@{5}: commit: add 3.txt
  2f0b0b7 (HEAD -> master) HEAD@{6}: commit (initial): init
  
  82341@bash MINGW64 /f/test (master)
  $ git reset d6bd2fc
  
  82341@bash MINGW64 /f/test (master)
  $ git status
  On branch master
  Untracked files:
    (use "git add <file>..." to include in what will be committed)
          debug.log
  
  nothing added to commit but untracked files present (use "git add" to track)
  
  ```

- `--soft`

  不修改暂存区(==包括当前和回退后暂存区中的内容==)和工作区，修改分支的HEAD的指针

  ```
  82341@bash MINGW64 /f/test (master)
  $ git status
  On branch master
  Untracked files:
    (use "git add <file>..." to include in what will be committed)
          debug.log
  
  nothing added to commit but untracked files present (use "git add" to track)
  
  82341@bash MINGW64 /f/test (master)
  $ git add debug.log
  warning: LF will be replaced by CRLF in debug.log.
  The file will have its original line endings in your working directory
  
  82341@bash MINGW64 /f/test (master)
  $ git status
  On branch master
  Changes to be committed:
    (use "git restore --staged <file>..." to unstage)
          new file:   debug.log
  82341@bash MINGW64 /f/test (master)
  $ git log
  commit d6bd2fc54ec3210f3b330237aedad777663985c9 (HEAD -> master)
  Author: cyberPelican <hostlockdown@gmail.com>
  Date:   Sat Dec 5 11:38:35 2020 +0800
  
      add 3.txt
  
  commit 2f0b0b7a5aa3555dccdc045da9e5239888ac3605
  Author: cyberPelican <hostlockdown@gmail.com>
  Date:   Sat Dec 5 11:36:58 2020 +0800
  
      init
  
  82341@bash MINGW64 /f/test (master)
  $ git reset --soft 2f0b0b7a5aa3555dccdc045da9e5239888ac3605
  
  82341@bash MINGW64 /f/test (master)
  $ git status
  On branch master
  Changes to be committed:
    (use "git restore --staged <file>..." to unstage)
          new file:   3.txt
          new file:   debug.log
  ```

- `--hard`

  修改工作区，暂存区，分支的HEAD的指针。如果文件没有`git add`或`git commit`，git不会修改它。

  ```
  82341@bash MINGW64 /f/test (master)
  $ git add debug.log
  warning: LF will be replaced by CRLF in debug.log.
  The file will have its original line endings in your working directory
  
  82341@bash MINGW64 /f/test (master)
  $ git status
  On branch master
  Changes to be committed:
    (use "git restore --staged <file>..." to unstage)
          new file:   debug.log
  
  
  82341@bash MINGW64 /f/test (master)
  $ git reflog
  d6bd2fc (HEAD -> master) HEAD@{0}: reset: moving to d6bd2fc
  2f0b0b7 HEAD@{1}: reset: moving to 2f0b0b7
  d6bd2fc (HEAD -> master) HEAD@{2}: reset: moving to d6bd2fc
  2f0b0b7 HEAD@{3}: reset: moving to 2f0b0b7a5aa3555dccdc045da9e5239888ac3605
  d6bd2fc (HEAD -> master) HEAD@{4}: reset: moving to HEAD
  d6bd2fc (HEAD -> master) HEAD@{5}: reset: moving to d6bd2fc
  2f0b0b7 HEAD@{6}: reset: moving to HEAD^
  d6bd2fc (HEAD -> master) HEAD@{7}: reset: moving to HEAD
  d6bd2fc (HEAD -> master) HEAD@{8}: commit: add 3.txt
  2f0b0b7 HEAD@{9}: commit (initial): init
  82341@bash MINGW64 /f/test (master)
  $ git reset  --hard 2f0b0b7
  HEAD is now at 2f0b0b7 init
  
  82341@bash MINGW64 /f/test (master)
  $ git status
  On branch master
  nothing to commit, working tree clean
  
  ```

  

