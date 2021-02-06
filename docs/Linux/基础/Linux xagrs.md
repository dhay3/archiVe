# Linux xagrs

参考：

https://www.ruanyifeng.com/blog/2019/08/xargs-tutorial.html

由于只有少数命令支持标准输出做为标准输入，只接收命令行参数，==`xargs`是将标准输出转为命令行参数==，一般和`rm，mkdir，ls，mv，cp`一起使用

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

#**参数**

- `-p`

  显示具体命名，并提示用户是否继续执行

  ```
  [root@chz Desktop]# echo 'xargs'|xargs -p mkdir 
  mkdir xargs ?...y
  ```

- `-i`

  https://www.cnblogs.com/paul8339/p/7521027.html

  将stdout放入`{}`中，`-I`需要只当标识符

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