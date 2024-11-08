# SSH 客户端

> 如果需要使用图形化界面编辑软件，例如gedit，mousepad 需要使用`-Y`或`-X`参数，开启X11 forward 
>
> 如果切换了用户，就不能使用X11 forward。
>
> 可以使用`ssh -vvvT hostname`来debug
>
> 建立连接后会在客户端创建一个known_host文件

## 远程连接

### login host

1. pattern：`ssh [user@]hostname`

   hostname可以时主机名，也可以是域名，也可以是IP地址或局域网内部主机名。==如果不指定用户名，将使用当前用户名，做为远程服务器的登录用户名。==

   ```
   PS C:\Users\82341> ssh root@192.168.80.200
    __             _
   /   \/|_  _  __|_) _  |  o  _  _ __
   \__ / |_)(/_ | |  (/_ |  | (_ (_|| |
   
   root@192.168.80.200's password:
   ```

2. pattern：`ssh://[user@]hostname[:port]`

   使用URI的形式，如果不指定端口，==默认使用22端口==

   ```
   PS C:\Users\82341> ssh ssh://root@192.168.80.200:22
    __             _
   /   \/|_  _  __|_) _  |  o  _  _ __
   \__ / |_)(/_ | |  (/_ |  | (_ (_|| |
   
   root@192.168.80.200's password:
   ```

### excute command

> 如果命令需要于用户交互，就需要使用`-t`参数
>
> ```
> PS C:\Users\82341> ssh -t ssh://root@192.168.80.200:22 "vim /opt/test.sh"
> ```

可以在login host形式末尾指定命令，直接在远程服务器上以==login-shell==的形式运行命令

```
PS C:\Users\82341> ssh ssh://root@192.168.80.200:22 "ls /root/"
 __             _
/   \/|_  _  __|_) _  |  o  _  _ __
\__ / |_)(/_ | |  (/_ |  | (_ (_|| |

root@192.168.80.200's password:
Desktop
Documents
Downloads
Music
Pictures
Public
revokey
src
Templates
Videos
```

由于login-shell没有`alias ll='ls -l'`所以不能使用`ll`

## 加密

SSH 连接握手阶段，客户端必须跟服务端约定加密参数集(cipher suite)。

加密参数集包含了若干不同的加密参数，它们之间使用下划线连接在一起，下面是一个例子。

```
TLS_RSA_WITH_AES_128_CBC_SHA
```

它的含义如下。

- TLS：协议
- RSA：密钥交换算法
- AES：加密算法
- 128：加密强度
- CBC：加密模式
- SHA：数字签名的 Hash 函数

## 参数

> X和Y参数好需要联系ForwardX11和ForwardX11Trusted
>
> https://askubuntu.com/questions/35512/what-is-the-difference-between-ssh-y-trusted-x11-forwarding-and-ssh-x-u

- `-i`

  指定连接使用的私钥，默认会在`~/.ssh/`下私钥。

- -X（也可以在配置文件中声明）

  开启X11 forwarding，remote machine会被作为untrusted X11 client，如果一些操作设计到安全问题，X11 server会拒绝显示X11 client的graphical output

- `-Y`（也可以在配置文件中声明）

  ==开启trusted X11 forwarding==，remote machine被作为trusted X11 client，==但是这涉及到安全问题，如果remote machine上的其他X11 client有问题，就会导致X11 server不安全或信息泄露==

  ```
  ssh -Y host
  ```

- `-f`

  以后台方式运行指定的命令，通常和X11一起使用

  ```
  ssh -p 65522 -f -o "ForwardX11 yes" ubuntu@82.157.1.137 mousepad
  ```

- `-Q`

  查询ssh支持的加密参宿集

  ```
  [root@cyberpelican ~]# ssh -Q cipher
  3des-cbc
  blowfish-cbc
  cast128-cbc
  arcfour
  arcfour128
  arcfour256
  aes128-cbc
  aes192-cbc
  aes256-cbc
  rijndael-cbc@lysator.liu.se
  aes128-ctr
  aes192-ctr
  aes256-ctr
  aes128-gcm@openssh.com
  aes256-gcm@openssh.com
  chacha20-poly1305@openssh.com
  ```

  查询支持的加密参数集

- `-q`

  quiet mode，不会输出报警和诊断信息

  ```
  root in /etc/ssh λ ssh -vvv -q localhost
  root in /etc/ssh λ 
  ```
  
- `-c`

  指定加密集数

  ```
   ┌─────( root)─────(~/.ssh) 
   └> $ ssh -c aes256-cbc root@192.168.80.143
  root@192.168.80.143's password: 
  Last login: Wed Dec 16 21:58:13 2020 from 192.168.80.200
  ```

- `-C`

  压缩数据

- `-o`

  指定参数，用于覆盖ssh客户端的配置文件

  ```
  ssh -o "ForwardX11 yes" host 
  ```

- `-p`

  指定连接的端口，默认为22号端口

- `-F`

  指定配置文件。默认如果有`~/.ssh/config`，优先级高于全局配置`/etc/ssh/ssh_config`

- `-T`

  认证成功后，不分配终端，用于校验ssh是否连接正常

- `-W host:port`

  将数据通过加密通道转发到指定主机的端口，用作代理

  ```
  ssh -q -W target_host:target_port proxy_host
  ```
  
  