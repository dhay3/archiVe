# linux lsof

> lsof 现在可以在 github 上看到源码
>
> https://github.com/lsof-org/lsof

## Digest

syntax:

```
lsof  [ -?abChlnNOPQRtUvVX ] [ -A A ] [ -c c ] [ +c c ] [ +|-d d ] [
+|-D D ] [ +|-e s ] [ +|-E ] [ +|-f [cfgGn] ] [ -F [f] ] [ -g [s]  ]
[  -i  [i] ] [ -k k ] [ -K k ] [ +|-L [l] ] [ +|-m m ] [ +|-M ] [ -o
[o] ] [ -p s ] [ +|-r [t[m<fmt>]] ] [ -s [p:s] ] [ -S [t] ] [ -T [t]
] [ -u s ] [ +|-w ] [ -x [fl] ] [ -z [z] ] [ -Z [Z] ] [ -- ] [names]
```

`lsof` 用于列出进程打开的文件(==这里的文件指所有的文件==，包括网络打开的socket，在linux上一切皆文件)

打开的文件可以是如下类型
• a regular file
• a directory
• a block special file
• a character special file
• a stream (socket)
• a executing file

`lsof` 如果没有带有参数，==默认输出当前所有进程打开的所有文件==。

## Output/Columns

```
COMMAND    PID TID            USER   FD      TYPE             DEVICE SIZE/OFF       NODE NAME
systemd      1                root  cwd       DIR              252,1     4096          2 /
```

lsof 输出默认有如下几个字段

1. COMMAND

   the first nine characters of the name of the UNIX comamd associated with the process

   如果需要输出关联进程的全部名字，需要使用`+c 0`扩展

2. PID

   process identification

3. TID

   task (thread) identification

   如果该值为空表示，当前这个条目的 COMMAND 是一个进程不是线程

4. USER

   uid or login name of the user to whom the process belongs

   可以通过`-l`实现 login name 转 uid

5. FD

   is the file descriptor number of the file or

   - cwd: current working directory, 即表示NAME 部分是当前进程打开的目录
   - Lnn: library references (AIX)
   - err: FD information error
   - jld: jail directory(FreeBSD)
   - ltx: shared library text(code and data)
   - Mxx: hex memeory-mapped type number xx
   - m86: DOS Merge mapped file
   - mem: memory-mapped file
   - mmap: memory-mapped device
   - pd: parent directory
   - rtd: root directory
   - tr: kernel trace file(OpenBSD)
   - txt: program text(code and data)
   - v86: VP/ix mapped file

   通常 FD 后面会有几个字符，用于表示 file 在什么 mode 下

   - r: for read access
   - w: for write access
   - u: for read and write access
   - space(表示空): mode unknown and no lock character follows
   - -: mode unknown and lock character follows

   mode 后面会有几个字符，用于表示 which lock applied to the file

   - N: for a Solaris NFS lock of unknown type;
   - r: for read lock on part of the file
   - R: for a read lock on the entire file
   - w: for a write lock on part of the file
   - W: for a write lock on the entire file
   - u: for a read and write lock of any length
   - U: for a lock of unknown type
   - x:  for  an SCO OpenServer Xenix lock on part of the file 
   - X: for an SCO OpenServer Xenix  lock  on  the  entire file 
   - space(空格): if there is no lock

6. TYPE

   is the type of the node associated with the file。==TYPE 值较多，具体参考 man page==

   - IPv4/IPv6: socket（包括 stream socket 和 datagram socket），可以参考 [Linux socket.md]()

   - STSO: stream socket

   - unix: unix domain socket

     如果在 NAME 字段中出现

     type=STREAM 表示使用 steam socket

     type=DGRAM 表示使用 datagram socket

   - sock: unknow domain socket

   - BLK: block special file

   - CHR: character special file

   - DEL: linux map file that has bee deleted

   - REG: regular file

   - DIR: a directory

   - PIPE: PIPE

   - FIFO: FIFO special file （具名管道符）

   - LINK: symbolic link file

   - PSXMQ: POSIX message queue file

7. DEVICE

   contains the device numbers

   ```
   systemd-j   389                             root  mem       REG              259,7  16777216    7995568 /var/log/journal/e628a91a78ee47159d95
   ```

   device number 可以使用`lsblk`来查看，`MAJ:MIN`部分

