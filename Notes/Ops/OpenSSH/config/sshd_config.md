---
createTime: 2024-11-19 14:02
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# sshd_config

> [!importantn]
> Contents

## 0x01 Preface

OpenSSH server `sshd` 的配置文件，如果 `sshd` 在没有使用 `-f` 的情况下，会默认使用 `/etc/ssh/sshd_config` 作为 `sshd` 的配置文件

`#` 开头表示 comments，以 `keyword arguments` 的形式配置，每个 `keyword arguments` 单独一行

例如

```
#this is a comment
Port 64422
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

### 0x01b Time Formats

如果 keyword 的 arguments 表示 time 时，`sshd_config` 支持使用 `time[qualifier]` 的格式

| qualifier | unit    | example                         |
| --------- | ------- | ------------------------------- |
| (none)    | seconds | 60  equals to 60 seconds        |
| s or S    | seconds | 60s or 60S equals to 60 seconds |
| m or M    | minutes | 60m or 60M equals to 60 minutes |
| h or H    | hours   | 60h or 60h equals to 60 hours   |
| d or D    | days    | 60d or 60D equals to 60 days    |
| w or W    | weeks   | 60w or 60W equals to 60 weeks   | 

还可以互相组合，例如

1h30m

### 0x01c Tokens

> [!note]
> 版本不同，支持的 Tokens 也不相同

`sshd_config` 中的一些 keywords 还支持使用 tokens，会在运行时扩展，类似于 placeholder

keywords 如下

- `AuthorizedKeysCommand` accepts the tokens `%%`, `%C`, `%D`, `%f`, `%h`, `%k`, `%t`, `%U`, and `%u`.

- `AuthorizedKeysFile` accepts the tokens `%%`, `%h`, `%U`, and `%u`.

- `AuthorizedPrincipalsCommand` accepts the tokens `%%`, `%C`, `%D`, `%F`, `%f`, `%h`, `%i`, `%K`, `%k`, `%s`, `%T`, `%t`, `%U`, and `%u`.

- `AuthorizedPrincipalsFile` accepts the tokens `%%`, `%h`, `%U`, and `%u`.

- `ChrootDirectory` accepts the tokens `%%`, `%h`, `%U`, and `%u`.

- `RoutingDomain` accepts the token `%D`.

支持的 token 含义如下

```
%%
    A literal ‘%’.
%C
    Identifies the connection endpoints, containing four space-separated values: client address, client port number, server address, and server port number.
%D
    The routing domain in which the incoming connection was received.
%F
    The fingerprint of the CA key.
%f
    The fingerprint of the key or certificate.
%h
    The home directory of the user.
%i
    The key ID in the certificate.
%K
    The base64-encoded CA key.
%k
    The base64-encoded key or certificate for authentication.
%s
    The serial number of the certificate.
%T
    The type of the CA key.
%t
    The key or certificate type.
%U
    The numeric user ID of the target user.
%u
    The username.
