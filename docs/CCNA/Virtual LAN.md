# Virtual LAN

## 概述

LAN 指的是 local area network，而VLAN指的是virtual LAN。

VLAN可以把一个实体LAN分割成多个虚拟的LAN，分割出来的VLAN各自独立，VLAN与VLAN之间默认无法互通。

VLAN使用的协议有两种，ISL(Inter-Switch Link)和IEEE 802.1q。ISL是Cisco专有的，而802.1q是公开的

## 为什么需要VLAN

需要VLAN的原因有两个，网络效能和安全

由于网络中会经常出现Broadcast封包，Broadcast 不仅需要通过Switch 传送到LAN里面的每一个Host，而且每一个收到Broadcast的Host都要花Computing Power去处理Broadcast，这对整个网络的效能都大打折扣。VLAN可以把实体LAN分割，==一个VLAN的Broadcast Traffic不会传到另一个VLAN==，每一个VLAN就变成是一个独立的Broadcast Domain，提升网络效能。==即划分广播域==

另一方面，Broadcast在网络上散播可能造成安保问题，只要下载packet capture，即可打开Broadcast 封包，窥探其他Host的信息（IP and MAC address），利用VLAN可以限制Broadcast Traffic只在信任的网络中散播

## 配置VLAN

我们使用下面一个例子，R1到R4表示四台Host，他们的e0/0分别设定成192.168.1.1/24 至 192.168.1.4/24

![Snipaste_2021-08-16_16-42-17](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20211130/Snipaste_2021-08-16_16-42-17.42zv5s5nt740.png)

```
hostname R1
 !
 interface Ethernet0/0
  ip address 192.168.1.1 255.255.255.0
```

```
hostname R2
 !
 interface Ethernet0/0
  ip address 192.168.1.2 255.255.255.0
```

```
hostname R3
 !
 interface Ethernet0/0
  ip address 192.168.1.3 255.255.255.0
```

```
hostname R4
 !
 interface Ethernet0/0
  ip address 192.168.1.4 255.255.255.0
```

当然R1和R2，R3，R4是互通的

```
R1#ping 192.168.1.2
 
 Type escape sequence to abort.
 Sending 5, 100-byte ICMP Echos to 192.168.1.2, timeout is 2 seconds:
 !!!!!
 Success rate is 100 percent (5/5), round-trip min/avg/max = 16/24/44 ms
 R1#ping 192.168.1.3
 
 Type escape sequence to abort.
 Sending 5, 100-byte ICMP Echos to 192.168.1.3, timeout is 2 seconds:
 !!!!!
 Success rate is 100 percent (5/5), round-trip min/avg/max = 16/20/24 ms
 R1#ping 192.168.1.4
 
 Type escape sequence to abort.
 Sending 5, 100-byte ICMP Echos to 192.168.1.4, timeout is 2 seconds:
 !!!!!
 Success rate is 100 percent (5/5), round-trip min/avg/max = 20/24/44 ms
```

现在把R1和R2 放进VLAN10，把R3和R4放进VLAN20。换句话，我们要把SW1的e0/0和 e0/1 设定VLAN 10，然后把e0/2 和 e0/3设定成VLAN20

![Snipaste_2021-08-11_20-17-25](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20211130/Snipaste_2021-08-11_20-17-25.2y91paxd4o00.png)



现在在SW1新增VLAN10和VLAN20，10和20换做VLAN ID，name指的是VLAN的名称，可有可无

```
SW1#conf t
 Enter configuration commands, one per line.  End with CNTL/Z.
 SW1(config)#vlan 10
 SW1(config-vlan)#name Yellow
 SW1(config-vlan)#vlan 20
 SW1(config-vlan)#name Green
 SW1(config-vlan)#exit
```

完成后用`show vlan`确认一下。VLAN 1,1002,1003,1004,1005都是default的，暂且不用理会，看到有10和20就可以了

