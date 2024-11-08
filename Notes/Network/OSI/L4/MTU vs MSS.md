# MTU vs MSS

![2021-09-06_15-42](https://github.com/dhay3/image-repo/raw/master/20210601/2021-09-06_15-42.6zgzrm9nbjo0.png)

## MTU

maximum transmission unit

帧或包能传输的最大值，在Ethernet中通常为1500byte。loopback默认为65536(a little bit of confuse)

```
cpl in ~ λ ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: wlp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 64:bc:58:bd:a6:19 brd ff:ff:ff:ff:ff:ff
    inet 30.131.88.17/24 brd 30.131.88.255 scope global dynamic noprefixroute wlp1s0
       valid_lft 5664sec preferred_lft 5664sec
    inet6 fe80::dc0f:b3f:1c2:d1c6/64 scope link noprefixroute 
       valid_lft forever preferred_lft forever

```

## MSS

maximum segment size

MTU除去包头，即数据的真实大小