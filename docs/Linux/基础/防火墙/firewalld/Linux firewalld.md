# Linux firewalld 

参考：

https://www.linuxprobe.com/chapter-08.html

https://blog.51cto.com/andyxu/2137046

https://wangchujiang.com/linux-command/c/firewall-cmd.html

https://github.com/taw00/howto/blob/master/howto-configure-firewalld.md

RHEL 7系统中使用firewalld取代了传统的iptables(删除了iptables service)，==但兼容iptables(任然可以使用iptables命令，但是不推荐)==

> 如果主机上存在iptables service，请将其置为disable，否则会影响firewalld。

## 概念

`firewall-cmd`是firewalld防火墙配置管理工具的CLI（命令行界面）版本。

所有的service配置文件以xml格式命名

- Zone

  简单来说就是firewalld预先准备了几套防火墙策略集合，用户可以根据生产场景的不同而选择合适的策略集合，从而实现防火墙策略之间的快速切换。

  ==firewalld中常用的区域名称以及策略规则，推荐采用blcok使用白名单机制==

| 区域     | 默认规则策略                                                 |
| -------- | ------------------------------------------------------------ |
| trusted  | Allow all network connections                                |
| home     | For trusted home network connections                         |
| internal | For internal network, restrict incoming connections          |
| work     | For trusted work areas                                       |
| public   | Public areas, do not trust other computers                   |
| external | For computers with masquerading enabled, protecting a local network |
| dmz      | For computers publicly accessible with restricted access.    |
| block    | Deny all incoming connections, with ICMP host prohibited messages issued. |
| drop     | Deny all incoming connections, outgoing ones are accepted.   |

- Services

  一组定义好的规则，包括端口，协议，src，dst。用于当前zone

- Direct interface

  规则使用的接口

## firewall-cmd

> firewall-cmd添加或修改的规则在restart或reload后就会失效，==使用`--permanent`让规则永久生效==。`firewall-cmd --reload`让==永久配置==立即生效，无需重启服务。==如果未使用`--zone`表示默认使用default zone==

### zone

- `firewall-cmd --get-zones`

  显示当前所有可用的zones

  ```
  [root@cyberpelican ~]# firewall-cmd --get-zones
  block dmz drop external home internal public trusted work
  ```

- `firewall-cmd --get-active-zones`

  显示当前==默认使用的zone==和网卡的名字

  ```
  [root@cyberpelican ~]# firewall-cmd --get-active-zones 
  public
    interfaces: ens33
  ```

- ==`firewall-cmd --list-all`==

  查看当前使用zone的具体设置

  ```
  [root@cyberpelican ~]# firewall-cmd --list-all
  public (active)
    target: default
    icmp-block-inversion: no
    interfaces: ens33
    sources: 
    services: dhcpv6-client ssh
    ports: 
    protocols: 
    masquerade: no
    forward-ports: 
    source-ports: 
    icmp-blocks: 
    rich rules: 
  
  ```

### service

- `firewall-cmd --get-services`

  显示当前所有可用的services

- `firewall-cmd --list-services`

  显示当开放的services

  ```
  [root@cyberpelican ~]# firewall-cmd --list-services 
  dhcpv6-client https ssh
  ```

- `firewall-cmd --info-service=<service>`

  显示service的具体内容

  ```
  [root@cyberpelican ~]# firewall-cmd --info-service=ssh
  ssh
    ports: 22/tcp
    protocols: 
    source-ports: 
    modules: 
    destination: 
  ```

### target

与iptables中的target相同，包括default，ACCEPT，DROP，REJECT

- `firewall-cmd --permanent --get-target`

  获取当前zone的target

  ```
  [root@cyberpelican ~]# firewall-cmd --permanent --get-target
  default
  ```

- `firewall-cmd --permanent --set-target=<target>`

  设置zone的target，带有permanent参数，需要使用`firewall-cmd --relaod`重新加载模块

  ```
  [root@cyberpelican ~]# firewall-cmd --permanent --set-target=DROP
  success
  [root@cyberpelican ~]# firewall-cmd --list-all
  public (active)
    target: default
    icmp-block-inversion: no
    interfaces: ens33
    sources: 
    services: dhcpv6-client ssh
    ports: 
    protocols: 
    masquerade: no
    forward-ports: 
    source-ports: 
    icmp-blocks: host-unknown
    rich rules: 
  ```

  

### 规则

允许规则为`--add-`，拒绝规则为`--remove-`，查询为`--query-`

