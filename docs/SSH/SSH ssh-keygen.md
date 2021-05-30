# SSH ssh-keygen

> ssh-keygen生成的key会更网络有关，==如果网卡的MAC地址改变了，就会失效。==所以git使用代理时，最好指定一个算法

## 概述

用于生成和管理（撤回）SSH需要的密钥对。所有生成的密钥对都会存储在`~/.ssh/`下

## 创建密钥对

> 这里由于使用VisualHostKey，会将指纹内容转为randomart
>
> 默认以rsa做为加密算法，当前主机的用的用户即`username@hostname`来生成密钥

```
 ┌─────( root)─────(~/.ssh) 
 └> $ ssh-keygen
Generating public/private rsa key pair.
Enter file in which to save the key (/root/.ssh/id_rsa): 
Enter passphrase (empty for no passphrase): 
Enter same passphrase again: 
Your identification has been saved in /root/.ssh/id_rsa
Your public key has been saved in /root/.ssh/id_rsa.pub
The key fingerprint is:
SHA256:64VOqnqtXpsdXh9fmvgSIX0F2vybNA04MlKV/LzrMA0 root@cyberpelican
The key's randomart image is:
+---[RSA 3072]----+
|          .o.... |
|         .  o=  .|
|        . o.+o+. |
|         ..oo.+o.|
|        S  .Eo +o|
|         o  .oo +|
|     .. = o +..+.|
|    ...X + ..*.+ |
|  .++o+ =   o+*  |
+----[SHA256]-----+
```

==这里需要记住passphase，如果遗忘密钥对就作废了==

在创建完密钥对后，为了保证安全，还需要设置对应的权限

```
 ┌─────( root)─────(~/.ssh) 
 └> $ chmod 700 id_rsa*
```

## 公钥格式

```
ssh-ed25519  AAAAC3NzaC1lZDI1NTE5AAAAIHo5M7e4p+lx7Krb3cS+ov9Ub1UEdMgCHyfhYCo7825S 82341@bash
```

1. 加密算法
2. Base64加密后的密钥内容
3. 注释无实际作用

## 参数

- `-b`

  指定加密的位数，1024bit - 3072bit

- `-C`

  提供注释

  ```
  ┌─────( root)─────(~/.ssh) 
   └> $ ssh-keygen -C "this is a comment" 
  Generating public/private rsa key pair.
  Enter file in which to save the key (/root/.ssh/id_rsa): 
  /root/.ssh/id_rsa already exists.
  Overwrite (y/n)? y
  Enter passphrase (empty for no passphrase): 
  Enter same passphrase again: 
  Your identification has been saved in /root/.ssh/id_rsa
  Your public key has been saved in /root/.ssh/id_rsa.pub
  The key fingerprint is:
  SHA256:SgeU+orHtjP/JOJb/W0x0v+GFUMvU+k+/wrVff5dc0k comment
  The key's randomart image is:
  +---[RSA 3072]----+
  |      ..        .|
  |     ..        o.|
  |     ..       o..|
  |    .  .      o=o|
  |     .. S  .  oE*|
  |     ..+  . +..++|
  |   o..+ o  ..+ =*|
  |  ..Bo o . ...o O|
  |   o+*... ... .+*|
  +----[SHA256]-----+
  
  ```

- `-t <alogrithm>`

  指定生成key的加密算法，具体查看命令参数，默认使用ras

  常用加密，dsa，esdsa，esdsa-sk

- `-E <algorithm>`

  指定指纹显示的加密算法，md5 或 sha256

- `-f`

  指定读取的文件

- `-F <hostname | [hostname]:port>`

  从`known_hosts`中查看是否有指定主机的Hash值

- `-f filename`

  指定查看的文件

- `-l`

  显示指定公钥的指纹信息，结合`-v`参数显示图形化信息

  ```
  ┌─────( root)─────(~/.ssh) 
  └> $ ssh-keygen -l
  Enter file in which the key is (/root/.ssh/id_rsa): 
  3072 SHA256:SgeU+orHtjP/JOJb/W0x0v+GFUMvU+k+/wrVff5dc0k comment (RSA)
  ┌─────( root)─────(~/.ssh) 
   └> $ ssh-keygen -l -v
  Enter file in which the key is (/root/.ssh/id_rsa): 
  3072 SHA256:SgeU+orHtjP/JOJb/W0x0v+GFUMvU+k+/wrVff5dc0k comment (RSA)
  +---[RSA 3072]----+
  |      ..        .|
  |     ..        o.|
  |     ..       o..|
  |    .  .      o=o|
  |     .. S  .  oE*|
  |     ..+  . +..++|
  |   o..+ o  ..+ =*|
  |  ..Bo o . ...o O|
  |   o+*... ... .+*|
  +----[SHA256]-----+
  ```

