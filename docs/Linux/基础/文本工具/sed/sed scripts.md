# sed scripts

参考：

https://www.gnu.org/software/sed/manual/sed.html#sed-scripts

如果没有指定`-f`或`-e`参数，sed默认第一个非参数选项做为sed script，由三部分组成：

```
[address]command[options]
```

- address：If address is specified, the command X will be executed only on the matched lines. 

  can be a ==single line number, a regular expression, or a range of lines==

- command：对匹配行的操作，有POSIX和GUN extensions两种。
- options：command的选项

sed scripts同样有注解，以`#`开头

```
[root@k8snode01 opt]# cat script.sed
#!/bin/sed -f

#this is a comment
/^label/a 3
[root@k8snode01 opt]# cat Dockerfile |  ./script.sed
FROM busybox
label:hello world
3
pwd
```















