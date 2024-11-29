---
createTime: 2024-11-19 11:29
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# sshd

## 0x01 Preface

OpenSSH 以 C/S 模式运行，而 `sshd`(OpenSSH Daemon) 是 OpenSSH server 用户态的工具。默认会监听从 22 端口过来的数据包，如果是可信的 OpenSSH Client 过来的请求，就会使用其安全可靠的信道传输数据

在 sysV 中 `sshd` 通常会从按照 `/etc/rc` 下的配置自启动，而现今大多数的 OS 都会使用 systemd 来自启动 `sshd`

## 0x02 Syntax

```
sshd 	[-46DdeGiqTtV] [-C connection_spec] [-c host_certificate_file] [-E log_file] [-f config_file] [-g login_grace_time] [-h host_key_file] [-o option] [-p port] [-u len]
```

一般很少会通过手动执行 `sshd` 的方式来启动 OpenSSH server，默认以 daemon 方式运行

## 0x03 Optional Args

> [!note]
> 具体请看 manual page[^1]

- `-D`

	以 foreground 形式运行 sshd

- `-d`

	以 foregrond 形式运行 sshd，且输出 debug 信息，多个 `-d` 可以增加详细度，最大 3

- `-E <log_file>`

	将 debug 信息输出到 log_file

- `-f <config_file>`

	指定 sshd 使用的配置文件，如果没有指定默认使用 `/etc/ssh/sshd_config`

- `-G`

	输出 `sshd` 使用的详细配置

- `-t`

	校验 `sshd` 配置文件是否正确

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [sshd(8) - OpenBSD manual pages](https://man.openbsd.org/sshd)


***References***

[^1]:[sshd(8) - OpenBSD manual pages](https://man.openbsd.org/sshd)

