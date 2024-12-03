---
createTime: 2024-11-19 10:34
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# OpenSSH 03 - Authentication Methods

> [!important]
> 不同 OS 的 OpenSSH 处理 authentication 的逻辑不同，所以本文中的例子可能会和你出现的现象有出入

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

client 和 server 默认开启，无需额外设置

为了方便记忆以及 debug，这里罗列出 client 和 server 开启 password authentication 需要的基础参数

#### Server `sshd_config`

```
PasswordAuthentication yes
AuthenticationMehods password
```

#### Client `ssh_config`

```
PasswordAuthentication yes
PreferredAuthentications password
```

#### Example

```
$ ssh -v -o PreferredAuthentications=password root@10.0.3.40
...
debug1: Authentications that can continue: publickey,password,keyboard-interactive
debug1: Next authentication method: password
root@10.0.3.40's password:
```

### Public Key Authentication

使用 asymmetric keypair 认证，通过 `PubkeyAuthentication` 设置是否开启

client 和 server 默认开启。但是 client 需要有 keypair，且 client public key 要在 server 的 authorized key files 中

为了方便记忆以及 debug，这里罗列出 client 和 server 开启 public key authentication 需要的基础参数

#### Server `sshd_config`

```
PubkeyAuthentication yes
AuthenticationMehods publickey
```

#### Client `ssh_config`

```
PubkeyAuthentication yes
PreferredAuthentications publickey
```

#### Example

要想使用 public key authentication，client 要有 keypair，如果 client 没有 keypair 可以使用 `ssh-keygen` 生成(OpenSSH 要求 RSA keypair min length 1024 bits)

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

当然 client 光有 keypair 还不行，还需要将 client public key 拷贝到 server authorized key files 中。这时我们可以通 `ssh-copy-id` 实现拷贝

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
> 在生成 keypair 时没有设置 passphrase，所以不需要提供任何 symmetric ciphers，直接可以登入

这里可以发现 client 会去尝试 `ssh_config` 中默认的 `IdentityFile` 即 private keys，然后提供对应的 public key 给 server 验证

因为在 asymmetric 中 public key 是基于 private key 的，所以可以从 private key 计算出 public key，但是你不能从 public key 计算出 private key。这也是 asymmetric 的一个特性

如果要想使用其他的 private keys 可以通过设置 `ssh_config` 中的`IdentityFile` 或者 `ssh -i` 指定来实现

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

#### User Certificate Authentication

