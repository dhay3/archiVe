---
createTime: 2024-07-09 16:20
tags:
  - "#hash1"
  - "#hash2"
---

# Let's Encrypt 02 - ACME Clients

## 0x01 Overview

Let‘s Encrypt 通过 ACME client 获取 cert，不同的语言的 client 不同，具体可以参考 [ACME Client Implementations -  Let's Encrypt](https://letsencrypt.org/docs/client-options/)

这里推荐 2 个 ACME client

1. [certbot](https://certbot.eff.org/)
2. [acme.sh](https://github.com/acmesh-official/acme.sh)

## 0x02 [Certbot](https://certbot.eff.org/instructions)

Certbot official site 默认只提供通过 snap 安装的方式
但是大多数 repo 都有对应的包，例如 centos 需要先安装 epel-release 然后再安装 certbot

```shell
sudo certbot certonly
```

## 0x03 [acem.sh](https://github.com/acmesh-official/acme.sh)

acem.sh 安装的方式很简单，只需要一行命令
```shell
curl https://get.acme.sh | sh -s email=my@example.com
```


---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]: