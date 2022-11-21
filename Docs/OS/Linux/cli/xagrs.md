# Linux xagrs

参考：

https://www.ruanyifeng.com/blog/2019/08/xargs-tutorial.html

由于只有少数命令支持标准输出做为标准输入，只接收命令行参数，==`xargs`是将标准输出转为命令行参数==，一般和`rm，mkdir，ls，mv，cp`一起使用(==提一嘴，有些命令是支持占位符的可以使用`-`来表示stdin的内容是从stdout中来的==)

```
[root@chz Desktop]# echo 'xargs'|xargs mkdir 
[root@chz Desktop]# ls
minikube  test  test.bak  test.ttt  v2ray-linux-64  xargs
```

将当前目录下所有符合条件的文件删除

```
[root@chz Desktop]# find . -name 'test*'|xargs rm
或
[root@chz Desktop]# ls|grep test
test1
test2
[root@chz Desktop]# ls|grep test|xargs rm 
[root@chz Desktop]# ls
file1  minikube  v2ray-linux-64  xargs

```

## 参数

- `-p`

  显示具体命名，并提示用户是否继续执行

  ```
  [root@chz Desktop]# echo 'xargs'|xargs -p mkdir 
  mkdir xargs ?...y
  ```

- `-i`

  https://www.cnblogs.com/paul8339/p/7521027.html

  将stdout放入`{}`中，`-I`需要指定标识符

  ```
  ############### 操作的目录下的文件###############
  [root@test05 ab]# ls
  1kk.zip  3kk.zip  5kk.zip  b.rar  d.rar  f.rar  h.rar  j.rar  mini.txt  ni.txt
  2kk.zip  4kk.zip  a.rar    c.rar  e.rar  g.rar  i.rar  k.rar  nii.txt
  ###################使用 i 参数 ##################
  [root@test05 ab]# find . -type f -name "*.txt" | xargs -i cp {}  /tmp/k/
  [root@test05 ab]# ls ../k/
  mini.txt  nii.txt  ni.txt
  [root@test05 ab]#
  ###################  使用 I  参数 ################
  [root@test05 ab]# find . -type f -name "*.txt" | xargs -I {} cp {}  /tmp/n/
  [root@test05 ab]# ls ../n/
  mini.txt  nii.txt  ni.txt
  ```

- `-0`

  以null为结尾，转义和引号变成字面意义

  ```
  [root@cyberpelican opt]# ls | xargs -0 ls
  ls: cannot access '1.txt'$'\n''a'$'\n''bak.xml'$'\n''Blasting_dictionary-master'$'\n''burpsuite pro'$'\n''containerd'$'\n''jdk-14.0.2'$'\n''jdk-14.0.2_linux-x64_bin.tar.gz'$'\n''lsd_0.18.0_amd64.deb.gz'$'\n''t1'$'\n': No such file or directory
  ```

- `-t`

  运行时将实际的命令输出

- `-n`

  表示将多少个参数为一组传给命令，如下表示一次传一个给traceroute。默认会将所有的输出传给xargs

  ```
  chz@cyberpelican:/opt$ { echo baidu.com; echo sohu.com; } | xargs  -n 1 traceroute
  
  ls *.js | xargs -t -n2 ls -al
  ```

- `-P`

  指定xargs使用的最大进程数，默认为1。如果为0表示调用可以使用的最大进程数

  ```
  chz@cyberpelican:/opt$ { echo baidu.com; echo sohu.com; } | xargs -n 1 -P 20 traceroute
  traceroute to sohu.com (211.159.191.77), 30 hops max, 60 byte packets
  traceroute to baidu.com (220.181.38.148), 30 hops max, 60 byte packets
   1  192.168.80.2 (192.168.80.2)  1.125 ms  1.029 ms  0.922 ms 1  192.168.80.2 (192.168.80.2)  0.093 ms  0.151 ms  0.154 ms
   2  * * *
   3  * * *
   4  * *
   * 2  *
  
  ```

  
