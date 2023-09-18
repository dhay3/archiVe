# Day19 - DTP & VTP

> DTP/VTP 已经不在 CCNA 200-301 考试提纲中了，但是还是有必要了解一下

## DTP

Dynamic Trunk Protocols

DTP is Cisco propreitary protocol(思科专有的协议) that allows Cisco switches to dynamically determine their interface status(access or trunk) without manual configuration

思科所有的 Switch interfaces 默认都开启 DTP，但是出于安全的考虑应该需要使用手动配置接口是否使用 acess 或者是 trunk，同时关闭 DTP

可以使用 `switchport mode dynamic` 来指定使用 DTP

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230531/2023-05-31_15-41.74ex2er3x20w.webp)

> 如果交换机同时支持 802.1Q 和 ISL，需要先使用 `switchport trunk encapsulation negotiate` 声明，才可以使用 `switchport mode dyanmic ...` 或者是 `switchport mode trunk`

如果想要取消端口 DTP negotiation，可以使用 `switchport nonegotiate`；也可以使用 `switchport mode access` 取消 DTP negotiation, 但是使用 `switchport mode trunk` 并不会取消 DTP negotiation 

### auto

A switchport in `dynamic auto` mode will not actively try to form a trunk with other Cisco switches, however it will form a trunk if the switch connected to it is actively trying to form a trunk. It will form a trunk with a switchport in the follwing modes

- `switchport mode trunk`
- `switchport mode dynamic desirable`

例如

SW1 G0/0 配置了 `switchport mode dynamic auto`, SW2 G0/0 配置了 `swithcport mode dynamic auto`，那么 SW1 G0/0 和 SW2 G0/0 就会使用 access mode

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230531/2023-05-31_15-54.3d0nks8ga4cg.webp)

如果 SW1 G0/0 配置了 `switchport mode dynamic auto`, SW2 G0/0 配置了 `switchport mode trunk`， 那么 SW1 G0/0 就会配置成 access mode

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230531/2023-05-31_16-15.46tm7v5azojk.webp)



> On older switches, `switchport mode dynamic desirable` is the ==default administrative mode== 

### desirable

A switchport in `dynamic desirable` mode will actively try to form a trunk with other Cisco switches.It will form a trunk if connected to another switchport in the following modes

- `switchport mode trunk`
- `switchport mode dynamic desirable`
- `switchport mode dynamic auto`

例如

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230531/2023-05-31_15-44.6tcc3rh6nog0.webp)

如果 SW1 G0/0 配置了 `switchport mode dynamic desirable`, SW2 G0/0 配置了 `switchport mode trunk` 那么 G0/0 会自动配置成 trunk mode

这里需要注意的是，如果 SW2 G0/0 配置的是 access port，那么 SW1 G0/0 也会是 access mode，因为是 try to

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230531/2023-05-31_15-51.6mrj3a03i5ts.webp)

> On newer switches, `switchport mode dynamic auto` is the ==default administrative mode== 	

## Relation between Switch ports

可以使用 `show interfaces <interface-name> switchport` 来查 L2 port 使用的 trunk 还是 access

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230531/2023-05-31_16-27.yfz21o3t780.webp)

- Adminstrative mode 是我们实际配置在接口上的状态
- Operational mode 是端口实际使用的状态

如果两个端口都手动配置，但是一个使用 trunk 一个使用 access, 就会有问题

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230531/2023-05-31_16-21.2me0c5oxweo0.webp)

接口 Administrative 之间的关系，以及最后使用 access还是 trunk 可以参考下表

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230531/2023-05-31_16-29.66nrvff5fklc.webp)

例如两个互联的接口，一个使用 dynamic desirable, 一个使用 dynamic auto 那么最后两个端口都会是 trunk mode

需要注意的一点是

*DTP will not form a trunk with a router, PC, etc. The switchport will be in access mode*

即如果你在和 Router on a stick 场景下，将 Switch 和 Router 互联的端口配置成了 dynamic desirable，那么 Switch 并不会宣告他的接口为 trunk，你必须要通过手动配置为 trunk

## VTP

VLAN Trunking Protocol

VTP 同样也是思科独有的协议

VTP allows you to configure VLANs on a central VTP server switch, and other switches(VTP clients) will synchronize their VLAN database to the server

逻辑和 NTP 类似,但是因为可能会导致数据丢失，一般不会使用 VTP

- It is designed for large networks with many VLANs, so that you don’t have to configure each VLAN on every switch

- There are three VTP versions: 1,2 and 3
- There are three VTP modes: server, client and transparent
- Cisco switches operate in VTP server mode by default

### VTP servers

可以对 VLAN 做 add/modify/delete 的操作，会将 VLAN database 存储在 non-volatile RAM 中

Will increase the revision number(直接理解成 VLAN 最新的快照版本) every time a VLAN is added/modified/deleted. Will advertise the lastest version of the VLAN database on **trunk interfaces**(只会通过 trunk port 发送 VTP 同步的消息), and the VTP clients will synchronize their VLAN database to it

