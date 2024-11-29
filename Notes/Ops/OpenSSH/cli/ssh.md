---
createTime: 2024-11-18 12:02
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# ssh

## 0x01 Preface

OpenSSH 以 C/S 模式运行，而 `ssh` 是 OpenSSH client 用户态的工具，用于 log into a remote machine and execute commands on a remote machine

## 0x02 Syntax

```
ssh   [-46AaCfGgKkMNnqsTtVvXxYy]   [-B   bind_interface]   [-b   bind_address]   [-c  cipher_spec]
[-D  [bind_address:]port]  [-E  log_file]  [-e  escape_char]  [-F  configfile]   [-I   pkcs11]
[-i  identity_file]  [-J  destination] [-L address] [-l login_name] [-m mac_spec] [-O ctl_cmd]
[-o   option]   [-P   tag]   [-p   port]   [-R   address]   [-S   ctl_path]   [-W   host:port]
[-w local_tun[:remote_tun]] destination [command [argument ...]]

ssh [-Q query_option]
```

通常会使用如下 EBNF

```
ssh [options] desination [command [argument ...]]
```

其中 `destination` 可以是 2 种格式

1. `[user@]hostname`
2. `ssh://[user@]hostname[:port]`

如果没有指定 `user@` 但是使用了 `-l <login_name>`，就会使用 `login_name` 作为 remote machine 登入名；如果也没有使用 `-l <login_name>` 默认就会使用环境变量 `${USER}` 作为 remote machine 登入名

其中 `command` 表示要在 remote mahcine 上执行的命令（如果指定了 `command` 就不会提供 login shell）
例如

```
ssh 192.168.2.1 opkg install luci-app-openclash
```

## 0x03 Optional Args

大多数参数都可以在 `ssh_config` 中配置

> [!note]
> 具体请看 manual page[^1]

其中一些 symbols EBNF 如下

- `destination = [user@]hostname, ssh://[user@]hostname[:port]`
- `ssh_config = /etc/ssh/ssh_config, ~/.ssh/config`

### 0x03a Connection Related

- `-l <login_name>`

	`User` 命令行的形式

	指定 remote machine 使用的用户名

- `-p <port>`

	`Port` 命令行的形式

	指定 remote machine 使用的端口

- `-B <bind_interface>`

	`BindInterface` 命令行的形式

	使用指定的 NIC 和 `destination` 做网络连接

- `-b <bind_address>`

	`BindAddress` 命令行的形式

	使用指定的 IP 和 `destination` 做网络连接

- `-i <identity_file>`

	`IdentityFile` 命令行的形式

	指定 PKI 使用的 private key，如果没有指定默认会使用 `~/.ssh/id_rsa`, `~/.ssh/id_ecdsa`, `~/.ssh/id_ecdsa_sk`, `~/.ssh/id_ed25519` 以及 `~/.ssh/id_ed25519_sk`

- `-C`

	`Compression yes` 命令行的形式

	对连接传输的数据做 compression，通常只有在 poor network 中需要使用该参数，如果在 fast network 中使用该参数只会让数据显示更慢

- `-c <cipher_spec>`

	`Ciphers` 命令行的形式

	session 加密的 cipher

- `-F <configfile>`

	指定 ssh 使用的配置文件，如果没有指定，会先从 `/etc/ssh/ssh_config` 读取，然后再读取 `~/.ssh/config`

- `-o <option>`

	指定当前连接使用的 option(优先使用命令行指定的，而非配置文件)，常用的有
	
	- ControlMaster
	- ControlPath
	
	所有可用 options 具体看 `ssh_config(5)`

### 0x03b Port Forwardings Related

> [!NOTE]
> 具体看 [[OpenSSH 0 - Port Forwardings]]

- `-D <[local_address:]local_port>`

	`DynamicForward` 命令行的形式(因为实际访问的目的不固定)

	发送到 OpenSSH Client `[local_address:]local_port` 的流量会通过 destination secure channel 代理出去(只支持 socks4 和 socks5)。即 `[local_address:]local_port` 为 socks client，而 destination 为 socks server。其他形式同理

	`local_address` 可以是空值或者是 `*` 表示本地的所有的 IP，但是如果 OpenSSH Client 没有设置 `GatewayPorts yes` 或者 指定 `-g` 默认只会绑定回环地址(为了安全)

	例如 在本地执行了 `ssh -D 1080 root@10.0.3.48`，那么访问本地的 `localhost:1080` 端口就会通过 `10.0.3.48` 代理出去

