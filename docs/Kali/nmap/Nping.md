# Nping

## 概述

Nping是Nmap中集成的utility，可以被用做普通的Ping命令，也可以用来生成raw packet进行network stack tests，ARP posioning，DDos attack

syntax：`Nping [options] <target>`

## options

### misc

- `-c | --count <n>`

  发包的个数

- -H | --hide-sent

  不输出发送的数据包(做flooding attack时使用)

- -N | --no-capture

  不输出回送的数据包(做flooding attack时使用)

### timing

- `--delay <time>`

  每执行一次probe间隔的时间 ms | s | m | h

- `--rate <rate>`

  每执行一次probe间隔的包数

### ipv4 options

- -S | --source-ip

- `--ttl <hops>`

  路由允许的最大跳数

- `--mtu <size>`

  传输介质会影响mtu

### payload options

- `--data-length <len>`

  ```
  cpl in /etc/X11/xinit/xinitrc.d λ sudo nping --tcp --data-length 200 taobao.com
  
  #tcp 首部40byte
  Starting Nping 0.7.91 ( https://nmap.org/nping ) at 2021-07-08 18:28 HKT
  #从0开始计时
  SENT (0.0773s) TCP 30.226.76.49:16913 > 140.205.220.96:80 S ttl=64 id=17913 iplen=240  seq=834338448 win=1480 
  #从0开始计时
  RCVD (0.0872s) TCP 140.205.220.96:80 > 30.226.76.49:16913 SA ttl=51 id=17913 iplen=240  seq=3784034306 win=1480 
  SENT (1.0783s) TCP 30.226.76.49:16913 > 140.205.220.96:80 S ttl=64 id=17913 iplen=240  seq=834338448 win=1480 
  RCVD (1.1268s) TCP 140.205.220.96:80 > 30.226.76.49:16913 SA ttl=51 id=17913 iplen=240  seq=3798473671 win=1480 
  ```

  send和recv为一组，seq值不会增加

- `--data-string <text>`

  ```
  cpl in /etc/X11/xinit/xinitrc.d λ sudo nping --tcp --data-string lucky taobao.com
  
  Starting Nping 0.7.91 ( https://nmap.org/nping ) at 2021-07-08 18:31 HKT
  SENT (0.0576s) TCP 30.226.76.49:61074 > 140.205.94.189:80 S ttl=64 id=53424 iplen=45  seq=3691758155 win=1480 
  RCVD (0.0632s) TCP 140.205.94.189:80 > 30.226.76.49:61074 SA ttl=53 id=53424 iplen=45  seq=1905578548 win=1480 
  SENT (1.0584s) TCP 30.226.76.49:61074 > 140.205.94.189:80 S ttl=64 id=53424 iplen=45  seq=3691758155 win=1480 
  RCVD (1.0671s) TCP 140.205.94.189:80 > 30.226.76.49:61074 SA ttl=53 id=53424 iplen=45  seq=1905578548 win=1480 
  ```

- `--data <hex string>`

### output

- -v | -v[level]

  输出信息的详细程度，level越高越详细

  ```
  nping --tcp -v4 taobao.com
  ```

## probe modes

nping默认发送数据20B，实际需要加上各协议报文的首部(tcp20，udp8，icmp8)

### icmp probe mode

如果没有指定probe mode默认使用icmp probe mode

```
cpl in ~ λ sudo nping --icmp -c 1 taobao.com

Starting Nping 0.7.91 ( https://nmap.org/nping ) at 2021-07-09 09:13 HKT
SENT (0.0767s) ICMP [30.226.76.49 > 140.205.94.189 Echo request (type=8/code=0) id=8158 seq=1] IP [ttl=64 id=48881 iplen=28 ]
 
Max rtt: N/A | Min rtt: N/A | Avg rtt: N/A
Raw packets sent: 1 (28B) | Rcvd: 0 (0B) | Lost: 1 (100.00%)
Nping done: 1 IP address pinged in 1.13 seconds
```

- `--icmp-type <type>`

  icmp type 参考https://erg.abdn.ac.uk/users/gorry/course/inet-pages/icmp-code.html#:~:text=Many%20of%20the%20types%20of,%2C%20Time%20Exceeded%20(11).&text=Many%20of%20these%20ICMP%20types%20have%20a%20%22code%22%20field.

- `--icmp-code <code>`

### tcp probe mode

使用`--tcp`切换到tcp probe mode