VTP server 也是一个 VTP client，意味着 VTP server 会和其他有高于自己 revision number 的 VTP server 做 VLAN database 同步

### VTP clients

不能对 VLAN 做 add/modify/delete 的操作，VLAN database 不存储在 non-volatile RAM 中(VTPv3 存储在 non-volatile RAM)

Will synchronize their VLAN database to the server with the highest revision number in their VTP domain. Will advertise their VLAN database, and forward VTP advertisements to other clients over their turnk ports

### How VTP works

有如下拓扑，互联的端口均配置成了 trunk mode

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230531/2023-05-31_17-18.65w7ct6h0zk0.webp)

因为 Cisco switches 默认是 VTP server，所以端口间，就会互相发送 VTP traffic

可以使用 `show vtp status` 来查看 Switch 上 VTP 的信息

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230531/2023-05-31_17-21.40kbxo95ehkw.webp)

- `VTP version capable` 表示当前 Switch 支持版本 1 to 3
- `VTP version running` 表示当前 Switch 使用的版本
- `VTP domain name` 如果需要使用 VTP synchronize 就需要配置成相同的 VTP domain name，默认为空(NULL)
- `VTP operating mode` 表示当前使用的 mode 是 server，默认 server
- `Maximum VLANs supported locally` 表示当前 VTP 标本支持的 VLAN 数，v1 和 v2 不支持 extended VLAN range(1006-4094)，所以是 1005，v3 支持 extended VLAN range
- `Number of existing VLANs` 表示当前 Switch 上有的 VLAN 数，因为默认有 VLAN1, 1002-1005，所以是 5
- `Configuration Revision` 表示当前 Switch 上 VLAN database 的版本，如果对 VLAN 做了 add/modify/delete 操作，就会加 1

现在使用 `vtp domain cisco` 修改 SW1 的 VTP domain name，然后创建一个 VLAN10

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230531/2023-05-31_17-30.g77vool20nc.webp)

那么我们就可以看见，VTP domain name 和 Configuration Revision 对应的字段值改变了，同时 Number of existing VLANs 值加 1，因为新增加了一个 VLAN10

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230531/2023-05-31_17-32.7c97qexifm2o.webp)

现在使用 `show vtp status` 看一下 SW2 的配置

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230531/2023-05-31_17-34.24l0bak4pkn4.webp)

可以看见竟然和 SW1 上的一样

VTP domain name 因为

*If a switch with no VTP domain(domain NULL) receives a VTP advertisement with a VTP domain name, it will automaatically join that VTP domain*

Number of existing VLANs 和 Configuration Revision 因为

*If a switch receives a VTP advertisement in the same VTP domain with a higher revision number, it will update it’s VLAN database to match*

如果在 SW2 上使用 `show vlan br` 就可以看到同步过来的 VLAN10

同理 SW3,SW4

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230531/2023-05-31_17-39.1hyez58e71k0.webp)

有一种情况，假设互联 SW1-SW2-SW3,通过 VTP 同步了

现在加入 SW4，revision number 50 比其他 3 台 SW 都大，那么其他 3 台 SW 就会通过 VTP 同步，但是 SW4 中 VLAN database 条目实际少于其他 3 台，所以同步后会导致 SW1/SW2/SW3 中和 SW4 中不一样的 VLAN 就会消失，那么对应的 VLAN 就不能通过 trunk port 传输，也就不通了 

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230531/2023-05-31_17-44.6bonhsx4r3b4.webp)

### VTP transparent

> 即只会做 advertise 但是不会做 sync

- Does not participate in the VTP domain(does not sync its VLAN database)
- Maintains its own VLAN database in NVRAM. It can add/modify/delete VLANs, but they won’t be advertised to other switches
- Will forward VTP advertisement that are in the **same** domain as it

对比 How VTP works 中的例子

SW2 如果使用 `vtp mode client`，然后创建一个 VLAN20 就会失败

SW3 使用 `vtp mode transparent`, 并将 SW2 domain name 改成 juniper，那么就不会和 cisco domain 做同步

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230531/2023-05-31_17-55.k7gvdru5m2o.webp)

然后在 SW1 中创建一个 VLAN20

> 这里的 Configuration Revision 逻辑上应该是 2，这里是 4，是为了匹配 LAB 

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230531/2023-05-31_18-01.zi6eh2ck5k0.webp)

因为 SW2 是 Client，所以会和 VTP server 同步

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230531/2023-05-31_18-03.4855yc4bzh0.webp)

因为 SW3 是 transparent mode，所以不会和 VTP server 同步

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230531/2023-05-31_18-04.7dobfo0lrds0.webp)

这里可以看见 Configuration Revision 值为 0，是因为

*Changeing the VTP domain to an unused domain will reset the revision number to 0. Changing the VTP mode to transparent will also reset the revision number to 0*

