---
createTime: 2024-07-09 15:50
tags:
  - "#hash1"
  - "#hash2"
---

# Let's Encrypt 01 - Overview

## 0x01 Overview

[Let's Encrypt](https://letsencrypt.org/) 是一个通过 ACME(Automatic Certificate Management Environment) Protocal 实现免费颁发 TLS Certificate 的 CA

Let's Encrypt 颁发的 TLS Certificate 有如下特点[^1]

1. 必须要通过 ACME client 获取 cert
2. TLS Cert 默认只有 90 天有效期，但是可以通过 ACME client 自动 renew cert
3. 支持生成 wildcard cert


---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[FAQ -  Let's Encrypt](https://letsencrypt.org/docs/faq/)
