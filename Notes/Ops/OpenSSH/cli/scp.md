---
createTime: 2024-12-03 15:36
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# scp

## 0x01 Preface

`scp` 是一个基于 sftp 协议的拷贝工具，可以在 OpenSSH client 和 server 之间完成单次拷贝文件

## 0x02 Syntax

```
scp 	[-346ABCOpqRrsTv] [-c cipher] [-D sftp_server_path] [-F ssh_config] [-i identity_file] [-J destination] [-l limit] [-o ssh_option] [-P port] [-S program] [-X sftp_option] source ... target
```

## 0x03 Positional Args

- `source ...`

	变参，需要拷贝的文件路径

- `target`

	拷贝的目的

	需要使用如下格式

	- `[user@]host:[path]`

		如果没有指定 `path`，默认为 user `${HOME}`

	- `scp://[user@]host[:port][/path]`

		如果没有指定 `/path`，默认为 user `${HOME}`

		*portable OpenSSH 不支持这种格式*

## 0x04 Optional Args

- `-F <ssh_config>`

	指定 client 使用的 ssh_config 配置路径

- `-i <identity_file>`

	指定 PKI 使用的 private key

- `-P <port>`

	指定 target server 监听的端口

- `-p`

	拷贝保留源文件的 mtime,atime 以及 file mode bits

- `-r`

	递归拷贝文件和目录

## 0x05 Example

```
scp -P 6022 -r ./*  root@10.0.3.49:/tmp
root@10.0.3.49's password:
user_key                                                                 100% 1679     4.5MB/s   00:00
user_key-cert.pub                                                        100% 2193     5.1MB/s   00:00
user_key.pub                                                             100%  400     1.3MB/s   00:00

```

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- `man scp`

***References***


