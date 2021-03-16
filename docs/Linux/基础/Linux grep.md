# Linux grep

参考：

http://www.zsythink.net/archives/1733

## 概述

grep 命令将匹配规则的file或是标准stdout中的内容，以一行的形式输出

pattern：`grep [options] pattern [file]`

可以使用`*`做为文件名的wildcard

## 参数

- -E

  使用extended regular，等价于egrep，对一些特殊的字符无需添加转义符

  ```
  root in /opt λ ls | grep -E "backup-[[:digit:]]{4}_[[:digit:]]{2}_[[:digit:]]{2}"
  backup-2021_03_01.zip
  backup-2021_03_010.zip
  backup-2021_03_02.zip
  backup-2021_03_03.zip
  backup-2021_03_04.zip
  backup-2021_03_05.zip
  backup-2021_03_06.zip
  backup-2021_03_07.zip
  backup-2021_03_08.zip
  backup-2021_03_09.zip
  ```

- -e

  使用basic regular

  ```
  root in /opt λ ls | grep -e "backup-[[:digit:]]\{4\}_[[:digit:]]\{2\}"
  backup-2021_03_01.zip
  backup-2021_03_010.zip
  backup-2021_03_02.zip
  backup-2021_03_03.zip
  backup-2021_03_04.zip
  backup-2021_03_05.zip
  backup-2021_03_06.zip
  backup-2021_03_07.zip
  backup-2021_03_08.zip
  backup-2021_03_09.zip
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

- -F 匹配固定

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
  
- -q

  用于检查是否有匹配项，如果有匹配项返回0

  ```
  root in /etc λ ll | grep -q host 
  root in /etc λ echo $?
  0
  root in /etc λ ll | grep -q aaaaaaaa
  root in /etc λ echo $?
  1
  ```

## Regular Expressions

> man grep -> /regural expressions
>
> 没有`\s`和`\d`

grep能识别三种不能的正则语法：Basic(BRE)，extended(ERE)，perl(PCRE)

在GUN grep中Basic和extended没有什么区别。如果是其他的grep

```
Basic vs Extended Regular Expressions
       In basic regular expressions the meta-characters ?, +, {, |, (, and ) lose their special
       meaning; instead use the backslashed versions \?, \+, \{, \|, \(, and \).
```

如果是extended不需要添加

`\<`表示匹配前置的空格，`\>`表示匹配后置的空格

