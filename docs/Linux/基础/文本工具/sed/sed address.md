# sed address

参考：

https://www.gnu.org/software/sed/manual/sed.html#Addresses-overview

address指定command的作用域，如果没有指定address默认对所有行生效

```
[root@k8snode01 opt]# sed 'i3' Dockerfile
3
FROM busybox
3
label:hello world
3
pwd
3
```

## sigleLine

**匹配单行**

```
[root@k8snode01 opt]# sed '1i3' Dockerfile
3
FROM busybox
label:hello world
pwd
```

`$`匹配最后一行

```
[root@k8snode01 opt]# sed '$i3' Dockerfile
FROM busybox
label:hello world
pwd
3
```

`first~step`，按照step迭代

```
[root@k8snode01 opt]# seq 10 | sed '3~2d'
1
2
4
6
8
10
```

## range

**linenumber,linumber**

```
[root@k8snode01 opt]# sed '1,3i3' Dockerfile
3
FROM busybox
3
label:hello world
3
pwd
```

**regex,linenumber**

```
[root@k8snode01 opt]# seq 10 | sed -n '4,/[0-9]/P'
4
5
```

**addr1,+N**

```
[root@k8snode01 opt]# seq 10 | sed -n '4,+3P'
4
5
6
7
```

## regex

```
[root@k8snode01 opt]# sed -n '/bash$/P' /etc/passwd
root:x:0:0:root:/root:/bin/bash
chz:x:1000:1000:chz:/home/chz:/bin/bash
```

**case-insensitive**

在regexp后面使用`I`表示忽略大小写

```
[root@k8snode01 opt]# printf '%s\n' a B c  | sed '/b/d'
a
B
c
[root@k8snode01 opt]# printf '%s\n' a B c  | sed '/b/Id'
a
c
```

**multi-line mode**

在regexp后面使用`M`表示使用multi-line mode(影响`$`和`^`)，==默认使用multi-line mode(i cannot tell it)==

```
[root@k8snode01 opt]# seq 5 | sed -n '/^[[:digit:]]/MP'
1
2
3
4
5
[root@k8snode01 opt]# seq 5 | sed -n '/^[[:digit:]]/P'
1
2
3
4
5
```

## not

```
[root@k8snode01 opt]# sed '/^label/!i3' Dockerfile
3
FROM busybox
label:hello world
3
pwd
3
```





