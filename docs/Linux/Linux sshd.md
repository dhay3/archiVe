# Linux  sshd

参考：

https://www.cnblogs.com/ftl1012/p/ssh.html

https://segmentfault.com/a/1190000014667067

## 概述

以守护进程模式开启ssh服务，允许rmote terminal登入。ssh默认端口是22，安全协议版本是SSH2。配置文件在`/etc/ssh/sshd_config`

### 工作机制

1. 远程主机收到用户的登陆请求，将自己的公钥发送给用户
2. ==用户使用公钥将登陆密码加密，发送回来==
3. 远程主机用自己的私钥解密登陆密码并校验

这里存在MITM，同样需要知道用户的登入密码，攻击者可以对用户进行爆破

## ssh

pattern：`ssh user@host`

```
[root@cyberpelican .ssh]# ssh root@192.168.80.200
The authenticity of host '192.168.80.200 (192.168.80.200)' can't be established.
ECDSA key fingerprint is SHA256:PrhyiqAgi2qz/sy2rmpB/r21Rj3i3mQkJ8ZrlpI7pW8.
ECDSA key fingerprint is MD5:cf:7e:83:5e:63:1e:95:c3:51:b6:45:c3:53:83:2f:df.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '192.168.80.200' (ECDSA) to the list of known hosts.
root@192.168.80.200's password: 
Linux chz 5.7.0-kali1-amd64 #1 SMP Debian 5.7.6-1kali2 (2020-07-01) x86_64

The programs included with the Kali GNU/Linux system are free software;
the exact distribution terms for each program are described in the
individual files in /usr/share/doc/*/copyright.

Kali GNU/Linux comes with ABSOLUTELY NO WARRANTY, to the extent
permitted by applicable law.
Last login: Wed Oct 14 22:33:02 2020 from 192.168.80.131
Sat 31 Oct 2020 10:53:55 PM EDT
Startup finished in 7.081s (kernel) + 48.105s (userspace) = 55.186s 
graphical.target reached after 48.091s in userspace
root@chz:~# 

```

当登入后会将密匙和host存放在`~/.ssh/know_hosts`

```
[root@cyberpelican .ssh]# cat known_hosts 
192.168.80.139 ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGMDVunwcREN3Zl7cQ2KWWzcR/+fupESottdQVvMb/BcxEDjs6mknNy9ArjBBdjfvhF1v3jR9UkWSvFS+e8jG5k=
192.168.80.200 ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBO46y3PnoJa4rm9SfdEDETZqjxC1990vkOqA+MKGeEXzzVi73a1nZO/FofSG9tQjQZK7zuJQ7iGQeQx9MVqJYiQ=

```

可以使用`-p`参数指定目标主机ssh服务器的端口

## ssh-keygen

`ssh-keygen`用于生成当前主机ssh的公钥和私钥，一般将生成的密钥存储在`~/.ssh`。公钥为`id_rsa.pub`，私钥为`id_rsa`

- `-t dsa | ecdsa | ed25519 | rsa | rsa1`

  指定生成密匙类型，默认使用rsa，也可直接使用`ssh-keygen`

  ```
  [root@cyberpelican .ssh]# ssh-keygen 
  Generating public/private rsa key pair.
  Enter file in which to save the key (/root/.ssh/id_rsa): 
  Enter passphrase (empty for no passphrase): 
  Enter same passphrase again: 
  Your identification has been saved in /root/.ssh/id_rsa.
  Your public key has been saved in /root/.ssh/id_rsa.pub.
  The key fingerprint is:
  SHA256:GkmIx7O1JzMBGCEdUN5u70ohrJ3yvduydtWzFbpWEP2o root@cyberpelican
  The key's randomart image is:
  +---[RSA 2048]----+
  |oo+*.            |
  | .oo+o           |
  |  .o= +     .    |
  |   ..= +     .   |
  |   .. B S . o .  |
  |. o .  B   o =   |
  |oo o.    .o o .  |
  | =.+ o. +..E     |
  |o =.=o o.oo      |
  +----[SHA256]-----+
  ```

  这里的passphrase指的是密码短语，近似为salt。指纹摘要`SHA256:GkmIx7O1JzMBGCEUN5u70ohrJ3yvduydtWzFbpWEP2o root@cyberpelican`

  ```
  [root@cyberpelican .ssh]# ls
  id_rsa  id_rsa.pub  known_hosts
  ```

