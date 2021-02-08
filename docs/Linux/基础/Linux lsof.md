# Linux lsof

> `root in /opt λ lsof -n | grep deleted`
>
> 列出被删除的文件

## 概述

lsof 默认用于展示当前所有进程打开的文件，如果携带了参数将不会显示完整的列数。

```
root in /opt λ lsof | more
COMMAND    PID TID            USER   FD      TYPE             DEVICE SIZE/OFF       NODE NAME
systemd      1                root  cwd       DIR              252,1     4096          2 /
systemd      1                root  rtd       DIR              252,1     4096          2 /
systemd      1                root  txt       REG              252,1  1616248    1190934 /lib/systemd/systemd
systemd      1                root  mem       REG              252,1  1700792    1188287 /lib/x86_64-linux-gnu/libm-2.27.so
systemd      1                root  mem       REG              252,1   121016    1180095 /lib/x86_64-linux-gnu/libudev.so.1.6.9
```

- COMMAND：进程的名称

- PID：进程标识符

- USER：进程所有者

- FD：文件描述符（一个索引用于表示被打开的文件，inode的接口）

  https://segmentfault.com/a/1190000009724931

  ```
  （1）cwd：表示current work dirctory，即：应用程序的当前工作目录，这是该应用程序启动的目录，除非它本身对这个目录进行更改
  （2）txt ：该类型的文件是程序代码，如应用程序二进制文件本身或共享库，如上列表中显示的 /sbin/init 程序
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
  ```

- TYPE：文件类型

  ```
  （1）DIR：表示目录
  （2）CHR：表示字符类型
  （3）BLK：块设备类型
  （4）UNIX： UNIX 域套接字
  （5）FIFO：先进先出 (FIFO) 队列
  （6）IPv4：网际协议 (IP) 套接字
  ```

- DEVICE：运行改进程的磁盘名称
- SIZE：文件的大小
- NODE：inode
- NAME：打开文件的确切名称

## 参数

- `-u`

  显示指定用户打开的文件

  ```
  root in /opt λ lsof -u root | more
  COMMAND    PID USER   FD      TYPE             DEVICE SIZE/OFF       NODE NAME
  systemd      1 root  cwd       DIR              252,1     4096          2 /
  systemd      1 root  rtd       DIR              252,1     4096          2 /
  systemd      1 root  txt       REG              252,1  1616248    1190934 /lib/systemd/systemd
  systemd      1 root  mem       REG              252,1  1700792    1188287 /lib/x86_64-linux-gnu/libm-2.27.so
  ```

- `-p`

  显示指定进程号的进程打开的文件

  ```
  root in /opt λ lsof -p 1 | more
  COMMAND PID USER   FD      TYPE             DEVICE SIZE/OFF       NODE NAME
  systemd   1 root  cwd       DIR              252,1     4096          2 /
  systemd   1 root  rtd       DIR              252,1     4096          2 /
  systemd   1 root  txt       REG              252,1  1616248    1190934 /lib/systemd/systemd
  ```

- `-i`

  显示网络连接(在linux中网络连接同样也是文件)

  ```
  root in /opt λ lsof -i  tcp #显示tcp连接的文件
  COMMAND    PID            USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
  systemd-r  490 systemd-resolve   13u  IPv4  14942      0t0  TCP localhost:domain (LISTEN)
  apache2    712            root    4u  IPv6  18438      0t0  TCP *:http (LISTEN)
  apache2    712            root    6u  IPv6  18442      0t0  TCP *:https (LISTEN)
  apache2    713        www-data    4u  IPv6  18438      0t0  TCP *:http (LISTEN)
  apache2    713        www-data    6u  IPv6  18442      0t0  TCP *:https (LISTEN)
  ```
  
  使用`lsof -i TCP:port`可以有效查看指定端口被占用的当前进程
  
  ```
  root in ~ λ lsof -i TCP:10086
  COMMAND   PID USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
  sshd    13848  cpl    9u  IPv4 888473      0t0  TCP localhost:10086 (LISTEN)  
  ```
  
  













