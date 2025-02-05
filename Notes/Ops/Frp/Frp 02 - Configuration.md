---
createTime: 2025-01-07 14:48
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# Frp 02 - Configuration

## 0x01 Preface

> 从 v0.52.0 开始，frp 支持以 toml/yaml/json 作为配置文件格式

frp 的配置分成 2 部分

- frps.toml
- frpc.toml


## 0x02 Environment Variables

frp 支持通过 Go template 获取环境变量

例如 

> 不建议直接将环境变量 export 到 parent shell

```
env FRP_SERVER_ADDR=192.168.56.1 FRP_SSH_REMOTE_PORT=6000 ./frpc -c ./frpc.toml
```

那么就可以通过如下格式获取环境变量

```
# frpc.toml
serverAddr = "{{ .Envs.FRP_SERVER_ADDR }}"
serverPort = 7000

[[proxies]]
name = "ssh"
type = "tcp"
localIP = "127.0.0.1"
localPort = 22
remotePort = {{ .Envs.FRP_SSH_REMOTE_PORT }}
```

## 0x03 Include Directive

> [!NOTE]
> 需要注意的是 includes 指定的文件中只能包含代理配置，通用参数的配置只能放在主配置文件中。

frp 支持通过 `includes` 指令引入其他配置

```
# frpc.toml
serverAddr = "x.x.x.x"
serverPort = 7000
includes = ["./confd/*.toml"]
```

```
# ./confd/test.toml
[[proxies]]
name = "ssh"
type = "tcp"
localIP = "127.0.0.1"
localPort = 22
remotePort = 6000
```


## 0x04 frps.toml

`frps.toml` 核心配置只有 2 个

```
#frps 监听的地址，默认 0.0.0.0
bindAddr = "192.168.2.1"
#frps 监听的端口，默认 7000
bindPort = 7000
```

其他配置按照场景来[^2]

## 0x05 frpc.toml

`frpc.toml` 核心配置如下

```
#frps 监听的地址
serverAddr = "192.168.56.1"
#frps 监听的端口，默认 7000
serverPort = 7000

#代理配置
[[proxies]]

name = "test-tcp"
type = "tcp"
localIP = "127.0.0.1"
localPort = 22
remotePort = 6000
```

其他配置按照场景来[^2]

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [示例 \| frp](https://gofrp.org/zh-cn/docs/examples/)
- [功能特性 \| frp](https://gofrp.org/zh-cn/docs/features/)
- [参考 \| frp](https://gofrp.org/zh-cn/docs/reference/)

***References***

[^1]:[配置文件 \| frp](https://gofrp.org/zh-cn/docs/features/common/configure/)
[^2]:[功能特性 \| frp](https://gofrp.org/zh-cn/docs/features/)