user certificate authentication 是 public key authentication 的一种变种具体看 [OpenSSH 0 - Certificates#0x03 User Certificates Authentication](OpenSSH%200%20-%20Certificates.md)

### Keyboard Interactive Authentication

使用 keyboard interactive 认证

通过 `KbdInteractiveAuthentication` 设置是否开启，历史上也可以通过 `ChallengeResponseAuthentication` 来设置

client 和 server 默认开启

> [!note]
> BSD OpenSSH 无需额外设置即可开启，PAM based OS(例如 CentOS/Ubuntu) 需要在 `sshd_config` 中设置 `UsePAM yes` 才会开启
>
> 在写这篇文章时，手上没有 BSD OS，所以按照 PAM based OS OpenSSH 记录 

为了方便记忆以及 debug，这里罗列出 client 和 server 开启 keyboard interactive authentication 需要的基础参数

#### Server `sshd_config`

```
kbdInteractiveAuthentication yes
AuthenticationMehods keyboard-interactive
```

#### Client `ssh_config`

```
kbdInteractiveAuthentication yes
PreferredAuthentications keyboard-interactive
```

#### Example

这里输入用户密码即可

```
$ ssh -v  -o PreferredAuthentications=keyboard-interactive  -o KbdInteractiveAuthentication=yes root@10.0.3.49
...
debug1: Authentications that can continue: keyboard-interactive
debug1: Next authentication method: keyboard-interactive
(root@10.0.3.49) Password:
```

#### Google Authenticator[^1]

这么看你是不是觉得 keyboard interactive authentication 就是 password authentication

是的，实际上如果没有其他配置 keyboard interactive authentication 的行为就和 password authentication 的一致，都是通过用户密码来认证

但是 keyboard interactive 顾名思义，就是通过键盘交互来验证。这也就支持直接通过第三方的 authenticator 提供的 2FA code 来验证，例如 Google Authenticator

要想使用 Google Authenticator 必须要先在 sever 上安装组件

```
$ yum install google-authenticator
```

然后配置 Google Authenticator Agent

```
$ google-authenticator

Do you want authentication tokens to be time-based (y/n) y
Warning: pasting the following URL into your browser exposes the OTP secret to Google:

QRCODE

Your new secret key is: KSNNV42DX3RQ5CCJIR3KU656LU
Your verification code is 013579
Your emergency scratch codes are:
  75866119
  95786494
  70505096
  94100689
  92323415

Do you want me to update your "/root/.google_authenticator" file? (y/n) y

Do you want to disallow multiple uses of the same authentication
token? This restricts you to one login about every 30s, but it increases
your chances to notice or even prevent man-in-the-middle attacks (y/n) y

By default, a new token is generated every 30 seconds by the mobile app.
In order to compensate for possible time-skew between the client and the server,
we allow an extra token before and after the current time. This allows for a
time skew of up to 30 seconds between authentication server and client. If you
experience problems with poor time synchronization, you can increase the window
from its default size of 3 permitted codes (one previous code, the current
code, the next code) to 17 permitted codes (the 8 previous codes, the current
code, and the 8 next codes). This will permit for a time skew of up to 4 minutes
between client and server.
Do you want to do so? (y/n) y

If the computer that you are logging into isn't hardened against brute-force
login attempts, you can enable rate-limiting for the authentication module.
By default, this limits attackers to no more than 3 login attempts every 30s.
Do you want to enable rate-limiting? (y/n) y
```

然后需要在 `/etc/pam.d/sshd` 中启用 google-authenticator shared object

```
auth       required     pam_google_authenticator.so
```

使用 `sshd` 创建一个 server 实例

```
$ /usr/sbin/sshd -d -p 6022 -o KbdInteractiveAuthentication=yes -o AuthenticationMethods=keyboard-interactive -o PasswordAuthentication=no -o PubkeyAuthentication=no -o UsePAM=yes
```

要想登入服务器，我们需要同时提供 Google Authenticator code 和 用户密码才可以

```
$ ssh -v  -o PreferredAuthentications=keyboard-interactive  -o KbdInteractiveAuthentication=yes root@10.0.3.49
...
debug1: Authentications that can continue: keyboard-interactive
debug1: Next authentication method: keyboard-interactive
(root@10.0.3.49) Verification code:
(root@10.0.3.49) Password:
Authenticated to 10.0.3.49 ([10.0.3.49]:22) using "keyboard-interactive".
...
```

> [!note]
> 因为 keyboard interactive authentication 同样也需要用户密码，所以通常 `KbdInteractiveAuthentication yes` 和 `PasswordAuthentication yes` 二选一

### HostBased Authentication

> [!note]
> 从 `rlogin` 中移植过来的逻辑

使用 host 白名单认证，通过 `HostbasedAuthentication` 设置是否开启

client 和 server 默认关闭。client public key 要在 server host key database files，且 server 要在 `/etc/hosts.equiv` 中对 client 开白名单



为了方便记忆以及 debug，这里罗列出 client 和 server 开启 host-based authentication 需要的基础参数

#### Server `sshd_config`

```
HostbasedAuthentication yes
AuthenticationMehods hostbased
```

#### Client `ssh_config`

```
HostbasedAuthentication yes
PreferredAuthentications hostbased
```

#### Example





如果 ``

```
/usr/sbin/sshd -D -p 6022 -o HostbasedAuthentication=yes -o MaxAuthTries=10
```

#### host.equiv



```
+|[-]hostname|+@netgroup|-@netgroup [+|[-]username|+@netgroup|-@netgroup]
```




### GSSAPI Authentication

使用 GSSAPI(Generic Security Service API[^2]) 认证，通过 `GSSAPIAuthentication` 设置是否开启

> [!note]
> GSSAPI

client 和 server 默认关闭

为了方便记忆以及 debug，这里罗列出 client 和 server 开启 gssapi authentication 需要的基础参数

> [!note]
> mic(message integrity code)

#### Server `sshd_config`

```
GSSAPIAuthentication yes
AuthenticationMehods gssapi-with-mic
```

#### Client `ssh_config`

```
GSSAPIAuthentication yes
PreferredAuthentications gssapi-with-mic
```

#### Example


### Kerberos Authentication

> [!note]
> kerberos 不存在单独的认证方式

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- `man hosts.equiv.5`
- `man ssh`
- `man sshd`
- `man ssh_config`
- `man sshd_config` 

***References***

[^1]:[Configure SSH to use two-factor authentication    | Ubuntu](https://ubuntu.com/tutorials/configure-ssh-2fa#1-overview)
[^2]:[Generic Security Services Application Program Interface - Wikipedia](https://en.wikipedia.org/wiki/Generic_Security_Services_Application_Program_Interface)

