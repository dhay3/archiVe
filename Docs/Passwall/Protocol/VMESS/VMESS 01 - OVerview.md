---
createTime: 2024-08-02 11:10
tags:
  - "#hash1"
  - "#hash2"
---

# VMESS 01 - OVerview

## 0x01 Overview

VMess 是一个基于 TCP 的加密协议，由 v2ray 原创。请求和响应无需握手直接可以传输数据(但是需要完成 TCP 握手)，请求报文和响应报文采用非对称格式

Request

|16 字节|X 字节|余下部分|
|---|---|---|
|认证信息|指令部分|数据部分|


Response
|1 字节|1 字节|1 字节|1 字节|M 字节|余下部分|
|---|---|---|---|---|---|
|响应认证 V|选项 Opt|指令 Cmd|指令长度 M|指令内容|实际应答数据|

[VMess 协议 | V2Fly.org](https://www.v2fly.org/developer/protocols/vmess.html#%E5%AE%A2%E6%88%B7%E7%AB%AF%E8%AF%B7%E6%B1%82)

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

