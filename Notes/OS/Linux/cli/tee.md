# Linux tee

## Digest

syntax

```
tee [OPTION]... [FILE]...
```

用于从 stdin 中读取内容，将内容输出到 stdout 和文件中

## Optional args

- `-a | --append`

  将 stdin 中的内容添加到指定文件的末尾，不会 overwrite

## Examples

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