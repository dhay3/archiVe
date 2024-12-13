# IPtables-extensions TEE Target

## Digest

TEE target 可以将本机收到的报文克隆到另外一台机器

## Optional args

- `--gateway ipaddr`

  send the cloned packet to the host reachable at the given IP address

  不能是 0.0.0.0 或者是 ::

## Examples

以 192.168.1.1 访问 192.168.3.1 流量镜像到 192.168.4.1 上

192.168.3.1 设置规则

```
[root@labos-2 /]# iptables -t mangle -A PREROUTING -s 192.168.1.1 -j TEE --gateway 192.168.4.1
```

192.168.1.1 ping 192.168.3.1

```
[root@labos-1 /]# ping 192.168.3.1
PING 192.168.3.1 (192.168.3.1) 56(84) bytes of data.
64 bytes from 192.168.3.1: icmp_seq=1 ttl=62 time=42.1 ms
64 bytes from 192.168.3.1: icmp_seq=1 ttl=62 time=51.9 ms (DUP!)
64 bytes from 192.168.3.1: icmp_seq=1 ttl=62 time=62.8 ms (DUP!
```