```
SW1#show vlan
 
 VLAN Name                             Status    Ports
 ---- -------------------------------- --------- -------------------------------
 1    default                          active    Et0/0, Et0/1, Et0/2, Et0/3
                                                 Et1/0, Et1/1, Et1/2, Et1/3
                                                 Et2/0, Et2/1, Et2/2, Et2/3
                                                 Et3/0, Et3/1, Et3/2, Et3/3
 10   Yellow                           active
 20   Green                            active
 1002 fddi-default                     act/unsup
 1003 token-ring-default               act/unsup
 1004 fddinet-default                  act/unsup
 1005 trnet-default                    act/unsup
 
 VLAN Type  SAID       MTU   Parent RingNo BridgeNo Stp  BrdgMode Trans1 Trans2
 ---- ----- ---------- ----- ------ ------ -------- ---- -------- ------ ------
 1    enet  100001     1500  -      -      -        -    -        0      0
 10   enet  100010     1500  -      -      -        -    -        0      0
 20   enet  100020     1500  -      -      -        -    -        0      0
 1002 fddi  101002     1500  -      -      -        -    -        0      0
 1003 tr    101003     1500  -      -      -        -    -        0      0
 1004 fdnet 101004     1500  -      -      -        ieee -        0      0
 1005 trnet 101005     1500  -      -      -        ibm  -        0      0
 
 Primary Secondary Type              Ports
 ------- --------- ----------------- ------------------------------------------
```

下一步把e0/0和e0/1放进VLAN 10

```
SW1(config)#int ethernet 0/0
 SW1(config-if)#switchport access vlan 10
 SW1(config-if)#int ethernet 0/1
 SW1(config-if)#switchport access vlan 10
```

switchport 将port的模式从access替换成vlan

VLAN 20的做法也一样

```
SW1(config)#int ethernet 0/2
 SW1(config-if)#switchport access vlan 20
 SW1(config-if)#int ethernet 0/3
 SW1(config-if)#switchport access vlan 20
```

再用`show vlan`来确认

```
SW1#show vlan
 
 VLAN Name                             Status    Ports
 ---- -------------------------------- --------- -------------------------------
 1    default                          active    Et0/0, Et0/1, Et0/2, Et0/3
                                                 Et1/0, Et1/1, Et1/2, Et1/3
                                                 Et2/0, Et2/1, Et2/2, Et2/3
                                                 Et3/0, Et3/1, Et3/2, Et3/3
 10   Yellow                           active 	 Et0/0, Et0/1
 20   Green                            active    Et0/2, Et0/3
 1002 fddi-default                     act/unsup
 1003 token-ring-default               act/unsup
 1004 fddinet-default                  act/unsup
 1005 trnet-default                    act/unsup
 
 VLAN Type  SAID       MTU   Parent RingNo BridgeNo Stp  BrdgMode Trans1 Trans2
 ---- ----- ---------- ----- ------ ------ -------- ---- -------- ------ ------
 1    enet  100001     1500  -      -      -        -    -        0      0
 10   enet  100010     1500  -      -      -        -    -        0      0
 20   enet  100020     1500  -      -      -        -    -        0      0
 1002 fddi  101002     1500  -      -      -        -    -        0      0
 1003 tr    101003     1500  -      -      -        -    -        0      0
 1004 fdnet 101004     1500  -      -      -        ieee -        0      0
 1005 trnet 101005     1500  -      -      -        ibm  -        0      0
 
 Primary Secondary Type              Ports
 ------- --------- ----------------- ------------------------------------------
```

这时候再试试R1 ping 就只能ping 通 R2 了，相反，R3 只能 ping 通R4。原因很简单，因为不同VLAN之间，packet不能互通

```
R1#ping 192.168.1.2
 
 Type escape sequence to abort.
 Sending 5, 100-byte ICMP Echos to 192.168.1.2, timeout is 2 seconds:
 !!!!!
 Success rate is 100 percent (5/5), round-trip min/avg/max = 20/266/776 ms
 R1#ping 192.168.1.3
 
 Type escape sequence to abort.
 Sending 5, 100-byte ICMP Echos to 192.168.1.3, timeout is 2 seconds:
 .....
 Success rate is 0 percent (0/5)
 R1#ping 192.168.1.4
 
 Type escape sequence to abort.
 Sending 5, 100-byte ICMP Echos to 192.168.1.4, timeout is 2 seconds:
 .....
 Success rate is 0 percent (0/5)
```

```
R3#ping 192.168.1.1
 
 Type escape sequence to abort.
 Sending 5, 100-byte ICMP Echos to 192.168.1.1, timeout is 2 seconds:
 .....
 Success rate is 0 percent (0/5)
 R3#ping 192.168.1.2
 
 Type escape sequence to abort.
 Sending 5, 100-byte ICMP Echos to 192.168.1.2, timeout is 2 seconds:
 .....
 Success rate is 0 percent (0/5)
 R3#ping 192.168.1.4
 
 Type escape sequence to abort.
 Sending 5, 100-byte ICMP Echos to 192.168.1.4, timeout is 2 seconds:
 !!!!!
 Success rate is 80 percent (4/5), round-trip min/avg/max = 16/19/20 ms
```

