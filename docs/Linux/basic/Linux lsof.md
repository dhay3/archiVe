# linux lsof

> lsof -nPi  tcp@172.16.253.1:3080

lsof 用于列出进程打开的文件(包括网络打开的socket，在linux上一切皆文件)
打开的文件可以是如下类型
• a regular file
• a directory
• a block special file
• a character special file
• a stream (socket)
• a executing file
lsof如果没有带有参数，默认输出当前所有进程打开的文件。默认字段如下
具体可以查看`manual page /^\s+COMMAND`

```
root in /opt λ lsof | more
COMMAND    PID TID            USER   FD      TYPE             DEVICE SIZE/OFF       NODE NAME
systemd      1                root  cwd       DIR              252,1     4096          2 /
systemd      1                root  rtd       DIR              252,1     4096          2 /
systemd      1                root  txt       REG              252,1  1616248    1190934 /lib/systemd/systemd
systemd      1                root  mem       REG              252,1  1700792    1188287 /lib/x86_64-linux-gnu/libm-2.27.so
systemd      1                root  mem       REG              252,1   121016    1180095 /lib/x86_64-linux-gnu/libudev.so.1.6.9
```

• COMMAND：进程的名称
• PID：进程标识符
• USER：进程所有者
• FD：文件描述符（一个索引用于表示被打开的文件，inode的接口）
https://segmentfault.com/a/1190000009724931
（1）cwd：表示current work dirctory，即：应用程序的当前工作目录，这是该应用程序启动的目录，除非它本身对这个目录进行更改
（2）txt ：该类型的文件是程序代码，如应用程序二进制文件本身或共享库，如上列表中显示的 systemd 程序
（3）lnn：library references (AIX);
（4）er：FD information error (see NAME column);
（5）jld：jail directory (FreeBSD);
（6）ltx：shared library text (code and data);
（7）mxx ：hex memory-mapped type number xx.
（8）m86：DOS Merge mapped file;
（9）mem：memory-mapped file;
（10）mmap：memory-mapped device;
（11）pd：parent directory;
（12）rtd：root directory;
（13）tr：kernel trace file (OpenBSD);
（14）v86  VP/ix mapped file;
（15）0：表示标准输入
（16）1：表示标准输出
（17）2：表示标准错误
一般在标准输出、标准错误、标准输入后还跟着文件状态模式：r、w、u等
（1）u：表示该文件被打开并处于读取/写入模式
（2）r：表示该文件被打开并处于只读模式
（3）w：表示该文件被打开并处于
（4）空格：表示该文件的状态模式为unknow，且没有锁定
（5）-：表示该文件的状态模式为unknow，且被锁定
同时在文件状态模式后面，还跟着相关的锁
（1）N：for a Solaris NFS lock of unknown type;
（2）r：for read lock on part of the file;
（3）R：for a read lock on the entire file;
（4）w：for a write lock on part of the file;（文件的部分写锁）
（5）W：for a write lock on the entire file;（整个文件的写锁）
（6）u：for a read and write lock of any length;
（7）U：for a lock of unknown type;
（8）x：for an SCO OpenServer Xenix lock on part      of the file;
（9）X：for an SCO OpenServer Xenix lock on the      entire file;
（10）space：if there is no lock.
• TYPE：文件类型
（1）DIR：表示目录
（2）CHR：表示字符类型
（3）BLK：块设备类型
（4）UNIX： UNIX 域套接字
（5）FIFO：先进先出 (FIFO) 队列
（6）IPv4：网际协议 (IP) 套接字
• DEVICE：运行该进程的磁盘名称
• SIZE：文件的大小
• NODE：inode
• NAME：打开文件的确切名称

## options

• -a 
在lsof默认输出列中添加字段，如果单独使用会报错，需要结合其他参数一起使用

```
root in /home/ubuntu λ lsof  -aU | head -1
COMMAND     PID            USER   FD   TYPE             DEVICE SIZE/OFF     NODE NAME
```

• -c
==既然有了该参数就不应该使用grep==
使用command字段(进程名)过滤lsof，可以使用多个-c来过滤多个，之间使用或逻辑

```
root in /home/ubuntu λ lsof -c systemd -c docker | more
```

可以使用^表示取非操作

```
root in /home/ubuntu λ lsof -c ^systemd -c ^docker
```

可以使用/regexp/格式表示regexp过滤

```
root in /home/ubuntu λ lsof -c /^systemd$/ 
```

