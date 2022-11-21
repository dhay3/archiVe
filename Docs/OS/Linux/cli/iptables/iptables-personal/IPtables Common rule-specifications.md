# IPtables Common rule-specifications

ref

https://www.cnblogs.com/sunsky303/p/12327863.html



- 机器能正常访问外部服务器，但是外部服务器不能访问除 22 端口以外的任何端口或者其他协议的报文

```
iptables -N INGRESS

iptables -t filter -A INGRESS -m state --state established,related -j ACCEPT
iptables -t filter -A INGRESS -p tcp --dport 22 -j ACCEPT
iptables -t filter -A INGRESS -j DROP

iptables -t filter -A INPUT -j INGRESS
```

- 限制 22 端口连接, 需要访问一下 65522 端口, 在 10 min 内可以连接 22 端口

```
iptables -N SSHOPEN

iptables -t filter -A SSHOPEN -m state --state established,related -j ACCEPT
iptables -t filter -A SSHOPEN -p tcp --dport 65522 -m recent --name SSHOPEN --set
iptables -t filter -A SSHOPEN -p tcp --dport 65522 -j DROP
iptables -t filter -A SSHOPEN -p tcp --dport 22 -m recent --name SSHOPEN --rcheck --seconds 600 -j ACCEPT
iptables -t filter -A SSHOPEN -j DROP

iptables -t filter -A INPUT -j SSHOPEN
```

- 对 HTTPS 报文负载均衡

```
#iptables -A PREROUTING -i eth0 -p tcp --dport 443 -m state --state NEW -m statistics --mode nth --packet 0 --every 3 --packet 0 -j DNAT --to-destination 192.168.1.101:443

#iptables -A PREROUTING -i eth0 -p tcp --dport 443 -m state --state NEW -m statistics --mode nth --packet 0 --every 3 --packet 0 -j DNAT --to-destination 192.168.1.102:443

#iptables -A PREROUTING -i eth0 -p tcp --dport 443 -m state --state NEW -m statistics --mode nth --packet 0 --every 3 --packet 0 -j DNAT --to-destination 192.168.1.103:443
```

- 对 80 端口服务做 Dos 防护

```
iptables -A INPUT -p tcp --dport 80 -m limit --limit 25/minute --limit-burst 100 -j ACCEPT
```

