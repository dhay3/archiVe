---
createTime: 2024-08-01 15:41
tags:
  - "#hash1"
  - "#hash2"
---

# SubConvert 01 - Overview

## 0x01 Overview

subconvert[^1] 是一个订阅地址转换工具(将一个 client 的配置转为另一个 client 可以使用的配置)

支持多种类型的订阅地址

> 实测发现也支持 singbox, targe=singbox

|类型|作为源类型|作为目标类型|参数|
|---|:-:|:-:|---|
|Clash|✓|✓|clash|
|ClashR|✓|✓|clashr|
|Quantumult (完整配置)|✓|✓|quan|
|Quantumult X (完整配置)|✓|✓|quanx|
|Loon|✓|✓|loon|
|Mellow|✓|✓|mellow|
|SS (SIP002)|✓|✓|ss|
|SS (软件订阅/SIP008)|✓|✓|sssub|
|SSD|✓|✓|ssd|
|SSR|✓|✓|ssr|
|Surfboard|✓|✓|surfboard|
|Surge 2|✓|✓|surge&ver=2|
|Surge 3|✓|✓|surge&ver=3|
|Surge 4|✓|✓|surge&ver=4|
|Trojan|✓|✓|trojan|
|V2Ray|✓|✓|v2ray|
|类 TG 代理的 HTTP/Socks 链接|✓|×|仅支持 `&url=` 调用|
|Mixed|×|✓|mixed|
|Auto|×|✓|auto|

## 0x02 Installation

subconvert 安装的方式很简单，具体可以看 README-docker.md[^2]

```sh
# run the container detached, forward internal port 25500 to host port 25500
docker run -d --restart=always -p 25500:25500 tindy2013/subconverter:latest
# then check its status
curl http://localhost:25500/version
# if you see `subconverter vx.x.x backend` then the container is up and running
```

或者使用 docker-compose

```yaml
version: '3'
services:
  subconverter:
    image: tindy2013/subconverter:latest
    container_name: subconverter
    ports:
      - "25500:25500"
    restart: always
```

可以通过如下命令来校验是否安装成功

```sh
curl http://localhost:25500/version
subconverter v0.9.0-6974910 backend
```

## 0x03 Simple Usage

> 要使用自建的 VPS 链接，具体看 subconvert README-cn.md[^1]

Simple Usage 只有 3 个参数

|调用参数|必要性|示例|解释|
|---|:-:|:--|---|
|target|必要|surge&ver=4|指想要生成的配置类型，详见上方 [支持类型](https://github.com/tindy2013/subconverter/blob/master/README-cn.md#%E6%94%AF%E6%8C%81%E7%B1%BB%E5%9E%8B) 中的参数|
|url|必要|https%3A%2F%2Fwww.xxx.com|指机场所提供的订阅链接或代理节点的分享链接，需要经过 [URLEncode](https://www.urlencoder.org/) 处理|
|config|可选|https%3A%2F%2Fwww.xxx.com|指 外部配置 的地址 (包含分组和规则部分)，需要经过 [URLEncode](https://www.urlencoder.org/) 处理，详见 [外部配置](https://github.com/tindy2013/subconverter/blob/master/README-cn.md#%E5%A4%96%E9%83%A8%E9%85%8D%E7%BD%AE) ，当此参数不存在时使用 程序的主程序目录中的配置文件|

假设你已经完成了 [0x02 Installation](#0x02%20Installation)

### 0x03a One Sub Convertion

现在有一个订阅地址需要转换成 clash 的格式

例如 `https://dler.cloud/subscribe/ABCDE?surge=ss`

1. 先将订阅地址做 URLencode
	可以通过 python 的 urllib 模块实现
	
	```sh
	python3 -c 'import urllib.parse; print(urllib.parse.quote("https://dler.cloud/subscribe/ABCDE?surge=ss"))'
	https%3A//dler.cloud/subscribe/ABCDE%3Fsurge%3Dss
	```

2. 拼接参数
	因为需要转为 clash ，所以 `target=clash` 
	原订阅地址为 `https://dler.cloud/subscribe/ABCDE?surge=ss`，encode 后所以为 `url=https%3A//dler.cloud/subscribe/ABCDE%3Fsurge%3Dss`
	
3. 拼接 new sub URL
	```
	http://127.0.0.1:25500/sub?target=clash&url=https%3A%2F%2Fdler.cloud%2Fsubscribe%2FABCDE%3Fsurge%3Dss
	```
	这个 URL 就可以作为新的订阅地址(subconvert container 必须是启动的)。如果想转存订阅地址中的配置直接使用 `curl -o` 即可

### 0x03b One More Sub Convertion

现在有多份订阅地址需要转换成 clash 的格式

例如

- `https://dler.cloud/subscribe/ABCDE?surge=ss`
- `https://rich.cloud/subscribe/ABCDE?clash=vmess`

1. 先使用管道符(`|`)将订阅地址分隔
	```
	https://dler.cloud/subscribe/ABCDE?surge=ss|https://rich.cloud/subscribe/ABCDE?clash=vmess
	```

2. 将使用管道符拼接的地址做 URLencode
	```
	python3 -c 'import urllib.parse; print(urllib.parse.quote("https://dler.cloud/subscribe/ABCDE?surge=ss|https://rich.cloud/subscribe/ABCDE?clash=vmess"))'
	https%3A//dler.cloud/subscribe/ABCDE%3Fsurge%3Dss%7Chttps%3A//rich.cloud/subscribe/ABCDE%3Fclash%3Dvmess
	```

3. 拼接参数
	`target=clash`
	`url=https%3A//dler.cloud/subscribe/ABCDE%3Fsurge%3Dss%7Chttps%3A//rich.cloud/subscribe/ABCDE%3Fclash%3Dvmess`
	
4. 拼接 new sub URL 
	```
	http://127.0.0.1:25500/sub?target=clash&url=https%3A%2F%2Fdler.cloud%2Fsubscribe%2FABCDE%3Fsurge%3Dss
	```
	使用方式通过 [0x03a One Sub Convertion](#0x03a%20One%20Sub%20Convertion)

## 0x04 Advanced Usage

具体看 subconvert README-cn.md[^1] 

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[subconverter/README-cn.md at master · tindy2013/subconverter · GitHub](https://github.com/tindy2013/subconverter/blob/master/README-cn.md)
[^2]:[subconverter/README-docker.md at master · tindy2013/subconverter · GitHub](https://github.com/tindy2013/subconverter/blob/master/README-docker.md)