---
createTime: 2024-11-19 14:37
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# ssh_config

## 0x01 Preface

OpenSSH Client `ssh` 的配置文件，默认会按照如下顺序(优先级)读取配置

1. command-line options
2. user's configuration file (`~/.ssh/config`)
3. system-wide configuration file (`/etc/ssh/ssh_config`)

`#` 开头表示 comments，以 `keyword arguments` 的形式配置

```
# this a comment
Host slokiv.com
```

其中 `keyword` 大小写不敏感，而 `arguments` 大小写敏感，每个 `keyword arguments` 单独一行

> [!importatn]
> 如果相同配置文件中或者是不同配置文件中有相同的 `keywords`，在没有特别说明的情况下(例如 `IdentityFile`，`Include` 就支持多个)，会使用第一个 `keywords` 的 `arguments`(没有 override 的逻辑)

### 0x01a Patterns

`ssh_config` 中的一些 keywords 会使用 patterns 做入参(例如 `Host`，`Match`)

patterns 由零个或者多个字符组成，还可以使用如下 wildcards

- `*`

	a wildcard that matches zero or more characters

	例如 `Host *.co.uk`

	匹配所有以 `co.uk` 结尾的 Host

- `？`  

	a wildcard that matches exactly one character

	例如 `Host 192.168.2.？`

	匹配 `192.168.2.[0-255]` 所有 Host

- `!`

	a wildcard that exclude matches

	例如 `Host !*.cn.gov`

	匹配所有不以 `cn.gov` 为结尾的 Host

patterns 也可以是一个 list

```
Host 10.0.1.? 10.0.3.*
```

### 0x01b Tokens

`ssh_config` 中的一些 `keywords` 还支持使用 tokens，会在运行时扩展，类似于 placeholder

keywords 如下

- `CertificateFile`, `ControlPath`, `IdentityAgent`, `IdentityFile`, `Include`, `KnownHostsCommand`, `LocalForward`, `Match exec`, `RemoteCommand`, `RemoteForward`, `RevokedHostKeys`, and `UserKnownHostsFile` accept the tokens `%%`, `%C`, `%d`, `%h`, `%i`, `%j`, `%k`, `%L`, `%l`, `%n`, `%p`, `%r`, and `%u`.

- `KnownHostsCommand` additionally accepts the tokens `%f`, `%H`, `%I`, `%K` and `%t`.

- `Hostname` accepts the tokens `%%` and `%h`.

- `LocalCommand` accepts all tokens.

- `ProxyCommand` and `ProxyJump` accept the tokens `%%`, `%h`, `%n`, `%p`, and `%r`.

支持的 token 含义如下

```
%%
    A literal ‘%’.
%C
    Hash of %l%h%p%r%j.
%d
    Local user's home directory.
%f
    The fingerprint of the server's host key.
%H
    The known_hosts hostname or address that is being searched for.
%h
    The remote hostname.
%I
    A string describing the reason for a KnownHostsCommand execution: either ADDRESS when looking up a host by address (only when CheckHostIP is enabled), HOSTNAME when searching by hostname, or ORDER when preparing the host key algorithm preference list to use for the destination host.
%i
    The local user ID.
%j
    The contents of the ProxyJump option, or the empty string if this option is unset.
%K
    The base64 encoded host key.
%k
    The host key alias if specified, otherwise the original remote hostname given on the command line.
%L
    The local hostname.
%l
    The local hostname, including the domain name.
%n
    The original remote hostname, as given on the command line.
%p
    The remote port.
%r
    The remote username.
%T
    The local tun(4) or tap(4) network interface assigned if tunnel forwarding was requested, or "NONE" otherwise.
%t
    The type of the server host key, e.g. ssh-ed25519.
%u
    The local username.
```

### 0x01c Environment Variables

以下 `keywords` 支持通过 `${}` 的形式获取系统变量

`CertificateFile`, `ControlPath`, `IdentityAgent`, `IdentityFile`, `Include`, `KnownHostsCommand`, `UserKnownHostsFile` 

例如

```
ControlPath ${HOME}/.ssh/%C
```

## 0x02 Keywords Arguments

> [!note]
> 具体看 manual page[^1]

其中一些 symbols EBNF 如下

- `destination = [user@]hostname, ssh://[user@]hostname[:port]`
- `ssh_config = /etc/ssh/ssh_config, ~/.ssh/config`