8. SIZE,SIZE/OFF or OFFSET

   is the size of the file or the file offset ==in bytes==

   the file size is displayed in decimal

   the offset is normally displayed in decimal with a leading `0t` if it contains 8 digits or less; in hexadecimal with a leading `0x` if it is longer than 8 digits

   该列会显示打开的文件 size 和 offset

   ```
   systemd-t   748                 systemd-timesync   15u     unix 0x00000000333b7388       0t0      11758 type=STREAM (CONNECTED)
   ```

9. NODE

   - the inode number of a local file
   - the internel protocol type, eg `TCP`
   - `STR` for a stream
   - `CCITT` for  a stream

10. NAME

    - the name of the mount point and file system on which the file resides

    - the name of a character special or block special device

    - the local and remote Internet addresses of a network file

      ```
      # lsof -nPi tcp@localhost
      COMMAND   PID USER   FD   TYPE  DEVICE SIZE/OFF NODE NAME
      qv2ray   1813  cpl   40u  IPv6  958250      0t0  TCP 127.0.0.1:47104->127.0.0.1:15490 (ESTABLISHED)
      ```

      由 local IP:port -> remote IP:port(state) 组成 （除处于 LISTEN 状态的）

    - the address or name fo a UNIX domain socket

      ```
      systemd-u   390                             root    5u     unix 0x00000000d3f17f38       0t0      12501 type=DGRAM (CONNECTED)
      ```

      这里表示 uninx socket 使用的是 datagram socket

    - `STR` followed by the straeam name

    - the system directory name

非默认字段

1. PPID

2. TASKCMD

   smae as the process named in COMMAND column

## Request option

如果`lsof`没有指定任何参数，默认会查看所有进程打开的所有文件。`lsof`可以一些 options 来查看特定的文件，这些 options 就被称为 request option 

如果指定的任一 request option，就只能查看该 request option 对应的文件类型，如果需要查看其他类型的文件，需要追加对应的 request option

eg.

if `-U` is specified for the listing of UNIX socket files, NFS files won't be listed unless `-N` is also specified;  or  if a user list is specified with the `-u` option, UNIX domain socket files, belonging to users not  in  the  list,  won't  be listed unless the `-U` option is also specified

常见的request option 有`-u`,`-i`, `-p`, `-g`, `-c`, `-s`,`+d`等等

### logical

#### default( or )

request options 通常使用 or 运算符进行逻辑拼接

eg.

specifying the `-i` option without an address  and  the  `-u foo`  option produces  a  listing of all network files OR files belonging to processes owned by user foo

#### negated

request options 可以和`^`一起使用表示逻辑非

1) the `^` (negated) login name or user ID (UID), specified with the 
  `-u` option; 

2) the `^` (negated) process ID (PID), specified with the `-p` option; 

3) the  `^` (negated) process group ID (PGID), specified with the `-g` 
  option; 

4) the `^` (negated) command, specified with the `-c*`option; 

5) the `^` (negated) TCP or UDP protocol state names, specified with the `-s [p:s]` option.

#### and

如果想要 request option 使用 and 运算符进行逻辑拼接，需要使用`-a`

eg.

specifying `-a`, `-U`, and `-u foo` produces a  listing  of  only  UNIX  socket files that belong to processes owned by user foo.

## Optional args

### mixed

- `-a` 

  使用 and 逻辑运算符对 request options 进行拼接

  ```
  #lsof -Uau cpl -p 1813
  
  COMMAND  PID USER   FD   TYPE             DEVICE SIZE/OFF  NODE NAME
  qv2ray  1813  cpl    5u  unix 0x000000005b22dedd      0t0 37133 type=STREAM (CONNECTED)
  ```

  这里查看 user cpl pid 1813 打开的 socket 文件

- ` -c word`

  selects the listing of files for processes executing the command that begins with the characters of word
  使用指定的 word 过滤 command 字段，可以使用多个`-c`来过滤多个，之间使用或逻辑

  ```
  root in /home/ubuntu λ lsof -c systemd -c docker | more
  ```

  可以使用`^`表示取非操作

  ```
  root in /home/ubuntu λ lsof -c ^systemd -c ^docker
  ```

  可以使用`/regexp/`格式表示 regexp 过滤

  ```
  root in /home/ubuntu λ lsof -c /^systemd$/ 
  ```

  可以在 closing slash 后面添加 modifier，可以是如下值：

  1. b    the regular expression is a basic one.
  2. i    ignore the case of letters.
  3. x  the regular expression is an extended one (default)

