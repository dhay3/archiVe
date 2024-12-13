# tar

## Digest

GNU /tar (tape archive)将文件保存为single tape或disk archive，也可以从archive中解压文件。解压缩`.tar`或`.tar.gz`文件

pattern：`tar [options] output input`

## Optional args

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
  cpl in ~ λ tar -xf 1.tar
  ```

- -f

  ==指定使用的压缩文件或压缩后的文件名==，一般与`-c`或`-x`参数一起使用

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
