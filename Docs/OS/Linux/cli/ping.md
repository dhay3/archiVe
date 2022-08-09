# Linux ping

ref:

https://www.ibm.com/docs/en/qsip/7.4?topic=applications-icmp-type-code-ids

https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol

https://gursimar27.medium.com/customizing-icmp-payload-in-ping-command-7c4486f4a1be

## Digest

syntax:

```
ping [-aAbBdDfhLnOqrRUvV46] [-c count] [-F flowlabel] [-i interval]
[-I interface] [-l preload] [-m mark] [-M pmtudisc_option]
[-N nodeinfo_option] [-w deadline] [-W timeout] [-p pattern]
[-Q tos] [-s packetsize] [-S sndbuf] [-t ttl]
[-T timestamp option] [hop...] {destination}
```

一个用来往 destination 发 ICMP type 8 echo_request 的网络工具

## Output

```
ubuntu@VM-12-16-ubuntu:~$ ping 39.156.66.10
PING 39.156.66.10 (39.156.66.10) 56(84) bytes of data.
64 bytes from 39.156.66.10: icmp_seq=1 ttl=249 time=37.2 ms
64 bytes from 39.156.66.10: icmp_seq=2 ttl=249 time=37.2 ms
64 bytes from 39.156.66.10: icmp_seq=3 ttl=249 time=37.2 ms
^C
--- 39.156.66.10 ping statistics ---
3 packets transmitted, 3 received, 0% packet loss, time 2002ms
rtt min/avg/max/mdev = 37.164/37.184/37.197/0.014 ms
```

ping 输出几部分:

1. first summary line 

   ```
   PING 39.156.66.10 (39.156.66.10) 56(84) bytes of data.
   ```

   描述了 ICMP request 发包的对象是 39.156.66.10

   发送的 ICMP payload 是 56 byte，封包 84 byte （56 + 8 ICMP header + 20 IP header）

2. panel

   ```
   64 bytes from 39.156.66.10: icmp_seq=1 ttl=249 time=37.2 ms
   ```

   `64byte`：表示从 39.156.66.10 收到的回包 64 byte (56 + 8, 这里 IP header 被 demultiplexing 了 )

   `icmp_seq=1`：表明这是发第 1 个包收到的回包

   `ttl=249`：说明从 39.156.66.10 经过的路由设备可能有 6 hops（255 - 249）

   `time=37.2ms`：表示 rrt 是 37.2 ms

3. finished summary line

   ```
   3 packets transmitted, 3 received, 0% packet loss, time 2002ms
   rtt min/avg/max/mdev = 37.164/37.184/37.197/0.014 ms
   ```

   `3 packets transmitted`：表示一共发了 3 个包

   `3 received`：表示一共收到的回包有 3 个

   `0% packet loss`：表示丢包率为 0%

   `time 2002ms`：表示从首包到尾包一共的时间

   `rtt min/avg/max/mdev = 37.164/37.184/37.197/0.014 ms`：round trip time(rtt)往返时延最小值，平均值，最大值，与平均值的偏差

## Optional args

### Input args

- `-A`

  以200ms间隔发包，几乎等价于 flood

- `-f `

  flood ping，不回显数据，需要 root 权限

- `-l preload`

  异步发包，ping 每次发送 preload 次数的包而不需要同步等回包

- `-B`

  ping 发包时只绑定一个 IP，不会变更

- `-I interface`

  设置发包的 IP 或者 iface

- `-w deadline`

  specify a timeout, in seconds, before ping exits regardless of how many packets have been sent or received

  ping 进程允许存活的最大时间

- `-W timeout`

  time to wait for a response

  等待回包的最大时间

- `-s packetsize`

  设置发包的大小（ICMP payload data），默认是 56 byte，以无意义的字符填充

- `-t ttl`

  设置默认发包的ttl，默认使用 OS 上的`net.ipv4.ip_default_ttl`

