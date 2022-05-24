# parameters proc-sys-net

ref：
[https://www.kernel.org/doc/html/latest/admin-guide/sysctl/index.html/proc/sys/net](https://www.kernel.org/doc/html/latest/admin-guide/sysctl/index.html/proc/sys/net)
[https://www.kernel.org/doc/html/latest/admin-guide/sysctl/net.html](https://www.kernel.org/doc/html/latest/admin-guide/sysctl/net.html)
[https://www.kernel.org/doc/html/latest/networking/ip-sysctl.html](https://www.kernel.org/doc/html/latest/networking/ip-sysctl.html)
[https://linuxconfig.org/how-to-turn-on-off-ip-forwarding-in-linux](https://linuxconfig.org/how-to-turn-on-off-ip-forwarding-in-linux)
[https://man7.org/linux/man-pages/man7/tcp.7.html](https://man7.org/linux/man-pages/man7/tcp.7.html)
[https://ixnfo.com/en/changing-gc_thresh-on-linux.html](https://ixnfo.com/en/changing-gc_thresh-on-linux.html)

## Dir

`/proc/sys/net/`下的文件主要是用来配置网络的，包含如下几个子文件

| subDir | brief |
| --- | --- |
| core | general parameter |
| unix | unix domain sockets |
| 802 | E802 protocol |
| ethernet | ethernet protocol |
| ipv4 | ip version 4 |
| bridge | bridging |
| ipv6 | ip version 6 |
| appletalk | appletalk protocol |
| netrom | net/rom |
| ax205 | ax25 |
| rose | x.25 plp player |
| decnet | dec net |
| tipc | tipc |
| x25 | x.25 protocol |


这里只记录常见的参数，其他参数可执行翻看内核文档

## core

-  rmem_default
the default setting of the socket receive buffer in bytes
提一嘴，套接字在 TCP/UDP 中使用，不仅仅只针对 TCP 
-  rmem_max
the maximum receive socket buffer size in bytes 
-  wmem_default
the default setting of the socket send buffer in bytes 
-  wmem_max
the maximum send socket buffer size in bytes 
-  message_burst / message_cost
these parameters are used to limit the warning message written to the kernel log from the networking. They enforce a rate limit to make denial-of-service attack impossible
用来限制内核记录网络相关的日志，通过调整这两个参数还可以限制 Dos 占用系统的资源(减少了因为Dos而产生的日志IO)
message_cost 的值越高，内核记录的日志就越少。mesage_burst 控制信息什么时候丢弃
the default settings limit warning message to one every five seconds 
```
net.core.message_burst = 10
net.core.message_cost = 5
```

不应该调整该参数即使出现 Dos 时，日志是关键信息 

-  netdev_max_backlog
maximum number of packets, queued on the INPUT side, when the interface receives packets faster than kernel can process them
INPUT 链中的等待队列的最大长度 

## ipv4

### grocery

-  ip_forward - boolean 
   - 0 : disable (default)
   - not 0 : enabled
   
   forward packets between interfaces
   如果linux server 扮演了 firewall，router，NAT device 时需要开启该参数 ( forward packets are meant for other destinations other than itself)
   注意如果使用了这个参数会导致其他配置变成默认值
   
- ip_default_ttl - integer 
   - default 64

   ttl 默认值，可以修改该值达到混淆扫描工具的目的(例如nmap就是通过ttl来判断target的OS) 

- fwmark_reflect - boolean
  如果生成和 socket 无关的回包( tcp reset or icmp replies)，是否设置 fwmark 标识 
  - 0：不设置，缺省值
  - 1：设置
- route.max_size -integer
  maximum number  of routes entry（路由条目） allowed in the kernel 

### ip fragmentation

-  ipfrag_high_thresh - integer
maximum memory used to reassemble ip fragments
ip 分片重组允许消耗的最大内存 
-  ipfrag_time - integer
time in seconds to keep an ip fragment in memory 
-  ipfrag_max_dist - integer
   reordering of  packets is not unusual, but if a large number of fragments arrive from a source ip address while a particular fragment queue remains incompelete, it probably indicated that one or more fragments beloging to that queue have been lost. 
   - 正数但是值比较小：result dropping fragment queues when normal reordering of packets occurs, which could lead to poor application performance
   - 正数但是比较大：increase the likelihood of incorrectly reassembling ip fragments that originate from different ip datagrams, which could result in data corruption
   - default: 64
-  inet_peer_threshold - integer
ip fragmets 在重组完成前能占用的最大内存，如果超过这个数值就会启动 GC 
-  inet_peer_maxttl - integer
ip fragments 在重组完成前能存活的最大时间 in seconds，如果超过这个时间就会过期直接被丢弃 
-  inet_peer_mintll - integer
ip fragments 在重组完成前至少能存活的时间 in seconds 

### mtu/mss

-  ip_no_pmtu_disc - integer
   disable [path mtu discovery(pmtud)](https://en.wikipedia.org/wiki/Path_MTU_Discovery)
   这是IPv4中的协议，mtu（以太网帧的最大值，不包含14字节的以太网帧头） = mss + tcp header + ip header，对以太网一般来说是 1500，但是有些广域网可能只有 576。
   假设中间链路的某台设备mtu只有576，源有1500，这时候就需要 pmtud。这台设备会把包丢弃然后回送一个 ICMP type 3 (fragmentation needed)，然后源调整发包的mtu后重新发包
   但是大多数防火墙都会禁止ICMP，所以可能会造成tcp握手成功，但是发数据包的时候出现超时重传 
   - 1: 如果收到 fragmentation need icmp 回包，源的 mtu 会设置成 the old mtu to this destination and `min_pmtu`
   - 2: fragmentation need icmp 回包会被丢弃，其他的和 1 一样
   - 3：...
   - default：0
-  min_pmtu - integer
   minium path mtu 
   - default: 552
-  mtu_expires - integer
time in seconds that cached pmtu informatino is kept 
-  route.min_adv_mss - integer
由第一跳的路由器推荐的最小 mss 

### arp

-  neigh.default.gc_thresh1 - integer
   minimum number of entries to keep. garbage collector will not purge enries if there are fewer than this number
   `ip neigh`用来管理arp信息，所以这里的 entries 指的就是 arp entires 。如果系统的 arp cache 条目小于该值就不会清理 
   - default：128
-  neigh.default.gc_thresh2 - integer
   threshold when garbage collector becomes more aggressive about purging entries. entries older than 5 seconds will be cleared when over this number 
   - default: 512
-  neigh.default.gc_thresh3 - integer
   maximum number of non-permanent neighbor entries allowed 
   - default: 1024

### tcp
具体可以 man tcp 来查看
#### 拥塞参数

- tcp_abc  - integer

  调整拥塞窗口

     - 0： increase cwnd once peer acknowledgment
     - 1：increase cwnd once peer acknowledgment of full sized segment
     - 2：allow increase cwnd by two if acknowledge is of two segments to compensate for delayed acknowledgments

- tcp_allowed_congestion_control - string

  show/set the congestion control choices available to non-privileged processes

- tcp_available_congestion_control

  show the available congestion control choices that are registered
  可以被 tcp_allowed_congestion_control 设置的 tcp 拥塞参数

- tcp_congestion_control - string

  set the default congestion-control algorithm to be used for new connections
  使用的 tcp 拥塞算法，默认会使用 reno，可以额外加参数

- tcp_ecn - integer

  是否显示声明拥塞通知，ECN

     - 0：关闭ECN功能
     - 1：允许ECN，入向和出向都可以发带有ECN的包
     - 2：允许ECN，入向回包允许发送ECN，出向不允许发ECN包

#### 重传参数

- tcp_frto - integer

  enable F-RTO, an enhanced recovery algorithm for TCP retransmission timeouts(RTOs)

     - 0：disabled
     - 1：basic frto is enabled
     - 2：enable sack-enhanced frto if flow uses SACK

#### 长连接参数

- tcp_keepalive_intvl - integer

  发送 tcp keepalive probe 的间隔秒，默认 75 sec

- tcp_keepalive_probes - integer

  the maximum number of TCP keepalive probes to send before giving up and killing the connection if no response is obtained from the other end，默认 9
  发送 tcp keepalive probes 的最大次数

- tcp_keepalive_probes - integer

  the number of seconds a connection needs to be idle before TCP begins sending out keep-alive probes.默认 2h，但是一般在 75*9 sec 如果没有回包，就会断连

#### 连接队列参数

- tcp_max_tw_buckets - integer

  the maximum number of sockets in TIME_WAIT state allowed in the system.
  通常用于防止Dos，但是不建议减小改值，反而需要增大该值如果有大量服务请求的话
  为什么能防止Dos，假设Dos使用 TCP SYN攻击，当服务器给对方回ACK包时，对方没有响应，服务器主动发送FIN包，然后 
  ffffffffffffffffffffffffffff

  

- tcp_max_syn_backlog - integer

  the maximum number of queued connection reqeusts which have still not received an acknowledgement from the connection client，即处于 SYN_RECV 状态的socket最大值(半连接队列)
  如果超过了改数字，内核就会开始丢包。当可用内存大于128MB是默认1024，如果小于32MB时会减小到128。通常应用程序也可以自己来设置半连接队列的最大值。但是需要注意一个socket 从 SYN_RECV 到 ESTABLISHED 需要消耗 304 bytes

- somaxconn - integer
  limit of socket listen() backlog，socket max connection 即server 全连接队列的最大值，详情可以查看listen函数()，backlog入参 
   - default: 4096 (was 128 before linux 5.4)

#### 窗口参数

- tcp_mem

- tcp_moderate_rcvbuf - boolean

  if enabled TCP performs receive buffer auto-tuning，but no greater than tcp_rmem
  是否开启recvbuf自动调整，不能操作tcp_rmem

- tcp_app_win - integer

  应用需保留的最小TCP窗口按照函数 max(window/2^tcp_app_win,mss)取，如果0表示不保留

- tcp_adv_win_scale - integer 

  The socket receive buffer space is shared between the  application  and  kernel.TCP  maintains part of the buffer as the TCP window, this is the size of the receive window advertised to the other end.  The rest of the space is used as  the "application"  buffer.计算公式`bytes/2^|tcp_adv_win_scale|`
  例如值为2，表示只有四分之一作为应用的TCP窗口





-  tcp_abort_on_overflow - boolean
   if listening service is too slow to accept new connections，reset them
   server 全连接队列满了是否丢包 
   - 0：server 直接丢包
   - 1：server 回送 reset 包
- tcp_low_lantency - boolean

  如果开启改参数，高吞吐对比低延迟，TCP 优先选择低延迟。如果关闭改参数，优先选择高吞吐。linux 4.14 后该参数过时失效，但是文件同样存在

- tcp_max_orphans - integer

  the maximum number of orphaned(not attached to any user file handle) TCP sockets allowd in the system. When this number is exceeded, the orphaned connection is reset and a warning is printed.
  将该值改小一般用于防止Dos，但是不推荐，如果调大改值，一个 orphaned socket 会占用大概64KB unswappable 内存

- tcp_dsack  - boolean

  是否支持 TCP Duplicate SACK

- tcp_fack - boolean

  是否支持 TCP Forward Acknowledgement

- tcp_fin_timeout - integer

  the length of time an orphaned (no longer referenced by any application) connection will remain in the FIN_WAIT_2(主动断开的发送了FIN，且收到ACK后的状态) state before it is aborted at the local end. 默认 60 sec 后如果没有收到对端 FIN 包，就会丢弃

- tcp_base_mss - intger

  if MTU probing(pmtud) is enabled, this is the initial MSS used by the connection

- tcp_mtu_probe_floor
