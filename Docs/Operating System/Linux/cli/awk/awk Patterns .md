# Patterns

awk Patterns 遵循如下几条规则

- awk先匹配或执行Patterns，然后执行action
- awk可以没有patterns或action，但是必须要有两者之一


## 模式

### BEGIN & END

BEGIN和END是两个特殊的patterns不会对输入的内容检查。且不管如何都会运行(action出错，END同样还是会运行)。==BEGIN通常用作初始化全局变量，END通常用作输出==

```
[root@chz opt]# awk 'BEGIN{print "123"}'
123

----

[root@chz opt]# df -hT | awk 'BEGIN{print "===start df==="}{printf "%s\n",$3}END{print "===done==="}'
===start df===
Size
470M
487M
487M
487M
17G
1014M
98M
4.4G
===done===
```

==BEGIN不能获取到内置的变量`$0`-`$..`，但是END可以==

```
root in /usr/local/\ λ df -hT | awk 'BEGIN{print "begin",$0,"\n"}{print $0}END{print "end",$0,"\n"}'
begin

Filesystem     Type      Size  Used Avail Use% Mounted on
udev           devtmpfs  2.0G     0  2.0G   0% /dev
tmpfs          tmpfs     395M  6.0M  389M   2% /run
/dev/vda1      ext4       40G  5.9G   32G  16% /
tmpfs          tmpfs     2.0G     0  2.0G   0% /dev/shm
tmpfs          tmpfs     5.0M     0  5.0M   0% /run/lock
tmpfs          tmpfs     2.0G     0  2.0G   0% /sys/fs/cgroup
tmpfs          tmpfs     395M     0  395M   0% /run/user/0
end tmpfs          tmpfs     395M     0  395M   0% /run/user/0
```

### BEGINFILE & ENDFILE



### /regular expression/

先对输入进行正则匹配(与高级编程语言中的正则不同)，正则规则具体查看`man /\<Regular Expressions\>`

需要注意的是，如果想要匹配`/`要按照正则的规则转义为`//`

```
[root@chz Desktop]# awk '/chz/{print $1}' /etc/passwd
chz:x:1000:1000:chz:/home/chz:/bin/bash

----

[root@chz t]# ip a | awk '/192\.168\.[0-9]{1,3}\.[0-9]{1,3}/{print$0}'
    inet 192.168.80.140/24 brd 192.168.80.255 scope global noprefixroute dynamic ens33
    inet 192.168.122.1/24 brd 192.168.122.255 scope global virbr0
```

默认使用POSIX regular expression和GNU regular expression，如果指定`--posix`参数，awk只使用POSIX regular expression。如下规则将不被使用

```
\y, \B, \<, \>, \s, \S, \w, \W, \`, and \'
```

## 组合

### relational expression

awk可以使用逻辑运算符`&&`，`||`，`!`，三目运算来连接pattern

```
[root@chz sda]# lsblk | awk 'NF<8 && NF >6{print $7}'
MOUNTPOINT
/boot
/
[SWAP]

----

[root@chz sda]# lsblk | awk 'NF==(NF>0?7:0){print $7}'
MOUNTPOINT
/boot
/
[SWAP]
```

打印特定行的内容

```
[root@chz t]# df -hT | awk 'NR>=2 && NR <=4 {print $0}'
devtmpfs                devtmpfs  470M     0  470M   0% /dev
tmpfs                   tmpfs     487M     0  487M   0% /dev/shm
tmpfs                   tmpfs     487M  8.6M  478M   2% /run
```

### pattern1,pattern2

这种方式会将从第一次匹配pattern1到第一次匹配pattern2之间的所有内容匹配

```
[root@chz t]# df -hT
Filesystem              Type      Size  Used Avail Use% Mounted on
devtmpfs                devtmpfs  470M     0  470M   0% /dev
tmpfs                   tmpfs     487M     0  487M   0% /dev/shm
tmpfs                   tmpfs     487M  8.6M  478M   2% /run
tmpfs                   tmpfs     487M     0  487M   0% /sys/fs/cgroup
/dev/mapper/centos-root xfs        17G  4.7G   13G  28% /
/dev/sda1               xfs      1014M  172M  843M  17% /boot
tmpfs                   tmpfs      98M   28K   98M   1% /run/user/0
/dev/sr0                iso9660   4.4G  4.4G     0 100% /run/media/root/CentOS 7 x86_64
-----

[root@chz t]# df -hT | awk '/devtmpfs/,/xfs/{print $0}'
devtmpfs                devtmpfs  470M     0  470M   0% /dev
tmpfs                   tmpfs     487M     0  487M   0% /dev/shm
tmpfs                   tmpfs     487M  8.6M  478M   2% /run
tmpfs                   tmpfs     487M     0  487M   0% /sys/fs/cgroup
/dev/mapper/centos-root xfs        17G  4.7G   13G  28% /
```

注意awk是以CR LF来标识，所以这里会输入两段

```
[root@k8smaster kubernetes]# awk '/^\s+client-certificate-data:/,/^\s+client-key-data:/ {print $0}' admin.conf 
    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURFekNDQWZ1Z0F3SUJBZ0lJY3lDVG41emJZaXN3RFFZSktvWklodmNOQVFFT...
    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBemVOMlZ3RVVkMG5DYkJEY3JOSVg4Zmw2WXJaeE9tT3g2T1R2UW5xU2Y5TzJjaHBvCiszRGI4ZG5xeVNVc0pML3VLRkp3Q28reXRiR1Z2bnF4MzdiVnFobUgzck9ZUlVXZHZUWjVJc2...
```

