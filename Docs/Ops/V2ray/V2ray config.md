# V2ray config

ref：

https://www.v2ray.com/chapter_02/

https://github.com/bannedbook/fanqiang/blob/master/v2ss/%E8%87%AA%E5%BB%BAV2ray%E6%9C%8D%E5%8A%A1%E5%99%A8%E7%AE%80%E6%98%8E%E6%95%99%E7%A8%8B.md

## 说明

> JSON可以使用`//`或`/**/`注释

V2ray 默认使用 Protobuf 配置，但是为了方便阅读也支持JSON。在运行之前V2ray会将JSON转为Protobuf。

==v2ray不分客户端和服务端的配置，但要区分inbound入站，和 outbound 出站在不同端的含义==

- 在v2ray客户端(被代理机)，inbound表示浏览器发送过来的数据包或是 v2ray服务器返回的数据包，outbound表示发送到的代理服务器。
- 在v2ray服务端(代理机)，inbound表示接受来自v2ray客户端的数据包，outbound表示实际访问的目标站点

配置文件都以`config.json`命名，通常存储在`/etc/v2ray/config.json`

## 配置文件

V2ray 主要包含如下几个 directives

```
{
  //日志配置
  "log": {},
  //RPC api
  "api": {},
  //DNS 配置，如果没有该 directive 默认使用本地 DNS
  "dns": {},
  //是否开启统计信息
  "stats": {},
  //路由策略
  "routing": {},
  //本地策略
  "policy": {},
  //方向代理配置
  "reverse": {},
  //入站配置
  "inbounds": [],
  //出站配置
  "outbounds": [],
  //配置和其他 v2ray 节点建立
  "transport": {}
}
```

### log

配置 snippet

```
{
  "access": "文件地址",
  "error": "文件地址",
  "loglevel": "warning"
}
```

- access

  访问日志目录，当该指令块值为空或者不存在时默认输出到 stdout

- error

  错误日志目录，当该指令块值为空或者不存在时默认输出到 stdout

- loglevel

  记录错误日志的等级

### api

目前没有使用到，所以未记录

### stats

目前没有任何配置项，只要有就表示开启

### dns

v2ray 内置了DNS 服务器，主要提供如下 2 个功能

1. 根据域名解析 IP 匹配路由规则
2. 像传统的 DNS 一样，解析地址进行 4 层连接

具体执行流程如下

