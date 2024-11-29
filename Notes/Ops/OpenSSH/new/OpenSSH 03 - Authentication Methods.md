---
createTime: 2024-11-19 10:34
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# OpenSSH 03 - Authentication Methods

## 0x01 Preface

OpenSSH 一个完整的 Authentication 分为 2 部分

- 



## Client

OpenSSH client 默认会按照如下的顺序去尝试鉴权

- GSSAPI-based authentication
- host-based authentication
- public key authentication
- keyboard-interactive authentication
- password authentication

也可以通过修改 `ssh_config` 中的 `PreferredAuthentications` 来修改鉴权方式或者是顺序

## Server

OpenSSH server 通过 

- `AuthenticationMethod`

## Authentication Methods

> [!note] 
> client 是否使用对应的 authentication methods 由 `<method>Authentication` 决定
> client 使用的 authentication methods 先后顺序，由 `PreferredAuthentications` 来决定

ssh 

### Password Authentication

使用账号密码认证，通过 `PasswordAuthentication` 设置是否开启

client 和 server 默认开启。无需额外设置

例如

```
ssh -v -o PreferredAuthentications=password root@10.0.3.40
...
debug1: Authentications that can continue: publickey,password,keyboard-interactive
debug1: Next authentication method: password
root@10.0.3.40's password:
```

### Public Key Authentication

使用 asymmetric keypair 验证，通过 `PubkeyAuthentication` 设置是否开启

client 和 server 默认开启。client 需要有 keypair，且 client public key 要在 server 的 authorized key files 中

如果 client 没有 keypair 可以使用 `ssh-keygen` 生成(OpenSSH 要求 RSA keypair min length 1024 bits)

例如

```
$ ssh-keygen -t rsa
Generating public/private rsa key pair.
Enter file in which to save the key (/root/.ssh/id_rsa):
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Your identification has been saved in /root/.ssh/id_rsa
Your public key has been saved in /root/.ssh/id_rsa.pub
The key fingerprint is:
SHA256:7ND90VNqGD7JAVbInWvT6F31D8ycPaR6cP1e+IqPvgY root@localhost.localdomain
The key's randomart image is:
+---[RSA 3072]----+
|         .o+..   |
|         .o.o  ..|
|            oB+o+|
|       o . +=OB=+|
|      . S .o@o=++|
|       o  Eo.=o.+|
|        .  .o  o.|
|            .o  o|
|           o=oo. |
+----[SHA256]-----+
```

然后可以通过 `ssh-copy-id` 将 client public key 以指定格式拷贝到 server authorized key files 中

例如

```
$ ssh-copy-id root@10.0.3.40
/usr/bin/ssh-copy-id: INFO: attempting to log in with the new key(s), to filter out any that are already installed
/usr/bin/ssh-copy-id: INFO: 1 key(s) remain to be installed -- if you are prompted now it is to install the new keys
root@10.0.3.40's password:

Number of key(s) added: 1

Now try logging into the machine, with:   "ssh 'root@10.0.3.40'"
and check to make sure that only the key(s) you wanted were added.
```

> [!NOTE]
> 这里需要借助任意一种 server 允许的 authentication method 完成拷贝，否则需要手动拷贝

然后就可以使用 publickey authentication 发起认证

例如

```
ssh -v -o PreferredAuthentications=publickey root@10.0.3.40
...
debug1: Authentications that can continue: publickey,gssapi-keyex,gssapi-with-mic,password
debug1: Next authentication method: publickey
debug1: Will attempt key: /root/.ssh/id_rsa RSA SHA256:7ND90VNqGD7JAVbInWvT6F31D8ycPaR6cP1e+IqPvgY
debug1: Will attempt key: /root/.ssh/id_ecdsa
debug1: Will attempt key: /root/.ssh/id_ecdsa_sk
debug1: Will attempt key: /root/.ssh/id_ed25519
debug1: Will attempt key: /root/.ssh/id_ed25519_sk
debug1: Will attempt key: /root/.ssh/id_xmss
debug1: Offering public key: /root/.ssh/id_rsa RSA SHA256:7ND90VNqGD7JAVbInWvT6F31D8ycPaR6cP1e+IqPvgY
debug1: Server accepts key: /root/.ssh/id_rsa RSA SHA256:7ND90VNqGD7JAVbInWvT6F31D8ycPaR6cP1e+IqPvgY
Authenticated to 10.0.3.40 ([10.0.3.40]:22) using "publickey".
...
```

