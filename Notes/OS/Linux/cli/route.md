# Linux route

## 网关

https://man.linuxde.net/route

https://blog.csdn.net/vevenlcf/article/details/48026965

### route

> 我们可以使用`traceroute`来查看路由的具体路径，等同于windows中的`tracert`，显示`***`表示响应超时。（==如果目标主机使用了防火墙设置了响应的规则，禁止ICMP　echo 或是禁用UDP端口，traceroute就不会生效==）可以使用`-T` 参数使用TCP SYN来发包
>
> ```
> [root@chz ~]# traceroute -T baidu.com
> traceroute to baidu.com (39.156.69.79), 30 hops max, 60 byte packets
>  1  gateway (192.168.80.2)  0.128 ms  0.080 ms  0.138 ms
>  2  39.156.69.79 (39.156.69.79)  35.046 ms  29.824 ms  33.684 ms
> ```

该命令已经过时，使用`ip route`替换。在命令行中使用route添加路由，不会永久生效，==当网卡重启或机器重启，就失效==。可以在`/etc/rc.local`中添加route命令来保证该路由设置永久生效

1. 不带任何参数显示当前的路由表，使用`-n`参数不会反向解析IP

   ```
   [root@chz html]# route -n
   Kernel IP routing table
   Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
   0.0.0.0         192.168.80.2    0.0.0.0         UG    100    0        0 ens33
   192.168.80.0    0.0.0.0         255.255.255.0   U     100    0        0 ens33
   ```

   在目标地址为`192.168.80.0`网段内不需要网关通信，其他IP使用`192.168.80.2`做为网关通信（==采用优先匹配==）

   |             |                                                          |
   | ----------- | -------------------------------------------------------- |
   | Destination | 目标网络或目标主机的IP，其他网络使用`default`或`0.0.0.0` |
   | Gateway     | 网关地址，使用`*`或`0.0.0.0`，表示不需要路由(同net-id)   |
   | Genmask     | 掩码，如果是默认路由为`0.0.0.0`                          |
   | Iface       | 数据包发送给那一张网卡进行转发，默认使用序号为0的网卡    |

   Flag

   - U：路由正在被使用
   - H：目标(Destination)是一台主机
   - G：使用网关
   - M：目标是网络，需要转发

   本地主机与Destinnation通信，通过本地主机Iface将数据包发送给Gateway

   ### 路由类型

   - 主机路由

     主机路由是路由选择表中==指向单个IP地址或主机名的路由记录==。主机路由的Flags字段为H。例如，在下面的示例中，本地主机通过IP地址192.168.1.1的路由器到达IP地址为10.0.0.10的主机。

     ```
     Destination    Gateway       Genmask Flags     Metric    Ref    Use    Iface
     10.0.0.10     192.168.1.1    255.255.255.255   UH       0    0      0    eth0
     ```

   - 网络路由

     网络路由是代表主机可以到达的网络。网络路由的Flags字段为M。例如，在下面的示例中，本地主机将发送到网络192.19.12(可以与Internet沟通的IP)的数据包转发到IP地址为192.168.1.1的路由器。

     ```
     Destination    Gateway       Genmask Flags    Metric    Ref     Use    Iface
     192.19.12     192.168.1.1    255.255.255.0      UN      0       0     0    eth0
     ```

   - 默认路由

     当主机不能再路由表中查找到目标主机的IP地址或网络时，数据包就被发送到默认路由上。默认路由的Flags字段为G

     ```
     Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
     default         gateway         0.0.0.0         UG    100    0        0 ens33
     ```

2. 添加路由

   pattern：`route add <flag> IP netmask <mask> dev <iface>`，dev可以省略

   - ==主机路由(-host)不需要指定子网掩码，因为route会自动根据IP识别是A，B，C类网==

     ```
     route add -host 192.168.80.200  ens33 #局域网内通信不需要gateway，类型为U
     route add -host 192.168.70.0 gw 192.168.80.1 ens33 #手动指定网关，类型为UGH
     ```

     example 1：

     使用外网IP，如果不指定网关无法访问该主机

     ```
     route add -host 39.156.69.79 ens33
     ```

     example 2：

     如果指定的网关，不是手动设置网关(这里我设置为`192.168.80.2`)的无法访问该主机

     ```
     route add -host 39.156.69.79 gw 192.168.80.5  ens33
     ```

   - ==网络路由(-net)需要指定子网掩码==

     ```
     route add -net 192.168.70.0 netmask 255.255.255.0 gw 192.168.80.2 ens33
     ```

   - 添加默认路由

     ```
     route add default gw 192.168.80.4
     ```

     

删除路由

pattern：`route del <flag> IP netmask <mask> dev <iface>`，dev可以省略

```
[root@chz ~]# route del -net  169.254.0.0/16 dev ens33
[root@chz ~]# route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
192.168.10.0    0.0.0.0         255.255.255.0   U     0      0        0 ens33
```

