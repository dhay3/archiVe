# Linux 文本处理工具

[TOC]

## tr

替换

```
root in /sys/block/vda/mq/0 λ echo "123" | tr 2 4
143
```

删除

```
root in /sys/block/vda/mq/0 λ echo "123" | tr -d 2
13
```

## awk

参考：

https://www.ruanyifeng.com/blog/2018/11/awk.html

> gawk是gnu开源的，==这里的awk就是gawk==

文本处理工具，==每次处理一行==

**pattern**

`awk '{[pattern]action}' {filname}`, pattern采用regex，和正则不同的是如果当前处理的行包含regex就会返回当前行

> 和shell脚本中的`$0,$1...`不同，awk中`$0`表示当前行，`$1`表示以制表符或是换行符（CR或是FL）分隔的第一个字段，例如：this is a test，`$0`表示this is a test，`$1`表示 this。通过`echo 'this is a test'|awk '{print $1}'`来校验

**常用参数**

- `-F`

  指定分隔符，默认以空格为separator

  `echo 'this is a test'|awk -F: '{print $1}'`输出this is a test，因为没有以`:`为分隔符

  `awk -F: '{print $1} /etc/passwd'`
  
  如果需要表明多个分隔符使用`[]`
  
  https://stackoverflow.com/questions/12204192/using-multiple-delimiters-in-awk
  
  ```
  /logs/tc0001/tomcat/tomcat7.1/conf/catalina.properties:app.env.server.name = demo.example.com
  /logs/tc0001/tomcat/tomcat7.2/conf/catalina.properties:app.env.server.name = quest.example.com
  /logs/tc0001/tomcat/tomcat7.5/conf/catalina.properties:app.env.server.name = www.example.com
  
  awk -F'[/=]' '{print $3 "\t" $5 "\t" $8}' file
  
  tc0001   tomcat7.1    demo.example.com  
  tc0001   tomcat7.2    quest.example.com  
  tc0001   tomcat7.5    www.example.com
  ```

**内建变量/built-in variables**

- NF

  表示所有的字段数

  `echo 'this is a test'|awk '{print $NF}'`所以这里就是打印出最后一个字段test，`$(NF-1)`表示倒数第二个字段

- NR

  获取行数

  ```
  awk -F: '{print NR")" $1}' /etc/passwd
  ...
  41)pcp
  42)sshd
  43)avahi
  44)postfix
  45)oprofile
  46)tcpdump
  47)chz
  ...
  ```

**Regex**

> 正则表达式需要在`//`之间，使用`~//`表示非运算
>
> ip a | awk '/inet /{print $2}'输出IP

```
[root@chz Desktop]# awk '/chz/{print $1}' /etc/passwd
chz:x:1000:1000:chz:/home/chz:/bin/bash
```

这里会匹配`/etc/passwd`包含chz的行，并打印出第一个字段

**BEGIN/END**

BEGIN block 会在action之前处理，END block会在action之后处理

```
[root@chz Desktop]# awk 'BEGIN{print "begin block"}/^chz/{print $0}END{print "end block"}' /etc/passwd
begin block
chz:x:1000:1000:chz:/home/chz:/bin/bash
end block
```

**printf**

如果要在awk中使用printf需要将参数分割

```
ps auxf | grep -v ps |awk '{printf "%s %s \n",$2,$8}' | grep Z | awk '{print $1}' | xargs kill -SIGCHLD

---

[root@chz opt]# blkid /dev/vg1/lv01 | awk -F '[= ]' '{printf "UUID=%s \t %s /opt\t defaults \t 0 \t 0\n",$3,$5 }' >> /etc/fstab
[root@chz opt]# cat /etc/fstab

#
# /etc/fstab
# Created by anaconda on Mon Aug 24 07:49:09 2020
#
# Accessible filesystems, by reference, are maintained under '/dev/disk'
# See man pages fstab(5), findfs(8), mount(8) and/or blkid(8) for more info
#
/dev/mapper/centos-root /                       xfs     defaults        0 0
UUID=52ff1027-e9d7-427d-9f43-3a98ba708796 /boot                   xfs     defaults        0 0
/dev/mapper/centos-swap swap                    swap    defaults        0 0
UUID="95a68178-1d5b-46aa-93fe-230c5664e1c3" 	 "ext4" 	 defualts 	 0 	 0
```

