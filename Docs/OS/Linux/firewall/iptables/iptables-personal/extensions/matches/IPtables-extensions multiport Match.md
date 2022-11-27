# IPtables-extensions multiport Match

## Digest

> 注意原生的`--dport`和`--sport`是不支持选择非连续端口的，所以需要使用 multiport match

matching a set of source or destination ports

只能和 4 层协议一起使用，即 `-p` 可以使用的值是 tcp, udp, udplite, dccp 和 sctp

## Optional args

- `[!] --source-ports | --sports port,[port|,port:port]...`
- `[!] --destination-ports,--dports port[,port|,port:port]...`
- `[!] --ports port[,port|,port:port]...`

## Exmaples

```
[root@labos-1 /]# iptables -t filter -A INPUT -p tcp -s 192.168.3.1 -m multiport --dports 1,2:65535 -j DROP
[root@labos-1 /]# iptables -nvL INPUT
Chain INPUT (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 DROP       tcp  --  *      *       192.168.3.1          0.0.0.0/0            multiport dports 1,2:65535

[root@labos-2 /]# nc -nv 192.168.1.1 65533
Ncat: Version 7.70 ( https://nmap.org/ncat )
Ncat: Connection timed out.
```

#### 