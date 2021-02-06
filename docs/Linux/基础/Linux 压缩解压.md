# Linux 压缩/解压

## gzip

### 概述

GUN/gzip用于压缩和解压`.gz`文件。==会替换原有的文件==

> gunzip用于==解压==`.tgz`，`taz`是`.tar.gz`缩写，同样也能解压`gz`
>
> zcat等价于`gunzip -c`，按照编码来识别压缩文件，无需`.gz`也可，输出stdout

### 参数

- gzip

  将文件压缩`.gz`，==不能压缩文件夹==

  ```
  [root@chz Desktop]# gzip minikube 
  [root@chz Desktop]# ls
  minikube.gz  test
  ```

  为tar文件做备份

  ```
  [root@chz test]# tar -zcvf test.tar test1
  test1
  [root@chz test]# ls
  test1  test2  test.tar
  [root@chz test]# gzip test.tar 
  [root@chz test]# ls
  test1  test2  test.tar.gz
  ```

- -d

  decompose

  解压

  ```
  [root@chz Desktop]# gzip -d minikube.gz 
  [root@chz Desktop]# ls
  minikube  test
  ```

- -f

  强制压缩或解压(如果生成的文件已尽存在)

- -v

  显示详细信息

  ```
  [root@chz Desktop]# gzip -v minikube 
  minikube:	 56.7% -- replaced with minikube.gz
  ```

- -r

  将文件夹中的文件压缩，==文件夹不会改变==

  ```
  [root@chz Desktop]# gzip -r test
  [root@chz Desktop]# ls
  minikube.gz  test
  [root@chz Desktop]# cd test
  [root@chz test]# ls
  tset1.gz  tset2.gz
  ```

  test目录下的所有文件都以`.gz`为扩展名

- -c

  将文件压缩后的内容输出，不会改变文件内容

  ```
  [root@chz test]# gzip -c test1 test2
  ��_test1�H���W(�/�I�-�
                           ��_test2+).1��ç�
  ```

  > 可以使用输出重定向
  >
  > ```
  > [root@chz test]# gzip -c test1 test2 > test.gz
  > [root@chz test]# ls
  > test1  test2  test.gz
  > ```
  >
  > ==使用cat方式压缩效率更高==
  >
  > ```
  > [root@chz test]# cat test1 test2| gzip > test.gz
  > [root@chz test]# ls
  > test1  test2  test.gz
  > ```
  >
  > 

## tar

### 概述

GNU /tar (tape archive)将文件保存为single tape或disk archive，也可以从archive中解压文件。解压缩`.tar`或`.tar.gz`文件

pattern：`tar [options] output input`

### 参数

- -z

  使用该参数表示调用gzip

  ```
  [root@chz test]# tar -xvf test.tar.gz
  ```

- -c

  create

  创建文件，一般与`-f`参数一起使用，生成文件不以压缩包形式结尾没有意义

  ```
  [root@chz Desktop]# tar -cf test.tar test1 test2
  ```

- -x

  extract

  解压文件，一般与`-f`参数一起使用

  ```
  
  ```

- -f

  使用压缩文件或是普通文件，一般与`-c`或`-x`参数一起使用

  ```
  [root@chz Desktop]# tar -cf test.tar test1 test2
  ```

- -v

  压缩和解压时列出文件
  
  ```
  [root@chz test]# tar -zxf test.tar.gz 
  [root@chz test]# ls
  test1  test2  test.tar.gz
  [root@chz test]# tar -zxvf test.tar.gz 
  test1
  test2
  [root@chz test]# ls
  test1  test2  test.tar.gz
  [root@chz test]# 
  ```

- -C

  指定文件的输出路径

  ```
  [root@chz test]# tar -xvf test.tar.gz -C /opt
  ```

## zip

将文件压缩成`.zip`

```
[root@chz test]# zip test.zip test1 
  adding: test1 (stored 0%)
[root@chz test]# ls
test1  test.zip
```

## unzip

解压`.zip`文件

```
[root@chz test]# unzip test.zip 
Archive:  test.zip
replace test1? [y]es, [n]o, [A]ll, [N]one, [r]ename: y
 extracting: test1    
```