### 0x02a General Relate

#### `Include <path>`

引入其他配置文件

- 如果是 relative path，会从 `~/.ssh` 和 `/etc/ssh` 中引入对应的配置
- 如果是 absolute path，直接引入 absolute path 的配置

#### `Tag <name>`

定义可以复用的指令块，例如

```
Tag router
	User admin
	Port 5122
	HostKeyAlgorithms +ssh-rsa

Host 10.0.1.*
	Tag router
```

#### `GlobalKnownHostsFile <path>`

指定 global host key database 的路径，默认为 `/etc/ssh/ssh_known_hosts /etc/ssh/ssh_known_hosts2`

#### `UserKnownHostsFile`

指定 user host key database 的路径，默认为 `~/.ssh/known_hosts ~/.ssh/known_hosts2`

#### `HashKnownHosts <yes|no>`

当 host key 被写入到 `~/.ssh/known_hosts` 时，是否将 host names 和 host addresses 哈希，默认为 no

#### `HostKeyAlias <alias>` 

Hostname 会以 `alias` 记录到 host key database files 中

例如 有 `ssh_config` 如下配置

```
Host 10.0.3.101
	HostKeyAlias meta
```

如果使用 `ssh root@10.0.3.101` 登入服务器，那么 `~/.ssh/known_hosts` 会以 meta 作为 10.0.3.101 的 Hostname

```
meta ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBElzkI3CiVNdMu2pf8eXwsVD0+blka7RsutyAK04VFZcOHO7xhXR3Z2RsGprp+KynVn2Ff+jO5kZO8sLCLjwwyY=
```

#### `StrictHostKeyChecking <yes|no>`

是否强制检查 remote host-key，可以是如下几个值

- yes

	client 不会自动将 host-key 写入到 host-key database 中，且拒绝连接 host-key 发生改变的 server

- accept-new

	client 在第一次连接 server 时自动将 host-key 写入到 host-key database，且拒绝连接 host-key 发生改变的 server

- no

	client 在第一次连接 server 时自动会将 host-key 写入到 host-key database，允许连接 host-key 发生改变的 server

- ask

	client 在第一次连接 server 时询问用户是否将 host-key 写入到 host-key database，且拒绝连接 host-key 发生改变的 server

	缺省值

#### `CertificateFile <path>`

指定 user certificate，必须要和 `IdentityFile` 一起使用

#### `VisualHostKey <yes|no>`

登入 OpenSSH server 时是否显示 ASCII remote host key，默认为 no(只有 unknown host 会显示 string rmeote host key)

#### `BatchMode <yes|no>`

如果为 yes，不会显示 interactive prompt（例如 password prompt，fingerprint prompt），默认为 no

通常为了在脚本中调用 `ssh` 才使用

> [!note]
> 需要配置例如 publickey authentication 等不需要 prompt 的认证

### 0x02b Connection Related

#### `Host <patterns...>`

`Host` 下面的指令(到 next `Host` 或者是 `Match` 为止)，只对匹配 patterns 的 hostname 生效

例如 有如下 `ssh_config` 配置

```
Host github.com
  HostName ssh.github.com
  User git
  Port 443


```

如果使用 `ssh github.com` 就会自动使用 `User git`,`Port 443`

如果通过空格分隔指定多个 patterns，只要配置任意一个 patterns，下面的指令就生效

例如 有如下 `ssh_config` 配置

```
Host 10.0.3.1 10.0.2.1 10.0.6.1
     HostKeyAlgorithms +ssh-rsa
```

如果使用 `ssh root@10.0.2.1`，`ssh root@10.0.3.1` 或者是 `ssh root@10.0.6.1`  就会在 SSH authentication 阶段额外添加 ssh-rsa HostKeyAlgorithms

#### `Match <criteria...>`

`Match` 下面的指令(到 next `Host` 或者是 `Match` 为止)，只对匹配 criteria 的 Host 生效

比 `Host` 的匹配细粒度更高，支持如下几个 criterias

- canonical
- final
- exec
- localnetwork
- host
- originalhost
- tagged
- user
- localuser
- all

例如 有如下 `ssh_config` 配置

```
Match exec sh /usr/local/bin/check_server
  User nginx
```

不管 hostname 是什么，只要 command 是 `/usr/local/bin/check_server` 就会使用 `User nginx`

