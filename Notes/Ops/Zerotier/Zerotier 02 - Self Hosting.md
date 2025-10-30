---
createTime: 2025-10-09 22:32
license: cc by 4.0
tags: 
 - "#hash1" 
 - "#hash2"
---

# Zerotier 02 - Self Hosting

## 0x01 Preface

zerotier 支持 Network Controller 以及 Root 的自建

## 0x02 Root Self Hosting[^1]

Root 承载 zerotier 实际的 p2p 流量，自建 Root（自建的 Root 也被称为 Moon）有几个好处

- 不会因为官方的 Root 节点 overload 而导致服务不可用
- 官方 Root 节点在海外，国内 client 连接 Root 延迟高，自建可以大幅度降低延迟

> [!note]
> Root 需要一个静态地址

Root self hosting 分为 2 部分

- Server Side
- Client Side

### 0x02a Server Side

在需要自建 Root 的服务器上操作

#### Creating a Moon Configuration

通过 zerotier 的公钥（公钥通常在配置目录中）生成 moon 配置文件，例如

```
sudo zerotier-idtool initmoon /var/lib/zerotier-one/identity.publi >> moon.json
```

修改生成的 `moon.json` 中的 `stableEndpoints` 为 `服务器静态地址/9993`，例如

> [!note]
> 云服务器需要放通 9993 UDP TCP

```
    {
      "id": "deadbeef00",
      "objtype": "world",
      "roots": [
        {
          "identity": "deadbeef00:0:34031483094...",
          "stableEndpoints": [ "8.154.81.120/9993" ]
        }
      ],
      "signingKey": "b324d84cec708d1b51d5ac03e75afba501a12e2124705ec34a614bf8f9b2c800f44d9824ad3ab2e3da1ac52ecb39ac052ce3f54e58d8944b52632eb6d671d0e0",
      "signingKey_SECRET": "ffc5dd0b2baf1c9b220d1c9cb39633f9e2151cf350a6d0e67c913f8952bafaf3671d2226388e1406e7670dc645851bf7d3643da701fd4599fedb9914c3918db3",
      "updatesMustBeSignedBy": "b324d84cec708d1b51d5ac03e75afba501a12e2124705ec34a614bf8f9b2c800f44d9824ad3ab2e3da1ac52ecb39ac052ce3f54e58d8944b52632eb6d671d0e0",
      "worldType": "moon"
    }
```

#### Configuring Your Moon Server

对 moon.json 签名生成全球唯一的 moon ID 文件

```
sudo zerotier-idtool genmoon moon.json
wrote 0000003f7285bd35.moon (signed world with timestamp 1760023168769)
```

将生成的 moon ID 文件拷贝到 `/var/lib/zerotier-one/moons.d` 下，如果没有对应的目录就手动创建

```
sudo cp 0000003f7285bd35.moon /var/lib/zerotier-one/moons.d
```

重启 zerotiter

```
sudo systemctl restart zerotier-one
```

### 0x02b Client Side

在 client 上操作

#### Configure Your Clients

##### Desktop/Server

有两种方式

**Option1（推荐）**: 将 server side 生成的 moon ID 文件拷贝到 client zerotier 配置目录，没有就创建

- Windows: `C:\ProgramData\ZeroTier\One\moons.d`
- Linux: `/var/lib/zerotier-one/moons.d`

重启 zerotier

**Option2**: 使用 `zerotier-cil orbit <moons ID> <moons ID>`

例如生成的 moon ID 文件名为 `0000003f7285bd35.moon` 那么就执行 `sudo zerotier-cli orbit 0000003f7285bd35 0000003f7285bd35`

重启 zerotier

---

完成后在 client 侧可以使用 `zerotier-cli peers` 查看 zerotier 是否成功将自建的 Root 作为中继，自建的 Root 会以 role MOON 标示。role LEAF 为当前网络下的其他 peers

```
sudo zerotier-cli peers
200 peers
<ztaddr>   <ver>  <role> <lat> <link>   <lastTX> <lastRX> <path>
5ede4bf411 1.14.2 LEAF      11 DIRECT   7411     7401     183.247.147.172/64342
3f7285bd35 1.16.0 MOON      11 DIRECT   2406     2393     8.154.81.120/9993
778cde7190 -      PLANET   644 DIRECT   2406     106863   103.195.103.66/9993
cafe04eba9 -      PLANET   502 DIRECT   2406     106943   84.17.53.155/9993
cafe80ed74 -      PLANET   587 DIRECT   2406     106920   185.152.67.145/9993
cafefd6717 -      PLANET   411 DIRECT   2406     107034   79.127.159.187/9993
d5e5fb6537 1.15.3 LEAF      -1 DIRECT   2406     13984    35.209.81.44/51982
```

server 侧连接当前 Root 的 client 会以 role LEAF 标示（可以核对 client ztaddr 来校验）

> [!note] 
> ztaddr 可以使用 zerotier-cli info 来查看 

```
zerotier-cli peers
200 peers
<ztaddr>   <ver>  <role> <lat> <link>   <lastTX> <lastRX> <path>
3f7285bd35 1.16.0 LEAF      15 DIRECT   468      468      183.247.147.173/17839
5ede4bf411 1.14.2 LEAF       8 DIRECT   410      410      183.247.147.172/36445
778cde7190 -      PLANET   234 DIRECT   5824     95723    103.195.103.66/9993
cafe04eba9 -      PLANET   167 DIRECT   5824     80779    84.17.53.155/9993
cafe80ed74 -      PLANET   161 DIRECT   819      80783    185.152.67.145/9993
cafefd6717 -      PLANET   214 DIRECT   5824     80732    79.127.159.187/9993
d5e5fb6537 1.15.3 LEAF     208 DIRECT   5824     11446    35.209.81.44/51982
```

##### Mobile

> [!note]
> IOS 未越狱的情况下不支持使用自建的 moon

Andriod 具体看文档

## 0x03 Network Controller Self Hosting[^1]

官方不提供 Web UI 部署的 Contrller 但是有开源的面板

- [GitHub - key-networks/ztncui: ZeroTier network controller UI](https://github.com/key-networks/ztncui)
- [GitHub - key-networks/ztncui-aio: Licensed Under AGPL v3](https://github.com/key-networks/ztncui-aio)
- [GitHub - dec0dOS/zero-ui: ZeroUI - ZeroTier Controller Web UI - is a web user interface for a self-hosted ZeroTier network controller.](https://github.com/dec0dOS/zero-ui) 推荐
 
---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [Private Root Servers \| ZeroTier Documentation](https://docs.zerotier.com/roots/)
- [Network Controller \| ZeroTier Documentation](https://docs.zerotier.com/controller/)


***References***

[^1]:[Private Root Servers \| ZeroTier Documentation](https://docs.zerotier.com/roots/)


