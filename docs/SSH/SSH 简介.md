# SSH 简介

参考：

https://wangdoc.com/ssh/basic.html

> 推荐使用public key authentication
>
> ==且为了安全性，最好在配置文件中关闭密码登入`PasswordAuthentication no`==

## 概述

SSH(Secure Shell)，是一种网络协议，用于加密两台计算机之间的通信，并支持各种身份验证机制。

## SSH vs 传统登入

**传统登入**

一个典型的例子就是服务器登录。登录远程服务器的时候，需要将用户输入的密码传给服务器，如果这个过程是明文通信，就意味着传递过程中，线路经过的中间计算机都能看到密码，这是很可怕的。

**SSH**

SSH 就是为了解决这个问题而诞生的，它能够加密计算机之间的通信，保证不被窃听或篡改。它还能对操作者进行认证（authentication）和授权（authorization）。明文的网络协议可以套用在它里面，从而实现加密。

## OpenSSH

> ssh协议默认端口为22

Linux所有的发行版，都使用OpenSSH，我们可以通过`-V`参数来查看

或是通过wireshark捕捉数据包查看客户端和服务端对应的版本

![Snipaste_2020-12-16_22-14-49](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/Snipaste_2020-12-16_22-14-49.5k89qclrdd40.png)

```
 ┌─────( root)─────(~) 
 └> $ ssh -V
OpenSSH_8.3p1 Debian-1, OpenSSL 1.1.1h  22 Sep 2020

```

OpenSSH中SSH client为客户端，而sshd为服务端。还用一些辅助工具`ssh-keygen`，`ssh-agent`，etc。

==所有的二进制文件一般存储在`/usr/bin/`下==

```
┌─────( root)─────(/usr/bin) 
 └> $ ll | grep ssh
lrwxrwxrwx root root        3 B  Sun Jun  7 08:44:04 2020 slogin ⇒ ssh
.rwxr-xr-x root root    762.8 KB Sun Jun  7 08:44:04 2020 ssh
.rwxr-xr-x root root    358.1 KB Sun Jun  7 08:44:04 2020 ssh-add
.rwxr-sr-x root ssh     334.1 KB Sun Jun  7 08:44:04 2020 ssh-agent
.rwxr-xr-x root root      1.4 KB Sun Jun  7 08:44:04 2020 ssh-argv0
.rwxr-xr-x root root     10.4 KB Tue May 26 20:38:00 2020 ssh-copy-id
.rwxr-xr-x root root    458.1 KB Sun Jun  7 08:44:04 2020 ssh-keygen
.rwxr-xr-x root root    442.1 KB Sun Jun  7 08:44:04 2020 ssh-keyscan
```

## 认证方式

OpenSSh 支持多种认证方式，host-based authentication（这种认证方式不安全，因该避免），public key authentication，password authentication。可以通过修改PreferredAuthentications 来修改默认的优先级

默认优先级

```
gssapi-with-mic,hostbased,publickey,
keyboard-interactive,password
```

- Public key authentication

  > 如果需要使用公钥认证，需要将连接者的公钥写入`~/.ssh/authorized_keys`中，

  ==使用这种方式可以免密登入，且安全==

  1. 服务器保存客户端的公钥，只有客户端自己知道自己的私钥。
  2. 服务器收到用户SSH登入请求，发送发送一些随机数据给用户，要求用户证明自己的身份。
  3. 客户端收到服务器发来的数据，==使用私钥对数据进行签名==，然后再发还给服务器。
  4. 服务器收到客户端发来的加密签名后，使用对应的公钥解密，然后跟原始数据比较。如果一致，就允许用户登录。

- password authentication

  只有所有认证都失败才会采用该方式。

  ==注意，这种方式存在MIMT，以但攻击者拿到服务器的公钥，就可以暴力破解服务器的密码==

  1. 远程主机将自己的公钥发送给客户端
  2. 用户使用服务器的公钥将登入密码加密，发送回来
  3. 服务器用自己的私钥解密登陆密码并校验

## 验证服务器

> 每台SSH服务器都有唯一的一对密钥对，fingerprint就是公钥的Hash值，用来识别服务器
>
> 密钥对一般存储在`/etc/ssh/`，不同的加密方式密钥不同

如果是第一次连接某一台服务器，会显示如下内容

```
The authenticity of host '192.168.10.200 (192.168.10.200)' can't be established.
ECDSA key fingerprint is SHA256:PrhyiqAgi2qz/sy2rmpB/r21Rj3i3mQkJ8ZrlpI7pW8.
Are you sure you want to continue connecting (yes/no)?
```

告诉用户这台服务器的fingerprint是陌生的，是否继续连接。

输入`yes`后就可以将服务器公钥的指纹存储在`~/.ssh/know_hosts`文件中

```
Warning: Permanently added '192.168.10.200' (ECDSA) to the list of known hosts.
```

### 密钥变更

如果服务器的密钥发送变更(比如重装的SSH服务器)，客户端再次连接时，就会发送公钥指纹不吻合的情况。

```
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@    WARNING: REMOTE HOST IDENTIFICATION HAS CHANGED!     @
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
IT IS POSSIBLE THAT SOMEONE IS DOING SOMETHING NASTY!
Someone could be eavesdropping on you right now (man-in-the-middle attack)!
It is also possible that the RSA host key has just been changed.
The fingerprint for the RSA key sent by the remote host is
77:a5:69:81:9b:eb:40:76:7b:13:04:a9:6c:f4:9c:5d.
Please contact your system administrator.
Add correct host key in /home/me/.ssh/known_hosts to get rid of this message.
Offending key in /home/me/.ssh/known_hosts:36
```

如果发送上述情况后，需要移除该服务器的fingerprint，才能再次连接

```
ssh-keygen -R host
```

同样的手动删除也可以

## 常用目录

> 具体查看 `man ssh` Files 模块
>
> 有些有文件有全局配置，但是个人配置优于全局配置

- `~/.rhosts`

  用于host-based authentication 校验

- `~/.ssh/`

  存放用户配置文件的默认路径，最好将该文件的用户权限设置为`r/w/x`，==但是不给予其他人任何权限==

- `~/.ssh/config`

  每个用户的配置文件

- `~/.ssh/environment`

  存放ssh额外配置的环境变量

- `~/.ssh/authorized_keys`

  ==存储在服务器上==表明用户身份的公钥

- ` ~/.ssh/id_<crypto>`

  存储用于认证的私钥

- `~/.ssh/id_<crypto>.pub`

  存储用于校验的公钥

- `~/.ssh/known_hosts`

  存储用户登入的主机的host keys

- `~/.ssh/rc`

  存储当用户登入后需要在执行的命令

  