![2022-09-18_15-36](https://git.poker/dhay3/image-repo/blob/master/20220918/2022-09-18_15-36.4s8l8sg27uo0.webp?raw=true)

配置 snippet

```
{
  "hosts": {
    "baidu.com": "127.0.0.1"
  },
  "servers": [
    {
      "address": "1.2.3.4",
      "port": 5353,
      "domains": [
        "domain:v2ray.com"
      ],
      "expectIPs": [
        "geoip:cn"
      ],
    },
    "8.8.8.8",
    "8.8.4.4",
    "localhost"
  ],
  "clientIp": "1.2.3.4",
  "tag": "dns_inbound"
}
```

- hosts：map{string:address}

  静态解析，其值为一组或者多组 `domain:address`。在做域名解析时，如果`domain`匹配了，则解析成`address`，并终止匹配下面内容。`address`的值也可以是域名

  域名使用如下格式

  1. 纯字符串: 当此域名完整匹配目标域名时，该规则生效

     例如`v2ray.com`匹配`v2ray.com`但不匹配 `www.v2ray.com`

  2. 正则表达式: 由`regexp:`开始，余下部分是一个正则表达式。当此正则表达式匹配目标域名时，该规则生效

     例如`regexp:.*\.goo.*\.com$`匹配`www.google.com`、`fonts.googleapis.com`，但不匹配 `google.com`

  3. 子域名 (推荐): 由 `domain:`开始，余下部分是一个域名。当此域名是目标域名或其子域名时，该规则生效

     例如`domain:v2ray.com`匹配`www.v2ray.com`、`v2ray.com`，但不匹配`xv2ray.com`

  4. 子串: 由`keyword:`开始，余下部分是一个字符串。当此字符串匹配目标域名中任意部分，该规则生效

     比如`keyword:sina.com`可以匹配`sina.com`、`sina.com.cn`和`www.sina.com`，但不匹配`sina.cn`

  5. 预定义域名列表：由`geosite:`开头，余下部分是一个名称

     如`geosite:google`或者`geosite:cn`。名称及域名列表参考[预定义域名列表](https://www.v2ray.com/chapter_02/03_routing.html#dlc)。

- servers: [string | [ServerObject](https://www.v2ray.com/chapter_02/04_dns.html#serverobject)]

- DNS 服务器，支持两种类型：DNS地址（字符串）和 [ServerObject](https://www.v2ray.com/chapter_02/04_dns.html#serverobject)

  1. 当它的值是一个 DNS IP 地址时，如果`8.8.8.8`，v2ray 会使用该地址的 53 端口进行 DNS 查询

  2. 当值为`localhost`时，表示用本机的 DNS

  3. 当值为`"https://dns.google/dns-query"`，V2Ray 会使用`DNS over HTTPS` (RFC8484, 简称DOH) 进行查询。有些服务商拥有IP别名的证书，可以直接写IP形式，比如`https://1.1.1.1/dns-query`。也可使用非标准端口和路径，如`"https://a.b.c.d:8443/my-dns-query"` (4.22.0+)

  5. ServerObject

     示例 snippet

     ```
     {
       "address": "1.2.3.4",
       "port": 5353,
       "domains": [
         "domain:v2ray.com"
       ],
       "expectIPs": [
         "geoip:cn"
       ]
     } 
     ```

     - address:address

       DNS 地址，可以使用 DOH 域名

     - port:number

       DNS 服务器端口，默认 53。但是用 DOH 时无效

     - domains:[string]

       一个域名列表，此列表包含的域名，将优先使用此服务器进行查询

       域名格式和[路由配置](https://www.v2ray.com/chapter_02/03_routing.html#ruleobject)中相同

     - expectIPs:[string]

       对 DNS 返回的 IP 进行校验，只返回包含 expectIPs。如果未配置此项，会原样返回 IP

- clientIp: string

  当前系统的 IP 地址，用于 DNS 查询时，通知服务器客户端的所在位置。不能是私有地址

- tag: string

  由此 DNS 发出的查询流量，除`localhost` 和 `DOHL_` 模式外，都会带有此标识，可在路由使用`inboundTag`进行匹配

### routing

v2ray 内建有路由功能，可以将入站数据按需有不同的出站连接发出，==常用与分流国内外流量==

示例 snippet

```
{
  "domainStrategy": "AsIs",
  "rules": [],
  "balancers": []
}
```

- domainStrategy: "AsIs" | "IPIfNonMatch" | "IPOnDemand"

  域名解析策略

  1. AsIs: 值使用域名进行路由选择。缺省值
  2. IPIfNonMatch: 当域名没有匹配任何规则时，将域名解析成 IP (A 记录或 AAA 记录)再次匹配
     - 当一个域名有多个 A 记录时，会尝试匹配所有的 A 记录，直到其中一个与某个规则匹配为止；
     - 解析后的 IP 仅在路由选择时起作用，转发的数据包中依然使用原始域名；
  3. IPOnDemand: 当匹配时碰到任何基于 IP 的规则，将域名立即解析为 IP 进行匹配；

- rules: [[RuleObject](https://www.v2ray.com/chapter_02/03_routing.html#ruleobject)]

  > AND 逻辑

  对应一个数组，数组中每个元素是一个规则。对于每一个连接，路由将根据这些规则依次进行判断，==当一个规则生效时，即将这个连接转发至它所指定的`outboundTag`(或`balancerTag`，V2Ray 4.4+)==。当没有匹配到任何规则时，流量默认由主出站协议发出

  示例 snippet

  ```
  {
    "type": "field",
    "domain": [
      "baidu.com",
      "qq.com",
      "geosite:cn"
    ],
    "ip": [
      "0.0.0.0/8",
      "10.0.0.0/8",
      "fc00::/7",
      "fe80::/10",
      "geoip:cn"
    ],
    "port": "53,443,1000-2000",
    "network": "tcp",
    "source": [
      "10.0.0.1"
    ],
    "user": [
      "love@v2ray.com"
    ],
    "inboundTag": [
      "tag-vmess"
    ],
    "protocol":["http", "tls", "bittorrent"],
    "attrs": "attrs[':method'] == 'GET'",
    "outboundTag": "direct",
    "balancerTag": "balancer"
  }
  ```

  - type

    目前只支持 field 一个值

  - domain: [string]

    数组，每一项是一个域名的匹配

    1. 纯字符串: 当此字符串匹配目标域名中任意部分，该规则生效。

       比如`"sina.com"`可以匹配`"sina.com"`、`"sina.com.cn"`和`"www.sina.com"`，但不匹配`"sina.cn"`。

    2. 正则表达式: 由`"regexp:"`开始，余下部分是一个正则表达式。当此正则表达式匹配目标域名时，该规则生效

       例如`"regexp:.*\.goo.*\.com$"`匹配`"www.google.com"`、`"fonts.googleapis.com"`，但不匹配`"google.com"`。

    3. 子域名 (推荐): 由`"domain:"`开始，余下部分是一个域名。当此域名是目标域名或其子域名时，该规则生效。

       例如`"domain:v2ray.com"`匹配`"www.v2ray.com"`、`"v2ray.com"`，但不匹配`"xv2ray.com"`。

    4. 完整匹配: 由`"full:"`开始，余下部分是一个域名。当此域名完整匹配目标域名时，该规则生效

       例如`"full:v2ray.com"`匹配`"v2ray.com"`但不匹配`"www.v2ray.com"`。

    5. 预定义域名列表：由`"geosite:"`开头，余下部分是一个名称，

       如`geosite:google`或者`geosite:cn`。名称及域名列表参考[预定义域名列表](https://www.v2ray.com/chapter_02/03_routing.html#dlc)。

    6. 从文件中加载域名: 形如`"ext:file:tag"`，必须以`ext:`（小写）开头，后面跟文件名和标签，文件存放在[资源目录](https://www.v2ray.com/chapter_02/env.html#asset-location)中，文件格式与`geosite.dat`相同，标签必须在文件中存在。

  - ip: [string]

    匹配目标 IP 时规则生效

    1. IP: 形如`"127.0.0.1"`。

    2. CIDR: 形如`"10.0.0.0/8"`.

    3. GeoIP: 形如 `"geoip:cn"` 必须以 `geoip:`（小写）开头，后面跟双字符国家代码，支持几乎所有可以上网的国家

       特殊值：`"geoip:private"` (V2Ray 3.5+)，包含所有私有地址，如`127.0.0.1`。

    4. 从文件中加载 IP: 形如`"ext:file:tag"`，必须以`ext:`（小写）开头，后面跟文件名和标签，文件存放在[资源目录](https://www.v2ray.com/chapter_02/env.html#asset-location)中，文件格式与`geoip.dat`相同标签必须在文件中存在。

  - port: number | string

    端口范围

    1. "a-b": a 和 b 均为正整数，且小于 65536。这个范围是一个前后闭合区间，当目标端口落在此范围内时，此规则生效。

    2. a: a 为正整数，且小于 65536。当目标端口为 a 时，此规则生效。

    3. (V2Ray 4.18+) 以上两种形式的混合，以逗号","分隔。形如：`"53,443,1000-2000"`

  - network: tcp | udp | tcp,udp

    匹配 4 层协议时 rule 生效

  - source:[string]

    一个数组，数组内每一个元素是一个 IP 或 CIDR。当某一元素匹配来源 IP 时，此规则生效

  - user: [string]

    一个数组，数组内每一个元素是一个邮箱地址。当某一元素匹配来源用户时，此规则生效。当前 Shadowsocks 和 VMess 支持此规则

  - inboundTag: [string]

    一个数组，数组内每一个元素是一个标识。当某一元素匹配入站协议的标识时，此规则生效

  - protocol: [ "http" | "tls" | "bittorrent" ]

    一个数组，数组内每一个元素表示一种协议。当某一个协议匹配当前连接的流量时，此规则生效。必须开启入站代理中的`sniffing`选项。

  - attrs: string

    (V2Ray 4.18+) 一段脚本，用于检测流量的属性值。当此脚本返回真值时，此规则生效。

    脚本语言为 [Starlark](https://github.com/bazelbuild/starlark)，它的语法是 Python 的子集。脚本接受一个全局变量`attrs`，其中包含了流量相关的属性。

    目前只有 http 入站代理会设置这一属性。

    示例：

    - 检测 HTTP GET: `"attrs[':method'] == 'GET'"`
    - 检测 HTTP Path: `"attrs[':path'].startswith('/test')"`
    - 检测 Content Type: `"attrs['accept'].index('text/html') >= 0"`

  - outboundTag: string

    对应一个[额外出站连接配置](https://www.v2ray.com/chapter_02/02_protocols.html)的标识

  - balancerTag: string

    对应一个负载均衡器的标识。`balancerTag`和`outboundTag`须二选一。当同时指定时，`outboundTag`生效。

- balancers:  [BalancerObject](https://www.v2ray.com/chapter_02/03_routing.html#balancerobject) 

  (V2Ray 4.4+)一个数组，数组中每个元素是一个负载均衡器的配置。当一个规则指向一个负载均衡器时，V2Ray 会通过此负载均衡器选出一个出站协议，然后由它转发流量

  示例 snippet

  ```
  {
    "tag": "balancer",
    "selector": []
  }
  ```

  - tag: string

    此负载均衡器的标识，用于匹配`RuleObject`中的`balancerTag`

  - selector: [ string ]

    一个字符串数组，其中每一个字符串将用于和出站协议标识的前缀匹配。在以下几个出站协议标识中：`[ "a", "ab", "c", "ba" ]`，`"selector": ["a"]`将匹配到`[ "a", "ab" ]`。

    如果匹配到多个出站协议，负载均衡器目前会从中随机选出一个作为最终的出站协议。

### policy

本地策略可以配置一些用户相关的权限，比如连接超时设置。V2Ray 处理的每一个连接，都对应到一个用户，按照这个用户的等级（level）应用不同的策略。本地策略可按照等级的不同而变化

示例 snippet

```
{
  "levels": {
    "0": {
      "handshake": 4,
      "connIdle": 300,
      "uplinkOnly": 2,
      "downlinkOnly": 5,
      "statsUserUplink": false,
      "statsUserDownlink": false,
      "bufferSize": 10240
    }
  },
  "system": {
    "statsInboundUplink": false,
    "statsInboundDownlink": false
  }
}
```

- level: map{string: [LevelPolicyObject](https://www.v2ray.com/chapter_02/policy.html#levelpolicyobject)}

  一组键值对，每个键是一个字符串形式的数字（JSON 的要求），比如 `"0"`、`"1"` 等，双引号不能省略，这个数字对应用户等级。每一个值是一个 [LevelPolicyObject](https://www.v2ray.com/chapter_02/policy.html#levelpolicyobject).

  LevelPolicyObject

  ```
  {
    "handshake": 4,
    "connIdle": 300,
    "uplinkOnly": 2,
    "downlinkOnly": 5,
    "statsUserUplink": false,
    "statsUserDownlink": false,
    "bufferSize": 10240
  }
  ```

  - handshake: number

    连接建立时的握手时间限制。单位为秒。默认值为`4`。在入站代理处理一个新连接时，在握手阶段（比如 VMess 读取头部数据，判断目标服务器地址），如果使用的时间超过这个时间，则中断该连接。

  - `connIdle`: number

    连接空闲的时间限制。单位为秒。默认值为`300`。在入站出站代理处理一个连接时，如果在 `connIdle` 时间内，没有任何数据被传输（包括上行和下行数据），则中断该连接。

  - `uplinkOnly`: number

    当连接下行线路关闭后的时间限制。单位为秒。默认值为`2`。当服务器（如远端网站）关闭下行连接时，出站代理会在等待 `uplinkOnly` 时间后中断连接。

  - `downlinkOnly`: number

    当连接上行线路关闭后的时间限制。单位为秒。默认值为`5`。当客户端（如浏览器）关闭上行连接时，入站代理会在等待 `downlinkOnly` 时间后中断连接。

  - `statsUserUplink`: true | false

    当值为`true`时，开启当前等级的所有用户的上行流量统计。

  - `statsUserDownlink`: true | false

    当值为`true`时，开启当前等级的所有用户的下行流量统计。

  - `bufferSize`: number

    每个连接的内部缓存大小。单位为 kB。当值为`0`时，内部缓存被禁用。

    默认值 (V2Ray 4.4+):

    - 在 ARM、MIPS、MIPSLE 平台上，默认值为`0`。
    - 在 ARM64、MIPS64、MIPS64LE 平台上，默认值为`4`。
    - 在其它平台上，默认值为`512`。

    默认值 (V2Ray 4.3-):

    - 在 ARM、MIPS、MIPSLE、ARM64、MIPS64、MIPS64LE 平台上，默认值为`16`。
    - 在其它平台上，默认值为`2048`。

- system: [SystemPolicyObject](https://www.v2ray.com/chapter_02/policy.html#systempolicyobject)

  SystemPolicyObject

  示例 snippet

  ```
  {
    "statsInboundUplink": false,
    "statsInboundDownlink": false
  }
  ```

  - `statsInboundUplink`: true | false

    当值为`true`时，开启所有入站代理的上行流量统计。

  - `statsInboundDownlink`: true | false

    当值为`true`时，开启所有入站代理的下行流量统计。

### reverse

https://www.v2ray.com/chapter_02/reverse.html

### inbounds

入站连接用于接收从客户端（浏览器或上一级代理服务器）发来的数据，可用的协议请见[协议列表](https://www.v2ray.com/chapter_02/02_protocols.html)

示例 snippet

```
{
  "port": 1080,
  "listen": "127.0.0.1",
  "protocol": "协议名称",
  "settings": {},
  "streamSettings": {},
  "tag": "标识",
  "sniffing": {
    "enabled": false,
    "destOverride": ["http", "tls"]
  },
  "allocate": {
    "strategy": "always",
    "refresh": 5,
    "concurrency": 3
  }
}
```

- `port`: number | "env:variable" | string

  端口。接受的格式如下:

  1. 整型数值: 实际的端口号。
  2. 环境变量: 以`"env:"`开头，后面是一个环境变量的名称，如`"env:PORT"`。V2Ray 会以字符串形式解析这个环境变量。
  3. 字符串: 可以是一个数值类型的字符串，如`"1234"`；或者一个数值范围，如`"5-10"`表示端口 5 到端口 10 这 6 个端口。

  当只有一个端口时，V2Ray 会在此端口监听入站连接。当指定了一个端口范围时，取决于`allocate`设置。

- `listen`: address

  监听地址，只允许 IP 地址，默认值为`"0.0.0.0"`，表示接收所有网卡上的连接。除此之外，必须指定一个现有网卡的地址。

- `protocol`: string

  连接协议名称，可选的值见[协议列表](https://www.v2ray.com/chapter_02/02_protocols.html)。

- `settings`: InboundConfigurationObject

  具体的配置内容，视协议不同而不同。详见每个协议中的`InboundConfigurationObject`。

  具体参考

  https://www.v2ray.com/chapter_02/02_protocols.html

- `streamSettings`: [StreamSettingsObject](https://www.v2ray.com/chapter_02/05_transport.html#perproxy)。

  [底层传输配置](https://www.v2ray.com/chapter_02/05_transport.html#perproxy)

  StreamSettingsObject

  示例 snippet

  ```
  {
    "network": "tcp",
    "security": "none",
    "tlsSettings": {},
    "tcpSettings": {},
    "kcpSettings": {},
    "wsSettings": {},
    "httpSettings": {},
    "dsSettings": {},
    "quicSettings": {},
    "sockopt": {
      "mark": 0,
      "tcpFastOpen": false,
      "tproxy": "off"
    }
  }
  ```

  - `network`: "tcp" | "kcp" | "ws" | "http" | "domainsocket" | "quic"

    数据流所使用的网络类型，默认值为 `"tcp"`

  - `security`: "none" | "tls"

    是否启入传输层加密，支持的选项有 `"none"` 表示不加密（默认值），`"tls"` 表示使用 [TLS](https://en.wikipedia.org/wiki/Transport_Layer_Security)。

  - `tlsSettings`: [TLSObject](https://www.v2ray.com/chapter_02/05_transport.html#tlsobject)

    TLS 配置。TLS 由 Golang 提供，支持 TLS 1.2，不支持 DTLS。

  - `tcpSettings`: [TcpObject](https://www.v2ray.com/chapter_02/transport/tcp.html)

    当前连接的 TCP 配置，仅当此连接使用 TCP 时有效。配置内容与上面的全局配置相同。

  - `kcpSettings`: KcpObject

    当前连接的 mKCP 配置，仅当此连接使用 mKCP 时有效。配置内容与上面的全局配置相同。

  - `wsSettings`: WebSocketObject

    当前连接的 WebSocket 配置，仅当此连接使用 WebSocket 时有效。配置内容与上面的全局配置相同。

  - `httpSettings`: HttpObject

    当前连接的 HTTP/2 配置，仅当此连接使用 HTTP/2 时有效。配置内容与上面的全局配置相同。

  - `dsSettings`: DomainSocketObject

    当前连接的 Domain socket 配置，仅当此连接使用 Domain socket 时有效。配置内容与上面的全局配置相同。

  - `quicSettings`: QUICObject

    (V2Ray 4.7+) 当前连接的 QUIC 配置，仅当此连接使用 QUIC 时有效。配置内容与上面的全局配置相同。

  - `sockopt`: SockoptObject

    连接选项

- `tag`: string

  此入站连接的标识，用于在其它的配置中定位此连接。当其不为空时，其值必须在所有`tag`中唯一。

- `sniffing`: [SniffingObject](https://www.v2ray.com/chapter_02/01_overview.html#sniffingobject)

  SniffingObject

  尝试探测流量的类型

  示例 snippet

  ```
  {
    "enabled": false,
    "destOverride": ["http", "tls"]
  }
  ```

  - `enabled`: true | false

    是否开启流量探测。

  - `destOverride`: ["http" | "tls"]

    当流量为指定类型时，按其中包括的目标地址重置当前连接的目标。

- `allocate`: [AllocateObject](https://www.v2ray.com/chapter_02/01_overview.html#allocateobject)

  端口分配设置

  AllocateObject

  示例 snippet

  ```
  {
    "strategy": "always",
    "refresh": 5,
    "concurrency": 3
  }
  ```

  - `strategy`: "always" | "random"

    端口分配策略。`"always"`表示总是分配所有已指定的端口，`port`中指定了多少个端口，V2Ray 就会监听这些端口。`"random"`表示随机开放端口，每隔`refresh`分钟在`port`范围中随机选取`concurrency`个端口来监听。

  - `refresh`: number

    随机端口刷新间隔，单位为分钟。最小值为`2`，建议值为`5`。这个属性仅当`strategy = random`时有效。

  - `concurrency`: number

    随机端口数量。最小值为`1`，最大值为`port`范围的三分之一。建议值为`3`

### outbounds

出站连接用于向远程网站或下一级代理服务器发送数据，可用的协议请见[协议列表](https://www.v2ray.com/chapter_02/02_protocols.html)。

示例 snippet

```
{
  "sendThrough": "0.0.0.0",
  "protocol": "协议名称",
  "settings": {},
  "tag": "标识",
  "streamSettings": {},
  "proxySettings": {
    "tag": "another-outbound-tag"
  },
  "mux": {}
}
```

- `sendThrough`: address

  用于发送数据的 IP 地址，当主机有多个 IP 地址时有效，默认值为`"0.0.0.0"`。

- `protocol`: string

  连接协议名称，可选的值见[协议列表](https://www.v2ray.com/chapter_02/02_protocols.html)。

- `settings`: OutboundConfigurationObject

  ==核心代理配置==

  具体的配置内容，视协议不同而不同。详见每个协议中的`OutboundConfigurationObject`。

  具体参考

  https://www.v2ray.com/chapter_02/02_protocols.html

- `tag`: string

  此出站连接的标识，用于在其它的配置中定位此连接。当其值不为空时，必须在所有 tag 中唯一。

- `streamSettings`: [StreamSettingsObject](https://www.v2ray.com/chapter_02/05_transport.html#perproxy)。

  [底层传输配置](https://www.v2ray.com/chapter_02/05_transport.html#perproxy)

  StreamSettingsObject

  示例 snippet

  ```
  {
    "network": "tcp",
    "security": "none",
    "tlsSettings": {},
    "tcpSettings": {},
    "kcpSettings": {},
    "wsSettings": {},
    "httpSettings": {},
    "dsSettings": {},
    "quicSettings": {},
    "sockopt": {
      "mark": 0,
      "tcpFastOpen": false,
      "tproxy": "off"
    }
  }
  ```

  - `network`: "tcp" | "kcp" | "ws" | "http" | "domainsocket" | "quic"

    数据流所使用的网络类型，默认值为 `"tcp"`

  - `security`: "none" | "tls"

    是否启入传输层加密，支持的选项有 Dokodemo`"none"` 表示不加密（默认值），`"tls"` 表示使用 [TLS](https://en.wikipedia.org/wiki/Transport_Layer_Security)。

  - `tlsSettings`: [TLSObject](https://www.v2ray.com/chapter_02/05_transport.html#tlsobject)

    TLS 配置。TLS 由 Golang 提供，支持 TLS 1.2，不支持 DTLS。

  - `tcpSettings`: [TcpObject](https://www.v2ray.com/chapter_02/transport/tcp.html)

    当前连接的 TCP 配置，仅当此连接使用 TCP 时有效。配置内容与上面的全局配置相同。

  - `kcpSettings`: KcpObject

    当前连接的 mKCP 配置，仅当此连接使用 mKCP 时有效。配置内容与上面的全局配置相同。

  - `wsSettings`: WebSocketObject

    当前连接的 WebSocket 配置，仅当此连接使用 WebSocket 时有效。配置内容与上面的全局配置相同。

  - `httpSettings`: HttpObject

    当前连接的 HTTP/2 配置，仅当此连接使用 HTTP/2 时有效。配置内容与上面的全局配置相同。

  - `dsSettings`: DomainSocketObject

    当前连接的 Domain socket 配置，仅当此连接使用 Domain socket 时有效。配置内容与上面的全局配置相同。

  - `quicSettings`: QUICObject

    (V2Ray 4.7+) 当前连接的 QUIC 配置，仅当此连接使用 QUIC 时有效。配置内容与上面的全局配置相同。

  - `sockopt`: SockoptObject

    连接选项

- `proxySettings`: [ProxySettingsObject](https://www.v2ray.com/chapter_02/01_overview.html#proxysettingsobject)

  出站代理配置。当出站代理生效时，此出站协议的`streamSettings`将不起作用。

  示例 snippet

  ```
  {
    "tag": "another-outbound-tag"
  }
  ```

  - `tag`: string

    当指定另一个出站协议的标识时，此出站协议发出的数据，将被转发至所指定的出站协议发出。

- `mux`: [MuxObject](https://www.v2ray.com/chapter_02/mux.html)

  [Mux 配置](https://www.v2ray.com/chapter_02/mux.html)。

### transport

https://www.v2ray.com/chapter_02/05_transport.html

底层传输方式（transport）是当前 V2Ray  节点和其它节点对接的方式。底层传输方式提供了稳定的数据传输通道。通常来说，一个网络连接的两端需要有对称的传输方式。比如一端用了  WebSocket，那么另一个端也必须使用 WebSocket，否则无法建立连接。

底层传输（transport）配置分为两部分，一是全局设置([TransportObject](https://www.v2ray.com/chapter_02/05_transport.html#transportobject))，二是分协议配置([StreamSettingsObject](https://www.v2ray.com/chapter_02/05_transport.html#streamsettingsobject))。分协议配置可以指定每个单独的入站出站协议用怎样的方式传输。通常来说客户端和服务器对应的出站入站协议需要使用同样的传输方式。当分协议传输配置指定了一种传输方式，但没有填写其设置时，此传输方式会使用全局配置中的设置。

















