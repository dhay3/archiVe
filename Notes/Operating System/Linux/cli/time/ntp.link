

# Linux ntp

##  ntpd

> 有的OS systemd以`ntpd.service`命名或`ntp.service` , 有的OS可能也不使用ntpd，而使用`systemd-timesyncd`或 `chrony`,

ntpd是linux上实现ntp4的工具(==ntpd.service是一个一次性的service==)，默认读取`ntp.conf`。如果系统的时间存在较大的偏差会导致一些协议无法正常工作例如Vmess。

大多数OS和hardware都会和**time-of-year(TOY) chip**协作，在系统断电后将时间信息写在TOY chip中。当OS引导完毕后TOY chip就会初始化OS的时间，当OS向ntp server同步完成后会修改TOY chip上的信息。当系统时差超过1000s就会被认为出现严重的错误ntpd自动退出，需要手动修改时间。

### options

- -g

  ntpd在系统时差大于1000s时可以设置一次，如果之后的时差还是大于1000ms就会退出

- -N

  ntpd以高优先级运行，cpu会优先处理

- -q

  set the time and quit. ntpd同步完时间后直接退出

## ntp.conf

### Digest

参考：

https://www.ntppool.org/en/use.html

https://www.thegeekstuff.com/2014/06/linux-ntp-server-client/

ntp(Network Time Protocol)，是用于synchronize all participating computers to within a few millisecods of UTC。

chrony是Linux中ntp实现的方式之一，systemd-timesynd 也是其中之一

### configuration

是ntp daemon的配置文件，ntp使用的mode根据command keyword和type of required IP address来决定。每一行由`<command> <address> [options]`组成

#### address

ntp可以使用的address分别有remote server(s)，broadcast address of a local interface(b,ABC类网段)，multicast address(m，D类网段)，reference clock address(127.127.x.x)

#### configuration commands

如下是几种commads是指定time server or address

- pool

  For  type  s  addresses,  this command mobilizes a persistent client mode association with a number of remote servers.  In this  mode  the  local  clock  can synchronized to the remote server, but the remote server can never  be  synchronized  to the local clock.

- server

  For  type s and r addresses, this command mobilizes a persistent client mode association with the specified remote server or  local radio clock.  In this mode the local clock can synchronized to the remote server, but  the  remote  server  can never  be  synchronized  to  the  local  clock.  This command should **not** be used for type b or m addresses.

- peer

  For type s addresses (only), this command mobilizes a persistent symmetric-active mode association with the specified remote peer.  In this mode the local clock can be  synchronized to  the remote peer or the remote peer can be synchronized to the local clock.  This is useful  in  a  network  of  servers where, depending on various failure scenarios, either the local or remote peer may be the better source  of  time.  This command should NOT be used for type b, m or r addresses.

#### options

- burst

  如果server is reachable使用8个数据包替代1个，可以保证timekeeping的质量，但是会增加网络的开销

- iburst

  如果server is unreachable使用8个数据包替代1个，用于安全减少同步的时间间隔

- minpoll | maxpoll

  从server上来去信息的间隔，2的次方秒

- noselect 

  server不在被使用

- prefer

  优先选中的server

#### monitoring commands

#### access conotorl commands

- restrict address flag

  对指定IP访问time server做限制，address的值可以是IP或default表示`0.0.0.0/0`。==如果没有flag表示没有限制==，flag通常可以分为两种restrict time service and restrict information queries。可以使用如下几种flag
  
  ignore：拒绝所有数据包，包括ntpq和ntpdc
  
  noserver：拒绝所有的数据包，但不包括ntpq和ntpdc的
  
  kod：如果发生访问冲突是限制数据包发送的速率
  
  limited：限制数据包如果违背discard command
  
  noepeer：拒绝短连接
  
  nomodify：拒绝ntpd和ntpdc修改time server的请求
  
  noquery：拒绝ntpq和ntpdc的请求
  
  nopeer：拒绝未授权的数据包
  
  notrust：拒绝服务除非数据包加密并授权
  
  notrap：拒绝向主机提供6中信息捕捉服务

