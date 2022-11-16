# Linux ss

参考：
[https://einverne.github.io/post/2013/01/ss-command-socket-statistics.html](https://einverne.github.io/post/2013/01/ss-command-socket-statistics.html)
[https://man7.org/linux/man-pages/man8/ss.8.html](https://man7.org/linux/man-pages/man8/ss.8.html)

https://wangchujiang.com/linux-command/c/ss.html

https://phoenixnap.com/kb/ss-command

## Digest

syntax：`ss [options] [filter]`

show socket 用来获取 unix 上的 socket，如果主机上不方便安装netstat可以使用该命令。如果没有带有参数默认只展示ESTABLISHED的socket
```
Netid State  Recv-Q Send-Q                     Local Address:Port      Peer Address:Port  Process
u_str ESTAB  0      0                                      * 176285               * 176284       
u_str ESTAB  0      0                                      * 32936                * 31884        
u_str ESTAB  0      0                     @/tmp/.X11-unix/X0 30784                * 30037
```
## Columns

一般 ss 会显示如下几个字段

1. Netid

   type of socket, 可以是 TCP, UDP, u_str (Unix Stream), u_seq(Unix sequence)

2. State

   State of the socket. Most commonly *ESTAB* (established), *UNCONN* (unconnected), *LISTEN* (listening).

3. Recv-Q

   Number of received packets in the queue

4. Send-Q

   Number of sent packets in the queue

5. Local address:port

    Address of local machine and port

6. Peer address:port

   Address of remote machine and port

## Optional args

由于版本不同，以下有些参数可能不同

- `-O | --online`

  以一行显示socket

  ```
  ubuntu@VM-12-16-ubuntu:~$ ss -lnptmO
  State            Recv-Q           Send-Q                                      Local Address:Port                     Peer Address:Port           Process                                                    
  LISTEN           0                4096                                              0.0.0.0:111                           0.0.0.0:*               skmem:(r0,rb131072,t0,tb16384,f0,w0,o0,bl0,d0)            
  ```

- `-r,--resolve`
  将数字IP解析为domain 

  ```
  ss -O | more
  udp   ESTAB  0      0                                  [::1]:47717            [::1]:47717       
  v_str ESTAB  0      0                             1621159833:1022                 0:976         
  v_str ESTAB  0      0                             1621159833:1023                 0:976  
  
  ---
  
  root in ~ λ ss -rO | more
  udp   ESTAB  0      0                              localhost:47717        localhost:47717       
  v_str ESTAB  0      0                             1621159833:1022                 0:976         
  v_str ESTAB  0      0                             1621159833:1023                 0:976
  ```

- `-l,--listening`
  只展示正在监听的套接字 ，缺省参数

- `-a,--all`
  展示正在监听和没有监听的套接字，如果指定`-t`参数只展示ESTABLISHED的socket 

- `-t|-u`
  tcp或udp使用的套接字 

- `-p,--process`
  显示socket的同时展示关联的进程 

  ```
  Netid State  Recv-Q Send-Q                     Local Address:Port      Peer Address:Port Process                                                                                                                                                                                                     
  u_str ESTAB  0      0                     @/tmp/.X11-unix/X0 86772                * 88006 users:(("Xorg",pid=1011,fd=59))
  ```

- `-n`
  显示IP和端口，而不显示解析后的主机名

  ```
  root in /usr/local/\ λ ss -ltpn
  State            Recv-Q            Send-Q                        Local Address:Port                       Peer Address:Port            
  LISTEN           0                 128                               127.0.0.1:6010                            0.0.0.0:*                users:(("sshd",pid=5995,fd=8))
  LISTEN           0                 128                           127.0.0.53%lo:53                              0.0.0.0:*                users:(("systemd-resolve",pid=2002,fd=13))
  LISTEN           0                 128                                 0.0.0.0:22                              0.0.0.0:*                users:(("sshd",pid=673,fd=3))
  LISTEN           0                 128                                       *:443                                   *:*                users:(("apache2",pid=5210,fd=6),("apache2",pid=5209,fd=6),("apache2",pid=855,fd=6))
  LISTEN           0                 128                                       *:80                                    *:*                users:(("apache2",pid=5210,fd=4),("apache2",pid=5209,fd=4),("apache2",pid=855,fd=4))
  ```

- `-o | --options`
  显示时间相关的信息 ，timer EBNF `timer:(<timer_name>,<expire_time>,<retrans>)`

  ```
  cpl in ~ λ ss -npto 
  State         Recv-Q      Send-Q                Local Address:Port                   Peer Address:Port       Process                                                                                                      
  ESTAB         0           0                         127.0.0.1:45392                     127.0.0.1:1089        users:(("chrome",pid=7767,fd=26)) timer:(keepalive,20sec,0)                                                 
  ESTAB         0           0                         127.0.0.1:15490                     127.0.0.1:37506       users:(("v2ray",pid=2406,fd=17)) timer:(keepalive,8.820ms,0)                                                
  ESTAB         0           0                         127.0.0.1:
  ```

  上述第一条表示使用 tcp keepalive time，20sec过期，没有重传。具体查看man page

- `-m | --memory`

  显示socket的内存使用详情，skmem EBNF`skmem:(r<rmem_alloc>,rb<rcv_buf>,t<wmem_alloc>,tb<snd_buf>,f<fwd_alloc>,w<wmem_queued>,o<opt_mem>,bl<back_log>,d<sock_drop>)`

  ```
  ubuntu@VM-12-16-ubuntu:~$ ss -lnptmO
  State            Recv-Q           Send-Q                                      Local Address:Port                     Peer Address:Port           Process                                                    
  LISTEN           0                4096                                              0.0.0.0:111                           0.0.0.0:*               skmem:(r0,rb131072,t0,tb16384,f0,w0,o0,bl0,d0)            
  ```

  主要观察几个参数

  1. rcv_buf：recive buffer，接受窗口占用的内存
  2. back_log：recive back log，建连未处理的包占用的内存

- `-s | --summary`

  显示summary statistics

  ```
  ubuntu@VM-12-16-ubuntu:~$ ss -s
  Total: 175
  TCP:   17 (estab 4, closed 2, orphaned 0, timewait 2)
  
  Transport Total     IP        IPv6
  RAW       1         0         1        
  UDP       16        9         7        
  TCP       15        10        5        
  INET      32        19        13       
  FRAG      0         0         0        
  ```

- `-K | --kill`
  按照 filter 和 expression 关闭指定的 socket，如果关闭成功会回显

## Filter
具体查看man page
Filter ::= [state STATE-FILTER] [EXPRESSION]

### State-filter

state-filter 用于过滤 socket 的状态

STATE-FILTER ::= [established | syn-sent | syn-recv | fin-wait-1 |  fin-wait-2 | time-wait | closed | close-wait | last-ack | listening | closing]

```
cpl in ~ λ ss -lnpt state LISTENING 
Recv-Q   Send-Q      Local Address:Port        Peer Address:Port   Process                                   
0        128             127.0.0.1:631              0.0.0.0:*                                                
0        4096            127.0.0.1:8889             0.0.0.0:*       users:(("v2ray",pid=2406,fd=4))          
0        4096            127.0.0.1:1089             0.0.0.0:*       users:(("v2ray",pid=2406,fd=8))          
0        4096            127.0.0.1:15490            0.0.0.0:*       users:(("v2ray",pid=2406,fd=10))         
0        50                      *:1716                   *:*       users:(("kdeconnectd",pid=1573,fd=15))   
0        128                 [::1]:631                 [::]:*
```
### Expression

> 需要注意的这些逻辑不在 shell 中有特殊含义，如果需要使用需要转义

expression 用于过滤指定条件的 socket，同时支持通过逻辑符`&&`，`||`，`!`拼接，如果没有使用逻辑符，默认使用`&&`拼接 

EXPRESSION 可以是如下的值(具体查看man page)

- {dst | src} [=]  HOST

过滤指定源目HOST(HOST 一般为 IP)关联的socket
```
ubuntu@VM-12-16-ubuntu:~$ ss state listening  src 127.0.0.1
Netid                  Recv-Q                  Send-Q                                   Local Address:Port                                      Peer Address:Port                  Process                  
tcp                    0                       10                                           127.0.0.1:domain                                         0.0.0.0:*                                              
tcp                    0                       4096                                         127.0.0.1:953                                            0.0.0.0:*                                              
```

- {dport | sport} [OP] [FAMILY] :PORT

OP表示 operator，可以是算数逻辑符。==需要注意的一点是如果使用了 operator 两侧必须要有空格==，否则会报错`Error: an inet prefix is expected rather than...`

```
ubuntu@VM-12-16-ubuntu:~$ ss -n state listening  sport = :53
Netid                Recv-Q                Send-Q                                                Local Address:Port                                 Peer Address:Port                Process                
tcp                  0                     10                                                       10.0.12.16:53                                        0.0.0.0:*                                          
tcp                  0                     10                                                        127.0.0.1:53                                        0.0.0.0:*                                          
tcp                  0                     4096                                                  127.0.0.53%lo:53                                        0.0.0.0:*                                          
tcp                  0                     10                                   [fe80::5054:ff:fe49:561a]%eth0:53                                           [::]:*                                          
tcp                  0                     10                                                            [::1]:53                                           [::]:*                                          
```
同时也需要注意另外一点，有些版本的 `ss` 必须指定 operator 否则同样会报错

## Example

-  `ss -tlp`等价于`netstat -lpt` 
```
State  Recv-Q Send-Q Local Address:Port       Peer Address:PortProcess                                                                                                                                                             
LISTEN 0      80         127.0.0.1:mysql           0.0.0.0:*    users:(("mysqld",pid=1154,fd=21))                                                                                                                                  
LISTEN 0      128          0.0.0.0:ssh             0.0.0.0:*    users:(("sshd",pid=1017,fd=3))                                                                                                                                     
LISTEN 0      244        127.0.0.1:postgresql      0.0.0.0:*    users:(("postgres",pid=1290,fd=4))
```

-  `ss dst ipaddr`
列出目标地址与本机打开的Socket 
```
root in ~ λ ss dst 115.233.222.34
Netid          State           Recv-Q           Send-Q                      Local Address:Port                       Peer Address:Port
tcp            ESTAB           0                0                           172.19.124.44:ssh                      115.233.222.34:32035
tcp            ESTAB           0                0                           172.19.124.44:ssh                      115.233.222.34:64452
tcp            ESTAB           0                160                         172.19.124.44:ssh                      115.233.222.34:64417
tcp            ESTAB           0                0                           172.19.124.44:ssh                      115.233.222.34:27469
tcp            ESTAB           0                0                           172.19.124.44:ssh                      115.233.222.34:32114
```

- `ss -K`

  关闭指定 socket

  ```
  cpl in ~/centos7 λ sudo ss -K state established  dport :50000
  Netid          Recv-Q          Send-Q                   Local Address:Port                      Peer Address:Port           Process          
  tcp            0               0                                [::1]:35260                            [::1]:50000
  
  cpl in ~/centos7 λ ncat -v localhost 50000
  Ncat: Version 7.92 ( https://nmap.org/ncat )
  Ncat: Connected to ::1:50000.
  
  Ncat: Software caused connection abort.
  ```

  
