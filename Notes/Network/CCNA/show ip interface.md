# show ip interface

reference:

https://www.cisco.com/E-Learning/bulk/public/tac/cim/cib/using_cisco_ios_software/cmdrefs/show_ip_interface.htm#:~:text=To%20get%20a%20detailed%20listing,the%20standard%20show%20interface%20output).

用于展示端口的信息，例如

```
Example
Router# show ip interface
Ethernet0 is administratively down, line protocol is down
  Internet address is 10.10.46.10, subnet mask is 255.0.0.0
  Broadcast address is 255.255.255.255
  Address determined by setup command
  MTU is 1500 bytes
  Helper address is not set
  Directed broadcast forwarding is enabled
  Multicast groups joined: 224.0.0.1 224.0.0.2
  Outgoing access list is not set
  Inbound access list is not set
  Proxy ARP is enabled
  Security level is default
  Split horizon is enabled
  ICMP redirects are always sent
  ICMP unreachables are always sent
  ICMP mask replies are never sent
  IP fast switching is enabled
  IP fast switching on the same interface is disabled
  IP SSE switching is disabled
  Router Discovery is disabled
  IP accounting is disabled
  TCP/IP header compression is disabled
  Probe proxy name replies are disabled
  Gateway Discovery is disabled
```

- `Ethernet0 is administratively down, line protocol is down`

  表示端口的状态，应用在端口上的配置

- `  Internet address is 10.10.46.10, subnet mask is 255.0.0.0
    Broadcast address is 255.255.255.255`

  表示端口配置的IP，subnet-mask，广播地址

- `MTU is 1500 bytes`

  这个端口上能处理的最大包

- `Helper address is not set`

  是否配置了helper address，用于处理destination不在路由中的包

- `Directed broadcast forwarding is enabled`

  配置这个接口上的subnet会进行广播

- `Multicast groups joined: 224.0.0.1 224.0.0.2`

  加入的组播组

- `  Outgoing access list is not set
    Inbound access list is not set`

  端口上配置的acl

- `Proxy ARP is enabled`

  会对ARP广播有回包（即使目标地址不是改端口的MAC，但是知道目标地址的MAC）

- `Split horizon is enabled`

    不会将从接口学到的路由转发回该接口

- `  ICMP redirects are always sent
  ICMP unreachables are always sent
  ICMP mask replies are never sent`

  告诉iface该怎么处理ICMP回包

- `  IP fast switching is enabled
    IP fast switching on the same interface is disabled
    IP SSE switching is disabled`

  表示router该怎么交换数据包

- `Router Discovery is disabled`

  router discovery protocol 对这个端口不适用

- `IP accounting is disabled`

  IP accounting 不对这个端口使用

-   `TCP/IP header compression is disabled`