除了 canonical,final,exec,localnetwork 外。其他 criteria 可以通过逗号组合，表示要匹配所有 criterias，下面的指令才会生效

例如 有如下 `ssh_config` 配置

```
Match host 10.0.3.*, user root
  Port 443
```

当使用 `ssh root@10.0.3.1` 时，就表示 Server 会使用 443 端口；但是不适用其他用户

#### `BindAddress <ip_address>`

连接 destination 使用指定的 source address

#### `BindInterface <nic_name>`

连接 destination 使用指定的 source NIC

#### `User <username>`

登入 OpenSSH server 的用户名

例如

```
Host github.com
  HostName ssh.github.com
  User git
  Port 443
```

#### `Hostname <Hostname>`

用于实际登入 Hostname

例如 有如下 `ssh_config`

```
Host vini
	Hostname 10.0.1.89
```

那么就可以使用 `ssh root@vini` 登入 10.0.1.89

#### `Port <number>`

指定连接 destination OpenSSH server 的端口

例如

```
Host github.com
  HostName ssh.github.com
  User git
  Port 443
```

#### `PermitLocalCommand <yes|no>`

是否允许使用 `LocalCommand`，默认 no

#### `LocalCommand <command>`

当和 OpenSSH server 成功建立信道后，在 OpenSSH Client 侧执行 `command`，`PermitLocalCommand` 的值需要为 `yes`

例如

```
ssh -o PermitLocalCommand=yes -o LocalCommand='whoami' root@10.0.3.49
```

#### `RemoteCommand <command>`

当和 OpenSSH server 成功建立信道后，在 OpenSSH server 侧以同步执行命令，等价于 `ssh desination command`(默认不会提供 login shell)

#### `ControlMaster <yes|no|ask|auto|autoask>`

是否允许 sessions 共用一个 SSH authentication connection（会共用一个 TCP 连接，逻辑类似于 TCP keep alive）。新的 sessions 无需再做 SSH authentication，即如果是 password authentication 就不需要再次输入密码。需要和 `ControlPath` 一起使用

可以是如下几个值

- yes

	OpenSSH Client 会按照 `ControlPath` 中指定的路径生成 socket

	再次使用该 socket 无需 SSH authentication

	只用于第一次登入，后面登入需要将值置为 no，否则还是需要 SSH authentication

- no

	OpenSSH Client 不会按照 `ControlPath` 中指定的路径生成 socket

	只用于后面的登入，缺省值

- ask

	OpenSSH Client 会复用 TCP 连接，但是还是会做 SSH authentication

- auto

	如果 `ControlPath` 指定的 socket 不存在就会创建，如果存在就不会创建，**`ControlMaster` 通常会使用该值**

- autoask

	如果 `ControlPath` 指定的 socket 不存在就会创建，如果存在就不会创建，但是会有一个 confirmation prompt

#### `ControlPath <path>`

`ControlMaster` 使用的 socket 路径，通常和 `%h`，`%p`，`%r` 或者 `%C` tokens 一起使用

如果想要让 `ControlMaster` 只在 OpenSSH Client 当前登入的 login session 中有效，可以使用 `/run/usr/${user_id}` 或者 `/tmp`（出于安全性考虑，也应该使用这种方式）

- `ControlPersist <yes|no|time>`

	需要和 `ControlMaster` 一起使用，initial connection(即创建 socket 的 connnection) 关闭时，是否保留 socket

	- yes

		保留 socket

	- no

		不保留 socket

	- time

		在 time 后自动销毁 socket

#### `ConnectionAttempts <integer>`

连接 OpenSSH server TCP timeout 时，在 1 sec 内最大的尝试次数，默认为 1

#### `ConnectTimeout <integer>`

在 integer seconds 后如果还没有连接到 OpenSSH server，就会判断为 TCP timeout，默认使用系统处理 TCP 的逻辑(`tcp_syn_linear_timeouts`)

#### `serverAliveCountMax <number>`

类似于 `tcp_keepalive_probes`，OpenSSH Client 会发送 keep alive message 来判断 OpenSSH server 还是否存活。该参数用于指定断开连接需要发送的 keep alive message 数(如果回包，重置为 0)

默认为 3

#### `serverAliveInterval <number>`

类似于 `tcp_keepalive_intvl`，如果 OpenSSH server 没有对 keep alive message 回包，OpenSSH Client 会在 number seconds 后发送 keep alive message

