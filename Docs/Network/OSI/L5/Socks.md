# Socks

ref

https://en.wikipedia.org/wiki/SOCKS

https://securityintelligence.com/posts/socks-proxy-primer-what-is-socks5-and-why-should-you-use-it/

## Digest

Socks（socket secure） is an Internet protocol that exchanges network packets between a client and server through a proxy

Socks 是一个应用在 proxy 上的 L5 协议( session layer) 用于交换 client 和 server 的流量。因为 Socks 工作中在 L5 ，==这也就意味 Socks 不仅能代理 HTTP ，DNS 流量也能带出其他协议的流量, 只要你的流量是向外发送的他就能代理==

## Configure socks

> 如果有软路由可以使用 passwall 快速配置

https://linuxize.com/post/how-to-setup-ssh-socks-tunnel-for-private-browsing/

https://www.simplified.guide/ssh/create-socks-proxy

https://www.comparitech.com/blog/vpn-privacy/how-to-set-up-a-socks5-proxy-on-a-virtual-private-server-vps/

在客户端上使用下面命令就可以直接建立 socks 通道

```
ssh -fCqN -D 9090 [USER]@[SERVER_IP]
```

- `-f`

  Requests ssh to go to background just before command execution

-  `-C`

  Requests compression of all data

- `-q`

  Quiet mode.  Causes most warning and diagnostic messages to be suppressed

- `-N` 

   Tells SSH not to execute a remote command.

- `-D 9090`

  Opens a SOCKS tunnel on the specified port number.

通过下面的命令来校验

```
curl -vx socks5://localhost:9090 ipinfo.io
```

