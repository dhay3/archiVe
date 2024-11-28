---
createTime: 2024-11-19 10:34
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# OpenSSH 07 - Authentication

## 0x01 Preface

OpenSSH 一个完整的 Authentication 分为 2 部分

- 



## Client

OpenSSH client 默认会按照如下的顺序去尝试鉴权

- GSSAPI-based authentication
- host-based authentication
- public key authentication
- keyboard-interactive authentication
- password authentication

也可以通过修改 `ssh_config` 中的 `PreferredAuthentications` 来修改鉴权方式或者是顺序

## Server

OpenSSH server 通过 

- `AuthenticationMethod`

## Authentication Methods

### Host-Based Authentication

### Public Key Authentication





---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***



***References***


