---
createTime: 2024-11-19 14:37
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# OpenSSH 0 - ssh_config

## 0x01 Preface

OpenSSH Client `ssh` 的配置文件，默认会按照如下顺序读取配置

1. command-line options
2. user's configuration file (`~/.ssh/config`)
3. system-wide configuration file (`/etc/ssh/ssh_config`)

`#` 开头表示 comments，以 `keyword arguments` 的形式配置，每个 `keyword arguments` 单独一行

> [!importatn]
> 如果相同配置文件中或者是不同配置文件中有相同的 keywords 会使用第一个 keywords 的 arguments(没有 override 的逻辑)

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

## 0x02 Keywords Arguments

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

chacha20-poly1305\@openssh.com,aes128-ctr,aes192-ctr,aes256-ctr,aes128-gcm@openssh.com,aes256-gcm@openssh.com

支持使用如下 ciphers

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

通常只有在一些 firmware 比较老的 routers 上可能需要使用该参数（因为不支持 openssh 的 cipher）

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


