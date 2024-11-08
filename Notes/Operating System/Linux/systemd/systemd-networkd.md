# systemd-networkd

参考：

https://linux.cn/article-6629-1.html

> `man systemd.network`

linux 通常有两个组件来管理网络，NetworkManager和systemd-networkd

我么可以使用下面的命令来查看具体使用的网络管理工具，关闭他们来避免冲突

```
root in ~ λ systemctl --type=service | grep -i network
cloud-init-local.service             loaded active exited  Initial cloud-init job (pre-networking)                  
networkd-dispatcher.service          loaded active running Dispatcher daemon for systemd-networkd                   
systemd-networkd-wait-online.service loaded active exited  Wait for Network to be Configured                        
systemd-networkd.service             loaded active running Network Service                                          
systemd-resolved.service             loaded active running Network Name Resolution      
```

## 配置文件

- `/lib/systemd/network/*.network`
- `/etc/systemd/network/*.network`

`*.network`配置文件还可以拥有`*.network.d/`，该目录下的所有`*.conf`都会对`*.network`生效。

systemd-networkd的配置文件可能会覆盖系统默认生成的配置文件，可以使用软链接指向`/dev/null`来让systemd-networkd的配置文件失效。

```
           # /etc/systemd/network/25-bridge-static.network
           [Match]
           Name=bridge0

           [Network]
           Address=192.168.0.15/24
           Gateway=192.168.0.1
           DNS=192.168.0.1

           # /etc/systemd/network/25-bridge-slave-interface-1.network
           [Match]
           Name=enp2s0

           [Network]
           Bridge=bridge0

           # /etc/systemd/network/25-bridge-slave-interface-2.network
           [Match]
           Name=wlp3s0

           [Network]
           Bridge=bridge0
```

## Match

用于匹配NIC

- MACAddress

- Name

  NIC的名字

## Network

用于设置匹配的NIC的信息

- DHCP yes | no | ipv4 | ipv6

  默认no，是否启用dhcp

- Address

  指定固定的IPv4或IPv6地址

- Gateway

  指定网关

- DNS

  指定DNS服务器，会被systemd-resolved.service读取

## Route

配置路由相关策略，具体查看manual page

## DHCP
