# wireshark filter syntax

参考：https://www.wireshark.org/docs/man-pages/wireshark-filter.html

> [cheat sheet](https://www.comparitech.com/net-admin/wireshark-cheat-sheet/)

## oprator

- eq, ==    Equal ，精确匹配  
- ne, !=    Not Equal    
- gt, >     Greater Than   
-  lt, <     Less Than    
- ge, >=    Greater than or Equal to    l
- e, <=    Less than or Equal to

### logical expression

- and, &&   Logical AND    
- or,  ||   Logical OR    
- not, !    Logical NOT

### search and match operators

- contains 使用模糊查询匹配
- matches,~ 使用regex匹配，忽略大小写

例如：匹配URL

```
http contains "https://www.wireshark.org"
```

### slice operator

和golang一样wireshark也可以使用slice，例如

```
eth.src[0:3] == 00:00:83
llc[0] eq aa
#也可以是数组
frame[4] == 0xff
```

### membership operator

可以使用curly bracket来描述范围，例如

```
tcp.port in {80 443 8080}
```

等价于

```
tcp.port == 80 or tcp.port == 443 or tcp.port == 8080
```

也可以通过这种方式来过滤请求指定的方法

```
http.request.method in {"HEAD" "GET"}
```

### Type conversions

wireshark 支持将字符转为ASCii的方式表达过滤式

```
http.request.method == "GET"
#等价
http.request.method == 47.45.54
```

## Comm filter syntax

| Usage                                   | Filter syntax                                                |
| :-------------------------------------- | :----------------------------------------------------------- |
| Wireshark Filter by IP                  | ip.addr == 10.10.50.1                                        |
| Filter by Destination IP                | ip.dst == 10.10.50.1                                         |
| Filter by Source IP                     | ip.src == 10.10.50.1                                         |
| Filter by IP range                      | ip.addr >= 10.10.50.1 and ip.addr <= 10.10.50.100            |
| Filter by Multiple Ips                  | ip.addr == 10.10.50.1 and ip.addr == 10.10.50.100            |
| Filter out/ Exclude IP address          | !(ip.addr == 10.10.50.1)                                     |
| Filter IP subnet                        | ip.addr == 10.10.50.1/24                                     |
| Filter by multiple specified IP subnets | ip.addr == 10.10.50.1/24 and ip.addr == 10.10.51.1/24        |
| Filter by Protocol                      | dnshttpftpssharptelneticmp                                   |
| Filter by port (TCP)                    | tcp.port == 25                                               |
| Filter by destination port (TCP)        | tcp.dstport == 23                                            |
| Filter by ip address and port           | ip.addr == 10.10.50.1 and Tcp.port == 25                     |
| Filter by URL                           | http.host == “host name”                                     |
| Filter by time stamp                    | frame.time >= “June 02, 2019 18:04:00”                       |
| Filter SYN flag                         | tcp.flags.syn == 1 tcp.flags.syn == 1 and tcp.flags.ack == 0 |
| Wireshark Beacon Filter                 | wlan.fc.type_subtype = 0x08                                  |
| Wireshark broadcast filter              | eth.dst == ff:ff:ff:ff:ff:ff                                 |
| WiresharkMulticast filter               | (eth.dst[0] & 1)                                             |
| Host name filter                        | ip.host = hostname                                           |
| MAC address filter                      | eth.addr == 00:70:f4:23:18:c4                                |
| RST flag filter                         | tcp.flags.reset == 1                                         |











