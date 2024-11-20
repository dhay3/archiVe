---
createTime: 2024-11-19 14:37
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# OpenSSH 03 - ssh_config

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
> 如果相同配置文件中或者是不同配置文件中有相同的 keywords，在没有特别说明的情况下(例如 `IdentityFile`，`Include` 就支持多个)，会使用第一个 keywords 的 arguments(没有 override 的逻辑)

### 0x01a Token

`ssh_config` 中的一些 keywords 还支持使用 tokens，会在运行时扩展，类似于 placeholder

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

### 0x01b Environment Variables



## 0x02 Keywords Arguments



> [!note]
> 具体看 manual page[^1]

#### `Host <Pattern>`

下面的指令(到 next `Host` 或者是 `Match` 为止)，只对匹配 [0x03 Patterns](#0x03%20Patterns) 的 Host 生效

#### `Match <Pattern>`

下面的指令(到 next `Host` 或者是 `Match` 为止)，只对匹配 [0x03 Patterns](#0x03%20Patterns) 的 Host 生效，比 `Host` 的匹配精度更高，支持如下几个维度

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

#### `BatchMode`


#### `BindAddress <ip_address>`

指定连接 destination 使用的 source address

#### `BindInterface <nic_name>`

指定连接 destination 使用的 source NIC

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


#### `Compression <yes|no>`

是否压缩传输的数据，默认为 no

#### `ConnectionAttempts <integer>`

连接 OpenSSH Server TCP timeout 时，在 1 sec 内最大的尝试次数，默认为 1

#### `ConnectTimeout`

连接 OpenSSH Server TCP timeout 的时间，默认使用系统处理 TCP 的逻辑(`tcp_syn_linear_timeouts`)

#### `ControlMaster <yes|no|ask|auto|autoask>`

是否允许 sessions 共用一个 SSH authentication connection（会共用一个 TCP 连接，逻辑类似于 TCP keep alive）。新的 sessions 无需再做 SSH authentication，即如果是 password authentication 就不需要再次输入密码。需要和 `ControlPath` 一起使用

可以是如下几个值

- yes

	OpenSSH Client 会按照 `ControlPath` 中指定的路径生成 socket

	再次使用该 socket 无需 SSH authentication

	只用于第一次登入，后面登入需要将值置为 no，否则还是需要 SSH authentication

- no

	OpenSSH Client 不会按照 `ControlPath` 中指定的路径生成 socket

	只用于后面的登入

- ask

	OpenSSH Client 会复用 TCP 连接，但是还是会做 SSH authentication

- auto

	如果 `ControlPath` 指定的 socket 不存在就会创建，如果存在就不会创建，**`ControlMaster` 通常会使用该值**

- autoask

	如果 `ControlPath` 指定的 socket 不存在就会创建，如果存在就不会创建，但是会有一个 confirmation prompt

#### `ControlPath <path>`

`ControlMaster` 使用的 socket 路径，通常和 `%h`，`%p`，`%r` 或者 `%C` tokens 一起使用

如果想要让 ControlMaster 只在 Client 当前登入的 login session 中有效，可以使用 `/run/usr/${user_id}` 或者 `/tmp`（出于安全性考虑，也应该使用这种方式）

- `ControlPersist <yes|no|time>`

	需要和 `ControlMaster` 一起使用，initial connection(即创建 socket 的 connnection) 关闭时，是否保留 socket

	- yes

		保留 socket

	- no

		不保留 socket

	- time

		在 time 后自动销毁 socket

#### `DynamicForward`


#### `GatewayPorts`



#### `GlobalKnownHostsFile <path>`

指定 global host key database 的路径，默认为 `/etc/ssh/ssh_known_hosts /etc/ssh/ssh_known_hosts2`

#### `HashKnownHosts <yes|no>`

当 host key 被写入到 `~/.ssh/known_hosts` 时，是否将 host names 和 host addresses 哈希，默认为 no

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

#### `HostKeyAlias <alias>` 

Hostname 会以 `alias` 记录到 host key databases files 中

例如 有 `ssh_config` 如下配置

```
Host 10.0.3.101
	HostKeyAlias meta
```

如果使用 `ssh root@10.0.3.101` 登入服务器，那么 `~/.ssh/known_hosts` 会以 meta 作为 10.0.3.101 的 Hostname

```
meta ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBElzkI3CiVNdMu2pf8eXwsVD0+blka7RsutyAK04VFZcOHO7xhXR3Z2RsGprp+KynVn2Ff+jO5kZO8sLCLjwwyY=
```


#### `Hostname <Hostname>`

用于实际登入 Hostname

例如 有如下 `ssh_config`

```
Host vini
	Hostname 10.0.1.89
```

那么就可以使用 `ssh root@vini` 登入 10.0.1.89

#### `IdentitiesOnly <yes|no>` 

OpenSSH Client 是否只使用 PKI authentication 默认为 no

#### `IdentityFile <path>`

OpenSSH Client 使用的 PKI private key

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

#### `Include <path>`

引入其他配置文件

1. 如果是 relative path，会从 `~/.ssh` 和 `/etc/ssh` 中引入对应的配置
2. 如果是 absolute path，直接引入 absolute path 的配置

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


#### `LocalCommand <command>`

当和 OpenSSH Server 成功建立信道后，在 OpenSSH Client 侧以同步执行命令，`PermitLocalCommand` 的值需要为 `yes`

例如

```
ssh -o PermitLocalCommand=yes -o LocalCommand='whoami' root@10.0.3.49
```

#### `LocalForward`

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

#### `NumberOfPasswordPrompts <number>`

password authentication 失败的重试次数，默认 3

#### `PasswordAuthentication <yes|no>`

是否使用 password authentication，默认 yes

如果为 no 就不能使用密码登入

```
ssh  -o PasswordAuthentication=no  root@10.0.3.49
root@10.0.3.49: Permission denied (publickey,gssapi-keyex,gssapi-with-mic,password).
```

#### `PermitLocalCommand <yes|no>`

是否允许使用 `LocalCommand`，默认 no

#### `PermitRemoteOpen`


#### `Port <number>`

指定连接 destination OpenSSH Server 的端口

例如

```
Host github.com
  HostName ssh.github.com
  User git
  Port 443
```

#### `PreferredAuthentications <authentications>`

支持 OpenSSH client 尝试 authentication 的先后顺序

默认为

`gssapi-with-mic,hostbased,publickey,keyboard-interactive,password`

即如果 OpenSSH Server 支持 PKI authentication(publickey) 那么就会先使用 PKI authentication 而不是 password authentication

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

#### ProxyJump

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

#### PubKeyAuthentication

#### RemoteCommand

#### RemoteForward

#### SendEnv

#### User

### X11 Related

> [!note]
> X11 是不安全的，如果启用了 `ForwardX11Trusted yes` ，你用 firefox 打开一个会监听 keystrokes 的挂马网站，那么攻击者就可以知道你在 X11 Server 上输入的内容
> 
> 所以应该要使用 Wayland 替代 X11，但是 OpenSSH 目前不支持 Wayland

#### `ForkAfterAuthentication <yes|no>`

执行 command 时，`ssh` 是否以 background 的形式运行（默认为 no）。通常和 X11 一起使用

例如

```
ssh -X -o ForkAfterAuthentication=yes root@10.0.3.49 firefox >& /dev/null
```

#### `ForwardX11 <yes|no>`

是否允许通过 OpenSSH 的信道传输 X11 数据以及是否自动设置 `DISPLAY` 环境变量(想要在 X11 server 显示 GUI 必须使用该参数或者是 `-X`)

#### `ForwardX11Timeout <time>` 

在 time 后 untrusted X11 forwarding 会自动端口，默认为 20 mins

如果值为 0 表示，永远不会自动断开

#### `ForwardX11Trusted <yes|no>`

X11 clients 是否可以完全控制 X11 display(server)，默认为 no，即 untrusted X11 forwarding


## 0x03 Patterns

patterns 由零个或者多个字符组成，可以使用 wildcards

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


---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [ssh\_config(5) - OpenBSD manual pages](https://man.openbsd.org/ssh_config)

***References***

[^1]:[ssh\_config(5) - OpenBSD manual pages](https://man.openbsd.org/ssh_config)