```
cpl in ~ λ sudo nping --tcp -c 1 taobao.com

Starting Nping 0.7.91 ( https://nmap.org/nping ) at 2021-07-09 09:13 HKT
#从0开始计时
SENT (0.0498s) TCP 30.226.76.49:53865 > 140.205.94.189:80 S ttl=64 id=7595 iplen=40  seq=2719624812 win=1480 
#从零开始计时
RCVD (0.0538s) TCP 140.205.94.189:80 > 30.226.76.49:53865 SA ttl=53 id=7595 iplen=40  seq=3696106700 win=1480 
 
Max rtt: 3.894ms | Min rtt: 3.894ms | Avg rtt: 3.894ms
Raw packets sent: 1 (40B) | Rcvd: 1 (46B) | Lost: 0 (0.00%)
Nping done: 1 IP address pinged in 1.11 seconds

```

- `-p | --dest-port <port spec>`

  默认dest 80端口

- `-g | --source-port <portnumber>`

  默认src随机端口

- `--seq <seqnumber>`

  指定数据包的seq序列号

- `--flags <flags list>`

  指定数据包标志位(ACK，SYN，FIN，RST)

  ```
  cpl in /etc/NetworkManager λ sudo nping --tcp --flags SYN --win 2048 --seq 1 taobao.com
  
  #这里可以看到seq没有增加，说明只是重复发送只有头部的tcp报文
  Starting Nping 0.7.91 ( https://nmap.org/nping ) at 2021-07-09 09:17 HKT
  SENT (1.1068s) TCP 30.226.76.49:55275 > 140.205.94.189:80 S ttl=64 id=63241 iplen=40  seq=1 win=2048 
  RCVD (1.6725s) TCP 140.205.94.189:80 > 30.226.76.49:55275 SA ttl=53 id=63241 iplen=40  seq=2118196352 win=2048 
  SENT (2.1069s) TCP 30.226.76.49:55275 > 140.205.94.189:80 S ttl=64 id=63241 iplen=40  seq=1 win=2048 
  RCVD (2.6706s) TCP 140.205.94.189:80 > 30.226.76.49:55275 SA ttl=53 id=63241 iplen=40  seq=2118196352 win=2048 
  SENT (3.1075s) TCP 30.226.76.49:55275 > 140.205.94.189:80 S ttl=64 id=63241 iplen=40  seq=1 win=2048 
  RCVD (3.6722s) TCP 140.205.94.189:80 > 30.226.76.49:55275 SA ttl=53 id=63241 iplen=40  seq=2118196352 win=2048 
  SENT (4.1078s) TCP 30.226.76.49:55275 > 140.205.94.189:80 S ttl=64 id=63241 iplen=40  seq=1 win=2048 
  RCVD (4.6741s) TCP 140.205.94.189:80 > 30.226.76.49:55275 SA ttl=53 id=63241 iplen=40  seq=2118196352 win=2048 
  SENT (5.1078s) TCP 30.226.76.49:55275 > 140.205.94.189:80 S ttl=64 id=63241 iplen=40  seq=1 win=2048 
  RCVD (5.6731s) TCP 140.205.94.189:80 > 30.226.76.49:55275 SA ttl=53 id=63241 iplen=40  seq=2118196352 win=2048 
   
  Max rtt: 566.230ms | Min rtt: 563.635ms | Avg rtt: 565.081ms
  Raw packets sent: 5 (200B) | Rcvd: 5 (230B) | Lost: 0 (0.00%)
  Nping done: 1 IP address pinged in 5.76 seconds
  ```

- `--ack <acknumber>`

  指定数据包的确认号

- `--win <size>`

  确认接收窗口的大小

- --badsum

  使用随机的CRC冗余码

==不会建立完整的tcp连接，探测后会发送RST包==

```
#发送单个tcp包
cpl in ~ λ sudo nping -c 1 --tcp 1.1.1.1

#这里可以看到1.1.1.1发送ACK SYN后主机响应了一个RST
cpl in ~ λ sudo tcpdump -nni wlp1s0 tcp and host 1.1.1.1
tcpdump: verbose output suppressed, use -v[v]... for full protocol decode
listening on wlp1s0, link-type EN10MB (Ethernet), snapshot length 262144 bytes
12:45:30.214545 IP 30.226.76.42.4244 > 1.1.1.1.80: Flags [S], seq 363736440, win 1480, length 0
12:45:30.457450 IP 1.1.1.1.80 > 30.226.76.42.4244: Flags [S.], seq 4235195346, ack 363736441, win 65535, options [mss 1460], length 0
12:45:30.457501 IP 30.226.76.42.4244 > 1.1.1.1.80: Flags [R], seq 363736441, win 0, length 0

```

