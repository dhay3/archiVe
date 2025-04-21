---
createTime: 2025-04-09 22:30
license: cc by 4.0
tags: 
 - "#hash1" 
 - "#hash2"
---

# Wireguard 01 - Overview

## 0x01 Preface

> Wireguard is an extremely simple and fast VPN utilizes

Wireguard 是一个简单的 VPN 工具，相比 OpenVPN 更轻量、更块，类似的工具还有 Zerotier、Tailscale 等等

## 0x02 Installation

> 具体查看 [Wireguard Installation](https://www.wireguard.com/install/)

Arch Based Kernel 直接执行 `pacman -S wireguard-tools` 即可

## 0x03 Tutorial

官方很好的提供了 [tutorial](https://www.wireguard.com/quickstart/)

假设

Host A Linux

```

```

Host B Windows

```

```

现在需要配置 100.69.0.0/24 段的 VPN

添加 VNIC

```
cc in ~ λ ip link add dev wg0 type wireguard
```

设置 VNIC IP address

```
cc in ~ λ ip address add dev wg0 100.69.0.1/24
```

如果只想要建立 P2P network 的话可以使用

```
cc in ~ λ ip address add dev wg0 100.69.0.1 peer 100.69.0.2
```

设置加密配置

```
# 设置 700 权限
cc in ~ λ umask 077
# 生成私钥
cc in ~ λ wg genkey > wg.privkey
# 生成公钥
cc in ~ λ wg pubkey < wg.privkey
TVI5H9w0lUOPaZFu/NyuPuGyWpxR0jqemMDhM3JWMwU=
```

应用密钥

```
cc in ~ λ sudo wg set wg0 private-key wg.privkey
```

设置 L2 链路 UP

```
cc in ~ λ sudo ip link set wg0 up
```

查看 wireguard 配置

```
cc in ~ λ sudo wg
interface: wg0
  public key: TVI5H9w0lUOPaZFu/NyuPuGyWpxR0jqemMDhM3JWMwU=
  private key: (hidden)
  listening port: 47494
```



## 0x04 How Wireguard Works

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [WireGuard: fast, modern, secure VPN tunnel](https://www.wireguard.com/)
- [Quick Start - WireGuard](https://www.wireguard.com/quickstart/)

***References***