在现实应用中，我们习惯把不通subnet ip 分配给不同的VLAN，例如VLAN 10用192.168.10.0/24，VLAN 20 用 192.168.20.0/24，很少会把同一个subnet设在多一个VLAN之中。

## Trunk Link

如果VLAN需要跨域多只switch，会出现什么问题？如果SW1把一些packet丢给SW2，SW2怎么分辨这些Packet来自哪一个VLAN？如果它不知道这些Packet来自哪个VLAN，它自然不知道应该把这个packet发到哪一个VLAN。解决方法很简单，只要在packet写上VLAN号码(VLAN ID)才把packet送走，其他switch就可以凭这个VLAN ID知道packet VLAN，这就是802.1 q VLAN Tag

VLAN Tag 是 Switch 在收到Packet时为他加上一个标识，目的是让packet在packet 在网络中传输是，所经过的switch 都可以查看这个packet是属于哪一个VLAN，从而把packet送到真正需要接收这个VLAN的port

switch之间需要传送这些VLAN Tag，我们需要把switch 与 switch 之间的Link设定成Trunk，因为只有Trunk Link 才可以容纳不同的VLAN。设定Trunk Link的方法有很多。==在这只使用Static 设定==，即强制使其成为Trunk Link

![Snipaste_2021-08-30_19-28-02](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20211130/Snipaste_2021-08-30_19-28-02.46up6yql0yi0.png)

设定方法不复杂，首先需要确认两只switch都已经建立需要处理的VLAN，留意VLAN ID必须相同，==VLAN名称可以不同==

```
SW1#show vlan
 
 VLAN Name                             Status    Ports
 ---- -------------------------------- --------- -------------------------------
 1    default                          active    Et0/3, Et1/0, Et1/1, Et1/2
                                                 Et1/3, Et2/0, Et2/1, Et2/2
                                                 Et2/3, Et3/0, Et3/1, Et3/2
                                                 Et3/3
                                                 
 10   Yellow                           active    Et0/0
 20   Green                            active    Et0/1
 1002 fddi-default                     act/unsup
 1003 token-ring-default               act/unsup
 1004 fddinet-default                  act/unsup
 1005 trnet-default                    act/unsup

<--Output Omitted-->
```

```
SW2#show vlan
 
 VLAN Name                             Status    Ports
 ---- -------------------------------- --------- -------------------------------
 1    default                          active    Et0/3, Et1/0, Et1/1, Et1/2
                                                 Et1/3, Et2/0, Et2/1, Et2/2
                                                 Et2/3, Et3/0, Et3/1, Et3/2
                                                 Et3/3
                                                 
 10   Yellow                           active    Et0/0
 20   Green                            active    Et0/1
 1002 fddi-default                     act/unsup
 1003 token-ring-default               act/unsup
 1004 fddinet-default                  act/unsup
 1005 trnet-default                    act/unsup

<--Output Omitted-->
```

然后把连着两只switch的e0/2 interface设定为802.1q Trunk

```
SW1(config)#int ethernet 0/2
 SW1(config-if)#switchport trunk encapsulation dot1q
 SW1(config-if)#switchport mode trunk
```

```
SW2(config)#int ethernet 0/2
 SW2(config-if)#switchport trunk encapsulation dot1q
 SW2(config-if)#switchport mode trunk
```

用`show interface trunk`可以确认那一条Link是Trunk Link

```
SW1#show interfaces trunk
 
 Port                Mode         Encapsulation  Status        Native vlan
 Et0/2               on           802.1q         trunking      1
 
 Port                Vlans allowed on trunk
 Et0/2               1-4094
 
 Port                Vlans allowed and active in management domain
 Et0/2               1,10,20
 
 Port                Vlans in spanning tree forwarding state and not pruned
 Et0/2               1,10,20
```

```
SW2#show interfaces trunk
 
 Port                Mode         Encapsulation  Status        Native vlan
 Et0/2               on           802.1q         trunking      1
 
 Port                Vlans allowed on trunk
 Et0/2               1-4094
 
 Port                Vlans allowed and active in management domain
 Et0/2               1,10,20
 
 Port                Vlans in spanning tree forwarding state and not pruned
 Et0/2               1,10,20
```

