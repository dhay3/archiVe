# Shell shopt 

## 概述

`shopt`命令用来调整 Shell 的参数，跟`set`命令的作用很类似。之所以会有这两个类似命令的主要原因是，`set`是从 Ksh 继承的，属于 POSIX 规范的一部分，==而`shopt`是 Bash 特有的。==

## 参数

- `shopt`

  显示所有Bash的参数，可以指定参数名

  ```
  [root@cyberpelican /]# shopt 
  autocd         	off
  assoc_expand_once	off
  cdable_vars    	off
  cdspell        	off
  checkhash      	off
  ...
  
  [root@cyberpelican /]# shopt globstar
  globstar       	off
  ```

- `-q`

  不输出内容，用于查询是否启用变量，使用`$?`校验

  ```
  [root@cyberpelican /]# shopt globstar
  globstar       	off
  [root@cyberpelican /]# shopt -q globstar
  [root@cyberpelican /]# echo $?
  1
  ```

  如果参数没有启用返回1

- `-s`

  启用该参数

  ```
  [root@cyberpelican /]# shopt -s  globstar
  [root@cyberpelican /]# shopt   globstar
  globstar       	o
  ```

- `-u`

  关闭该参数

  ```
  [root@cyberpelican /]# shopt -u  globstar
  ```

  