- `-p`

  使用新的passphrase替代旧的，不会重新生成私钥。用于私钥安全

  ```
  [root@cyberpelican cron.hourly]# ssh-keygen 
  Generating public/private rsa key pair.
  Enter file in which to save the key (/root/.ssh/id_rsa): 
  Enter passphrase (empty for no passphrase): 
  Enter same passphrase again: 
  Passphrases do not match.  Try again.
  Enter passphrase (empty for no passphrase): 
  Enter same passphrase again: 
  Your identification has been saved in /root/.ssh/id_rsa.
  Your public key has been saved in /root/.ssh/id_rsa.pub.
  The key fingerprint is:
  SHA256:m7nmEvP608d1FhU62n5Qxd+wdNa+qAB6geu5DmpwE6gBg root@cyberpelican
  The key's randomart image is:
  +---[RSA 2048]----+
  |               o |
  |              =o=|
  |E            + *B|
  |o.    . .     *oo|
  |=    .  So   +...|
  |o. .o o .+o . ...|
  |o =. o ++o = .   |
  | + o+ . +.. +    |
  |  oo...*+.       |
  +----[SHA256]-----+
  
  [root@cyberpelican cron.hourly]# ssh-keygen -p
  Enter file in which the key is (/root/.ssh/id_rsa): 
  Enter old passphrase: 
  Enter new passphrase (empty for no passphrase): 
  Enter same passphrase again: 
  ```

## ssh-copy-id

将本地的==加密后的公钥==拷贝到目标主机，以后登入目标主机无需使用密码

```
[root@cyberpelican .ssh]# ssh-copy-id -i ~/.ssh/id_rsa.pub chz@192.168.80.200
/usr/bin/ssh-copy-id: INFO: Source of key(s) to be installed: "/root/.ssh/id_rsa.pub"
/usr/bin/ssh-copy-id: INFO: attempting to log in with the new key(s), to filter out any that are already installed
/usr/bin/ssh-copy-id: INFO: 1 key(s) remain to be installed -- if you are prompted now it is to install the new keys
chz@192.168.80.200's password: 

Number of key(s) added: 1

Now try logging into the machine, with:   "ssh 'chz@192.168.80.200'"
and check to make sure that only the key(s) you wanted were added.
```

本机192.168.80.138，目标主机192.168.80.200。`-i`将本机指定的公钥发送给目标主机

## 配置文件

参考：https://blog.csdn.net/kalision/article/details/7481508

`etc/ssh/sshd_config`

- Banner

  指定连接ssh服务器时显示的banner

- Port

  > 如果修改了端口，连接SFTP服务时也需要修改端口

  修改Port可以起到一定的安全作用，指定连接ssh服务的端口，需要修改SELinux  domain

  ```
  [root@cyberpelican opt]# semanage port -a -t ssh_port_t -p tcp 122
  [root@cyberpelican opt]# semanage port  -l |grep ssh_port_t
  ssh_port_t                     tcp      122, 22
  [root@cyberpelican ssh]# iptables -F INPUT
  [root@cyberpelican ssh]# iptables -F OUTPUT
  [root@cyberpelican ssh]# netstat -lnpt |grep ssh
  tcp        0      0 0.0.0.0:122             0.0.0.0:*               LISTEN      17320/sshd          
  tcp6       0      0 :::122                  :::*                    LISTEN      17320/sshd   
  ```

  这里为了方便直接删除了防火墙中INPUT和OUTPUT链中的规则

- LoginGraceTime

  用户在登陆界面超过指定等待时间后，自动断开，如果没有单位默认为秒

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

- PermitRootLogin

  是否允许root登陆，默认yes，==为了安全需要将其设置为no==

- PermitEmptyPasswords

  是否允许空密码登入，默认no

- PermitAuthentication

  是否允许密码校验登入，默认yes

- UseDNS

  是否使用DNS解析主机名，默认yes，设置为no会加快登入

- MaxStartups

  允许未登入界面的并法数，10:30:60

  当到达10时，有30%的几率拒绝连接，到达60时全部拒绝

- MaxSessions

  连接的session到达指定数量后，直接拒绝连接

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

  