- `-c count`

  waits  for count ICMP echo reply packets, util the timeout exprires

  原文并不是指发送几个包，但是实际测试即 ICMP 发包是同步的，即就是指定发几个包

  ```
  ubuntu@VM-12-16-ubuntu:~$ ping -c 3 93.46.8.90
  PING 93.46.8.90 (93.46.8.90) 56(84) bytes of data.
  
  --- 93.46.8.90 ping statistics ---
  3 packets transmitted, 0 received, 100% packet loss, time 2049ms
  
  ubuntu@VM-12-16-ubuntu:~$ sudo tcpdump -nni eth0 icmp and host 93.46.8.90
  tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
  listening on eth0, link-type EN10MB (Ethernet), capture size 262144 bytes
  13:25:59.883495 IP 10.0.12.16 > 93.46.8.90: ICMP echo request, id 11, seq 1, length 64
  13:26:00.908205 IP 10.0.12.16 > 93.46.8.90: ICMP echo request, id 11, seq 2, length 64
  13:26:01.932188 IP 10.0.12.16 > 93.46.8.90: ICMP echo request, id 11, seq 3, length 64
  ```

- `-i interval`

  每次发包的间隔时间， 默认 1s。root 可以设置c成小于 2 ms 即  0.02

- `-m mark`

  https://unix.stackexchange.com/questions/281015/m-option-does-not-work-in-ping-command

  use mark to tag the packets going out

- `-n`

  numeric output

- `-p pattern`

  ==可以用做打标==

  指定 3 层包的 payload, 需要使用 hexcimal

  ```
  ubuntu@VM-12-16-ubuntu:~$ sudo ping -p ffff  39.156.66.10
  PATTERN: 0xffff
  PING 39.156.66.10 (39.156.66.10) 56(84) bytes of data.
  64 bytes from 39.156.66.10: icmp_seq=1 ttl=249 time=37.1 ms
  
  ubuntu@VM-12-16-ubuntu:~$ sudo tcpdump -nnvvvXi eth0 icmp and host 39.156.66.10
  tcpdump: listening on eth0, link-type EN10MB (Ethernet), capture size 262144 bytes
  13:51:45.445537 IP (tos 0x0, ttl 64, id 60867, offset 0, flags [DF], proto ICMP (1), length 84)
      10.0.12.16 > 39.156.66.10: ICMP echo request, id 25, seq 5, length 64
          0x0000:  4500 0054 edc3 4000 4001 cd2f 0a00 0c10  E..T..@.@../....
          0x0010:  279c 420a 0800 c3bc 0019 0005 f1f5 f162  '.B............b
          0x0020:  0000 0000 4acc 0600 0000 0000 ffff ffff  ....J...........
          0x0030:  ffff ffff ffff ffff ffff ffff ffff ffff  ................
          0x0040:  ffff ffff ffff ffff ffff ffff ffff ffff  ................
          0x0050:  ffff ffff                                ....
  13:51:45.482719 IP (tos 0xa0, ttl 249, id 60867, offset 0, flags [DF], proto ICMP (1), length 84)
      39.156.66.10 > 10.0.12.16: ICMP echo reply, id 25, seq 5, length 64
          0x0000:  45a0 0054 edc3 4000 f901 138f 279c 420a  E..T..@.....'.B.
          0x0010:  0a00 0c10 0000 cbbc 0019 0005 f1f5 f162  ...............b
          0x0020:  0000 0000 4acc 0600 0000 0000 ffff ffff  ....J...........
          0x0030:  ffff ffff ffff ffff ffff ffff ffff ffff  ................
          0x0040:  ffff ffff ffff ffff ffff ffff ffff ffff  ................
          0x0050:  ffff ffff                                ....
  ```

  数据填充部分会变成`ffff`