默认为 0 表示永远也不会发送 keep alive packet

> [!note]
> OpenSSH 默认会使用 `TCPKeepAlive` 即系统的 tcp keep alive 逻辑

#### `TCPKeepAlive <yes|no>`

系统是否允许发送 tcp keep alive message，来判断对端是否存活

默认为 yes

#### `Compression <yes|no>`

是否压缩传输的数据，默认为 no

#### `RequestTTY <no|yes|force|auto>`

连接后是否分配 pseudo-TTY，等价于 `-t | -T`

- no

	不分配 TTY

- yes

	分配 TTY

- force

	强制分配 TTY

- auto

	自动判断是否需要分配 TTY

#### `SendEnv <Patterns>`

本地发送那些 environment variables 至 OpenSSH server，需要有 OpenSSH server `AcceptEnv` 的支持

值为 Patterns，默认只接收 `${TERM}`，但是如果 server 没有对应的 terminfo 就会导致 pseudo-TTY 乱码

例如 kitty 默认会使用 `TERM=kitty-xterm`。在你 SSH 登入 server 后，会设置 server 环境变量 `TERM=kitty-xterm`。而大多数 server 并不会安装 kitty，`terminfo` 也就没有 `kitty-xterm`，所以 pseduo-TTY 会出现乱码(kitty-ssh 会通过拷贝 terminfo 下的文件解决)

#### `Ciphers <ciphers...>`

指定加密的 ciphers，如果有多个 ciphers 需要使用 comma 分隔，从左到右优先级逐级递减

- 如果在 ciphers 前使用 `+` 表示在默认的 ciphers 后添加指定的 ciphers
- 如果在 ciphers 前使用 `-` 表示从默认的 ciphers 中移除指定的 ciphers
- 如果在 ciphers 前使用 `^` 表示在默认的 Ciphers 前添加指定的 ciphers

默认会使用如下 ciphers

`chacha20-poly1305@openssh.com,aes128-ctr,aes192-ctr,aes256-ctr,aes128-gcm@openssh.com,aes256-gcm@openssh.com`

支持使用如下 ciphers，也可以使用 `ssh -Q Ciphers` 查看

- 3des-cbc
- aes128-cbc
- aes192-cbc
- aes256-cbc
- aes128-ctr
- aes192     
- aes256-ctr
- aes128-gcm\@openssh.com
- aes256-gcm\@openssh.com
- chacha20-poly1305\@openssh.com

> [!note]
> 通常只有在一些 firmware 比较老的 routers 上可能需要使用该参数（因为不支持 openssh 的 new Ciphers） 

就可以在 `ssh_config` 中使用类似如下配置

```
Host 10.0.2.254
	Ciphers +aes128-cbc
```

#### `HostKeyAlgorithms <host-key-signature-algo ...>`

指定加密的 host-key-signature-algo，如果有多个 host-key-signature-algo 需要使用 comma 分隔，从左到右优先级逐级递减

- 如果在 host-key-signature-algo 前使用 `+` 表示在默认的 host-key-signature-algo 后添加指定的 host key signature aglorithm
- 如果在 host-key-signature-algo 前使用 `-` 表示从默认的 host-key-signature-algo 中移除指定的 host-key-signature-algo
- 如果在 host-key-signature-algo 前使用 `^` 表示在默认的 host-key-signature-algo 前添加指定的 host-key-signature-algo

默认为

`ssh-ed25519-cert-v01@openssh.com,ecdsa-sha2-nistp256-cert-v01@openssh.com,ecdsa-sha2-nistp384-cert-v01@openssh.com,ecdsa-sha2-nistp521-cert-v01@openssh.com,sk-ssh-ed25519-cert-v01@openssh.com,sk-ecdsa-sha2-nistp256-cert-v01@openssh.com,rsa-sha2-512-cert-v01@openssh.com,rsa-sha2-256-cert-v01@openssh.com,ssh-ed25519,ecdsa-sha2-nistp256,ecdsa-sha2-nistp384,ecdsa-sha2-nistp521,sk-ecdsa-sha2-nistp256@openssh.com,sk-ssh-ed25519@openssh.com,rsa-sha2-512,rsa-sha2-256`

可以使用 `ssh -Q HostKeyAlgorithms` 查看所有的 host key signature aglorithm

