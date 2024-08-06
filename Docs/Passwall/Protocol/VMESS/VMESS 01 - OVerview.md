---
createTime: 2024-08-02 11:10
tags:
  - "#hash1"
  - "#hash2"
---

# VMESS 01 - OVerview

## 0x01 Overview

VMess 是一个基于 TCP 的加密协议，由 v2ray 原创。

## 0x02 Request/Response

VMess 请求和响应无需握手直接可以传输数据(但是需要完成 TCP 握手)。当 VMess client 发起一次请求，VMess server 会判断请求是否来自一个合法的 client。如验证通过，则转发该请求，并把获得的响应发回给客户端。 VMess 报文使用非对称格式，即客户端发出的请求和服务器端的响应使用了不同的格式。

3. VMess 使用了 UUID 替代了 Shadowsocks 中的密码
4. VMess 基于时间，需要确保 Client/Server 差值在 正负 30 秒内

### 0x02a Client Request

Request 由 3 部分组成

| 16 字节  | X 字节   | 余下部分 |
| -------- | -------- | -------- |
| 认证信息 | 指令部分 | 数据部分 | 


####  Certification Information

Certification Information 在 VMess 也被称为 request header，是一个通过 HMAC 按照 UUID, UTC Timestamp ± 30 使用 MD5 计算出来的 16byte Hash

这串 Hash 会在 VMess Server 用于确认用户是否合法

入口逻辑 [session.EncodeRequestHeader(request, writer)](https://github.com/v2ray/v2ray-core/blob/master/proxy/vmess/outbound/outbound.go#L57)

```go
func (h *Handler) Process(ctx context.Context, link *transport.Link, dialer internet.Dialer) error {
	...
	// DefaultIDHash hmac.New(md5.New, key)
	session := encoding.NewClientSession(isAEAD, protocol.DefaultIDHash, ctx)
	...
	requestDone := func() error {
		defer timer.SetTimeout(sessionPolicy.Timeouts.DownlinkOnly)
		writer := buf.NewBufferedWriter(buf.NewWriter(conn))
		if err := session.EncodeRequestHeader(request, writer); err != nil {
			return newError("failed to encode request").Base(err).AtWarning()
		}
	...
	}
```

生成认证信息 [EncodeRequestHeader(header \*protocol.RequestHeader, writer io.Writer)](https://github.com/v2ray/v2ray-core/blob/master/proxy/vmess/encoding/client.go#L76)

```go
func (c *ClientSession) EncodeRequestHeader(header *protocol.RequestHeader, writer io.Writer) error {
	timestamp := protocol.NewTimestampGenerator(protocol.NowTime(), 30)()
	account := header.User.Account.(*vmess.MemoryAccount)
	if !c.isAEAD {
		//在 alterId 为 0 时，对应 vmess 配置文件中的 uuid MD5 值
		idHash := c.idHash(account.AnyValidID().Bytes())
		// 往 uuid MD5 值中拼接 Timestamp ± 30
		common.Must2(serial.WriteUint64(idHash, uint64(timestamp)))
		//4bf0baa4-fc14-4ebb-81e7-59cc903d89cc
		common.Must2(writer.Write(idHash.Sum(nil)))
	}
```

### 0x02b Response

| 1 字节     | 1 字节   | 1 字节   | 1 字节     | M 字节   | 余下部分     |
| ---------- | -------- | -------- | ---------- | -------- | ------------ |
| 响应认证 V | 选项 Opt | 指令 Cmd | 指令长度 M | 指令内容 | 实际应答数据 |

[VMess 协议 | V2Fly.org](https://www.v2fly.org/developer/protocols/vmess.html#%E5%AE%A2%E6%88%B7%E7%AB%AF%E8%AF%B7%E6%B1%82)

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

