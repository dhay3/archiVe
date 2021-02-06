# Shell eval

参考：

https://blog.csdn.net/baidu_37964071/article/details/80930704

eval 关键字用于扫描commandline多次（会解析字符串中的commandline包括变量）

```
[root@cyberpelican opt]# a="cat test.sh"
[root@cyberpelican opt]# cat test.sh 
#!/bin/sh
* * * * * echo "test for crond"
[root@cyberpelican opt]# echo $a
cat test.sh
[root@cyberpelican opt]# eval $a
#!/bin/sh
* * * * * echo "test for crond"
[root@cyberpelican opt]# $a
error

```

## 特殊用法

使用eval可以作为map处理

```
[root@localhost ~]# cat file
NAME ZHANG
AGE 20
SEX NV
[root@localhost ~]# cat test.sh
#!/bin/bash

while read KEY VALUE
do
    eval "${KEY}=${VALUE}"
done < file
echo "$NAME $AGE $SEX"
[root@localhost ~]# ./test.sh
ZHANG 20 NV
```

