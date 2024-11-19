---
createTime: 2024-11-18 12:02
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# OpenSSH 02 - ssh cli

## 0x01 Preface

OpenSSH 以 C/S 模式运行，而 `ssh` 是 OpenSSH Client 用户态的工具，用于 log into a remote machine and execute commands on a remote machine

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
ssh desination [command [argument ...]]
```

其中 `destination` 可以是 2 种格式

1. `[user@]hostname`
2. `ssh://[user@]hostname[:port]`

如果没有指定 `user@` 但是使用了 `-l <login_name>`，就会使用 `login_name` 作为 remote machine 登入名；如果也没有使用 `-l <login_name>` 默认就会使用环境变量 `${USER}` 作为 remote machine 登入名

如果指定了 `command` 就不会提供 login shell，相对地会执行 `command`

例如

```
ssh 192.168.2.1 opkg install luci-app-openclash
```

## 0x02 Optional Args

> [!note]
> 具体请看 manual page[^1]

### Log into Relate

- `-l <login_name>`

	指定 remote machine 使用的用户名

- `-p <port>`

	指定 remote machine 使用的端口

- `i <identity_file>`

	指定 PKI 使用的 private key，如果没有指定默认会使用 `~/.ssh/id_rsa`, `~/.ssh/id_ecdsa`, `~/.ssh/id_ecdsa_sk`, `~/.ssh/id_ed25519` 以及 `~/.ssh/id_ed25519_sk`

- `-B <bind_interface>`

	使用指定的 NIC 和 `destination` 做网络连接

- `-b <bind_address>`

	使用指定的 IP 和 `destination` 做网络连接

### Config Related

- `-C`

	对连接传输的数据做 compression，通常只有在 poor network 中需要使用该参数，如果在 fast network 中使用该参数只会让数据显示更慢

- `-F <configfile>`

	指定 ssh 使用的配置文件，如果没有指定，会先从 `/etc/ssh/ssh_config` 读取，然后再读取 `~/.ssh/config`

- `-o <option>`

	指定当前连接使用的 option(优先使用命令行指定的，而非配置文件)，常用的有
	
	- ControlMaster
	- ControlPath
	
	所有可用 options 具体看 `ssh_config(5)`

### Port Forward Related

- `-D`
- `-g`
- `-J`
- `-L`
- `-N`
	不执行 `command` 也提供 login shell，通常和 [Prot Forward Related]() 的参数一起使用
- `-R`
- `-W`
- `-A`

### Debug Related

- `-G`

	会输出连接 `destination` 使用的配置详细

- `-E <log_file>`

	将 debug logs 输出到 log_file 而不是 stderr

- `-T`

	不提供 login shell

- `-V`

	打印版本信息并退出

- `-v`

	输出 debug 信息，多个 `-v` 可以增加详细度，最多 3

### Misc Related

- `-X`

	允许 X11 Forwarding(client 作为 X11 Server)

- `-Y`

	允许 trusted X11 Forwarding(client 作为 X11 Server)

- `-f`

	`command` 会以异步的方式在 remote machine 中运行，而不会阻塞 client current shell

## 0x03 Authentication



---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [ssh(1) - OpenBSD manual pages](https://man.openbsd.org/ssh)

***References***

[^1]:[ssh(1) - OpenBSD manual pages](https://man.openbsd.org/ssh)
