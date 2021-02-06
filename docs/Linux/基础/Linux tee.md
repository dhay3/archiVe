# Linux tee

将stdin输出到stdout和file中，使用`-a`参数追加写入不会覆盖源文件内容，使用`ctrl + D`表示终止输入

```
[root@chz Desktop]# tee file1
hello world
hello world
java
java
c++
c++
[root@chz Desktop]# cat file1 
hello world
java
c++
```

可以使用管道符，将`ls`显示的内容写入管道符中，做为stdout

```
ls | tee out.txt
ls > out.txt
```



