# Linux chsh

## 概述

`chsh`命令用于修改login shell 

pattern：`chsh [options] [Login]`

login指代用户名

### 目录

- `/etc/passwd`

  存储用户的信息，包括用户当前使用的shell。使用chsh会修改该文件中的值

- `/etc/shells`

  存储当前可用的所有shell

## 参数

- `-s`

  指定用户使用的shell，必须是绝对路径
  
  ```
  chsh -s /bin/bash
  ```
  
  