### 求和

https://blog.csdn.net/csCrazybing/article/details/52594989

## cut

截取文本

- `-d`

  指定dleimiter， ==默认以制表符为分隔==

- `-f`

  指定截取的字段序号，从1开始

example 1

```
[root@chz Desktop]# cat test
1 a hello
2 b banna
3 c cat
4 d dog
[root@chz Desktop]# cut -f 1 -d " " test 
1
2
3
4
```

example 2

> 注意grep是截取一行，支持正则

```
[root@chz Desktop]# cat test|grep dog
4 d dog
[root@chz Desktop]# cat test|grep dog|cut -d ' ' -f 2 
d
```

## sed

> 删除所有注解` sed -i '/^[[:space:]]*#/'d nginx.conf`
>
> 所有操作都可以使用`/regex/`来匹配

流式编辑器，不会对文本中的实际内容生效

#**pattern** 

`sed [options] [command...] [input-file]`

#**commad**

> ```
> [root@chz Desktop]# cat test
> 1 a hello
> 2 b banna
> 3 c cat
> 4 d dog
> ```

- `a\`append

  在当前行的下一行新增字符串

  ```
  [root@chz Desktop]# sed a\ok test
  1 a hello
  ok
  2 b banna
  ok
  3 c cat
  ok
  4 d dog
  ok
  ```

- `i\`insert

  在当前行的上一行新增字符串

  ```
  [root@chz Desktop]# sed i\ok test
  ok
  1 a hello
  ok
  2 b banna
  ok
  3 c cat
  ok
  4 d dog
  ```

  可以在指定位置执行操作，这里的backslash可以替换成空格

  ```
  [root@chz Desktop]# sed '2i\tea' test
  1 a hello
  tea
  2 b banna
  3 c cat
  4 d dog
  ```

  使用正则

  ```
  [root@cyberpelican opt]# sed '/^hello/i\ni hao' test 
  ni hao
  hello world
  ni hao
  hello 
  hey
  hi 
  
  #删除最后一行
  root@VM-0-4-centos rc.d]# sed -i $\d  ~/.bashrc
  
  ```

- `c\`replace

  替换所有的内容

  ```
  [root@chz Desktop]# sed c\ok test
  ok
  ok
  ok
  ok
  ```

  将[2,4]内容替换为tea

  ```
  [root@chz Desktop]# sed '2,4c\a cup of tea' test
  1 a hello
  a cup of tea
  [root@chz Desktop]# 
  ```

- `d`delete

  删除指定行

  ```
  [root@chz Desktop]# sed 2,4d  test
  1 a hello
  ```

  这里删除[2,4]行

- `p`只做打印操作

  p操作会先打印出所有，然后再打印出匹配的内容, 可以使用`-n`参数来忽略自动打印的内容

  ```
  [root@chz Desktop]# nl /etc/passwd|sed -n '/chz/p'
      47	chz:x:1000:1000:chz:/home/chz:/bin/bash
  ```

- `s/old/new/`

  使用正则将old替换为new

  ```
  [root@chz Desktop]# sed 's/dog/pig/' test
  1 a hello
  2 b banna
  3 c cat
  4 d pig
  ```

  > 通过` -e` 参数来使用多条 command expression，如果使用` -i `参数会对源文件生效，可以指定suffix

  ```
  [root@chz Desktop]# sed -e '1a\sed' -e 's/dog/pig/' test
  1 a hello
  sed
  2 b banna
  3 c cat
  4 d pig
  -------------------
  [root@chz Desktop]# sed -i .bak 's/i/t/' test
  [root@chz Desktop]# cat test.bak 
  1 a hello
  2 b banna
  3 c cai
  4 d dog
  [root@chz Desktop]# cat test
  1 a hello
  2 b banna
  3 c cat
  4 d dog
  [root@chz Desktop]# 
  
  ```

==如果需要替换的是路径，带有斜杠需要使用转义符==

```
[root@chz opt]# cat t
/a/b/c
[root@chz opt]# sed 's/\/a\/b\/c/\/abc/' t
/abc

```

#**特殊案例**

> 区别于s//中正则的`$`，在其他command中`$`表示最后一行

```
[root@chz Desktop]# sed '$a runoob' test
1 a hello
2 b banna
3 c cat
4 d dog
runoob
```