> [!note]
> 通常只有在一些 firmware 比较老的 routers 上可能需要使用该参数（因为不支持 openssh 的 new HostKeyAlgorithms）

就可以在 `ssh_config` 中使用类似如下配置

```
Host 10.0.2.254
	HostKeyAlgorithms +ssh-rsa
```

#### `KexAlgorithms <key-exchange-algorithms...>`

指定允许的 key-exchange-algorithms，如果有多个 key-exchange-algorithms 需要使用 comma 分隔，从左到右优先级逐级递减

- 如果在 key-exchange-algorithms 前使用 `+` 表示在默认的 key-exchange-algorithms 后添加指定的 host key-exchange-algorithms
- 如果在 key-exchange-algorithms 前使用 `-` 表示从默认的 key-exchange-algorithms 中移除指定的 key-exchange-algorithms
- 如果在 key-exchange-algorithms 前使用 `^` 表示在默认的 key-exchange-algorithms 前添加指定的 key-exchange-algorithms

默认为

`sntrup761x25519-sha512@openssh.com,curve25519-sha256,curve25519-sha256@libssh.org,ecdh-sha2-nistp256,ecdh-sha2-nistp384,ecdh-sha2-nistp521,diffie-hellman-group-exchange-sha256,diffie-hellman-group16-sha512,diffie-hellman-group18-sha512,diffie-hellman-group14-sha256casignaturealgorithms ssh-ed25519,ecdsa-sha2-nistp256,ecdsa-sha2-nistp384,ecdsa-sha2-nistp521,sk-ssh-ed25519@openssh.com,sk-ecdsa-sha2-nistp256@openssh.com,rsa-sha2-512,rsa-sha2-256`

可以使用 `ssh -Q KexAlgorithms` 查看所有的 KexAlgorithms

> [!note]
> 通常只有在一些 firmware 比较老的 routers 上可能需要使用该参数（因为不支持 openssh 的 new KexAlgorithms） 

就可以在 `ssh_config` 中使用类似如下配置

```
Host 10.0.2.254
	KexAlgorithms +diffie-hellman-group-exchange-sha1
```

#### `CheckHostIP <yes|no>`

如果为 yes，OpenSSH Client 会在第一次登入时，将对端的 hostname,resolved IP, figerprint 记录到 `~/.ssh/known_hosts` 中(默认只会记录 hostname)。在下次登入相同 hostname 时校验 resolved IP，如果不匹配，就不会和 OpenSSH server 建立连接，防止 MITMA。默认为 no

例如 `~/.ssh/known_hosts` 如下

```
openwrt.local ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIPZyMT01AnVsdtV+TCw3ZEclLe43e8MVFcc08cQMjFib
openwrt.local ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBDoZcuTW7jEu+8gvXzAi/rB/RwMFYiq361Amay5UjmIPKX6HoH+ntWYtsSknty4RaA0jLpUZ4uAFpo3MV6jf18k=
10.0.1.10 ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIPZyMT01AnVsdtV+TCw3ZEclLe43e8MVFcc08cQMjFib
10.0.1.10 ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBDoZcuTW7jEu+8gvXzAi/rB/RwMFYiq361Amay5UjmIPKX6HoH+ntWYtsSknty4RaA0jLpUZ4uAFpo3MV6jf18k=
```

如果 `CheckHostIP=yes`，那么就会出现如下错误

```
ssh -o CheckHostIP=yes  root@openwrt.local
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@       WARNING: POSSIBLE DNS SPOOFING DETECTED!          @
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
The ED25519 host key for openwrt.local has changed,
and the key for the corresponding IP address 10.0.1.101
is unknown. This could either mean that
DNS SPOOFING is happening or the IP address for the host
and its host key have changed at the same time.
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@    WARNING: REMOTE HOST IDENTIFICATION HAS CHANGED!     @
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
IT IS POSSIBLE THAT SOMEONE IS DOING SOMETHING NASTY!
Someone could be eavesdropping on you right now (man-in-the-middle attack)!
It is also possible that a host key has just been changed.
The fingerprint for the ED25519 key sent by the remote host is
SHA256:WebMA3wHEL1HRbcgawJfCUGCAOmUEhzRJY6ruISxH8I.
Please contact your system administrator.
Add correct host key in /root/.ssh/known_hosts to get rid of this message.
Offending ECDSA key in /root/.ssh/known_hosts:7
Host key for openwrt.local has changed and you have requested strict checking.
Host key verification failed.
```

