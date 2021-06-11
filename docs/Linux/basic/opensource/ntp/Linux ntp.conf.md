# ntp.conf

参考：

https://www.ntppool.org/en/use.html

https://www.thegeekstuff.com/2014/06/linux-ntp-server-client/

ntp(Network Time Protocol)，是用于synchronize all participating computers to within a few millisecods of UTC。

chrony是Linux中ntp实现的方式之一

## ntp.conf

是ntp daemon的配置文件，ntp使用的mode根据command keyword和type of required IP address来决定。每一行由`<command> <address> [options]`组成

### address

ntp可以使用的address分别有remote server(s)，broadcast address of a local interface(b,ABC类网段)，multicast address(m，D类网段)，reference clock address(127.127.x.x)

### configuration commands

如下是几种commads是指定time server or address

- pool

  For  type  s  addresses,  this command mobilizes a persistent client mode association with a number of remote servers.  In this  mode  the  local  clock  can synchronized to the remote server, but the remote server can never  be  synchronized  to the local clock.

- server

  For  type s and r addresses, this command mobilizes a persistent client mode association with the specified remote server or  local radio clock.  In this mode the local clock can synchronized to the remote server, but  the  remote  server  can never  be  synchronized  to  the  local  clock.  This command should **not** be used for type b or m addresses.

- peer

  For type s addresses (only), this command mobilizes a persistent symmetric-active mode association with the specified remote peer.  In this mode the local clock can be  synchronized to  the remote peer or the remote peer can be synchronized to the local clock.  This is useful  in  a  network  of  servers where, depending on various failure scenarios, either the local or remote peer may be the better source  of  time.  This command should NOT be used for type b, m or r addresses.

### options

- burst

  如果server is reachable使用8个数据包替代1个，可以保证timekeeping的质量 

- iburst

  如果server is unreachable使用8个数据包替代1个，用于加快同步的速度当ntpd使用`-q`的参数

- minpoll | maxpoll

  从server上来去信息的间隔，2的次方秒

- noselect 

  server不在被使用

- prefer

  优先选中的server

### monitoring commands

### access conotorl commands

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

### miscellaneous commands

- driftfile

  记录时间偏差的文件，会在ntpd启动是。如果文件不在默认0偏差

- interface [listen | ignore | drop] [all | wildcard | address]

  控制ntpd打开主机上的那些iface，listen表示打开iface，ignore表示ntpd忽略iface，drop表示ntpd会打开iface并将所有的数据包发送到该iface不会经过检验。==可以有多条，最后一条决定该参数==

## 例子

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

