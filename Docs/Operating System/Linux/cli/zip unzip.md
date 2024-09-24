# zip/unzip

将文件压缩成`.zip`

```
[root@chz test]# zip test.zip test1 
  adding: test1 (stored 0%)
[root@chz test]# ls
test1  test.zip
```

解压`.zip`文件

```
[root@chz test]# unzip test.zip 
Archive:  test.zip
replace test1? [y]es, [n]o, [A]ll, [N]one, [r]ename: y
 extracting: test1    
```
