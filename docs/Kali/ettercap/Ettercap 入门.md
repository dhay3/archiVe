# Ettercap 入门

## syntax

`ettercap [options] [target1] [target2]`

没有目标地址或是源地址之分, 因为通信是双向的

> ==如果接口支持IPv6，target 以这样的形式 MAC/IPs/IPv6/Ports否则以MAC/IPs/Ports的形式。==
>
> /10.0.0.1/表示所有mac地址，所有端口，只要ip地址是10.0.0.1
>
> //表示所有mac地址，所有ip，所有端口
>
> 也可以指定范围
>
> /10.0.0.1-5;10.0.1.33/

## sniffing option

### unified

入侵式的嗅探局域网, 如果需要嗅探网关最好设置为unoffensive mode 。可以施实中间攻击

### briged

嗅探两台主机之间的数据包，需要手动添加IP，不能施实MITM，但是可以和unified一起使用。

## MITM

`-M --mitm <METHOD:ARGS>`

### arp([remote],[oneway])

欺骗目标机的arp缓存，将mac地址解析成攻击机

- remote：如果想要嗅探与当前使用的接口不在同一网段的IP需要指明该参数。可以同时欺骗目标机和网关

- oneway：target1 to target2 单向欺骗。target1的会受影响，但是target2不会。一般用于让网关正常转发

如果没有指定target将嗅探整个局域网

```
ettercap -T -i eth0 -M arp 
```

当目标机使用`arp /d`就会失效

### icmp

将局域网真实的IP的Mac地址做为参数，其他用户会将信息发送给攻击机，攻击机然后将信息传送给网关。需要设置gw mac地址和gw ip

```
ettercap -T -i eth0 -M icmp:00:11:22:33:44:55/10.0.0.1 /// ///
```

### dhcp

将攻击机伪装成DHCP服务器，可以强制目标机接收攻击机的回应。需要指定dhcp pool，dns服务器子网掩码，dns服务器ip。当关闭ettercap时，受害机还是持有被指定IP，只有当租期过期时才会失效。可以不提供dhcp pool这样也就不会应答dhcp offer。无需提供有效的真实IP

```
ettercap -T -i eth0 -M dhcp:192.168.30.100/255.255.255.0/192.168.30.1 ///  ///
```

目标机可以主动使用`ipconig /release && ipconfig /renew`重新生成ip。下图是在攻击中抓取的dhcp packets

<img src="..\..\..\imgs\_Kali\ettercap\Snipaste_2020-09-18_10-54-31.png"/>

- `-o --only-mitm`

  只使用mitm，不抓取流量。可以让wireshark抓取流量。需要开启ip forward

  具体查看https://www.linuxprobe.com/ubuntu-ip-forward.html

- `-B`

  需要双网卡，ettercap会将一个接口的流量转到另外一个。目标机无法有效的直接通过物物理层查看攻击机，但是不能使用mitm。会抓取攻击机的两张网卡对应LAN的所有流量
  
  ```mermaid
  graph TD
  b(proxy interface) -->a(true attacker interface)
  c(target1) & d(target2)-->b
  ```

```
ettercap -T -i eth0 -B eth1
```

## offline sniffing

- `-r, --read <FILE>`

  从pcap file中读取数据包并嗅探

- `-w --write <FILE>`

  将嗅探到数据包写入pcap文件，之后可以使用wireshark分析数据包

## user interfaces option

- `-T, --text`

  使用控制台模式开启ettercap

- `-C --curses`

  使用基本图形化窗口

- `-G --gtk`

  使用GTK2 interface

- `-q --quiet`

  只能在控制台模式下使用，ettercap将不打印出数据包

- `-D --daemonize`

  以守护线程的模式用于后台记录日志

## General options

- `-i --iface <IFACE>`

  指定攻击机的接口

- `-I --iflist`

  列出能使用的接口

- `-A --address <ADRESS>`

  当接口有多个IP时，通过该参数指定IP

- `-z --silent`

  不扫描当前局域网，需要手动指定target IP

- `-p --nopromisc`

  ettercap默认会嗅探指定网段的所有流量，通过该参数可以只嗅探被引导到攻击的流量

- `-u --unoffensive`

  ettercap通过自己转发流量，当想开启多个实例时，需要设置该参数

- `-k --save-hosts <FILENAME>`

  如果不想让ettercap每次启动时都引起广播风暴，可以使用该参数将这次扫描的host写入到指定文件。==需要修改`/etc/ettercap/etter.conf`，将ec_uid 和 ec_gid 置为用户id==。

  参考：https://github.com/Ettercap/ettercap/issues/752

- `-j --load-hosts <FILENAME>`

  加载通过`-k`参数生成的host文件

  ```
  ettercap -T -j filename -M arp
  ```

- `-P --plugin <PLUGIN>`

  加载指定插件，可以通过`ettercap -P list`查看所有插件

## logging options

- `-L --log <LOGFILE>`

  将所有的数据包保存到指定文件，生成eci(保存数据包)和ecp(保存信息)文件

- `-l --log-info <LOGFILE>`

  只保留账号和密码

- `-m --log-msg <LOGFILE>`

  保存所有被ettercap打印的信息