### 0x02c Authentication Related

#### `HostbasedAuthentication <yes|no>`

是否允许 hostbased authentication，默认为 no

#### `KbdInteractiveAuthentication <yes|no>`
   
是否允许 keyboard interactive authentication，默认为 yes

#### `PubKeyAuthentication <yes|no|unbound|host-bound>`

是否开启 publickey authentication，默认为 yes

unbound 和 host-bound 表示是否允许在 publickey authentication 中使用 host-bound authentication protocol extension

#### `IdentitiesOnly <yes|no>` 

OpenSSH Client 是否只使用 public authentication 默认为 no

#### `IdentityFile <path>`

OpenSSH Client 使用的 private key

可以指定多个 `IdentityFile`，从上往下优先级递减

默认为

```
identityfile ~/.ssh/id_rsa
identityfile ~/.ssh/id_ecdsa
identityfile ~/.ssh/id_ecdsa_sk
identityfile ~/.ssh/id_ed25519
identityfile ~/.ssh/id_ed25519_sk
identityfile ~/.ssh/id_xmss
```

#### `PasswordAuthentication <yes|no>`

是否使用 password authentication，默认 yes

如果为 no 就不能使用密码登入

```
ssh  -o PasswordAuthentication=no  root@10.0.3.49
root@10.0.3.49: Permission denied (publickey,gssapi-keyex,gssapi-with-mic,password).
```

#### `NumberOfPasswordPrompts <number>`

password authentication 失败的重试次数，默认 3

#### `PreferredAuthentications <authentications>`

> [!NOTE]
> 无需在 OpenSSH Client 设置该参数，通过控制 `sshd_config` 中对应的配置即可

支持 OpenSSH client 尝试 authentication 的先后顺序

默认为

`gssapi-with-mic,hostbased,publickey,keyboard-interactive,password`

即如果 OpenSSH server 支持 publickey authentication 那么就会先使用 publickey authentication 而不是 password authentication

### 0x02d Port Forwardings Related

#### `DynamicForward <[local_address:]local_port>`

发送到 OpenSSH Client `[local_address:]local_port` 的流量会通过 destination secure channel 代理出去(只支持 socks4 和 socks5)。即 `[local_address:]local_port` 为 socks client，而 destination 为 socks server。其他形式同理
   
`local_address` 可以是空值或者是 `*` 表示本地的所有的 IP，但是如果 OpenSSH Client 没有设置 `GatewayPorts yes` 默认只会绑定回环地址(为了安全)
   
例如 在本地执行了 `ssh -D 1080 root@10.0.3.48`，那么访问本地的 `localhost:1080` 端口就会通过 `10.0.3.48` 代理出去

#### `LocalForward`

发送到 OpenSSH Client `[local_address:]local_port` 的流量会通过 destination secure shell 代理到 `remote_address:remote_port`。其他形式同理

`local_address` 可以是空值或者是 `*` 表示本地的所有的 IP，但是如果 OpenSSH Client 没有设置 `GatewayPorts yes`  或者 指定 `-g` 默认只会绑定回环地址(为了安全)

例如 在本地执行了 `ssh -L 1080:10.0.3.35:3306 root@10.0.3.36`，那么访问本地的 `localhost:1080` 端口就会通过 `10.0.3.36` 代理到 `10.0.3.35:3306`

#### `RemoteForward`

发送到 `[remote_address1:]remote_port1` 的流量会通过 destination secure shell 代理到 `remote_address2:remote_port2`(remote_address1 要和 destination address 一致)

`remote_address1` 可以是空值或者是 `*` 表示本地的所有的 IP，但是如果 OpenSSH server 没有设置 `GatewayPorts yes`  或者 指定 `-g` 默认只会绑定回环地址(为了安全)

例如 在 `10.0.3.40` 执行了 `ssh -R 1080：10.0.3.40:80 root@10.0.3.41`，那么 `10.0.3.41` 访问本地的 1080 就会通过 `10.0.3.49` secure shell 代理到 `10.0.3.40:80`

#### `GatewayPorts <yes|no>`

ssh 默认不允许 remote hosts 直接对执行 `-D`,`-L`,`-R` 的主机发起连接（只绑定回环地址），默认为 no
   	
