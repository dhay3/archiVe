# pgrep pkill

pgrep用于过滤特定的进程，pkill过滤并杀掉进程。两者参数差不多

syntax：`pgrep|pkill [options] pattern`

- -u

  过滤指定用户打开的进程可以是numeric或symbolic value

  ```
  root in ~ λ pgrep -u root ssh
  9290
  9292
  9438
  9440
  9645
  9647
  9772
  ```

- -g

  过滤指定进程组

  ```
  root in ~ λ pgrep -ag 1 | head -4
  1 /sbin/init noibrs splash
  ```

- -t 

  过滤指定tty打开的进程，==同样可以过滤pts==

  ```
  root in /dev/pts λ pgrep -t tty7
  818
  root in /dev/pts λ ps -ef | grep -v grep | grep 818
  root         818     789  0 10:58 tty7     00:00:36 /usr/lib/xorg/Xorg :0 -seat seat0 -auth /var/run/lightdm/root/:0 -nolisten tcp vt7 -novtswitch
  
  root in ~ λ pgrep -at pts/3
  10532 -zsh
  10564 cat
  ```

- -d

  按照指定delimiter分隔输出，pgrep特有

  ```
  root in ~ λ pgrep -d : -u root | head -2
  1:2:4:6:7:8:9:10:11:12:13:14:15:16:18:19:20:21:22:25:26:27:28:29:30:31:32:33:34:35:36:37:38:41:42:43:85:86:87:88:89:90:96:105:122:177:179:180:223:239:274:349:681:682:684:723:728:729:736:743:747:774:780:785:847:889:935:1283:2277:7348:8660:8661:9290:9292:9341:9389:9438:9440:9509:9538:9637:9644:9645:9647:9695:9743:9772:9784:9842:9910:10022:10078
  root in ~ λ
  ```

- -a

  显示所有列

  ```
  root in ~ λ pgrep -au root | head -4
  1 /sbin/init noibrs splash
  2 kthreadd
  4 kworker/0:0H
  6 mm_percpu_wq
  ```

- -l 

  显示进程号的同时显示进程名，pgrep特有

  ```
  root in ~ λ pgrep -lx -u root | head -5
  1 systemd
  2 kthreadd
  4 kworker/0:0H
  6 mm_percpu_wq
  7 ksoftirqd/0
  ```

- -x

  精确匹配

  ```
  root in ~ λ pgrep -xl sshd
  9290 sshd
  9292 sshd
  9438 sshd
  9440 sshd
  9645 sshd
  9647 sshd
  9772 sshd
  ```

- -POSIXSIG | --signal POSIX signal

  发送指定的POSIXSIG给进程，默认发送SIGTERM(无法杀掉pts，需要使用SIGKILL)

  ```
  root in ~ λ pkill -9 -t tty7 -u root
  root in ~ λ pkill -9 -t pts/3
```
  
  