- `-L`

  用于显示证书，一般与`-f`一起使用

  ```
  [root@ssh-server .ssh]# ssh-keygen -L -f host_key-cert.pub 
  host_key-cert.pub:
          Type: ssh-rsa-cert-v01@openssh.com host certificate
          Public key: RSA-CERT SHA256:13l7ndxFGBZk/shhFsAyxBI6dKEz5dEzFlIcpgrgALo
          Signing CA: RSA SHA256:gLc56ouH3TC2taKlNpAQnv5D0bj0tU2o1Kqffz7V/BQ
          Key ID: "host_ca"
          Serial: 0
          Valid: from 2020-12-29T13:15:00 to 2022-05-17T13:16:55
          Principals: 
                  192.168.80.200
          Critical Options: (none)
          Extensions: (none)
  ```

- `-i`

  输入未加密公钥(私钥)，输出对应的私钥(公钥)

- `-N <new_passphrase>`

  重新分配一个密码

  ```
  ┌─────( root)─────(~/.ssh) 
  └> $ ssh-keygen -N 123
  Generating public/private rsa key pair.
  Enter file in which to save the key (/root/.ssh/id_rsa): 
  /root/.ssh/id_rsa already exists.
  Overwrite (y/n)? y
  Your identification has been saved in /root/.ssh/id_rsa
  Your public key has been saved in /root/.ssh/id_rsa.pub
  The key fingerprint is:
  SHA256:Qa9StxBre2P7pmjnwR0INX3QK9RfmSKp0ZQeDxOgrH0 root@cyberpelican
  The key's randomart image is:
  +---[RSA 3072]----+
  |        o.*+=+  o|
  |      ...* Xo.+o.|
  |       oB *.*..o.|
  |      oo O +... .|
  |     ...SE* ..   |
  |       ..+ + .   |
  |          + .    |
  |        ...o.    |
  |       ..ooo.    |
  +----[SHA256]-----+
  
  ```

- `-R <hostname | [hostname]:port>`

  删除指定主机在客户端上的指纹

  ```
   ┌─────( root)─────(~/.ssh) 
   └> $ ssh-keygen -R 192.168.80.143
  # Host 192.168.80.143 found: line 3
  /root/.ssh/known_hosts updated.
  Original contents retained as /root/.ssh/known_hosts.old
  ```

- `-P passphrase`

  生成一个以该passphrase的私钥

  ```
   ┌─────( root)─────(~/.ssh) 
   └> $ ssh-keygen -P 123
  Generating public/private rsa key pair.
  Enter file in which to save the key (/root/.ssh/id_rsa): 
  /root/.ssh/id_rsa already exists.
  Overwrite (y/n)? y
  Your identification has been saved in /root/.ssh/id_rsa
  Your public key has been saved in /root/.ssh/id_rsa.pub
  The key fingerprint is:
  SHA256:sK9/sIAer9mHIqscDZQwRac50/1O/qOnIC9vv4bjDeI root@cyberpelican
  The key's randomart image is:
  +---[RSA 3072]----+
  |ooo .            |
  |...= .           |
  | o= . o          |
  |.  o   +         |
  | .   .. S        |
  |  o o .=.        |
  | . oooo++o       |
  |. o.+**=+.+      |
  |.o.oEB**B*..     |
  +----[SHA256]-----+
  ```

- `-p`

  修改指定私钥的passphrase

  ```
   ┌─────( root)─────(~/.ssh) 
   └> $ ssh-keygen -p
  Enter file in which the key is (/root/.ssh/id_rsa): 
  Enter old passphrase: 
  Key has comment 'root@cyberpelican'
  Enter new passphrase (empty for no passphrase): 
  Enter same passphrase again: 
  Your identification has been saved with the new passphrase.
  ```

- `-I`

  签名时指定key id，相当于注释，后面可以用作撤证书的ID

- `-h`

  签发证书时，指定该证书用于服务器

- `-n`

  指定生成证书

- `-S`

  签发证书时指定私钥，撤回证书时指定公钥

- `-V start-time:end-time`

  指定证书的有效时长

  **start-time**

  always：没有指定的起始时间

  **end-time**

  forever：永不失效

  **通用时间格式**

  YYYYMMDD | YYYYMMDD[SS]

  +当前时间向后，-当前时间向前

  w:week，d:day，m:month

  **例子**

  `+52w1d`：有效时间为现在52周1天

  `20201228:20201230`：有效时间为2020.12.28到2020.12.30

  `always:forever`：永不失效

  