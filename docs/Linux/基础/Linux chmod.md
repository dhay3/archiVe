# Linux chmod

## 概述

Linux chmod（英文全拼：change mode）命令是控制用户对文件的权限的命令

Linux/Unix 的文件调用权限分为三级 : 文件所有者（Owner）、用户组（Group）、其它用户（Other Users）。

<img src="https://www.runoob.com/wp-content/uploads/2014/08/file-permissions-rwx.jpg"/>

只有文件所有者和超级用户可以修改文件或目录的权限。可以使用绝对模式（八进制数字模式），符号模式指定文件的权限。

<img src="https://www.runoob.com/wp-content/uploads/2014/08/rwx-standard-unix-permission-bits.png"/>

## pattern

> 使用`-R`递归

`chmod [ugoa...][[+-=][rwxX]...][,...]`

- ugoa

  ==如果没有给定ugoa默认为a==

  1. u：the user who owns it
  2. g：other users in the file's group 
  3. o：other users not in the file's group
  4. a：all users

- operator
  1. +：增加权限
  2. -：去除权限
  3. =：设置指定用户权限的设置，即将用户类型的所有权限重新设置

- permission
  1. r：读，4
  2. w：写，2
  3. x：执行，1

## 例子

1. 将文件 file1.txt 设为所有人皆可读取 

   ```
   chmod a+r file1.txt
   ```

2. 将目前目录下的所有文件与子目录皆设为任何人可读取 

   ```
   chmod -R a+r *
   ```

3. 此外chmod也可以用数字来表示权限如 

   ```
   chmod 777 file
   ```

   

