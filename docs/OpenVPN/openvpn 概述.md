# openvpn 概述

参考：

https://openvpn.net/community-resources/how-to/

openvpn支持SSL/TLS，ethernet bridging，TCP or UDP tunnel，dynamic IP address and DHCP 是一款灵活的VPN软件

openvpn支持将选项放置在配置文件中或是命令行中，当在配置文件中double-leading-dash(--)需要被去除，`--config file`使用配置文件替代命令行参数(--config可以被去除)，`#`表示注释。syntax：`openvn [options...]`

## Tunnel options

- --mode m

  指定openvpn运行的模式，默认以`p2p`(point-to-point)模式，在openvpn2.0后可以使用`server`表示服务端连接多个客户端

- --dev tunX | tapX | null

  指定使用的虚拟设备，X可以被忽略自动使用动态设备。TUN在OSI layer 3工作，TAP在OSI Layer 2工作。参考：https://segmentfault.com/a/1190000009249039

- --user user | --group group

  指定openvpn tunnel所属用户和组

- --daemon [progname]

  以守护进程的方式启动openvpn，可以指定进程的progname来区分，所有的日志都会记录在syslog，也可以通过`--log file`将日志存储在一个具体的位置

- --local host

  将本地的一个host name 或 IP address绑定，如果没有指定openvpn将所有的IP都绑定

- --remote host [port] [proto]

  openvpn client 会通过指定的port 和 protocol连接remote host。如果想要指定多个可以通过client connection profile 来指定。==如果没有指定--remote==，openvpn会监听所有的IP，但是不会对数据包进行操作，只有通过authentication(static key or TLS)才会转发数据包。

- --proto-forece p

  当使用client connection profile时只使使用特定的protocol

- --proto p

  ==默认使用UDP通过客户端1194 port和服务端1194 port连接==，如果想要使用TCP连接，需要在客户端和服务端分别使用tcp-client和tcp-client，当连接失败时可以指定--connect-retry和--connect-retry-max。因为TCP面向连接，确保数据完整达到。

- --conect-retry n |--conect-retry-max

  指定尝试重连的时间间隔，尝试重连的次数

- --http-proxy server port [auth-file | auto | nct] [auth-method]

  通过http proxy 连接remote host

- --socks-proxy server [port] [authfile]

  通过socks5连接remote host

- --port port | --lport  | --rport

  指定local prot 和 rmote port 使用的端口，可以通过 --lport 或 --rport来分别指定

- --nobind

  使用动态端口绑定local address

- --topology mode

  只有在`--dev tun`才有意义，mode具体参考manual page

- --ifconfig local rn

  设置TUN/TAP配置器参数。

  1. local表示local VPN endpoint
  2. 对于TUN device p2p mode(默认模式)，rn表示表示remote VPN endpoint
  3. 对于TAP device或是 TUN device 使用 `--topology subnet`，rn表示连接的网络的subnet mask   

- --verb n

  openvpn 日志输出的等级

- --compression [lz4 | lzo]

  使用指定的加密算法，客户端和服务端需要相同

- --ping n

  在n seconds内没有发送ping packet，就会向remote server发送ping请求

- --ping-exit n

  如果在n seconds内没有收到ping packet，openvpn就会退出。和--ping-restart 不能一起使用

- --ping-restart n

  如果在n seconds内没有收到ping packet，openvpn就会重启。如果remote host用DDNS时但是TTL设置的较小时，就可以使用该参数

- --keepalive interval timeout

  等价与--ping和--png-restart，服务端的配置优先于客户端配置。例如--keepalive 10 60，等价于--ping 10，--ping-restart 60

- --persist-tun

  当ifcae的状态为up/down时不关闭TUN/TAP设备，例如指定了--ping-restart

- --persist-key

  确保openvpn不会重新读校验文件，例如指定了--ping-restart

## client mode

在openvpn 客户端使用

- --client 

  等价于`--pull`加`--tls-client`

- --pull

  表示连接一个multi-client server，可以将server的route table推送给client

- --pull-filter accept | ignore | reject text

  client过滤从server推送过来的optiono starts with text

- --auth-user-pass [up]

  up是一个文件，包含username/password on 2 lines 用于校验用户

## server mode

在openvpn 服务端使用

- --server network netmask

  等价于`--mode server`加`--tls-server`。服务端提供一对多功能，服务端通过network netmask自动为客户端分配IP。从第一个可用的host-id开始分配

- --push option

  将服务端的中option推送给客户端，例如`--push "port 1134"`即客户端使用1134端口进行通信，option必须在quote中

- --duplicate-cn

  允许openvpn 客户端以相同的common name连接服务端，默认如果之前有相同common name连接，新连接就会断开

## TLS mode 

- --tls-server

  表示是tls握手的服务端方

- --tls-client

  表示是tls握手的客户端方

- --ca file

  指定CA file，只能是PEM格式

- --cert file

  指定使用的crt文件，两端都需要

- --key file

  生成crt时使用的私钥

- --dh file

  使用Diffie Hellman而不是RSA

## connection

一组`<connection></connection>`用于定义client connection profile(openvpn client 如何连接opnevpn server，可以理解就是一个--remote)。当一组client connection profile连接失败后使用下一组继续尝试

```
#在<connection>之外都作为全局配置，在每一个<connection>中都生效
dev tun
#使用udp连接198.19.34.56:1194
<connection>
remote 198.19.34.56 1194 udp
</connection>

<connection>
remote 198.19.34.56 443 tcp
</connection>
#通过HTTP proxy 192.168.0.8:8080使用tcp连接198.19.34.56:443
<connection>
remote 198.19.34.56 443 tcp
http-proxy 192.168.0.8 8080
</connection>

<connection>
remote 198.19.36.99 443 tcp
http-proxy 192.168.0.8 8080
</connection>
#全局配置
persist-key
persist-tun
pkcs12 client.p12
remote-cert-tls server
verb 3
```

每一组client connection profile中可用参数具体参考manual page