### upd probe mode

使用`--udp`切换到udp probe mode

```
cpl in /etc/NetworkManager λ nping --udp -c 1 taobao.com

Starting Nping 0.7.91 ( https://nmap.org/nping ) at 2021-07-09 09:20 HKT
SENT (0.0113s) UDP packet with 4 bytes to taobao.com:40125 (140.205.94.189:40125)
 
Max rtt: N/A | Min rtt: N/A | Avg rtt: N/A
UDP packets sent: 1 | Rcvd: 0 | Lost: 1 (100.00%)
Nping done: 1 IP address pinged in 1.01 seconds
```

- `-g | --source-port <portnumber>`
- `-p | --dest-port <port spec>`
- --badsum

### arp/rarp probe mode

使用`--arp`切换到udp probe mode

```
root in /home/ubuntu λ nping --arp taobao.com

Starting Nping 0.7.60 ( https://nmap.org/nping ) at 2021-07-09 09:27 CST
SENT (0.1326s) ARP who has 140.205.94.189? Tell 172.21.16.3
RCVD (0.3337s) ARP reply 140.205.94.189 is at FE:EE:4E:0D:BD:1E
SENT (1.1328s) ARP who has 140.205.94.189? Tell 172.21.16.3
RCVD (1.1497s) ARP reply 140.205.94.189 is at FE:EE:4E:0D:BD:1E
SENT (2.1347s) ARP who has 140.205.94.189? Tell 172.21.16.3
RCVD (2.1697s) ARP reply 140.205.94.189 is at FE:EE:4E:0D:BD:1E
^C 
Max rtt: N/A | Min rtt: N/A | Avg rtt: N/A
Raw packets sent: 3 (126B) | Rcvd: 3 (150B) | Lost: 0 (0.00%)
Nping done: 1 IP address pinged in 2.34 seconds
```

由于目的地址不在同一网段，所以会返回网关的MAC

```
#使用arp命令需要安装net-tools
root in /home/ubuntu λ arp -n
Address                  HWtype  HWaddress           Flags Mask            Iface
172.17.0.2               ether   02:42:ac:11:00:02   C                     docker0
172.21.16.1              ether   fe:ee:4e:0d:bd:1e   C                     eth0

root in /home/ubuntu λ route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         172.21.16.1     0.0.0.0         UG    100    0        0 eth0
172.16.253.0    172.16.253.2    255.255.255.0   UG    0      0        0 tun1194
172.16.253.2    0.0.0.0         255.255.255.255 UH    0      0        0 tun1194
172.17.0.0      0.0.0.0         255.255.0.0     U     0      0        0 docker0
172.21.16.0     0.0.0.0         255.255.240.0   U     0      0        0 eth0
192.168.122.0   0.0.0.0         255.255.255.0   U     0      0        0 virbr0
```

- `--arp-type <type>`

  ARP | ARP-reply | RARP | RARP-reply

- `arp-sender-mac <mac>`

- `--arp-sender-ip <addr>`

- `--arp-target-mac <mac>`

- `--arp-target-ip <addr>`

```
#伪造arp请求
root in /home/ubuntu λ nping --arp --arp-sender-mac 02:42:ac:11:00:02 --arp-sender-ip 1.1.1.1 taobao.com

Starting Nping 0.7.60 ( https://nmap.org/nping ) at 2021-07-09 09:30 CST
SENT (0.0040s) ARP who has 140.205.94.189? Tell 1.1.1.1
RCVD (0.2060s) ARP reply 140.205.94.189 is at FE:EE:4E:0D:BD:1E
SENT (1.0042s) ARP who has 140.205.94.189? Tell 1.1.1.1
RCVD (1.0220s) ARP reply 140.205.94.189 is at FE:EE:4E:0D:BD:1E
SENT (2.0060s) ARP who has 140.205.94.189? Tell 1.1.1.1
RCVD (2.0420s) ARP reply 140.205.94.189 is at FE:EE:4E:0D:BD:1E
^C 
Max rtt: N/A | Min rtt: N/A | Avg rtt: N/A
Raw packets sent: 3 (126B) | Rcvd: 3 (150B) | Lost: 0 (0.00%)
Nping done: 1 IP address pinged in 2.56 seconds
```

## 例子

小包攻击

```
nping -HN --tcp --flag SYN --data-length 0 --delay 100ms -c 10000 taobao.com
```

