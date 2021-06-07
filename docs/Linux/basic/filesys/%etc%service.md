# /etc/service 

> man services

记录着网络打开服务的列表，port由IANA(Internet Assigned Numbers Authority)分配，每个服务都会默认按照改文件分配端口。

每一行由`service-name  port/protocol  [aliases...]`组成

查看具体的端口的服务，并不一定打开

```
[root@k8snode01 ~]# grep -w '80/tcp' /etc/services
http            80/tcp          www www-http    # WorldWideWeb HTTP

[root@k8snode01 ~]# ss -lnpt | grep -w 80
```

