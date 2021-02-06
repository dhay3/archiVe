# Git restore

## 概述

git restore 用于撤销操作

pattern：`git restore [<options>] [--source=<tree>] [--staged] [--worktree] [--] <pathspec>`

## 常用参数

- `--worktree`

  如果不使用`staged`，默认使用`worktree`。将在工作空间但是不在暂存区的文件撤销更改

  ```
  82341@bash MINGW64 /f/test (master)
  $ cat 4.txt
  world
  
  82341@bash MINGW64 /f/test (master)
  $ vim 4.txt
  
  82341@bash MINGW64 /f/test (master)
  $ cat 4.txt
  hello
  
  82341@bash MINGW64 /f/test (master)
  $ git status
  On branch master
  Changes not staged for commit:
    (use "git add <file>..." to update what will be committed)
    (use "git restore <file>..." to discard changes in working directory)
          modified:   4.txt
  
  no changes added to commit (use "git add" and/or "git commit -a")
  
  82341@bash MINGW64 /f/test (master)
  $ git restore 4.txt
  
  82341@bash MINGW64 /f/test (master)
  $ cat 4.txt
  world
  ```

- `--staged`

  从暂存区撤销指定文件

  ```
  82341@bash MINGW64 /f/test (master)
  $ git add 4.txt
  warning: LF will be replaced by CRLF in 4.txt.
  The file will have its original line endings in your working directory
  
  82341@bash MINGW64 /f/test (master)
  $ git restore --staged 4.txt
  
  82341@bash MINGW64 /f/test (master)
  $ git status
  On branch master
  Untracked files:
    (use "git add <file>..." to include in what will be committed)
          4.txt
  
  nothing added to commit but untracked files present (use "git add" to track)
  ```

  
