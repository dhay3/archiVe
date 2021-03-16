# awk 变量

## 内建变量

> 使用`$num`来获取特定列field的值，`$0`表示整行。也可以使用`$NF`来获取最后一列filed

- IGNORECASE

  正则是否忽略大小写，默认为0表示不忽略大小写。

- FILENAME

  awk当前输入的文件名，如果输入的不是文件将显示`-`

  ```
  [root@chz t]# awk '{print FILENAME}' ip.src
  ip.src
  
  [root@chz t]# ip a | awk '{print FILENAME}' 
  -
  ```

- FS

  分隔符

  ```
  [root@chz t]# ip a | awk -F '[:]' '{print FS}' 
  [:]
  [:]
  ```

- OFS

  输出的分隔符，我们可以使用`-v`参数来对原有的输出分隔符覆盖

  ```
  [root@chz t]# cat /etc/passwd | awk -v OFS=:  -F':' '{print $1,$2}'
  root:x
  bin:x
  daemon:x
  ```

- ORS

  一条记录的分隔符，默认CR，LF

  ```
  [root@chz t]# cat /etc/passwd | awk -F : -v 'OFS=#' -v 'ORS=\t' '{print $1,$2}'
  root#x	bin#x	daemon#x	adm#x	lp#x	sync#x	shutdown#x	halt#x
  ```

- ==NF==

  number of fields 

  ```
  [root@chz t]# cat /etc/passwd | awk -F':' '{print NF,$0}'
  7 root:x:0:0:root:/root:/bin/bash
  7 bin:x:1:1:bin:/bin:/sbin/nologin
  ```

- ==NR==

  number of records 当前输入的行

  ```
  [root@chz t]# cat /etc/passwd | awk -F':' '{print NR,$0}'
  1 root:x:0:0:root:/root:/bin/bash
  2 bin:x:1:1:bin:/bin:/sbin/nologin
  3 daemon:x:2:2:daemon:/sbin:/sbin/nologin
  ```

## 自定义变量

参考：

https://www.zsythink.net/archives/1374

**方法一：**

使用`-v var=value`，==这种方式可以直接引用shell中变量==

```
[root@chz t]# awk -v var=value 'BEGIN{print var}'
value
```

**方法二：**

```
[root@chz t]# awk 'BEGIN{var1="value1";var2="value2";print var1,var2}'
value1 value2
[root@chz t]# awk 'BEGIN{var1=value1;var2=value2;print var1,var2}'
 
```

==下面没有带双引号的输出空，因为读取value1的值默认为空==

如果是数字类型的，同C#一样直接引用

```
[root@chz t]# awk 'BEGIN{sum=0;for(i=0;i<=5;i++){sum+=i};print sum}'
15
```

==如果变量不存在，直接打印==

```
[root@chz opt]# ip a | awk '{print $i}'
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: ens33: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdis
```

awk无法取到外界的变量，需要通过`-v`参数传入。这里`$i`变量不存在，直接打印

```
[root@chz opt]# echo 1 > jobfile 
[root@chz opt]# for((i=0;i<3;i++));do cat jobfile | awk '{if(i<2)print i,$i}';done;
 1
 1
 1
```