于是R1可以ping通相同VLAN的R2，却无法与其他VLAN沟通

```
R1#ping 192.168.1.2
 
 Type escape sequence to abort.
 Sending 5, 100-byte ICMP Echos to 192.168.1.2, timeout is 2 seconds:
 !!!!!
 Success rate is 100 percent (5/5), round-trip min/avg/max = 100/150/212 ms
 R1#ping 192.168.1.3
 
 Type escape sequence to abort.
 Sending 5, 100-byte ICMP Echos to 192.168.1.3, timeout is 2 seconds:
 .....
 Success rate is 0 percent (0/5)
 R1#ping 192.168.1.4
 
 Type escape sequence to abort.
 Sending 5, 100-byte ICMP Echos to 192.168.1.4, timeout is 2 seconds:
 .....
 Success rate is 0 percent (0/5)
```

如果在Trunk Link 进行packet Capture 的会清楚看到VLAN ID

![Snipaste_2021-08-11_20-17-25](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20211130/Snipaste_2021-08-11_20-17-25.9qqtepd6lis.png)

### Allowed VLAN

使用Trunk Link时，可以设定只允许特定VLAN放行，预设 1 - 4094 全部可以通过，如果想要更改设定，可以在Trunk Interface 使用 `switchport trunk allowed vlan <vlan id>`

```
SW1(config-if)#switchport trunk allowed vlan 10,20,30
 SW1(config-if)#end
 SW1#show interfaces trunk
 
 Port                Mode         Encapsulation  Status        Native vlan
 Et0/2               on           802.1q         trunking      1
 
 Port                Vlans allowed on trunk
 Et0/2               10,20,30
 
 Port                Vlans allowed and active in management domain
 Et0/2               10,20
 
 Port                Vlans in spanning tree forwarding state and not pruned
 Et0/2               10,20
```

更改设定后使用`show interfaces trunk`查看结果：

- vlans allowed on trunk

  表示这条Trunk allow什么VLAN通过，如上图，方通了VLAN10,20,30

- vlans allowed and active in management domain

  在这只switch 上就只有VLAN 10 和 20，虽然允许了VLAN 10,20 ,30，但是实际上VLAN30是通不过的

- vlans in spanning tree forwarding state and not pruned

  真正可以通过的VLAN，即VLAN 10和 VLAN 20

还有add，all，except，none，remove等操作，具体可以查看man page

## VLAN 1

VLAN 1是一个预设的VLAN，所有Cisco Switch皆有VLAN 1，所有port也都预设在VLAN 1。

VLAN 1和其他VLAN一样可以传送data之外，还可以传送Control Plane Traffic，例如 VTP，CDP，PAgP

基于安全考虑，VLAN 1应避免给一般HOST使用

## Native VLAN

Native VLAN的意思是switch把这个VLAN的packet送上Trunk Link时，是不会放VLAN Tag。VLAN  1 就是一个预设的Native VLAN

但是和刚才的说的不一样，但是可以仔细想想，如果2 - 4096的VLAN都有VLAN ID，只要Trunk Link 两边的Switch 都协议没有VLAN ID的packet 就是VLAN 1，那么VLAN 1就算没有VLAN ID 也可以被区别出来。

所以当VLAN 1的packet 通过Trunk Link，用Packet capture软件也不会看到VLAN ID 1，只会看见一个没有VLAN ID的封包（即没有Tag）

Native VLAN 的 ID是可以设定的，`switchport trunk native vlan <vlan_id>`，==trunk link 两边 interface 的 native VLAN 必须相同，否则会造成Native VLAN mismatch 的问题==

```
SW1(config-if)#switchport trunk native vlan 100
 SW1(config-if)#exit
 SW1#show int trunk
 
 Port                Mode         Encapsulation  Status     Native vlan
 Et0/1               on           802.1q         trunking      100   
```

另外Native VLAN 需要 和 分配给Port 的 VLAN 不同，为了避免Double tagging attack。

除了更改Native VLAN外，较为简单的方法就是不要吧VLAN 1 分配给port使用

## VLAN internal usage

在multilayer Switch 使用系统会偷偷得使用一些VLAN做内部用途

先用`show vlan internal usage`确认此刻没有VLAN被做internal usage

