---
createTime: 2024-11-19 14:02
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# OpenSSH 0 - sshd_config

## 0x01 Preface

OpenSSH Server `sshd` 的配置文件，如果 `sshd` 在没有使用 `-f` 的情况下，会默认使用 `/etc/ssh/sshd_config` 作为 `sshd` 的配置文件

`#` 开头表示 comments，以 `keyword arguments` 的形式配置，每个 `keyword arguments` 单独一行

例如

```
#this is a comment
Port 64422
```

> [!note]
> keyword 和 arguments 大小写敏感

### 0x01a Time Formats

如果 keyword 的 arguments 表示 time 时，`sshd_config` 支持使用 `time[qualifier]` 的格式

| qualifier | unit    | example                         |
| --------- | ------- | ------------------------------- |
| (none)    | seconds | 60  equals to 60 seconds        |
| s or S    | seconds | 60s or 60S equals to 60 seconds |
| m or M    | minutes | 60m or 60M equals to 60 minutes |
| h or H    | hours   | 60h or 60h equals to 60 hours   |
| d or D    | days    | 60d or 60D equals to 60 days    |
| w or W    | weeks   | 60w or 60W equals to 60 weeks   | 

还可以互相组合，例如

1h30m

### 0x01b Tokens

> [!note]
> 版本不同，支持的 Tokens 也不相同

`sshd_config` 中的一些 keywords 还支持使用 tokens，会在运行时扩展，类似于 placeholder

keywords 如下

- `AuthorizedKeysCommand` accepts the tokens `%%`, `%C`, `%D`, `%f`, `%h`, `%k`, `%t`, `%U`, and `%u`.

- `AuthorizedKeysFile` accepts the tokens `%%`, `%h`, `%U`, and `%u`.

- `AuthorizedPrincipalsCommand` accepts the tokens `%%`, `%C`, `%D`, `%F`, `%f`, `%h`, `%i`, `%K`, `%k`, `%s`, `%T`, `%t`, `%U`, and `%u`.

- `AuthorizedPrincipalsFile` accepts the tokens `%%`, `%h`, `%U`, and `%u`.

- `ChrootDirectory` accepts the tokens `%%`, `%h`, `%U`, and `%u`.

- `RoutingDomain` accepts the token `%D`.

支持的 token 含义如下

```
%%
    A literal ‘%’.
%C
    Identifies the connection endpoints, containing four space-separated values: client address, client port number, server address, and server port number.
%D
    The routing domain in which the incoming connection was received.
%F
    The fingerprint of the CA key.
%f
    The fingerprint of the key or certificate.
%h
    The home directory of the user.
%i
    The key ID in the certificate.
%K
    The base64-encoded CA key.
%k
    The base64-encoded key or certificate for authentication.
%s
    The serial number of the certificate.
%T
    The type of the CA key.
%t
    The key or certificate type.
%U
    The numeric user ID of the target user.
%u
    The username.
```


## 0x02 Keywords Arguments

- `AcceptEnv`

	从 client 发送过来的 environment variables 可以在 

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [sshd\_config(5) - OpenBSD manual pages](https://man.openbsd.org/sshd_config)

***References***


