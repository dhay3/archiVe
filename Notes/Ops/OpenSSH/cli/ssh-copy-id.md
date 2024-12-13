---
createTime: 2024-12-03 15:07
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# ssh-copy-id

## 0x01 Preface

`ssh-copy-id` 是一个用于将 client public key 拷贝至 server authorized_keys(默认为 `~/.ssh/authorized_keys` 如果不存在会自动创建) 的工具

## 0x02 Syntax

```
ssh-copy-id   [-f]   [-n]  [-s]  [-x]  [-i  [identity_file]]  [-t  target_path]  [-F  ssh_config]
[[-o ssh_option] ...] [-p port] [user@]hostname
ssh-copy-id -h | -?
```

## 0x03 Opotional Args

- `-n`

	dry-run

- `-t <path>`

	server authorized_keys 的路径，默认为 `~/.ssh/authorized_keys`

- `-p port`

	server 监听的端口

- 

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***



***References***


