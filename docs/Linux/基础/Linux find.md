# Linux find

## 概述

pattern：`find [path] [option] [action]`

find命令直接搜索硬盘速度较慢

## 基础条件

- `-name pattern`

  匹配文件名，可以使用shell的模式扩展。==如果没有使用模式扩展表示精确匹配==

  ```
  root in /usr/local/\/test1 λ find . -name "*2020-*"
  ./error.xin.2020-12-01.log.gz  
  ```

- `-perm pattern`

  匹配有指定权限的文件

  ```
  find . -perm 755 
  ```

- `-user username`

  按照用户的名字来查找文件

  ```
  find ~ -user sam
  ```

- `-regex pattern`

  使用正则表达式来匹配文件，如果使用了regex

- `-type c`

  1. b：块文件
  2. c：字符设备文件
  3. d：目录
  4. p：具名管道符
  5. f：普通文件
  6. l：链接文件
  7. s：socket文件

  

## 时间条件

- `-amin n `

  匹配被n分钟之前查看过的文件

- `-atime n`

  匹配被n*24小时之前查看过的文件

- `-cmin n`

  匹配n分钟之前文件状态被修改的文件

- `-ctime n`

  匹配n*24小时之前文件状态被修改的文件

- `-mmin`

  匹配n分钟之前被修改的文件

- `-mtime`

  匹配n*24小时之前被修改的文件

## 动作

- `-delete`

  找到文件后删除

- `-exec`

  找到文件后执行

- `-ok command`

- `-print`

  将匹配后的文件输出，默认动作

## 案例

1. ```
   find /tmp -name core -type f -print | xargs /bin/rm -f
   ```

