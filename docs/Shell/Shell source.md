# Shell source

## 概述

==在当前shell环境下读取并执行filename中的命令，==source命令也被称为“点命令”，通常用于重新执行刚修改的初始化文件，使之立即生效，而不必注销并重新登入（但是如果关闭tty就会失效）。该命令通常用`.`来替代

```
source filename
. filename
```

例如

```
[root@chz Desktop]# echo 'pwd'>testsource 
[root@chz Desktop]# nl testsource 
     1	pwd
[root@chz Desktop]# . testsource 
/root/Desktop
```

通过短路与可以执行多条命令

```
[root@chz Desktop]# cat testsource 
pwd &&
ls &&
echo source
[root@chz Desktop]# . testsource 
/root/Desktop
minikube  test  test.bak  testsource  test.ttt  v2ray-linux-64  xargs
source
```

## source filname VS bash filename

1. shell脚本需要执行权限，而使用source不需要执行权限

2. source 不会解释shebang

   ```
   root in /usr/local/\ λ cat shebang.sh
   #!/bin/cat
   df -hT
   
   root in /usr/local/\ λ source shebang.sh
   Filesystem     Type      Size  Used Avail Use% Mounted on
   udev           devtmpfs  2.0G     0  2.0G   0% /dev
   tmpfs          tmpfs     395M  6.0M  389M   2% /run
   /dev/vda1      ext4       40G  5.9G   32G  16% /
   tmpfs          tmpfs     2.0G     0  2.0G   0% /dev/shm
   tmpfs          tmpfs     5.0M     0  5.0M   0% /run/lock
   tmpfs          tmpfs     2.0G     0  2.0G   0% /sys/fs/cgroup
   tmpfs          tmpfs     395M     0  395M   0% /run/user/0
   
   root in /usr/local/\ λ ./shebang.sh
   #!/bin/cat
   df -hT
   ```

3. bash 会建立一个新的子shell，但是source不会。==所以如果在shell脚本中是创建变量的话，就等于在当前shell中创建变量。==see defferent pid 

   ```
   root in /usr/local/\ λ cat shebang.sh
   echo "pid=$$"
   root in /usr/local/\ λ echo $$
   8679
   root in /usr/local/\ λ source shebang.sh
   pid=8679
   root in /usr/local/\ λ ./shebang.sh
   pid=9166
   ```

   
