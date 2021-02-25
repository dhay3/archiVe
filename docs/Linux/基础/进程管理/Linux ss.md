# Linux ss

参考：

https://einverne.github.io/post/2013/01/ss-command-socket-statistics.html

## 概述

用来获取unix上的socket，==如果主机上不方便安装netstat可以使用该命令==。如果没有带有参数默认只展示ESTABLISHED的socket

```
Netid State  Recv-Q Send-Q                     Local Address:Port      Peer Address:Port  Process
u_str ESTAB  0      0                                      * 176285               * 176284       
u_str ESTAB  0      0                                      * 32936                * 31884        
u_str ESTAB  0      0                     @/tmp/.X11-unix/X0 30784                * 30037  
```

## 参数

- `-O,--oneline`

  以一行展示套接字情况

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

  只展示正在监听的套接字

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

- `-k`

  关闭指定socket

  ```
  
  ```

  

## Example

- `ss -tlp`等价于`netstat -lpt`

  ```
  State  Recv-Q Send-Q Local Address:Port       Peer Address:PortProcess                                                                                                                                                             
  LISTEN 0      80         127.0.0.1:mysql           0.0.0.0:*    users:(("mysqld",pid=1154,fd=21))                                                                                                                                  
  LISTEN 0      128          0.0.0.0:ssh             0.0.0.0:*    users:(("sshd",pid=1017,fd=3))                                                                                                                                     
  LISTEN 0      244        127.0.0.1:postgresql      0.0.0.0:*    users:(("postgres",pid=1290,fd=4))      
  ```

- `ss dst ipaddr`

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