> [!note]
> 在生成 keypair 时没有设置 passphrase，所以不需要提供任何 symmetric ciphers

这里可以发现 client 会去尝试 `ssh_config` 中默认的 `IdentityFile` 即 private keys，然后提供对应的 public key 给 server 验证

因为在 asymmetric 中 public key 是基于 private key 的，所以可以从 private key 计算出 public key，但是你不能从 public key 计算出 private key

所以如果要想使用其他的 private keys 可以通过设置 `ssh_config` 中的`IdentityFile` 或者 `ssh -i` 指定来实现

例如

```
ssh -v -i /tmp/id_rsa -o PreferredAuthentications=publickey root@10.0.3.40
...
debug1: Authentications that can continue: publickey,gssapi-keyex,gssapi-with-mic,password
debug1: Next authentication method: publickey
debug1: Will attempt key: /tmp/id_rsa RSA SHA256:7ND90VNqGD7JAVbInWvT6F31D8ycPaR6cP1e+IqPvgY explicit
debug1: Offering public key: /tmp/id_rsa RSA SHA256:7ND90VNqGD7JAVbInWvT6F31D8ycPaR6cP1e+IqPvgY explicit
debug1: Server accepts key: /tmp/id_rsa RSA SHA256:7ND90VNqGD7JAVbInWvT6F31D8ycPaR6cP1e+IqPvgY explicit
Authenticated to 10.0.3.40 ([10.0.3.40]:22) using "publickey".
...
```

#### Authorized-Keys Files

authorized keys files 是 publickey authentication 中认证的核心。server 会校验 client 发送过来的 public key 是否在 authorized keys files 中，如果在则认证通过，否则失败

server 默认会从如下路径读取 authorized keys files

- `~/.ssh/authorized_keys`
- `~/.ssh/authorized_keys2`

也可以通过设置 `sshd_config` 中的 `AuthorizedKeysFile` 来自定义

authorized key 由几部分组成，通过空格分隔

- options(optional)

	通常不会指定，具体看 `sshd.8`

- keytype

	生成 client keypair 的加密算法

	只能是如下几个值

	- `sk-ecdsa-sha2-nistp256@openssh.com`
	- `ecdsa-sha2-nistp256`
	- `ecdsa-sha2-nistp384`
	- `ecdsa-sha2-nistp521`
	- `sk-ssh-ed25519@openssh.com`
	- `ssh-ed25519`
	- `ssh-rsa`

- base64-ecoded key

	client public key 实际上是 raw binary，为了方便识别，对 client public key 做 base64 encode；也被称为 client public key fingerprint(a unique id represent that keypair)

- comment(optional)

	备注

例如

```
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCax3m23XdPFMY0K/HdiQ5N+HqzbEP85suloUxiBqVOE88NLqVQFXmCc87E/4eMoYztGh69jhpHHUc+msjPGST/tH7BHf0wCpvJR+HaAEnQ1gxkNmgRTTu8bUxX6nmrbBS06tpFJF8WOveSrWTieLkCkqdvc6qNQaDElepquSihtVt2NkIiG20OaTlLILWCWzzEYEUnxBWoR3cCssya0Qsf8d0OWky7VNnkVmTW3x3fFFaBNnrAtTrmi0Dhxpg2UNJb1otA2Xc3vpSXSFuZ1ILUi2qP82tcR1d+UymNaPJdDyiNQGduoimNxomObAspwwuGXURCZ5qtp2YaZF6bvhSPeI/CHWmXEioyGshtcY6dT6yfjWBg5U2ot8n2ebf2INuwGOZjn7LQI5ubOYUJRjQKhOFNYcokZbl/w06LNmKnWRIc5EuwN1BtX50ISMYl7Uoka74DA7MjSm8w9MpGwzxg0TbNnE/6mYJ6iqeJ6QhhlO87znFgI5I+Tb4c/tNpvJ7bU= (none)
```

### Host-Based Authentication

> [!note]
> 从 `rlogin` 中移植过来的逻辑

根据 `/etc/hosts.equiv` 


hostname 和 username 来认证

如果 ``

```
/usr/sbin/sshd -D -p 6022 -o HostbasedAuthentication=yes -o MaxAuthTries=10
```

#### host.equiv



```
+|[-]hostname|+@netgroup|-@netgroup [+|[-]username|+@netgroup|-@netgroup]
```



### Keyboard Authentication




### GSSAPI Authentication





---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- `man hosts.equiv.5`

***References***