```
SW1#show vlan internal usage
 
 VLAN Usage
 ---- --------------------
```

把port 设定成Layer 3 port。可以看到VLAN 1006被使用了。为什么是1006？因为系统预设会有1006开始递增使用VLAN ID 作为Internal Usage

```
SW1(config)#int ethernet 0/1
 SW1(config-if)#no switchport
 SW1(config-if)#end
 *Dec  3 11:15:24.478: %LINK-3-UPDOWN: Interface Ethernet0/1, changed state to up
 SW1#show vlan internal usage
 
 VLAN Usage
 ---- --------------------
 1006 Ethernet0/1
```

如果这时候选要创建VLAN 1006 就会报错

```
SW1#conf t
 Enter configuration commands, one per line.  End with CNTL/Z.
 SW1(config)#vlan 1006
 VLAN id: 1006 is an internal vlan id - cannot use it to create a VTP VLAN.
```

可以使用`vlan interanl allocation policy descending`将递增改为递减，下次就会从4094开始

```
SW1(config)#vlan internal allocation policy descending
 SW1(config)#int range ethernet 1/1 - 2
 SW1(config-if-range)#no switchport
 SW1(config-if-range)#end
 *Dec  3 11:27:17.794: %LINK-3-UPDOWN: Interface Ethernet1/1, changed state to up
 *Dec  3 11:27:17.798: %LINK-3-UPDOWN: Interface Ethernet1/2, changed state to up
 SW1#show vlan internal usage
 
 VLAN Usage
 ---- --------------------
 1006 Ethernet0/1
 4093 Ethernet1/2
 4094 Ethernet1/1
```



## show vlan

https://www.cisco.com/c/en/us/td/docs/ios/lanswitch/command/reference/lsw_book/lsw_s2.html

用于查看VLAN

```
VLAN Name                             Status    Ports
---- -------------------------------- --------- -------------------------------
1    default                          active    Gi1/5, Gi1/6, Gi1/7, Gi1/8, Gi1/9, Gi1/10
                                                Gi1/11, Gi1/12, Gi1/13, Gi1/14, Gi1/15
                                                Gi1/16, Gi1/17, Gi1/18, Gi1/19, Gi1/20
                                                Gi1/21, Gi1/22, Gi1/23, Gi1/24, Gi1/25
                                                Gi1/26, Gi1/27, Gi1/28, Gi1/29, Gi1/30
                                                Gi1/31, Gi1/32, Gi1/33, Gi1/34, Gi1/35
                                                Gi1/36, Gi1/37, Gi1/38, Gi1/39, Gi1/40
                                                Gi1/41, Gi1/42, Gi1/43, Gi1/44
1002 fddi-default                     act/unsup
1003 token-ring-default               act/unsup
1004 fddinet-default                  act/unsup
1005 trnet-default                    act/unsup

VLAN Type  SAID       MTU   Parent RingNo BridgeNo Stp  BrdgMode Trans1 Trans2
---- ----- ---------- ----- ------ ------ -------- ---- -------- ------ ------
1    enet  100001     1500  -      -      -        -    -        0      0
1002 fddi  101002     1500  -      -      -        -    -        0      0
1003 tr    101003     1500  -      -      -        -    -        0      0
1004 fdnet 101004     1500  -      -      -        ieee -        0      0
1005 trnet 101005     1500  -      -      -        ibm  -        0      0

Remote SPAN VLANs
------------------------------------------------------------------------------

Primary Secondary Type              Ports
------- --------- ----------------- ------------------------------------------

```

- 1 - VLAN

  VLAN number

- default - NAME

  VLAN的名字，如果没有配置

- active - status

  VLAN的状态

  可以是如下几个值

  1. active or suspend

  2. act/unsup

     https://community.cisco.com/t5/switching/what-is-the-difference-btw-active-and-act-unsup/td-p/904152

     VLAN is correct but is unsupported for that switch

  3. sus/unsup

- Gi1/5 … - Ports

  ports that belong to the VLAN

- enet - Type

  Media type of the VLAN

  1. enet = Ethernet
  2. fddi = FDDI
  3. tnet = Token Ring

- 100001 - SAID

  security association ID value for the VLAN

- 1500 - MTU

  maximum transmission unit size for the VLAN

- Parent

  Parent VLAN

- RingNo

- BridgeNo

- Stp

  Spanning Tree Protocol type that is used on the VLAN

- BrdgMode
- Trans1/2