- `-Q tos`

  ==可以用做打标==

  set the tos field ( quality of service ), tos can be decimal or hex number

  ```
  ubuntu@VM-12-16-ubuntu:~$ sudo ping -Q 0xff  39.156.66.10
  PING 39.156.66.10 (39.156.66.10) 56(84) bytes of data.
  64 bytes from 39.156.66.10: icmp_seq=1 ttl=249 time=38.9 ms
  
  ubuntu@VM-12-16-ubuntu:~$ sudo tcpdump -nnvi eth0 icmp and host 39.156.66.10
  tcpdump: listening on eth0, link-type EN10MB (Ethernet), capture size 262144 bytes
  13:58:29.213444 IP (tos 0xff,CE, ttl 64, id 48646, offset 0, flags [DF], proto ICMP (1), length 84)
      10.0.12.16 > 39.156.66.10: ICMP echo request, id 28, seq 6, length 64
  13:58:29.250585 IP (tos 0xa0, ttl 249, id 48646, offset 0, flags [DF], proto ICMP (1), length 84)
      39.156.66.10 > 10.0.12.16: ICMP echo reply, id 28, seq 6, length 64
  ```

- `-r`

  bypass the normal routing tables and send directly to a host on an attached interface

  只看同 LAN 中 3 层 ICMP 是否可达

  ```
  ubuntu@VM-12-16-ubuntu:~$ ping -r baidu.com
  PING baidu.com (110.242.68.66) 56(84) bytes of data.
  ping: sendmsg: Network is unreachable
  ping: sendmsg: Network is unreachable
  ```

  如果路由不可达，会显示 network is unreachable

### Output args

- `-q`

  quiet output

  only summary lines an startup time and when finished are displayed

- `-O`

  如果包被丢了会被打印，一般和`-D`一起使用用做 debug
  
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
  
- `-D`

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

- `-R`

  record route

  和 traceroute 一样显示 3 层的 route path

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

## Size/data

```
ubuntu@VM-12-16-ubuntu:~$ ping 39.156.66.10
PING 39.156.66.10 (39.156.66.10) 56(84) bytes of data.
64 bytes from 39.156.66.10: icmp_seq=1 ttl=249 time=37.2 ms

ubuntu@VM-12-16-ubuntu:~$ sudo tcpdump -nnvi eth0 icmp and host 39.156.66.10
tcpdump: listening on eth0, link-type EN10MB (Ethernet), capture size 262144 bytes
14:14:48.496306 IP (tos 0x0, ttl 64, id 35092, offset 0, flags [DF], proto ICMP (1), length 84)
    10.0.12.16 > 39.156.66.10: ICMP echo request, id 34, seq 1, length 64
14:14:48.533451 IP (tos 0xa0, ttl 249, id 35092, offset 0, flags [DF], proto ICMP (1), length 84)
    39.156.66.10 > 10.0.12.16: ICMP echo reply, id 34, seq 1, length 64
```

![](https://www.researchgate.net/profile/Md-Nazmul-Islam-2/publication/316727741/figure/fig5/AS:614213521268736@1523451323001/ICMP-packet-structure.png)

ping 默认发 56 byte 作为 ICMP payload data，加 8 byte ICMP header

其中根据不同 OS 上 ping 实现的逻辑方式不同，ICMP payload data 部分不同

Linux 上一般填充会填充

```
空字符 !"#$%&'()*+,-./01234567
```

Windows 上一般会填充

```
abcdefghijklmnopqrstuvwabcdefg hi
```

## Cautions

1. 如果ICMP不通但是TCP可以同可能是路由器的策略，主机iptables，内核参数(`/proc/sys/net/ipv4/icmp_ignore_all==1`)

## Examples

### 0x01 ICMP 常用探测命令

```
ubuntu@VM-12-16-ubuntu:~$ ping -DOs 0 baidu.com
PING baidu.com (110.242.68.66) 0(28) bytes of data.
[1660028438.535737] 8 bytes from 
```

### 0x02 ICMP 打标

同时打标 ICMP payload data 和 IP tos

```
ubuntu@VM-12-16-ubuntu:~$ ping -Q 0xff -p ff baidu.com
PATTERN: 0xff
PING baidu.com (110.242.68.66) 56(84) bytes of data.
64 bytes from 110.242.68.66 (110.242.68.66): icmp_seq=1 ttl=51 time=26.9 ms
```