- `+c word`

  COMMAND字段能显示 char 的个数，默认9。如果 word 置位 0 表示显示全部

- `+d directory`

  lsof 查找文件的父目录是 directory 的 

  ```
  # lsof +d /etc
  COMMAND   PID  USER   FD   TYPE DEVICE SIZE/OFF     NODE NAME
  avahi-dae 770 avahi  cwd    DIR  259,7     4096 12850912 /etc/avahi
  avahi-dae 770 avahi  rtd    DIR  259,7     4096 12850912 /etc/avahi
  cupsd     815  root    5r   REG  259,7       68 12848164 /etc/papersize
  ```

- `+D directory`

  lsof 查找文件的子目录是 directory 的 

- `+ | -r [t[c<N>][m<fmt>]]`

  以 repeat mode（直接可以理解成监控模式） 运行 lsof，每 **t** sec 执行一次

  ```
  ➜  test lsof -w -r 1 -p 44856
  COMMAND     PID USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
  kworker/0 44856 root  cwd    DIR  259,7     4096    2 /
  kworker/0 44856 root  rtd    DIR  259,7     4096    2 /
  =======
  COMMAND     PID USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
  kworker/0 44856 root  cwd    DIR  259,7     4096    2 /
  kworker/0 44856 root  rtd    DIR  259,7     4096    2 /
  =======
  ```

  和`-c`，`-p`，`-i`一起使用，可以做到细粒度监控

- `-l`

  将USER 字段转为uid输出，默认以 username 输出

  ```
  
  root in /home/ubuntu λ lsof -l  -i @172.16.253.1
  COMMAND    PID     USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
  gns3serve 1139     1003    7u  IPv4  24324      0t0  TCP gns3vm:3080 (LISTEN)
  gns3serve 1139     1003    8u  IPv4  24329      0t0  TCP gns3vm:41224->gns3vm:3080 (ESTABLISHED)
  gns3serve 1139     1003    9u  IPv4  24330      0t0  TCP gns3vm:3080->gns3vm:41224 (ESTABLISHED)
  root in /home/ubuntu λ id 1003
  uid=1003(gns3) gid=1003(gns3) groups=1003(gns3),999(docker),120(kvm),121(ubridge)
  ```

-  `-u user`

  查看指定用户打开的文件，user 可以是 username 也可以是 uid

- `+ | -w`

  关闭或打开warning message，例如：

  如果用户被删除了，但是系统上该用户启动的进程还未结束，lsof就会报错`lsof: no pwd entry for UID 1000`，可以使用该参数关闭告警

- `-Q`

  ignore failed search terms.

  如果报错 exit code 返回 0

### process

- `-R`

  output 增加一列 PPID

- `-p pid`

  输出和 pid 相关进程打开的文件

  ```
  root in /home/ubuntu λ lsof -p 1139,1123
  COMMAND    PID USER   FD      TYPE             DEVICE SIZE/OFF    NODE NAME
  gns3serve 1139 gns3  cwd       DIR              252,1     4096       2 /
  gns3serve 1139 gns3  rtd       DIR              252,1     4096       2 /
  ```

- `-g pgid`

  输出和 pgid 相关进程打开的文件

- `-t`

  默认会使用-w参数，只打印PID，不会输出其他。用在脚本中，通常结合其他参数或者条件一起使用

  ```
  # lsof -ti tcp@localhost
  815
  1813
  29016
  39143
  ```

### network

- `-U `
  只列出UNIX socket相关的信息

  ```
  root in /home/ubuntu λ lsof -U | head -2
  COMMAND     PID            USER   FD   TYPE             DEVICE SIZE/OFF     NODE NAME
  systemd       1            root   17u  unix 0xffff8f1d774aa000      0t0    13657 /run/systemd/private 
  ```

