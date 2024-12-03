---
createTime: 2024-12-02 11:42
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# ssh-keygen

## 0x01 Preface

`ssh-keygen` 是 OpenSSH suite 中用于管理 keypair 的一个工具

- private key 默认会以 `id_<type>` 存储在 `~/.ssh`，例如 `~/.ssh/id_rsa`

> [!NOTE]
> OpenSSH 要求 private key 的权限为 `[6|7]00`

- public key 默认会以 `id_<type>.pub` 存储在 `~/.ssh`，例如 `~/.ssh/id_rsa.pub`

type 通过 `-t` 指定，如果没有指定默认会使用 ed25519

例如

```
$ ssh-keygen
Generating public/private ed25519 key pair.
Enter file in which to save the key (/root/.ssh/id_ed25519):
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Your identification has been saved in /root/.ssh/id_ed25519
Your public key has been saved in /root/.ssh/id_ed25519.pub
The key fingerprint is:
SHA256:vEFicjLOsWx/OjKhUPdOoCNNQyfDI+ECBftw1dfZK/w root@localhost.localdomain
The key's randomart image is:
+--[ED25519 256]--+
|+=. ..   . o     |
|+.*.. . . o .    |
|++.== +...   .   |
|.+++oO +  o .    |
| +.+*o  S  o     |
|o +...o  o  E    |
| o o +. o        |
|  . o oo         |
|     o..         |
+----[SHA256]-----+
```

如果使用的 passphrase 不为空，那么在使用类似 public key authentication 时，需要提供 passpharse 才会认证通过

```
$ ssh root@10.0.3.40
Enter passphrase for key '/root/.ssh/id_ed25519':
```

如果 passphrase 遗忘了，需要重新生成 keypair，无法重置。生成的 private key 会以 base64 armor 编码

例如

```
-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACAHENNEae//Cq/Ji3Cj5BzaDnKMw+KIhgY/6wqbNTHeXQAAAKDHwdivx8HY
rwAAAAtzc2gtZWQyNTUxOQAAACAHENNEae//Cq/Ji3Cj5BzaDnKMw+KIhgY/6wqbNTHeXQ
AAAEB6PjM1RRZYHQiXMb0L93mJBbkP9wC+HmbwuXUCYNgsogcQ00Rp7/8Kr8mLcKPkHNoO
cozD4oiGBj/rCps1Md5dAAAAGnJvb3RAbG9jYWxob3N0LmxvY2FsZG9tYWluAQID
-----END OPENSSH PRIVATE KEY-----
```

而 public key 由几部分组成

- keytype
- base64-ecoded key
- comment

例如

```
ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIAcQ00Rp7/8Kr8mLcKPkHNoOcozD4oiGBj/rCps1Md5d root@localhost.localdomain
```

## 0x02 Syntax

```
ssh-keygen [-q] [-a rounds] [-b bits] [-C comment] [-f output_keyfile] [-m format] [-N new_passphrase] [-O option] [-t ecdsa | ecdsa-sk | ed25519 | ed25519-sk | rsa] [-w provider] [-Z cipher]
ssh-keygen -p [-a rounds] [-f keyfile] [-m format] [-N new_passphrase] [-P old_passphrase] [-Z cipher]
ssh-keygen -i [-f input_keyfile] [-m key_format]
ssh-keygen -e [-f input_keyfile] [-m key_format]
ssh-keygen -y [-f input_keyfile]
ssh-keygen -c [-a rounds] [-C comment] [-f keyfile] [-P passphrase]
ssh-keygen -l [-v] [-E fingerprint_hash] [-f input_keyfile]
ssh-keygen -B [-f input_keyfile]
ssh-keygen -D pkcs11
ssh-keygen -F hostname [-lv] [-f known_hosts_file]
ssh-keygen -H [-f known_hosts_file]
ssh-keygen -K [-a rounds] [-w provider]
ssh-keygen -R hostname [-f known_hosts_file]
ssh-keygen -r hostname [-g] [-f input_keyfile]
ssh-keygen -M generate [-O option] output_file
ssh-keygen -M screen [-f input_file] [-O option] output_file
ssh-keygen -I certificate_identity -s ca_key [-hU] [-D pkcs11_provider] [-n principals] [-O option] [-V validity_interval] [-z serial_number] file ...
ssh-keygen -L [-f input_keyfile]
ssh-keygen -A [-a rounds] [-f prefix_path]
ssh-keygen -k -f krl_file [-u] [-s ca_public] [-z version_number] file ...
ssh-keygen -Q [-l] -f krl_file file ...
ssh-keygen -Y find-principals [-O option] -s signature_file -f allowed_signers_file
ssh-keygen -Y match-principals -I signer_identity -f allowed_signers_file
ssh-keygen -Y check-novalidate [-O option] -n namespace -s signature_file
ssh-keygen -Y sign [-O option] -f key_file -n namespace file ...
ssh-keygen -Y verify [-O option] -f allowed_signers_file -I signer_identity -n namespace -s signature_file [-r revocation_file]
```