可以在closing slash后面添加modifier

1. b    the regular expression is a basic one.

2. i    ignore the case of letters.

3. x    the regular expression is an extended one (default)\
    • -U 
    只列出UNIX socket相关的信息

  ```
  root in /home/ubuntu λ lsof -U | head -2
  COMMAND     PID            USER   FD   TYPE             DEVICE SIZE/OFF     NODE NAME
  systemd       1            root   17u  unix 0xffff8f1d774aa000      0t0    13657 /run/systemd/private 
  ```

  type=STREAM
  • `-i  <host>`
  输出和host相关进程打开的文件，host由`[4|6][protocol][@host][:service|port]`
  46表示host是V4还是V6，service使用/etc/services中列出的service

  ```
  root in /home/ubuntu λ lsof -i tcp@172.16.253.1:3080
  COMMAND    PID USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
  gns3serve 1139 gns3    7u  IPv4  24324      0t0  TCP gns3vm:3080 (LISTEN)
  gns3serve 1139 gns3    8u  IPv4  24329      0t0  TCP gns3vm:41224->gns3vm:3080 (ESTABLISHED)
  gns3serve 1139 gns3    9u  IPv4  24330      0t0  TCP gns3vm:3080->gns3vm:41224 (ESTABLISHED)
  ```

  如果只想指定IP必须使用@

  ```
  root in /home/ubuntu λ lsof -i @172.16.253.1
  COMMAND    PID USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
  gns3serve 1139 gns3    7u  IPv4  24324      0t0  TCP gns3vm:3080 (LISTEN)
  gns3serve 1139 gns3    8u  IPv4  24329      0t0  TCP gns3vm:41224->gns3vm:3080 (ESTABLISHED)
  gns3serve 1139 gns3    9u  IPv4  24330      0t0  TCP gns3vm:3080->gns3vm:41224 (ESTABLISHED)
  ```

  • -l

  ```
  将username转为uid输出，默认以username输出
  root in /home/ubuntu λ lsof -l  -i @172.16.253.1
  COMMAND    PID     USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
  gns3serve 1139     1003    7u  IPv4  24324      0t0  TCP gns3vm:3080 (LISTEN)
  gns3serve 1139     1003    8u  IPv4  24329      0t0  TCP gns3vm:41224->gns3vm:3080 (ESTABLISHED)
  gns3serve 1139     1003    9u  IPv4  24330      0t0  TCP gns3vm:3080->gns3vm:41224 (ESTABLISHED)
  root in /home/ubuntu λ id 1003
  uid=1003(gns3) gid=1003(gns3) groups=1003(gns3),999(docker),120(kvm),121(ubridge)
  ```

  • -n

  将hostname转为ip输出，默认以hostname输出

  ```
  root in /home/ubuntu λ lsof -ln  -i @172.16.253.1
  COMMAND    PID     USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
  gns3serve 1139     1003    7u  IPv4  24324      0t0  TCP 172.16.253.1:3080 (LISTEN)
  gns3serve 1139     1003    8u  IPv4  24329      0t0  TCP 172.16.253.1:41224->172.16.253.1:3080 (ESTABLISHED)
  gns3serve 1139     1003    9u  IPv4  24330      0t0  TCP 172.16.253.1:3080->172.16.253.1:41224 (ESTABLISHED)
  ```

  • -P

  以numeric的形式显示端口号

  • -s [p:s]
  如果没有[p:s]显示size，而不显示offset

  ```
  root in /home/ubuntu λ lsof -lns  -c openvpn
  COMMAND  PID     USER   FD   TYPE DEVICE    SIZE    NODE NAME
  openvpn 1064        0  cwd    DIR  252,1    4096 1321051 /etc/openvpn
  openvpn 1064        0  rtd    DIR  252,1    4096       2 /
  openvpn 1064        0  txt    REG  252,1  768272  411400 /usr/sbin/openvpn
  ```

  如果有[p:s]表示过滤指定protocol 指定state

  ```
  root in /home/ubuntu λ lsof -i tcp -s TCP:LISTEN
  COMMAND     PID            USER   FD   TYPE   DEVICE SIZE/OFF NODE NAME
  systemd-r   968 systemd-resolve   13u  IPv4    17468      0t0  TCP localhost:domain (LISTEN)
  gns3serve  1139            gns3    7u  IPv4    24324      0t0  TCP gns3vm:3080 (LISTEN)
  apache2    1317            root    4u  IPv6    21687      0t0  TCP *:http (LISTEN)
  master     1503            root   13u  IPv4    22269      0t0  TCP *:smtp (LISTEN)
  master     1503            root   14u  IPv6    22270      0t0  TCP *:smtp (LISTEN)
  dnsmasq    1709 libvirt-dnsmasq    6u  IPv4    23391      0t0  TCP gns3vm:domain (LISTEN)
  docker-pr  1903            root    4u  IPv4    24422      0t0  TCP *:ssh (LISTEN)
  docker-pr  1908            root    4u  IPv6    24428      0t0  TCP *:ssh (LISTEN)
  sshd       5055            root    3u  IPv4 35447544      0t0  TCP *:65522 (LISTEN)
  apache2    7452        www-data    4u  IPv6    21687      0t0  TCP *:http (LISTEN)
  apache2    7453        www-data    4u  IPv6    21687      0t0  TCP *:http (LISTEN)
  iperf3    21949            root    3u  IPv6 41291064      0t0  TCP *:10086 (LISTEN)
  ```

  • -p
  输出和pid相关进程打开的文件

  ```
  root in /home/ubuntu λ lsof -p 1139
  COMMAND    PID USER   FD      TYPE             DEVICE SIZE/OFF    NODE NAME
  gns3serve 1139 gns3  cwd       DIR              252,1     4096       2 /
  gns3serve 1139 gns3  rtd       DIR              252,1     4096       2 /
  ```

  • -u
  指定用户打开的文件
  • -t
  默认会使用-w参数，只打印PID，不会输出其他

  ```
  root in /home/ubuntu λ lsof -t | head -3
  1
  2
  4
  ```

  • + | -w
  关闭或打开warning message，如果用户被删除了，但是系统上该用户启动的进程还未         结束，lsof就会报错`lsof: no pwd entry for UID 1000`，可以使用该参数关闭告警

