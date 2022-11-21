# IPtables-extensions TTL Target

## Digest

设置 TTL 值，只在 mangle tables 中生效

## Optional args

- `--ttl-set value`

  set the TTL value to `value`

- `--ttl-dec value`

  decrement the TTL value `value` times

- `--ttl-inc value`

  Increment the TTL value `value` times

## Examples

例如 192.168.1.1 访问 192.168.4.1

192.168.1.1 设置规则

```
[root@labos-1 /]# iptables -t mangle -A OUTPUT -j TTL --ttl-set 1
```

从 192.168.1.1 ping 192.168.4.1

```
[root@labos-1 /]# ping 192.168.4.1
PING 192.168.4.1 (192.168.4.1) 56(84) bytes of data.
From 192.168.1.2 icmp_seq=1 Time to live exceeded
From 192.168.1.2 icmp_seq=2 Time to live exceeded
^C
--- 192.168.4.1 ping statistics ---
2 packets transmitted, 0 received, +2 errors, 100% packet loss, time 1002ms
```

192.168.1.1 抓包

可以看到 192.168.1.1 发出的报文的 TTL 值已经为 1 了

```
04:56:34.728643 da:97:7a:8f:b0:f5 > ca:01:09:a1:00:00, ethertype IPv4 (0x0800), length 98: (tos 0x0, ttl 1, id 10770, offset 0, flags [DF], proto ICMP (1), length 84)
    192.168.1.1 > 192.168.4.1: ICMP echo request, id 13, seq 1, length 64
04:56:34.739261 ca:01:09:a1:00:00 > da:97:7a:8f:b0:f5, ethertype IPv4 (0x0800), length 70: (tos 0xc0, ttl 255, id 9, offset 0, flags [none], proto ICMP (1), length 56)
    192.168.1.2 > 192.168.1.1: ICMP time exceeded in-transit, length 36
        (tos 0x0, ttl 1, id 10770, offset 0, flags [DF], proto ICMP (1), length 84)
    192.168.1.1 > 192.168.4.1: ICMP echo request, id 13, seq 1, length 64
```