```

## 0x02 Keywords Arguments

### 0x02a General Related 

#### `Include <path>`

引入其他配置文件

- 如果是 relative path，会从 `/etc/ssh` 中引入对应的配置
- 如果是 absolute path，直接引入 absolute path 的配置

#### `Hostkey <path...>`

指定 OpenSSH server 使用的 private key

可以指定多个 `Hostkey`，从上往下优先级递减

默认为

```
hostkey /etc/ssh/ssh_host_rsa_key
hostkey /etc/ssh/ssh_host_ecdsa_key
hostkey /etc/ssh/ssh_host_ed25519_key
```

#### `StrictModes <yes|no>`

OpenSSH server 是否检查文件的 ownership，默认 yes

例如 

ssh 要求 `~/.ssh/authorized_keys` 的权限要为 600，如果 `StrictModes yes` 而 `~/.ssh/authorized_keys` 为 655 就会报错(noob should leave it)

### 0x02 Connection Related

#### `Match <criteria...>`

类似于 `ssh_config` `Match`

支持如下几个 criterias

single token criteria 直接使用

- all
- invalid-user

criteria-pattern pairs 需要给值

- user
- group
- host
- localaddress
- localport
- rdomain
- address

例如 有如下 `sshd_config` 配置

criteria 还可以通过逗号互相组合，表示要匹配所有 criterias，下面的指令才会生效

> [!NOTE]
> `Match` 下支持的 keywords 具体看 manual page

#### `ListenAddress <{hostname | address}[:port]>`

指定 OpenSSH server 监听的 IP/port，如果没有指定的 port，server 会监听 `Port` 指定的端口

默认

```
listenaddress [::]:22
listenaddress 0.0.0.0:22
```

即监听本机所有绑定的 IP

#### `Port <1-65535>`

指定 OpenSSH server 监听的端口，默认 22

#### `AllowUsers <patterns>`

允许那些用户登入，值为 patterns(不能使用 numerical user ID)，默认允许所有

如果使用类似 `USER@HOST` 的格式，会同时比对两个值(HOST 为 server 上的 IP address)

OpenSSH server 会先处理 `DenyUsers` 然后再处理 `AllowUsers`

#### `DenyUsers <patterns>`

`AllowUsers` 取反

#### `AllowGroups <patterns>`

允许那些 Groups 的用户登入，值为 patterns(不能使用 numerical group ID)，默认允许所有

OpenSSH server 会先处理 `DenyGroups` 然后再处理 `AllowGroups`

#### `DenyGroups <patterns>`

`AllowGroups` 取反

#### `AcceptEnv <patterns>`

接收 OpenSSH Client `SendEnv` 的那些 environment variables

值为 patterns，默认只接收 `${TERM}`

#### `MaxAuthTries <number>`

指定单次连接所有 authentication method 最大的尝试次数，默认为 6

例如 

使用如下指令开启一个 `sshd` 实例

```
/usr/local/sbin/sshd -p 6022 -o MaxAuthTries=1 -D
```

那么 `ssh` 连接到这个实例只要尝试任意一种 authentication method 失败，server 就会断开连接(所以 `ssh_config` 中的 `PreferredAuthentications` 很重要)

#### `MaxSessions <number>`

每个 network connection 允许的 SSH session 数，默认为 10

如果 `MaxSessions 0` 就表示拒绝所有的 SSH session 但是不包括 forwarding

如果 `MaxSessions 1` 就表示每个 network connection 允许 1 个 SSH session，如果没有关闭 session 的情况下就会拒绝新的 session

#### `LoginGraceTime <number>`

在 number(单位看 [0x01b Time Formats](#0x01b%20Time%20Formats)) 后如果用户还没有 successfully login，server 就会自动断开，默认 120s

#### `MaxStartups <number:percentage:number>`

OpenSSH server 同一时间允许的最大 unauthenticated connections，默认为 10:30:100

10 表示允许连接的数量，如果大于等于 10 有 30% 的几率拒绝连接，当大于等于 100 时直接拒绝连接

#### `ChannelTimeout <type=number>`

指定 OpenSSH server 在 number(单位看 [0x01b Time Formats](#0x01b%20Time%20Formats)) 后关闭 inactive type channel，默认为 none 表示所有 inactivte channels 都不会关闭

channel 如下

- agent-connect

	和 `ssh-agent` 建立的信道

- direct-tcpip

	由 client `DynamicForward`,`LocalForward` 打开的信道

- forwarded-tcpip

	由 client `RemoteForward` 打开的信道

- session

	login shell, command execution(`scp`,`sftp`) 打开的信道

- tun-connection

	由 client `TunnelFoward` 打开的信道

- x11-connection

	由 client `X11Forward`,`X11TrustedForward` 打开的信道

如果 number 为 0 则表示该 type channel 不会断开

#### `TCPKeepAlive <yes|no>`

OpenSSH server 是否向 OpenSSH Client 发送 TCP keep alive messages，默认为 yes

如果 OpenSSH Client 在指定时间(使用系统发送 TCP keep alive messages 的逻辑)内没有回包，OpenSSH server 就会主动断开连接

#### `ClientAliveCountMax <number>`

在收到 OpenSSH Client 发送了 number keep alive messages 后 OpenSSH server 断开 SSH Session(说明在 $number \times serverAliveInterval$ 内 server 到 Client 中的包都丢了)

默认为 3

#### `ClientAliveInteral <number>`

如果在 number seconds 内没有收到 OpenSSH Client 发送过来的报文，OpenSSH server 会要求 OpenSSH Client 回送报文

默认为 0，表示 OpenSSH server 不做要求

如果 `ClientAliveInterval 15`，那么如果 OpenSSH Client 在 45 seconds 内没有回包，就会断开 

#### `Compression <yes|no>`

是否允许传输压缩的数据(是否启用压缩由 OpenSSH Client 决定)，默认为 yes

#### `PermitTTY <yes|no>`

是否允许分配 pseudo-TTY，默认 yes

#### `PrintLastLog <yes|no>`

当 authentication passed，是否输出 last login 信息（例如 `Last login: Tue Nov 26 15:14:43 2024 from 10.100.13.47`），默认 yes

#### `PrintMotd <yes|no>`

当 authentication passed，是否输出 `/etc/motd`(message of today) 信息，默认 yes

#### `Banner <path>`

连接 OpenSSH server 时显示的 banner 路径

#### `Ciphers <ciphers...>`

指定允许的 ciphers，如果有多个 ciphers 需要使用 comma 分隔，从左到右优先级逐级递减

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

#### `HostKeyAlgorithms <host-key-signature-algo ...>`

指定加密的 host-key-signature-algo，如果有多个 host-key-signature-algo 需要使用 comma 分隔，从左到右优先级逐级递减

- 如果在 host-key-signature-algo 前使用 `+` 表示在默认的 host-key-signature-algo 后添加指定的 host key signature aglorithm
- 如果在 host-key-signature-algo 前使用 `-` 表示从默认的 host-key-signature-algo 中移除指定的 host-key-signature-algo
- 如果在 host-key-signature-algo 前使用 `^` 表示在默认的 host-key-signature-algo 前添加指定的 host-key-signature-algo

默认为

`ssh-ed25519-cert-v01@openssh.com,ecdsa-sha2-nistp256-cert-v01@openssh.com,ecdsa-sha2-nistp384-cert-v01@openssh.com,ecdsa-sha2-nistp521-cert-v01@openssh.com,sk-ssh-ed25519-cert-v01@openssh.com,sk-ecdsa-sha2-nistp256-cert-v01@openssh.com,rsa-sha2-512-cert-v01@openssh.com,rsa-sha2-256-cert-v01@openssh.com,ssh-ed25519,ecdsa-sha2-nistp256,ecdsa-sha2-nistp384,ecdsa-sha2-nistp521,sk-ecdsa-sha2-nistp256@openssh.com,sk-ssh-ed25519@openssh.com,rsa-sha2-512,rsa-sha2-256`

可以使用 `ssh -Q HostKeyAlgorithms` 查看所有的 host key signature aglorithm

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

### Authentication Related

#### `HostbasedAuthentication`

是否允许 host-based authentication，默认为 no

#### `GSSAPIAuthentication <yes|no>`

是否允许 gssapi authentication，默认为 no

#### `KerberosAuthentication <yes|no>`

是否允许 kerberos authentication，默认为 yes

#### `KbdInteractiveAuthentication <yes|no>`

是否允许 keyboard interactive authentication(早的时候被称为 `ChallengeResponseAuthentication`)，默认为 yes

#### `PubkeyAuthentication <yes|no>`

是否允许 publickey authentication，默认为 yes

#### `AuthorizedKeysFile <path>`

指定 publickey authentication 中用于验证用户 publickey 的文件，如果为相对路径，基点为 `${HOME}`，默认为 `.ssh/authorized_keys .ssh/authorized_keys2`

#### `RevokedKeys <path>`

path 文件中的 host keys 对应的用户会被拒绝 public key authentication

文件必须是可读的，否则会拒绝所有用户的 public key authentication 

#### `PasswordAuthentication <yes|no>`

是否允许 password authentication，默认为 yes

#### `PermitEmptyPasswords <yes|no>`

是否允许以空密码登入，默认为 no

#### `AuthenticationMethods <method...>`

用户必须要完成的 Authentication，以 comma 分隔或者是 list

例如 `publickey,password publickey,keyboard-interactive`

表示 用户必须要先完成 publickey authentication，然后再完成 password authentication 或者是 keyboard-interactive authentication；如果 publickey authentication 失败就不会做 password authentication 或者是 keyboard-interactive authentication

支持的 authentication methods 有

- gssapi-with-mic
- hostbased
- keyboard-interactive
- none

	在 `PermitEmptyPassword yes` 的情况下，允许使用空密码验证

- password
- publickey

其中 publickey 可以被连续使用，例如 `publickey,publickey` 表示使用 2 个不同的 publickey 验证

还有一个特殊值 any 表示，任意一个 authentication method 验证通过即可，也是缺省值 

#### `PermitRootLogin <yes|no|prohibit-password|forced-commands-only>`

是否允许以 root 的身份登入，默认为 prohibit-password

- yes

	允许以任何 authentication methods 登入

- no

	不允许以任何 authentication methods 登入

- prohibit-password

	不允许以 password 以及 keyboard authentication methods 登入

- forced-commands-only

	只允许 publickey authentication methods，但是要求 OpenSSH Client 要提供 command

### 0x02 Port Forwarding Related

#### `DisableForwarding <yes|no>`

是否关闭所有的 forwarding(X11，tcp，ssh-agent，stream)，默认为 no

#### `PermitOpen`

port forwarding 中 server 允许转发的目的端口，可以使用空格分隔指定多个端口，除了 `[host:]port` 外，还可以是

- none

	表示 server 限制 port forwarding 所有请求

- *

	可以用在 host 或者是 port，表示 server 允许 port forwarding 到所有 host 或者是 port

- any

	表示 server 不对 port forwarding 限制任何端口或者是地址，缺省值

#### `AllowTcpForwarding <yes|no|local|remote>`

是否允许 port fowarding

可以是如下几个值

- yes

	允许所有的 port fowarding，缺省值

- no

	不允许所有的 port forwarding

- local

	只允许 local port forwarding

- remote

	只允许 remote port forwarding

即使设置为 no 也不能保证安全，因为默认会分配 Shell，攻击者可以自己安装其他 port forward 的软件

#### `GatewayPorts <yes|no>`

在 `RemoteForward` 的场景，server 默认只会监听回环端口，所以 remote hosts 不能通过 server port forwarding 访问到 client

开启该参数 remote hosts 可以访问到 server port forwarding

#### `PermitListen <[host:]port>`

> [!NOTE]
> `sshd_config` `GatewayPorts` 也会影响 host

remote port forwarding 中 server 允许监听的端口(`remote_port1`)，可以使用空格分隔指定多个端口，除了 `[host:]port` 外，还可以是

- none

	表示 server 不允许 remote port forwarding 监听任何端口

- *

	表示 server 允许 remote port forwarding 监听所有端口

- any

	表示 server 不对 remote port forwarding 限制任何端口或者是地址，缺省值

### X11 Related

#### `X11Forwarding <yes|no>`

是否允许 X11 forwarding，默认为 no

### Debug Related

#### `LogLevel <level>`

指定 server 输出日志的详细程度

- QUIET
- FATAL
- ERROR
- INFO
- VERBOSE
- DEBUG
- DEBUG1
- DEBUG2
- DEBUG3

## 0x03 Example of sshd_config

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [sshd\_config(5) - OpenBSD manual pages](https://man.openbsd.org/sshd_config)

***References***


