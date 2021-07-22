# SSH 客户端配置文件

> `man ssh_config`
>
> 如果客户端与服务端配置文件中参数相同，服务端配置优先启用。

- 全局配置`/etc/ssh/ssh_config`
- 个人配置`~/.ssh/config`

全局配置优于个人配置，使用配置文件也可以使用`-o`参数指定

## 语法

```
#comment
option value
等价
option = value
```

## tokens

一些关键字，可以在运行时扩展，类似于占位符。

```
           %%    A literal ‘%’.
           %C    Hash of %l%h%p%r.
           %d    Local user's home directory.
           %h    The remote hostname.
           %i    The local user ID.
           %L    The local hostname.
           %l    The local hostname, including the domain name.
           %n    The original remote hostname, as given on the command line.
           %p    The remote port.
           %r    The remote username.
           %T    The local tun(4) or tap(4) network interface assigned if tunnel forwarding was requested, or "NONE" otherwise.
           %u    The local username.

```

## 参数

> man ssh_config

- Host 

  用于匹配ssh中的hostname，可以使用通配符，`*`，`?`，`!`。如果给出了多个pattern，需要用空格隔开

  ```
  Host *.baidu.com *.sohu.com
  ```

  例如`ssh root@a.baidu.com`会自动匹配该规则

- AddKeysToAgent

  是否将keys 添加到ssh-agent

- AddressFamily

  允许连接的IP类型

- BindAddress

  当主机存在多IP时，指定IP做为源地址，默认第一个接口

- BindInterface

  当主机存在多NIC时，指定NIC做为源地址，默认第一张NIC

- CheckHostIP

  是否检查IP有效性，可以防止DNS spoofing，默认为yes

- Ciphers

  指定加密参数集，可以使用`+`，`-`，`^`等符号，优先级从高到低

- Compression

  是否压缩信息，默认no

- ConnectionAttempts

  客户端进行连接时，最大的尝试次数。默认为1

- ConnectTimeout

  客户端进行连接时，服务器在指定秒数内没有回复，则中断连接尝试

- EnableSSHKeysign

  在HostbasedAuthentication期间是否允许使用ssh-keysign来辅助，默认no

- FingerprintHash

  显示公钥指纹时使用的算法，默认sha256

- ==ForwardX11==

  允许X11，但是有安全限制

- ==ForwardTrustedX11==

  允许X11，无安全限制

- GlobalKnownHostsFile

  指定存储服务器公钥指纹的位置

- HostbaseAuthentication

  是否对远程主机进行校验，默认no

- HostKeyAlgorithms

  指定host key的加密算法，优先级从高到低

- HostkeyAlias

  为主机指定一个别名

- Hostname

  主机的真实域名或IP地址

- IdentityFile

  指定用户用于通讯的私钥，如果CertificateFile没有指定，默认从IdentityFile同一目录下的`-cert.pub`文件来校验

- KexAlgorithms

  指定密钥交换的算法

- LocalCommand

  ssh连接成功后，本地执行的命令，需要开启PermitLocalCommand

- `LocalForward <[address]:localport> <host:port>`

  address可以使用`*`通配符，port可以使用`0`通配符。指定本地TCP port转发到指定的host:port

- LogLevel

  ssh客户端记录日志的等级

- MACs

  指定数据校验的加密算法

- NumberofPasswordPrompts

  用户输入错误密码的最大尝试次数，只针对客户端。==在服务端设置才有意义==

- PasswordAuthentication

  是否允许密码认证登入，默认为yes

- PermitLocalCommand

  结合LocalCommand，默认no

- Port

  远程服务器连接的端口

- ==ProxyCommand==

  默认情况下ssh直接与目标主机的22号端口连接，使用该参数后将通过指定命令来建立连接。通常与netcat一起使用

  这里的%h表示的是当前指令块中的主机名，%p表示的是当前指令快中的端口

  ```
  ProxyCommand /usr/bin/nc -X 5 -x 127.0.0.1 %h %p
  ```

- ProxyJump

  指定中继

- PubkeyAcceptedKeyTypes

  指定服务器公钥认证方法

- PubkeyAuthentication

  是否允许使用公钥登入，默认no

- RemoteCommand

  登入成功后在，远程服务器上执行的命令

- `RemoteForward <[address]:localport> <host:port>`

  远程转发

- SendEnv

  将本地的环境变量发送到远程主机

- SeverAliveCountMax

  如果没有收到服务器的回应，客户端连续发送多少次keepalive信号，才断开连接。需要结合ServerAliveInterval，默认值为3

- ServerAliveInterval

  客户端发送keepalive信号的间隔

- StrickHostKeyChecking

  如果该值设置为yes，ssh不会自动将host keys添加到`known_hosts`中，拒绝连接host key改变的服务器，==能有效的防止MITM==；如果设置为accept-new会自动将host key添加到`known_hosts`中，但是拒绝连接host key改变的服务器；如果设置为ask，ssh会询问是否将host key 添加到`known_hosts`中，但是拒绝连接host key 改变的服务器

- TCPKeepAlive

  是否发送keepalive信号来检验服务器是否存活或网络中断

- User

  指定登入的用户

- UserKnownHostsFile

  指定known_host，默认`/.ssh/known_hosts, ~/.ssh/known_hosts2`

- VisualHostKey

  将服务器的公钥指纹以图形化的形式显示，默认no。==常用于校验服务器的公钥指纹是否有改变==

## ProxyCommand

如果使用的是公钥认证，使用的还是被代理的主机公钥，而不是代理主机的公钥。

如果是虚拟机于公网主机交互，默认会使用宿主机的username和host做为comment

```
Include /etc/ssh/ssh_config.d/*.conf

Host proxy
	Hostname 192.168.80.139
	Port 22
	User root

Host target
	Hostname 8.135.0.171
	Port 22
	User cpl
	#这里表示的是通过proxy转发到%h:%p
	ProxyCommand ssh -q -W %h:%p proxy
	
----

root in /etc/ssh λ ssh cpl@target
root@192.168.80.139's password: 
cpl@8.135.0.179's password: 
Welcome to Ubuntu 18.04.5 LTS (GNU/Linux 4.15.0-132-generic x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage

root in ~ λ cat .ssh/authorized_keys
ssh-rsa q7NaaGey2F+qp1rwWPhE96WsH65YcwhwjtuAWeWqD1oUkWdwo363pYgtLxZUa5epF9G8prkipPQpPig4hjlgLnfOq3coH/v55vGFJyYURNQkSZghMNT2h+KWCbE1Z8fAuzbDKdc/Skbx5ZlkIfxGUkPyoUuUshk= 82341@bash
ssh-rsa 0SVZmXnnTKl6LOXfEbScA5XsBz9+5ELTGKaOnz2H7Vpo9XKuDX7yBfFYk/4v5MhBPN9C5A7ZJ7s40mYigh2bdNt46+QYpZIu1mC095v+ty0mfW85EaYDTewL9Vrmm5fPf9dUU3FkU9UyPUZKY49P/OxJ5bMZbJiiFs= 82341@bash      
```

## example

```
#可以设置全局变量，对所有指令快生效
Compresssion yes
Host *
     Port 2222

Host remoteserver
     HostName remote.example.com
     User neo
     Port 2112
```

Host remoteserver 设置的port 会覆盖

```
ssh neo@remoteserver
```

