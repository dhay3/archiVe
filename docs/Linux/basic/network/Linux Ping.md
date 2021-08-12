# Ping

> 常用`ping -Dnc 3 <host>`



如果ICMP不通但是TCP可以同可能是路由器的策略，主机iptables，内核参数(/proc/sys/net/ipv4/icmp_ignore_all 1)

ping强制使用ICMP echo(8)发送数据包，默认带有IP首部(60 byte)和ICMP首部(8 byte)

```
root in /tmp λ tcpdump -i eth0 -n -vvv host baidu.com                                                                   
tcpdump: listening on eth0, link-type EN10MB (Ethernet), capture size 262144 bytes                                      
11:41:26.435929 IP (tos 0x0, ttl 64, id 49677, offset 0, flags [DF], proto ICMP (1), length 84)                             
172.21.16.3 > 220.181.38.148: ICMP echo request, id 21080, seq 1, length 64          
11:41:26.439521 IP (tos 0xa0, ttl 251, id 49677, offset 0, flags [DF], proto ICMP (1), length 84)                           220.181.38.148 > 172.21.16.3: ICMP echo reply, id 21080, seq 1, length 64
```

syntax：`ping [options] destination`

## options

- -A

  以200ms间隔发包，等价于flood

- -f 

  flood ping，不回显数据

- -c count

  发送指定个数的数据包

- -i <interval>

  每次发包的间隔时间，默认1s

- -n

  numeric output

- -O

  如果前一个包不可用，这打印一般和-D一起使用

```
root in /tmp λ ping -OD dtstack.com
PING dtstack.com (114.55.58.88) 56(84) bytes of data.
[1625809131.843766] no answer yet for icmp_seq=1
[1625809132.867795] no answer yet for icmp_seq=2
[1625809133.891769] no answer yet for icmp_seq=3
^C
--- dtstack.com ping statistics ---
4 packets transmitted, 0 received, 100% packet loss, time 3059ms
```

- -r
- -R

  record route

```
root in /tmp λ ping -R taobao.com
PING taobao.com (140.205.94.189) 56(124) bytes of data.
64 bytes from 140.205.94.189 (140.205.94.189): icmp_seq=1 ttl=41 time=78.8 ms
NOP
RR:     gns3vm (172.21.16.3)
        9.61.119.35 (9.61.119.35)
        10.196.65.242 (10.196.65.242)
        0.0.0.0
        10.200.1.42 (10.200.1.42)
        119.38.219.45 (119.38.219.45)
        42.120.241.42 (42.120.241.42)
        42.120.239.181 (42.120.239.181)
        10.86.122.34 (10.86.122.34)

64 bytes from 140.205.94.189 (140.205.94.189): icmp_seq=2 ttl=41 time=46.1 ms
NOP
```

- -s <packagesize>
- -t <ttl>
- -D 

  print timestamp

```
root in /tmp λ ping -D taobao.com
PING taobao.com (140.205.220.96) 56(84) bytes of data.
[1625808887.143203] 64 bytes from 140.205.220.96 (140.205.220.96): icmp_seq=1 ttl=251 time=27.3 ms
[1625808888.038573] 64 bytes from 140.205.220.96 (140.205.220.96): icmp_seq=2 ttl=251 time=27.3 ms
[1625808889.039930] 64 bytes from 140.205.220.96 (140.205.220.96): icmp_seq=3 ttl=251 time=27.3 ms
^C
--- taobao.com ping statistics ---
3 packets transmitted, 3 received, 0% packet loss, time 2002ms
rtt min/avg/max/mdev = 27.308/27.347/27.369/0.193 ms
```

- -W <timeout>