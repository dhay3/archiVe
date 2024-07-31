---
createTime: 2024-07-29 11:19
tags:
  - "#Passwall"
  - "#Clash"
---

# Mihomo 06 - Proxy Providers

## 0x01 Overview

从指定 URL(一般是订阅地址) 或者是 文件 中获取 proxies 部分，具体配置看如下文档

[代理集合配置 - 虚空终端 Docs](https://wiki.metacubex.one/config/proxy-providers/#_2)

[代理集合内容 - 虚空终端 Docs](https://wiki.metacubex.one/config/proxy-providers/content/)

> [!NOTE] 
> 如果使用了 proxy providers，在 clash verge rev 中要在 Proxies tab 中的 Proxy Provider 按钮中点 update all 才会更新节点信息

```yaml
proxy-providers:
  provider1:
    type: http
    url: "https://example.org"
    path: ./proxy_providers/Wallless.yaml
    interval: 3600
    proxy: DIRECT
    health-check:
      enable: false
  provider2:
    type: http
    url: "https://example.org"
    path: ./proxy_providers/YiYuan.yaml 
    interval: 3600
    proxy: DIRECT
    health-check:
      enable: false
```



---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**