- `-L <[local_address:]local_port:remote_address:remote_port>`

	`LocalForward` 命令行的形式

	发送到 OpenSSH Client `[local_address:]local_port` 的流量会通过 destination secure shell 代理到 `remote_address:remote_port`。其他形式同理
	
	`local_address` 可以是空值或者是 `*` 表示本地的所有的 IP，但是如果 OpenSSH Client 没有设置 `GatewayPorts yes` 或者 指定 `-g` 默认只会绑定回环地址(为了安全)

	例如 在本地执行了 `ssh -L 1080:10.0.3.35:3306 root@10.0.3.36`，那么访问本地的 `localhost:1080` 端口就会通过 `10.0.3.36` 代理到 `10.0.3.35:3306`

- `-R <[remote_address1:]remote_port1:remote_address2:remote_port2>`

	`RemoteForward` 命令行的形式

	发送到 `[remote_address1:]remote_port1` 的流量会通过 destination secure shell 代理到 `remote_address2:remote_port2`(remote_address1 要和 destination address 一致)

	`remote_address1` 可以是空值或者是 `*` 表示本地的所有的 IP，但是如果 OpenSSH server 没有设置 `GatewayPorts yes` 或者 指定 `-g` 默认只会绑定回环地址(为了安全)

	例如 在 `10.0.3.40` 执行了 `ssh -R 1080：10.0.3.40:80 root@10.0.3.41`，那么 `10.0.3.41` 访问本地的 1080 就会通过 `10.0.3.49` secure shell 代理到 `10.0.3.40:80`

- `-g`

	`GatewayPorts yes` 命令行的形式

	ssh 默认不允许 remote hosts 直接对执行 `-D`,`-L`,`-R` 的主机发起连接（只绑定回环地址）
	
	使用该参数 remote hosts 可以连接本地的 port forwarding，类似于 clash 中的 `allow-lan`

	例如 在本地执行了 `ssh -g -D 10.0.3.47:1080 root@10.0.3.48 -N`，那么其他的主机就可以通过访问 `10.0.3.48:1080` 代理到 `10.0.3.48`，然后通过 `10.0.3.48` 代理取出

- `-N`

	不执行 `command` 只做 SSH 隧道，通常和 `-D`,`-L`,`-R` 一起使用

- `-J <proxy...>`

	`ProxyJump` 命令行的形式

	先和 proxy 建立连接，再通过最后一个 proxy 和 destination 建立连接

### 0x03c X11 Related

- `-X`

	`ForwardX11 yes` 命令行的形式

	允许通过 OpenSSH 的信道传输 X11 数据(X11 本身是文明的)以及是否自动设置 `DISPLAY` 环境变量

- `-Y`

	`ForwardX11Trusted yes` 命令行的形式

	X11 clients(OpenSSH server) 可以完全控制 X11 server display(OpenSSH Client)，涵盖 `-X` 

### 0x03d Debug Related

- `-G`

	输出连接 `destination` 时 `ssh` 使用的详细配置

- `-E <log_file>`

	`LogLevel` 命令行的形式

	将 debug logs 输出到 log_file 而不是 stderr

- `-t`

	`RequestTTY force` 命令行的形式

	强制提供 login shell

- `-T`

	`RequestTTY no` 命令行的形式

	不提供 login shell

- `-V`

	打印版本信息并退出

- `-v`

	输出 debug 信息，多个 `-v` 可以增加详细度，最多 3

### 0x03e Misc Related

- `-f`

	`command` 会以异步的方式在 remote machine 中运行，而不会阻塞 client current shell

## 0x04 Authentication

OpenSSH 

- GSSAPI-based authentication
- host-based authentication
- public key authentication
- keyboard-interactive authentication
- password authentication

也可以通过修改 `PreferredAuthentications`


Host-Based Authentication

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [ssh(1) - OpenBSD manual pages](https://man.openbsd.org/ssh)

***References***

[^1]:[ssh(1) - OpenBSD manual pages](https://man.openbsd.org/ssh)
