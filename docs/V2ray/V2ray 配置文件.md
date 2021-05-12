# V2ray 配置文件

参考：

https://www.v2ray.com/chapter_02/

V2ray默认使用Protobuf配置，但是为了方便阅读也支持JSON。在运行之前V2ray会将JSON转为Protobuf。

> JSON可以使用`//`或`/**/`注释

## 配置

==v2ray不分客户端和服务端的配置，但要区分inbound入站，和outbound出站在不同端的含义==

- 在v2ray客户端(被代理机)，inbound表示浏览器发送过来的数据包或是v2ray服务器返回的数据包，outbound表示发送到的代理服务器。
- 在v2ray服务端(代理机)，inbound表示接受来自v2ray客户端的数据包，outbound表示实际访问的目标站点

配置文件都以`config.json`命名

### 客户端

```json
{
  "log": {
    "access": "/tmp/v2ray/_access.log",
    "error": "/tmp/v2ray/_error.log",
    "loglevel": "info"
  },
  "stats": null,
  "api": null,
  "policy": {
    "system": {
      "statusUserUplink": false,
      "statusUserDownlink": false
    },
    "levels": {
      "0": {
        "handshake": 4,
        "connIdle": 300,
        "uplinkOnly": 0,
        "downlinkOnly": 0,
        "statusUserUplink": false,
        "statusUserDownlink": false,
        "bufferSize": 10240
      },
      "1": {
        "handshake": 5,
        "connIdle": 300,
        "uplinkOnly": 2,
        "downlinkOnly": 2,
        "statusUserUplink": false,
        "statusUserDownlink": false,
        "bufferSize": 10240
      }
    }
  },
  "dns": {
    "tag": "dns-server",
    "hosts": {
      "github.com": "13.250.177.223",
      "geosite:category-ads": "127.0.0.1"
    },
    "servers": [
      "localhost",
      {
        "address": "https://8.8.8.8/dns-query",
        "port": "853"
      },
      "https://dns.google:853/dns-query"
    ]
  },
  "inbounds": [
    {
      "tag": "socks-inbound-noauth",
      "port": 1080,
      "listen": "0.0.0.0",
      "protocol": "socks",
      "settings": {
        "auth": "noauth",
        "accounts": null,
        "udp": true,
        "ip": "127.0.0.1",
        "userLevel": 1
      },
      "sniffing": {
        "enabled": true,
        "destOverride": [
          "http",
          "tls"
        ]
      },
      "allocate": {
        "strategy": "always",
        "refresh": null,
        "concurrency": null
      },
      "streamSettings": {
        "network": "tcp",
        "security": "none",
        "tlsSettings": null,
        "tcpSettings": null,
        "wsSettings": null,
        "httpSettings": null,
        "dsSettings": null,
        "quicSettings": null,
        "sockopt": {
          "mark": 0,
          "tcpFastopen": false,
          "tproxy": "off"
        }
      }
    }
  ],
  "outbounds": [
    {
      "sedThrough": "0.0.0.0",
      "protocol":"vmess",
    }
  ]
}
```

### 服务端

















