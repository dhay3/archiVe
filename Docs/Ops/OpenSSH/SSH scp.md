# SSH scp

## 概述

OpenSSH secure file copy

pattern：`scp [option] <source> <target>`

使用ssh做为通道，用于安全传输数据。拷贝两台电脑之间的文件，如果文件存在会自动覆盖。如果使用publickey authentication同样也会跳过密码。

source和target需要如下格式，如果source只给出路径表示当前主机

`[user@]hostname:path`或`scp://[user@]host[:port][/path]`

## 参数

- `-B`

  不使用password authentication，使用其他认证方式

  ```
   ┌─────( root)─────(/opt) 
   └> $ scp -B root@192.168.80.143:/opt/test.sh /opt
  Host key fingerprint is SHA256:S6TRfoi/8wkrM74w95gjaTnPZApIKB3W2xJ9Pbghlyk
  +---[ECDSA 256]---+
  |                 |
  |   . . . =       |
  |  o o E X o      |
  |.o . + @ + .     |
  |o.. o + S .      |
  |o.   . o o       |
  |. . oo+ +        |
  |   .*O=+.+ .     |
  |   ..=B*+oo      |
  +----[SHA256]-----+
  root@192.168.80.143: Permission denied (publickey,gssapi-keyex,gssapi-with-mic,password).
  ```

- `-C`

  压缩传输

- `-p`

  从源文件保留时间戳，模式等。不使用该参数，等于拷贝生成一个新文件

- `-r`

  递归拷贝文件和目录

- `-v`

  输出详细信息

- `-F`

  指定ssh的配置文件，默认使用全局配置文件

## 案例

```
$ scp ~/.ssh/ssh_host_rsa_key-cert.pub root@host.example.com:/etc/ssh/
```

