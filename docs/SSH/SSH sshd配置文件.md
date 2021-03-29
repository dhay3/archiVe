# SSH ssh服务器

> `man sshd_config`
>
> 服务端配置文件`/etc/ssh/sshd_config`，同时会继承`/etc/ssh/sshd_config.d/*.conf`中所有的配置文件
>
> ==如果使用云主机且有公网IP，一定要关闭密码校验登入==

## 概述

OpenSSH 以C/S模式，ssh做为客户端，sshd做为服务器(统一由systemd管理)。

## 配置文件

- ==TCPKeepAlive==

  保持长连接，服务端不会与客户端断开。需要在客户端同样设置才会有效。

- AllowUsers

  使用白名单机制，使用Pattern形式，支持CIDR地址，同理DenyUsers

- AllowGroups

  使用白名单机制，使用Pattern形式，同理DenyGroups

- AllowTcpForwarding

  是否允许TCP流量转发

- AuthenticationMethods

  指定认证的方式顺序，多种认证方式之间用空格隔开

- AuthorizedKeysFile

  指定存储用户公钥的文件位置，默认`.ssh/authorized_keys`，`.ssh/authorized_keys2`

- Banner

  指定客户端连接服务器时显示的banner

- ChrootDirectory

  指定chroot的目录，修改该参数可能会造成ssh无法连接

- Ciphers

  指定加密参数集

- ==ClientAliveCountMax==

  服务器发送keepalive信号到达指定次数后，服务器主动断开连接，默认为3

- ==ClientAliveInterval==

  服务器发送keepalive的时间间隔，默认为0

- Compression

  是否采用数据压缩

- DisableForwarding

  是否关闭所有的流量转发

- FingerprintHash

  指定登录时显示指纹的加密算法

- HostKey

  指定SSH服务器的密钥

- KdInteractiveAuthentication

  是否可以从键盘输入认证，默认与ChallengeResponseAuthentication值相同

- `ListenAddress hostname:port`

  指定sshd监听的地址和端口，默认为`127.0.0.1:22`，监听本机所有IP的22端口

- LoginGraceTime

  用户在登入界面超过指定等待时间后，服务器自动断开，如果没有单位默认为秒

  ```
  #LoginGraceTime 2m
  LoginGraceTime 3
  #PermitRootLogin yes
  #StrictModes yes
  #MaxAuthTries 6
  #MaxSessions 10
  ---
  root@chz:~# ssh -p 122 root@192.168.80.139
   __             _                   
  /   \/|_  _  __|_) _  |  o  _  _ __ 
  \__ / |_)(/_ | |  (/_ |  | (_ (_|| |
  
  root@192.168.80.139's password: 
  Connection closed by 192.168.80.139 port 122
  ```

- LogLevel

  设置服务器记录日志的等级，可以通过journal模块查看，不建议设置较高的日志等级，不利于安全

- MaxAuthTries

  登入认证最大次数(密码输入错误)，如果超过该指定次数自动断开连接，默认为6

- MaxSessions

  指定允许连接的最大session数，连接的session到达指定数量后，直接拒绝连接

  ```
  MaxSessions 1
  ---
  root@chz:~# ssh -p 122 root@192.168.80.139
   __             _                   
  /   \/|_  _  __|_) _  |  o  _  _ __ 
  \__ / |_)(/_ | |  (/_ |  | (_ (_|| |
  
  root@192.168.80.139's password: 
  Last login: Sun Nov  1 12:59:18 2020 from 192.168.80.200
  Sun Nov  1 13:04:30 CST 2020
  Startup finished in 683ms (kernel) + 1.628s (initrd) + 35.467s (userspace) = 37.779s
  [root@cyberpelican ~]# ssh -p 122 root@192.168.80.139
   __             _                   
  /   \/|_  _  __|_) _  |  o  _  _ __ 
  \__ / |_)(/_ | |  (/_ |  | (_ (_|| |
  
  root@192.168.80.139's password: 
  Authentication failed.
  
  ```

- MaxStartups

  指定为认证的界面的最大并发数，默认为10:30:100。如果设置为0表示没有限制

  10表示允许连接的数量，如果超过10后有30%的几率拒绝连接，当超过100时直接拒绝连接

- PasswordAurthentication

  是否允许密码校验，默认为yes。==为了安全需要将其设置为no，统一使用公钥校验==

- PermitEmptyPasswords

  是否允许空密码登入，默认为no

- PermitRootLogin

  是否允许root用户登入，默认为prohibit-password，==为了安全因该将其设置为no==