#### miscellaneous commands

- driftfile

  记录时间偏差的文件，会在ntpd启动是。如果文件不在默认0偏差

- interface [listen | ignore | drop] [all | wildcard | address]

  控制ntpd打开主机上的那些iface，listen表示打开iface，ignore表示ntpd忽略iface，drop表示ntpd会打开iface并将所有的数据包发送到该iface不会经过检验。==可以有多条，最后一条决定该参数==

### 例子

```
➜  /etc cat /etc/ntp.conf

#  NTP pool
server 0.arch.pool.ntp.org iburst
server 1.arch.pool.ntp.org iburst
server 2.arch.pool.ntp.org iburst
server 3.arch.pool.ntp.org iburst
server 0.cn.pool.ntp.org iburst prefer

#restrictions
restrict -4 default kod limited nomodify nopeer noquery notrap
restrict -6 default kod limited nomodify nopeer noquery notrap
restrict 127.0.0.1
restrict ::1

# Location of drift file
driftfile /var/lib/ntp/ntp.drift
```

## ntpq

参考：

https://detailed.wordpress.com/2017/10/22/understanding-ntpq-output/



syntax：`ntpq [options] [host...]`

ntpq是NTP的查询工具有interactive和cli两种方式，会将请求发送到ntp服务器。默认使用localhost，如果没有指定host。==NTP是一个UDP协议，所以连接不是可靠的==

### options

- n | --numeric

  对host不做解析，以数字的格式输出

- -p | --peers

  当前使用的ntp server的信息

  ```
  cpl in /var/lib/ntp λ ntpq -pn 127.0.0.1                 
       remote           refid      st t when poll reach   delay   offset  jitter
  ==============================================================================
  *84.16.73.33     .GPS.            1 u  978 1024    7  222.338  +10.820   1.841
  -116.203.151.74  131.188.3.222    2 u  948 1024    7  262.759  +21.756  13.279
  +139.199.215.251 100.122.36.4     2 u  959 1024    7   33.316   +6.734   9.906
  +5.79.108.34     130.133.1.10     2 u  947 1024    7  231.657  +28.945  12.221
  ```

  1. remote：同步服务器
  2. refid：同步服务器参考的服务器
  3. st：sratum，0表示根服务器，1表示直接参考跟服务器的服务器，以此类推16表示无同步的服务器
  4. t：服务器的种类
  5. when：最后一个数据包接受的时间，如果为`-`表示从未收到
  6. poll：向服务器请求的间隔
  7. reach：到达位移寄存器
  8. delay：往返时延
  9. offset：服务器和主机的偏移量
  10. jitter：估计错误的偏移量

  ```
  * Synchronized to this peer
  # Almost synchronized to this peer
  + Peer selected for possible synchronization
  – Peer is a candidate for selection
  ~ Peer is statically configured 
  ```

  ### error

  参考：

  https://unix.stackexchange.com/questions/345778/why-does-ntpq-pn-report-connection-refused

  需要先开启`ntp.service`或`ntpd.service`，ntpq才能使用

## ntpdate

参考：

https://linux.die.net/man/8/ntpdate

syntax：`ntpdate [options] <ntp-server>`

ntpdate用于从ntp server上同步时间信息，必须以root身份在localhost上运行

`ntpdate`可以手动设置也可以通过脚本在boot阶段设置，==如果ntpd在运行ntpdate会拒绝设置时间==

```
root in /home/ubuntu λ ntpdate cn.pool.ntp.org 
14 Jun 18:48:12 ntpdate[28540]: the NTP socket is in use, exiting
root in /home/ubuntu λ systemctl stop ntp.service 
root in /home/ubuntu λ ntpdate cn.pool.ntp.org   
14 Jun 18:48:32 ntpdate[28592]: adjust time server 84.16.73.33 offset -0.000849 sec
```
