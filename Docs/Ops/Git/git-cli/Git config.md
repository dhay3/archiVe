# Git config

参考：https://www.atlassian.com/git/tutorials/setting-up-a-repository/git-config

> 如果git 使用 `-c`，置顶配置就会被强制替换，例如：
>
> `git -c foo=bar`

## 概述

git config 用于查询和设置git的配置文件。git的配置文件有三个等级

1. `--local`

   本地级别，文件在`.git/config`，缺省值

2. `--global`

   针对某一用户的全局级别，文件在`~/.gitconfig`

3. `--system`

   针对所有用户的系统级别

## 查询

- `git config key`

  ```
  82341@bash MINGW64 /d/asset/note/docs/Git (master)
  $ git config user.email
  hostlockdown@gmail.com
  ```

## 设置

- `git config [level] <key> <vale>`

  ```
  git config --global user.email "your_email@example.com"
  ```

## 可配置参数

> https://git-scm.com/docs/git-config#_variables
>
> 所有可用的参数如上
>
> 使用`git help command`查看具体的命令用法

- core.editor

  有一些命令会让git打开文本编辑器，设定默认的文本编辑器。可选值，atom，emacs，vim，nano，subl(sublimetext)

- alias.user_alias

  为一些常的命令设置快捷命令

  ```
  git config --global alias.ci commit
  ```

  这里为git commit设置了快捷键ci，使用`git ci`调用

- http.proxy

  配置http和https的代理

  ```
  git config --global http.proxy 'socks5://127.0.0.1:1080' 
  ```

  

  



