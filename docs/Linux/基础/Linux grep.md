# Linux grep

参考：

http://www.zsythink.net/archives/1733

## 概述

grep 命令将匹配规则的file或是标准stdout中的内容，以一行的形式输出

pattern：`grep [options] pattern [file]`

可以使用`*`做为文件名的wildcard

## 参数

- -E

  使用正则表达式匹配，等价于egrep

  ```
  [root@chz Desktop]# ls | egrep 'test*'
  test1
  test2
  [root@chz Desktop]# ls | grep  '^test*'
  test1
  test2
  ```

- -i

  忽略Pattern 和 file 或是 stdout中的大小写

  ```
  root@chz:/etc# grep -i debian passwd
  Debian-snmp:x:117:122::/var/lib/snmp:/bin/false
  debian-tor:x:133:141::/var/lib/tor:/bin/false
  ```

- -w

  使用word-regexp，匹配全词

  ```
  [root@chz etc]# grep -w post  passwd
  [root@chz etc]# grep post passwd
  postfix:x:89:89::/var/spool/postfix:/sbin/nologin
  ```

- -n

  显示内容的同时，显示行数

  ```
  [root@chz etc]# grep -n chz passwd
  47:chz:x:1000:1000:chz:/home/chz:/bin/bash
  ```

  可以通古vim的`set nu`来显示行数

- --color

  高亮显示，设置`/etc/bashrc`来默认使用

- -o

  不以行显示内容，只显示匹配的内容

  ```
  [root@cyberpelican \]# cat testfile 
    File: testfile
    a111
    22
    a111
    22
    2
    1
  [root@cyberpelican \]# grep -o 11 testfile 
  11
  11
  ```

- -An,-Bn,-Cn

  > 用于想要匹配的内容单独成行，且有不同的匹配。内容之间以`---`分隔

  -A在匹配内容输出n行，-B在匹配内容之前输出n行，-C是-B和-A的结合

  ```
  [root@chz etc]# grep -C1 root passwd
  root:x:0:0:root:/root:/bin/bash
  bin:x:1:1:bin:/bin:/sbin/nologin
  --
  mail:x:8:12:mail:/var/spool/mail:/sbin/nologin
  operator:x:11:0:operator:/root:/sbin/nologin
  games:x:12:100:games:/usr/games:/sbin/nologin
  ```

- -v

  查询不匹配规则的

  ```
  [root@chz etc]# grep -v root passwd
  ...输出的内容不包含root
  ```

- -w

  全词匹配

  ```
  [root@cyberpelican bin]# ll | grep -w ls
  lrwxrwxrwx root root        6 B  Tue Mar 31 07:06:41 2020 hwloc-ls ⇒ lstopo
  .rwxr-xr-x root root    135.6 KB Wed Aug  7 02:45:27 2019 ls
  ```

  







