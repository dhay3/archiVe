---
createTime: 2024-08-05 19:06
tags:
  - "#hash1"
  - "#hash2"
---

# V2ray 02 - How V2ray Works

## 0x01 Overview

V2ray 以 CS 模式工作[^1]

![](https://mermaid.ink/img/pako:eNpt0D0LwjAQBuC_Em6qUJHaThUEJW7i0G6S5WxOGzRJSZNBxP9uVMSvbi_PvRzcXaCxkqCEg8OuZetqJswiS5TZ2WDkiI3Hc8YTrvoOfdOSYxNW2eAfgW_q0b0-_anfLR-w4t_4Iy2zJO58jt44HcJ8CIsPhBQ0OY1KxqMuwjAmwLekSUAZo0R3FCDMNfYweFufTQOld4FSCJ1ET1xh_IV-YYdma63-7qyk8tZBucdTT9cb9YZmmA?type=png)


1. 对 Client(被代理服务器)而言，inbound 处理应用或者是系统发送过来的数据包，由 Dispatcher/Router/DNS 决定发送到那个 outbound，然后 outbound 将数据包发送到 Server
2. 对 Server(代理服务器)而言，inbound 处理 Client 发送过来的数据包，由 Dispatcher/Router/DNS 决定发送到那个 outbound，然后 outbound 将数据包发送到实际的目标站点

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[工作机制 | V2Fly.org](https://www.v2fly.org/guide/workflow.html)