- `-i  <host>`

  输出和 host 相关打开的文件

  host 使用`[4|6][protocol][@host][:service|port]` 格式。`-i` 可以被重复使用多次，以 or 逻辑显示
  其中 46表示host是IPv4还是IPv6，service使用/etc/services中列出的service

  ```
  root in /home/ubuntu λ lsof -i tcp@172.16.253.1:3080
  COMMAND    PID USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
  gns3serve 1139 gns3    7u  IPv4  24324      0t0  TCP gns3vm:3080 (LISTEN)
  gns3serve 1139 gns3    8u  IPv4  24329      0t0  TCP gns3vm:41224->gns3vm:3080 (ESTABLISHED)
  gns3serve 1139 gns3    9u  IPv4  24330      0t0  TCP gns3vm:3080->gns3vm:41224 (ESTABLISHED)
  ```

  使用`-i`时，必须提供 `[4|6][protocol][@host][:service|port]` 其中的一个单元（`[]`），如果只想指定IP必须使用@

  ```
  root in /home/ubuntu λ lsof -i @172.16.253.1
  COMMAND    PID USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
  gns3serve 1139 gns3    7u  IPv4  24324      0t0  TCP gns3vm:3080 (LISTEN)
  gns3serve 1139 gns3    8u  IPv4  24329      0t0  TCP gns3vm:41224->gns3vm:3080 (ESTABLISHED)
  gns3serve 1139 gns3    9u  IPv4  24330      0t0  TCP gns3vm:3080->gns3vm:41224 (ESTABLISHED)
  ```

  port 部分可以使用`port1-port2`表示 range

- `-s [p:s]`

  如果没有`[p:s]` 显示 size，而不显示 offset

  ```
  root in /home/ubuntu λ lsof -lns  -c openvpn
  COMMAND  PID     USER   FD   TYPE DEVICE    SIZE    NODE NAME
  openvpn 1064        0  cwd    DIR  252,1    4096 1321051 /etc/openvpn
  openvpn 1064        0  rtd    DIR  252,1    4096       2 /
  openvpn 1064        0  txt    REG  252,1  768272  411400 /usr/sbin/openvpn
  ```

  如果有`[p:s]`表示查找指定协议状态机下打开的文件，其中的表示 protocol , s 表示 protocol state

  ```
  root in /home/ubuntu λ lsof -i tcp -s TCP:LISTEN
  COMMAND     PID            USER   FD   TYPE   DEVICE SIZE/OFF NODE NAME
  systemd-r   968 systemd-resolve   13u  IPv4    17468      0t0  TCP localhost:domain (LISTEN)
  gns3serve  1139            gns3    7u  IPv4    24324      0t0  TCP gns3vm:3080 (LISTEN)
  apache2    1317            root    4u  IPv6    21687      0t0  TCP *:http (LISTEN)
  ```

  state 和 UNIX dialects 不同，可以是如下值

  1. CLOSED
  2. IDLE
  3. BOUND
  4. LISTEN
  5. ESTABLISHED
  6. SYN_SENT
  7. SYN_RCDV
  8. CLOSE_WIAT
  9. FIN_WAIT1
  10. CLOSING
  11. LAST_ACK
  12. FIN_WAIT_2
  13. TIME_WAIT

  同样的也可以结合`^`表示非运算

  ```
  lsof -s tcp:^established
  ```

- `-n`

  将 hostname 转为 ip 输出，默认以 hostname 输出

  ```
  root in /home/ubuntu λ lsof -ln  -i @172.16.253.1
  COMMAND    PID     USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
  gns3serve 1139     1003    7u  IPv4  24324      0t0  TCP 172.16.253.1:3080 (LISTEN)
  gns3serve 1139     1003    8u  IPv4  24329      0t0  TCP 172.16.253.1:41224->172.16.253.1:3080 (ESTABLISHED)
  gns3serve 1139     1003    9u  IPv4  24330      0t0  TCP 172.16.253.1:3080->172.16.253.1:41224 (ESTABLISHED)
  ```

- `-P`

   以 numeric 的形式显示端口号

## Examples

使用逻辑与查看 IPv4 pid 为 1234 进程打开的文件

```
lsof -i 4 -a -p 1234
```

查看和`wonderland.cc.puredue.edu`端口513-515关联打开的文件

```
lsof -i @wonderland.cc.purdue.edu:513-515
```

使用逻辑或查看 pid 456 123 用户为 1234, abe 打开的文件

```
lsof -p 456,123,789 -u 1234,abe
```

查看块文件打开的所有文件

```
lsof /dev/hd4
```

查看打开`/ect`的进程

```
➜  /etc lsof -w /etc
COMMAND   PID USER   FD   TYPE DEVICE SIZE/OFF     NODE NAME
zsh     52523 root  cwd    DIR  259,7    12288 12845057 /etc
```

## Specail examples

### 0x001   

如果后面直接跟一个device file，必须是mount中可以查到的(挂载后的)。lsof会列出device file对应file system 所有打开的文件。
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

这时可以使用 lsof 查看运行a.sh打开的文件

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

