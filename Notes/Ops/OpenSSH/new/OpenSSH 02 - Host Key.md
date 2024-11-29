---
createTime: 2024-11-26 18:02
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# OpenSSH 02 - Host-Key

## 0x01 Preface

当 OpenSSH client 连接 OpenSSH server 时，server 会回送自己的 public key 给 client(和 TLS handshake 的逻辑类似)；然后由 client 决定是否把 server public key 按照一定的格式记录到 host key database 中。server public key 也被称为 host-key，是 server 的 unique id

client 默认只会在第一次连接时，将 server host-key 记录到 host-key database 中。后续如果 client 和 server 要建立 SSH session 时，就会校验 host-key database 中对应 server host-key 是否匹配，来决定是否允许建立 SSH session

这在一定程度上避免了 MITMA，但是如果 server 因为 IP 变更或者是 key pair 变更，就会出现 `WARNING: REMOTE HOST IDENTIFICATION HAS CHANGED! `

这时可以通过修改(删除) host-key database 中对应条目的内容，或者是 `ssh_config` 中的 `StrictHostKeyChecking` 来解决。前提需要确定 server 就是你想要连接的

## 0x02 Host-Key Database

host-key database 实际上就是一个文件(OpenSSH 要求 host-key database 只有 root/owner 有 w 权限)，默认定义如下路径

- `~/.ssh/known_hosts`

	当前用户记录的 host key

- `/etc/ssh/known_hosts`

	系统全局记录的 host key

也可以通过设置 `ssh_config` 中的 `GlobalKnownHostsFile`, `UserKnownHostsFile` 来自定义

## 0x03 Host-Key Database Format

host-key database 由很多条目组成，每条条目由几部分组成，通过空格分隔

- marker(optional)

	只能是 2 个值

	- `@cert-authority`

		表示 OpenSSH server 的 keypair 由 CA 签发

	- `@revoked`

		标示 OpenSSh server 的 keypair 已失效

- hostnames

	server hsotname

- keytype

	生成 server keypair 的加密算法

- base64-encoded

	server public key 实际上是 raw binary，为了方便识别，对 server public key 做 base64 encode；也被称为 server public key fingerprint(a unique id represent that keypair)

- comment(optional)

	备注

例如

```
github.com ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIOMqqnkVzrm0SdG6UOoqKLsabgH5C9okWi0dh2l9GKJl
github.com ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBEmKSENjQEezOmxkZMy7opKgwFB9nkt5YRrYMjNuG5N87uRgg6CLrbo5wAdT/y6v0mKV0U2w0WZ2YB/++Tpockg=
```

## 0x04 Server Keypair

keypair 就是 server 的密钥对，通常由 OpenSSH 自动生成，路径如下

- /etc/ssh/ssh_host_ed25519_key
- /etc/ssh/ssh_host_ed25519_key.pub
- /etc/ssh/ssh_host_ecdsa_key
- /etc/ssh/ssh_host_ecdsa_key.pub
- /etc/ssh/ssh_host_rsa_key
- /etc/ssh/ssh_host_rsa_key.pub

也可以通过 `sshd_config` 的 `HostKey` 指定

且 OpenSSH 允许不同的 server 可以使用相同的 server keypair，如果 server keypair 相同，就会出现如下提示

```
$ ssh root@10.0.3.40
The authenticity of host '10.0.3.40 (10.0.3.40)' can't be established.
ED25519 key fingerprint is SHA256:WfL7sVGttLvi8v330i5aBcnlqoSNM5oGPaHv2TLlziE.
This host key is known by the following other names/addresses:
    ~/.ssh/known_hosts:6: 10.0.3.41
    ~/.ssh/known_hosts:7: 10.0.3.42
    ~/.ssh/known_hosts:8: 10.0.3.43
    (52 additional names omitted)
Are you sure you want to continue connecting (yes/no/[fingerprint])? yes
Warning: Permanently added '10.0.3.40' (ED25519) to the list of known hosts.
```

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [ssh(1) - OpenBSD manual pages](https://man.openbsd.org/ssh)
- [sshd(8) - OpenBSD manual pages](https://man.openbsd.org/sshd)

***References***

[^1]:[ssh(1) - OpenBSD manual pages](https://man.openbsd.org/ssh#VERIFYING_HOST_KEYS)
