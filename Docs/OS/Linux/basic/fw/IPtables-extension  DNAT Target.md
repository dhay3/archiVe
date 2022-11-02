# IPtables-extension  DNAT Target

ref

https://www.zsythink.net/archives/1764

## Digest

DNAT target 只能在 nat table 中的 PREROUTING 和 OUTPUT 中使用

## Optional args

- `--to-destination [[ipaddr[-ipaddr]][:port[-port[/baseport]]]`

  which can specify a single new destination IP address, an inclusive range of IP addresses. Optionally a port range, if the rule also specifies  one of the following protocols: tcp, udp, dccp or sctp.  If no port range is specified, then the destination port will never be modified. If no IP address is specified then only the destination port will be modified.  If baseport is given, the difference of the  original destination port and its value is used as offset into the mapping port range. This allows to create shifted portmap ranges and is available since kernel version 4.18.  For a single port or baseport, a service name as listed in /etc/services may be used.

  根据指定的参数决定是否使用 PNAT

## Example

以访问 39.156.66.10 DNAT 成 180.97.251.233 为例子

设置 DNAT 规则

```
sudo iptables -t nat -A OUTPUT -d 39.156.66.10 -j DNAT --to-destination 180.97.251.233

sudo iptables -t nat -nvL OUTPUT
Chain OUTPUT (policy ACCEPT 350 packets, 37995 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    0   180 DNAT       all  --  *      *       0.0.0.0/0            39.156.66.10         to:180.97.251.233
```

这是使用 curl 访问 39.156.66.10

```
curl -svL 39.156.66.10
*   Trying 39.156.66.10:80...
* Connected to 39.156.66.10 (39.156.66.10) port 80 (#0)
> GET / HTTP/1.1
> Host: 39.156.66.10
> User-Agent: curl/7.84.0
> Accept: */*
> 
```

这里显示的是 `connected to 39.156.66.10` 实际应该是 180.97.251.233。因为应用输出到 stdout 的内容是最优先的(在 OSI7 之前)，但是可以通过抓包看到报文是已经 DNAT 的了

![2022-10-31_23-16](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221031/2022-10-31_23-16.1noi1a38ay3k.webp)

因为 nfilter 在 tcpdump 或者 wireshark 之前也在 NIC 处理报文之前

另外这里还需要注意的一点是一次请求的所有报文 IPtables pkts 只会记录一次，并没按照报文个数计算

```
cpl in ~ λ sudo iptables -t nat -nvL OUTPUT
Chain OUTPUT (policy ACCEPT 13 packets, 2311 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    1    60 DNAT       all  --  *      *       0.0.0.0/0            39.156.66.10         to:180.97.251.233
```