- `firewall-cmd --set-default-zones=<zone>`

  设置默认使用的zone，无需使用`--permanent`同样永久生效

  ```
  [root@cyberpelican ~]# firewall-cmd --set-default-zone=trusted 
  success
  [root@cyberpelican ~]# firewall-cmd --get-default-zone 
  trusted
  ```

- `firewall-cmd --add-service=<services>`

  指定service允许通过防火墙

  ```
  [root@cyberpelican ~]# firewall-cmd --add-service=https
  success
  ```

- `firewall-cmd --add-port=<port/proto>`

  指定端口/协议流量允许通过防火墙，可以指定范围

  ```
  [root@cyberpelican ~]# firewall-cmd --add-port=122/tcp
  success
  ```

- `firewall-cmd --add-protocal=<proto>`

  指定协议流量允许通过防火墙

  ```
  [root@cyberpelican ~]# firewall-cmd --add-protocol=icmp
  success
  ```

- `firewall-cmd --add-source=<IP>`

  将当前主机的traffic转到指定的IP，可以使用tcpdump来校验

  ```
  [root@cyberpelican ~]# firewall-cmd --add-source=192.168.80.200
  success
  [root@cyberpelican ~]# firefox &
  [1] 2987
  [root@cyberpelican ~]# firewall-cmd --remove-source=192.168.80.200
  success
  [root@cyberpelican ~]# firewall-cmd --add-source=192.168.80.0/24
  success
  ```

- `firewall-cmd --add-interface=<interface>`

  将指定iface绑定到档期那使用的zone

  ```
  [root@chz ~]# firewall-cmd --add-interface==ens33
  success
  [root@chz ~]# firewall-cmd --list-all
  public (active)
    target: default
    icmp-block-inversion: no
    interfaces: =ens33 ens33
    sources: 
    services: dhcpv6-client ssh
    ports: 
    protocols: 
    masquerade: no
    forward-ports: 
    source-ports: 
    icmp-blocks: 
    rich rules: 
  
  ```

- `firewall-cmd --add-forward-port=<port>:proto=<proto>:toprot=<forward-port>:toaddr=<forward-addr>`

  将指定的端口/协议的流量转发到指定IP/端口

  ```
  [root@linuxprobe ~]# firewall-cmd --permanent --zone=public --add-forward-port=port=888:proto=tcp:toport=22:toaddr=192.168.10.10
  success
  ```

- `firewall-cmd --add-icmp-block=<icmp-type>`

  > 如果拦截了icmp，攻击者同样能通过嗅探获取目标机的IP，将target设置为DROP

  firewalld拦截指定的icmp-type，通过`firewall-cmd --get-icmptypes`来查看所有的icmp-type

  ```
  [root@cyberpelican ~]# firewall-cmd --add-icmp-block=echo-request 
  success
  [root@cyberpelican ~]# firewall-cmd --list-all
  public (active)
    target: default
    icmp-block-inversion: no
    interfaces: ens33
    sources: 
    services: dhcpv6-client ssh
    ports: 
    protocols: 
    masquerade: no
    forward-ports: 
    source-ports: 
    icmp-blocks: echo-request
    rich rules: 
  ```

> 添加zone和service后需要使用`firewall-cmd --reload`重新加载firewalld

### 自定义zone

```
[root@cyberpelican ~]# firewall-cmd --permanent --new-zone=cyber
success
[root@cyberpelican ~]# firewall-cmd --permanent  --zone=cyber --set-target=DROP 
success
[root@cyberpelican ~]# firewall-cmd --reload 
success
[root@cyberpelican ~]# firewall-cmd --zone=cyber --list-all
cyber
  target: DROP
  icmp-block-inversion: no
  interfaces: 
  sources: 
  services: 
  ports: 
  protocols: 
  masquerade: no
  forward-ports: 
  source-ports: 
  icmp-blocks: 
  rich rules: 
```

### 自定义service

```
[root@cyberpelican ~]# firewall-cmd --permanent --new-service=cus_sshd
success
[root@cyberpelican ~]# firewall-cmd --permanent --service=cus_sshd --set-description='cus_sshd 122 port'
success
[root@cyberpelican ~]# firewall-cmd --service=cus_sshd --get-description  --permanent 
cus_sshd 122 port
[root@cyberpelican ~]# firewall-cmd --list-all
cyber (active)
  target: DROP
  icmp-block-inversion: no
  interfaces: ens33
  sources: 
  services: cus_sshd
  ports: 
  protocols: 
  masquerade: no
  forward-ports: 
  source-ports: 
  icmp-blocks: 
  rich rules: 
```