开启该参数 remote hosts 可以连接本地的 port forwarding，类似于 clash 中的 `allow-lan`
   
例如 在本地执行了 `ssh -g -D 10.0.3.47:1080 root@10.0.3.48 -N`，那么其他的主机就可以通过访问 `10.0.3.48:1080` 代理到 `10.0.3.48`，然后通过 `10.0.3.48` 代理取出

#### `ProxyCommand <command>`

> [!note]
> 如今的代理工具(singbox, clash, etc.)几乎都使用了 pbr(policy-based-route) + tun/tap 实现了系统层面的代理，所以可以不需要该配置

指定用于连接 destination 前的代理命令（通常为 `nc`），需要 absolute path

> [!note]
> 有一点要注意的是 bsd netcat 和 nmap netcat 参数不同，具体看 manual page

例如

```
$ ssh -o ProxyCommand='/usr/bin/ncat -v --proxy 127.0.0.1:37898 --proxy-type socks5 %h %p' ssh://git@ssh.github.com:443
Ncat: Version 7.95 ( https://nmap.org/ncat )
Ncat: Connected to proxy 127.0.0.1:7898
Ncat: No authentication needed.
Ncat: Host ssh.github.com will be resolved by the proxy.
Ncat: connection succeeded.
PTY allocation request failed on channel 0
Hi dhay3! You've successfully authenticated, but GitHub does not provide shell access.
Connection to ssh.github.com closed.
```

通常会和 `%h` `%p` tokens 一起使用

#### `ProxyJump <SSH>`

按照链式顺序通过 SSH 代理连接 destination

通过如下格式

- `[user@]host[:port]`
- `ssh://[user@]hostname[:port]`

指定代理

例如

```
$ ssh -o ProxyJump=root@10.0.3.101,root@10.0.3.102 root@10.0.3.49
root@10.0.3.101's password:
root@10.0.3.102's password:
root@10.0.3.49's password:
```

`ProxyJump` 和 `ProxyCommand` 互斥，那个先出现就先使用那个

### 0x02e X11 Related

> [!note]
> X11 是不安全的，如果启用了 `ForwardX11Trusted yes` ，你用 firefox 打开一个会监听 keystrokes 的挂马网站，那么攻击者就可以知道你在 X11 server 上输入的内容
> 
> 所以应该要使用 Wayland 替代 X11，但是 OpenSSH 目前不支持 Wayland

#### `ForkAfterAuthentication <yes|no>`

执行 command 时，`ssh` 是否以 background 的形式运行（默认为 no）。通常和 X11 一起使用

例如

```
ssh -X -o ForkAfterAuthentication=yes root@10.0.3.49 firefox >& /dev/null
```

#### `ForwardX11 <yes|no>`

是否允许通过 OpenSSH 的信道传输 X11 数据(X11 本身是文明的)以及是否自动设置 `DISPLAY` 环境变量，默认为 no

#### `ForwardX11Trusted <yes|no>`

X11 clients(OpenSSH server) 可以完全控制 X11 server display(OpenSSH Client)，涵盖 `ForwardX11 yes`，默认为 no

例如 执行如下命令

```
ssh -o ForwardX11Trusted=yes root@10.0.3.40 firefox
```

那么 firefox 就可以读取 X11 server 的 clipboard 或者是对 X11 server 截屏，甚至还可以通过 JS 来监听键盘的输入。而 X11 clients 是可以获取到 firefox 对应的信息的，这些信息并不会加密

#### `ForwardX11Timeout <time>` 

在 time 后 untrusted X11 forwarding 会自动端口，默认为 20 mins

如果值为 0 表示，永远不会自动断开

### 0x02f Debug Related

#### `LogLevel <level>`

类似于 `ssh -v[v[v]]`，输出 `ssh` 连接 destination 时的详细信息，信息按照如下顺序逐级增加

- QUIET
- FATAL
- ERROR
- INFO
- VERBOSE
- DEBUG
- DEBUG1
- DEBUG2
- DEBUG3

例如

```
ssh -o LogLevel=DEBUG3 root@10.0.3.48
```

## 0x03 Example of ssh_config

```

```

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [ssh\_config(5) - OpenBSD manual pages](https://man.openbsd.org/ssh_config)

***References***

[^1]:[ssh\_config(5) - OpenBSD manual pages](https://man.openbsd.org/ssh_config)