> 所以如果需要将新的 Switch(VLAN database 条目比当前 VTP server 少) 加入到 VTP 组网中，需要通过下面两条命令确保 revision number 为 0
>
> - `vtp mode transparent`
> - `vtp domain <domain-name>`

因为 SW1 VTP server 有 VLAN20，那么 SW4 会有 VLAN20 吗？ 答案是不会，因为 SW3 是 transparent mode，domain name 不一样，所以不会转发 SW1 的 VTP advertise 

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230531/2023-05-31_18-12.2jla3skonudc.webp)

那么如何让 SW3 转发 SW1 VTP advertise 呢，可以将 SW3 domain name 修改成和 SW1 相同的 domain name 即可

> 这里，增加了 SW1 VLAN numbers 到 11，所以 SW4 Number of existing VLANs 和 Configuration Revision 都增加了。逻辑上应该是 7 和 4

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230531/2023-05-31_18-19.4n0iv7pcojr4.webp)

### VTP version

如果需要修改 SW 使用的 VTP 可以使用 `vtp version <1-3>`

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230531/2023-05-31_18-21.4od9hub7jy80.webp)

修改 version 后，Configuration Revision 同样会做加 1，同时其他的 SW client 会同步，version 同样也会改变

## LAB

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230531/2023-05-31_20-40.52cbe0avsi9s.webp)

### 0x01

```
SW1>en
SW1#conf t
SW1(config)#int g0/1
SW1(config-if)#switchport mode trunk 
SW1(config-if)#switchport nonegotiate 
SW1(config-if)#do sh int g0/1 switchport

SW2>en
SW2#conf t
SW2(config)#int g0/1
SW2(config-if)#switchport mode trunk 
SW2(config-if)#switchport nonegotiate 
SW2(config-if)#do sh int g0/1
SW2(config)#int g0/2
SW2(config-if)#switchport mode trunk 
SW2(config-if)#switchport nonegotiate 
SW2(config-if)#do sh int g0/2

SW3>en
SW3#conf t
SW3(config)#int g0/1
SW3(config-if)#switchport mode trunk 
SW3(config-if)#switchport nonegotiate 
SW3(config-if)#do sh int g0/1 switchport
```

如果 DTP Disable 后使用 `sh int <interface-name> switchport` 对应 Negotiation of Trunking 会显示 off，例如

```
SW1(config-if-range)#do sh int g0/1 sw
Name: Gig0/1
Switchport: Enabled
Administrative Mode: trunk
Operational Mode: trunk
Administrative Trunking Encapsulation: dot1q
Operational Trunking Encapsulation: dot1q
Negotiation of Trunking: Off
Access Mode VLAN: 1 (default)
Trunking Native Mode VLAN: 1 (default)
Voice VLAN: none
Administrative private-vlan host-association: none
Administrative private-vlan mapping: none
Administrative private-vlan trunk native VLAN: none
Administrative private-vlan trunk encapsulation: dot1q
Administrative private-vlan trunk normal VLANs: none
Administrative private-vlan trunk private VLANs: none
Operational private-vlan: none
Trunking VLANs Enabled: All
Pruning VLANs Enabled: 2-1001
Capture Mode Disabled
Capture VLANs Allowed: ALL
Protected: false
```

### 0x02

```
SW1(config)#vtp domain CCNA
SW1(config)#vlan 10
SW1(config-vlan)#vlan 20
SW1(config-vlan)#vlan 30
```

使用 `sh vtp status` 来查看 domain name 是否修改

SW2 and SW3 added VLAN10,20 and 30，use `show vlan br` to check



### 0x03

```
SW2(config)#vtp mode transparent
SW2(config)#vlan 40
```

VLAN40 will not be added to the VLAN database of SW1/SW3

### 0x04

```
SW3(config)#vtp mode client
SW3(config)#vlan 50
VTP VLAN configuration not allowed when device is in CLIENT mode.
```

VLAN50 will not be added on SW3

### 0x05

```
SW1(config)#int range f0/1-2
SW1(config-if-range)#switchport mode access 
SW1(config-if-range)#switchport access vlan 10
SW1(config)#int range f0/3
SW1(config-if)#switchport mode access 
SW1(config-if)#switchport access vlan 20

SW2(config)#int range f0/1-2
SW2(config-if-range)#switchport mode access 
SW2(config-if-range)#switchport access vlan 40

SW3(config)#int f0/1
SW3(config-if)#switchport mode access
SW1(config-if)#switchport access vlan 10
SW3(config)#int f0/4
SW3(config-if)#switchport mode access
SW1(config-if)#switchport access vlan 20
SW2(config)#int range f0/2-3
SW2(config-if-range)#switchport mode access 
SW2(config-if-range)#switchport access vlan 30
```

SW 默认使用 `switchport mode dynamic auto`，所以会开启 DTP, 在使用 `switchport mode access` 后才会 disbable DTP，可以使用 `sh int <interface-name> switchport` 来查看 

**references**

[^jeremy’s IT Lab]:https://www.youtube.com/watch?v=JtQV_0Sjszg&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=36l