## 0x02 Optional Args

> [!note]
> 只记录常用的参数，具体看 man page

- `-A`

	以默认的值生成所有支持的 keypair（rsa,ecdsa,ed25519）

- `-t <ecdsa|ecdsa-sk|ed25519|ed25519-sk|rsa>`

	指定生成 keypair's type，默认 ed25519

- `-b <bits>`

	指定生成 keypari's bits length 

	- rsa 1024 - 3072
	- ecdsa 256 | 384 | 521
	- ecdsa-sk/ed25519/ed25519-sk ignored

- `-C <comment>`

	指定生成 keypair's comment，默认 `user@hostname`

- `-c`

	修改指定 keypair's comment

- `-N <passpharse>`

	指定生成 private key(new) 使用的 passpharse

- `-P <passpharse>`

	指定 private key(old) 使用的 passpharse 

- `-p`

	修改 private key passpharse

- `-f <path>`

	使用指定的 private key 或者是指定生成 private key 的路径

- `-y`

	输出 private key 对应的 public key

- `-l`

	输出 keypair fingerprint

- `-F {hostname|[hostname]:port}`

	从 host key database files 查找指定的 hostname 或者是 \[hostname\]:port 的 host key

- `-R <hostname|[hostname:]port>`

	从 host key database files 中删除指定的 hostname 或者是 \[hostname\]:port 的 host key 

- `-H`

	将 host key database files 中的内容完全 Hash

- `-s <ca_private_key>`

	使用指定的 CA private key 签名

- `-I <identifier>`

	CA 签发生成 certificate 的标识符

- `-h`

	CA 签发生成 host certificate 而不是默认的 user certificate

- `n`

	指定 CA 签发的 certificate 只对指定 username 或者是 hostname 生效(支持 [ssh_config#0x01a Patterns](../config/ssh_config.md))，类似于 X509 中的 DN domain(只对指定的 domain 生效)

	host certificate 如果没有指定，类似于泛域名，对 server 所有的 username 和 hostname 都生效

	user certificate 必须要指明，否则无法使用 certificate 完成认证

- `-V <expire_time>`

	指定 CA 签发的 certifcate 的有效时间

	支持 2 种格式

	- relative `+time`

		从当前时间开始计算相对值

		time 按照 [sshd_config#0x01b Time Formats](../sshd_config) 指定

		例如

		- `-v +52w1d`
		- `-v -1d`

	- absolute `start:end`

		指定开始时间和过期时间

		- start 可以使用 always 表示没有特定的开始时间
		- end 可以使用 forever 表示没有特定的结束时间
		- start/end 支持 YYMMDD 或者  YYYYMMDDHHMM\[SS\]
		- start/end 支持 epoch time
		- start/end 同样支持 relative 格式

		例如

		- `-4w:+4w`
		- `20201101:20211101`
		- `always:forever`

- `-O <option>`

	指定 user cert 的权限

	常用的 option 有

	- clear
	- no-port-forwarding
	- no-pty
	- no-x11-forwarding
	- permit-port-forwarding
	- permit-pty
	- permit-x11-forwarding
	- source-address=address_list

- `-L`

	查看 CA 签发证书的详细信息，和 `-f <cert>` 一起使用

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- `man ssh-keygen`
- [How to configure SSH Certificate-Based Authentication](https://goteleport.com/blog/how-to-configure-ssh-certificate-based-authentication/)

***References***