## 例子

### **0x001**

如果后面直接跟一个device file，必须是mount中可以查到的(挂载后的)。lsof会列出device file对应file system所有打开的文件。
所以当磁盘在忙的时候无法卸载，可以通过lsof查看占用的文件，然后关闭进程

```
cpl in /mnt λ sudo umount /dev/sda4
umount: /mnt/win: target is busy.
cpl in /mnt λ lsof /dev/sda4
COMMAND   PID USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
zsh      8323  cpl  cwd    DIR    8,4     4096    2 /mnt/win
zsh     61660  cpl  cwd    DIR    8,4     4096    2 /mnt/win
vim     61819  cpl  cwd    DIR    8,4     4096    2 /mnt/win
zsh     70464  cpl  cwd    DIR    8,4     4096    2 /mnt/win
zsh     77976  cpl  cwd    DIR    8,4     4096    2 /mnt/win
man     99630  cpl  cwd    DIR    8,4     4096    2 /mnt/win
man     99639  cpl  cwd    DIR    8,4     4096    2 /mnt/win
less    99640  cpl  cwd    DIR    8,4     4096    2 /mnt/win
cpl in /mnt λ lsof /dev/sda4  | sed -n '2,$p' | awk '{print($2)}' | xargs -i kill -9 {}
cpl in /mnt λ sudo umount /dev/sda4
```

### 0x002

linux上如果文件被删除了，但是当前有进程在使用该文件就不会讲文件真正的删除，只有将进程终止才会删除文件

```
root in /tmp λ cat a.sh;bash a.sh &
#!/usr/bin sh
while true;do echo 1 >& /dev/null;done
```

这时可以使用lsof查看运行a.sh打开的文件

```
root in /tmp λ lsof -w | grep a.sh
bash       9384                  root  255r      REG              252,1         53     411838 /tmp/a.sh
```

删除a.sh，可以看到系统还是保留了lsof打开的文件，但是加了一个deleted标识

```
root in /tmp λ rm -f a.sh;lsof -w | grep a.sh
bash       9384                  root  255r      REG              252,1         53     411838 /tmp/a.sh (deleted)
```

所以只有总之对应的进程，文件才是真正的删除，且进程不是zombie

```
root in /tmp λ ps -ef9384
  PID TTY      STAT   TIME COMMAND
 9384 pts/1    RN     4:46 bash a.sh SSH_CONNECTION=115.206.246.212 11738 172.21.16.3 65522 LANG=en_US.utf8 DISPLAY=localhost:10.0 HISTTIMEFO
```

