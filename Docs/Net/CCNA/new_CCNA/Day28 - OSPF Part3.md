

# Day28 - OSPF Part3

## Loopback Interface

在说 OSPF 之前，先简短的介绍一下 loopback interface

- A loopback interface is a virtual interface in the router
- It is always up/up(unless you manually shut it down)
- It is not dependent on physical interface
- So, it provides a consistent IP address that can be use to reach/identify the router

在一些场景下，router 之间需要相互通信，例如 OSPF neighbors 互相发送 hello message 等

例如下图

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-14_17-48.50t9weazcry8.webp)

假设 R4 需要发包到 R1，Dst IP 为 10.0.13.1，如果这时 R1 G1/0 因为端口的原因导致链路 down 了，因为 via R3 的路由是 preferred route，那么 R4 就不能发包到 R1 G1/0

如果 R1 使用了 loopback，R4 使用 R1 loopback 地址作为目的 IP，并且 R4 有到 R1 loopback 的路由(这里需要使用 dynamic routing protocol，如果使用了 static route 指定了一定要走 R3 那么还是有问题)，即使 G1/0 端口有问题， R1 同样可以通过 via R2 的方式到 R1

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-14_17-53.648j9k630dq8.webp)

但是如果 R1 G0/0 也有问题，那么 R4 就不能正常到 R1 了，即使 R1 loopback interface 加入到了 dynamic protocol network，因为 loopback 只是虚拟的接口，实际还是通过 router 的物理接口走的

> 可以抽象的把 loopback interface 对应的 IP address 看成 physical interfaces 额外的一个 IP address，到 loopback interface 的报文，会经过任意一个 physical interface

## OSPF Network Types

OSPF networks types 指的是 OSPF neighbors 之间的连接(connection)类型，主要有 3 种 OSFP network types

1. Broadcast

   enable by default on Ethernet and FDDI(Fiber Distributed Data Interfaces) interfaces

2. Point-to-point

   enabled by default on PPP(Point-to-Point Protocol) and HDLC(High-Level Data Link Control) interfaces

3. Non-broadcast

   enabled by default on Frame Relay and X.25 interfaces

> 在 CCNA 中只需要关注 Broadcast 和 Point-to-point

### Broadcast

- Enable on Ethernet and FDDI interfaces by default

  > 现在的网络拓扑中一般都使用 Ethernet，所以一般默认是 broadcast type

- Routers dynamically discover neighbors by sending/listening for OSPF Hello messages using multicast address 224.0.0.5

- A **DR**(designated  router) and **BDR**(backup designated router) must be elected on each subnet(only DR if there are no OSPF neighbors, ie R1’s G1/0 interface)

  > 在每一个 subnet 中 Router 的角色并不一定
  >
  > 例如 R1 G1/0 和 R2 G1/0 互联， R1 G2/0 和 R3 G1/0 互联
  >
  > 在 R1 G1/0 - R2 G1/0 subnet 中判断 R1 是 DR，R2 是 BDR
  >
  > 在 R1 G2/0 - R3 G1/0 subnet 中判断 R1 是 BDR，R3 是 DR
  >
  > R1 既是 DR 又是 BDR，两者不冲突
  >
  > 
  >
  > 在每一个通过 router 组成的 subnet 中，都需要有一个 DR 和 一个 BDR；如果是通 router 和 end host 组成，只需要有一个 DR，并不需要 BDR (显然这个是 OSPF 中的名词，需要组合 neighbors，end host 和 router 并不能组成 OSPF neighbors)
  >
  > 如果 R3 在一个 subnet 中没有和其他 router 互联，那么在这个 subnet 中 R3 就是 DR

  例如下图

  ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-15_13-46.63ma5uj93vgg.webp)

  以 R1 为例子，在 10.0.1.0/24 subnet 中是 DR；而在 192.1681.0/30 subnet 中，假设 R2 是 DR，那么 R1 是 BDR

- Routers which aren’t the DR or BDR become a **DROther**

#### DR/BDR election

选择 DR/BDR 的逻辑如下 

1. 先比较 OSPF interface priority，高的作为 BR，低的作为 BDR

   > *The default OSPF interface priority is 1 on all interfaces*

2. 如果 OSPF interface priority 一样，就会比较 Router ID，Router ID 高的作为 BR，第二高的作为 BDR

   > 每个 subnet 中，都有一个 DR 和 BDR

3. 其他剩下的 Router 都是 DROther

下图中，假设 R1 router ID 1.1.1.1, R2 router ID 2.2.2.2, etc

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-15_13-53.6klxv4kihoxs.webp)

那么在 R1 - R2 subnet 中 R2 是 DR；R2/3/4/5 subnet 中 R5 是 DR，R4 router ID 第 2 高所以 R4 是 BDR，剩下的 R2/3 都是 DROther

如果在 R5 中使用 `show ip ospf interface g0/0`

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-15_13-58.25eucxxkawgw.webp)

- `State DR`

  表示 R5 在当前 subnet 中是 DR

- `Priority 1`

  表示当前接口的 pritority

- `Designated Router`

  表示在当前 subnet 中的 DR 是谁，这里是 R5

- `Backup Designated Router`

  表示在当前 subnet 中的 BDR 是谁，这里是 R4

对比一下 R2

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-15_14-01.1en2smor4tr4.webp)

- `State DROTHER`

  表示 R2 在当前 subnet 中是 DROTHER

- `Designated Router`

  表示在当前 subnet 中的 DR 是谁，这里是 R5

- `Backup Designated Router`

  表示在当前 subnet 中的 BDR 是谁，这里是 R4

所有 interface ospf priority 默认都为 1，如果想要手动修改可以使用 `ip ospf priority <priority>` 命令

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-15_14-05.2o4j840gixa8.webp)

> *if you set the OSPF interface priority 0, the router CANNOT be the DR/DBR for the subnet*

如果使用了上面的命令，现在看一下 R2 `show ip ospf int g0/0`

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-15_14-08.6flcvkvrrjls.webp)

逻辑上 R2 g0/0 的 priority 现在应该是最大的，在 R2/3/4/5 中，R2 应该是 DR；但是这里任然显示 R2 是 DROther，即使 priority 已经显示为 255

*This is because of the DR/BDR election is ‘non-preemptive’(非抢占式的).Once the DR/BDR are selected they will keep their role until OSPF is reset, the interface fails/is shut down, etc*

如果这时在 R5 使用 `clear ip ospf process` 让 OSPF reset，并使用 `show ip ospf neighbor`

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-15_14-13.3urvsp3ri0g0.webp)

拓扑类似下图

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-15_14-35.15htayrjkio0.webp)

会发现只有 R2(2.2.2.2) 和 R4(2.2.2.2) OSPF state 是 Full

同时 R2 并没有因为 g0/0 priority 值最高变成 DR，而是 DBR；R4 从 BDR 变为 DR

> *When the DR goes down, the BDR becomes the new DR(even if it does have the highest priority). Then an election is held for the next BDR*

此外 R3 任然保持是 DROther，而 R5 从 DR 变为 DROther(DR 会直接变成 DROther) 并且 R3(3.3.3.3) OSPF state 并不是 Full 而是 2-way

> *DROthers(R3 and R5 in this subnet)will only move to the FULL state with the DR and BDR. The neighbor state with other DROthers will be 2-way*
>
> DROther 只有和 DR/DBR 之间才会形成 Full state
>
> 同时也说明了
>
> ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-15_14-45.13uu0yr55d7k.webp)
>
> 1. Down -> 2-way State 是建立 OSPF neighbors
> 2. Exstart -> Full 是交换 LSAs，只会在 DR/BDR 中

即在 broadcast network type 中，只有 DR 和 BDR 之间才会形成 Full state OSPF neighbors

*Therefore, routers only exchange LSAs with DR and BDR. DROther will not exchange LSAs with each other*

> 简单的理解就是只有 DR 和 DBR 会发送 LSAs

*All routers will still have the same LSDB, but this reduces the amount of LSAs flooding the network*

假设不区分 DR/BDR/DROther，每台 Router 都会发送 LSAs，流量拓扑类似下图

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-15_14-39.6z69ewuxufwg.webp)

如果区分了 DR/BDR/DROther，只有 DR 和 BDR 会相发送 LSAs，那么拓扑就会如下

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-15_14-40.54f8wqmi66m8.webp)

大大的减少了网络中的流量，但是每台 Router 一样可以收到 LSAs

> 需要注意的一点是，DR/BDR 之间使用 multicast，地址为 224.0.0.6 和 OSPF 发送 hello message 的 multicast address 224.0.0.5 不同

除了上面使用的 `show ip ospf neighbor` 或者是 `show ip ospf interface <interface-id>` 来查看 router 是否是 DR/BDR/DROther 外，还可以使用 `show ip ospf interface brief` 来查看

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-15_14-51.352m7n2acg3k.webp)

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-15_14-35.15htayrjkio0.webp)

`Nbrs F/C` 其中的 F 表示当前接口 Full state neighbors 的数量，这里能和 R3 DROther 形成 full state 的只有 R4 DR 和 R2 BDR；其中的 C 表示当前接口所有状态 neighbors 的数量，这里有 R2/4/5

上面显示的 `Nbrs F/C` 部分，在 `show ip ospf interface <interface-id>` 中出现下图红框中

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-15_14-55.32gdxhr8leio.webp)

#### Point-to-Point

- Enabled on serial interfaces using the PPP or HDLC encapsulations by default

- Routers dynamically discover neighbors by sending/listening for OSPF hello messages using multicast address 224.0.0.5

- A DR and BDR are not elected

  > 在 broadcast type 中因为一个 subnet 中必然有 DR 和 BDR，在点对点的链路中，既然都是两台 router 互联，那么谁是 DR 或者是 DBR 就没有意义，反正都要发送 LSAs 

例如下图，R1 和 R2 使用了 Serial connection 构成 Point-to-Point type

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-15_15-46.14r88fw21duo.webp)

如果这时在 R2 上使用 `show ip ospf neighbor`

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230615/2023-06-15_16-16.1qluvj88fr6o.webp)

可以看到 R1 并不是 DR/DBR/DROTher，因为 R1 - R2 subnet 是 Point-to-Point type

##### Serial interfaces

Serial interfaces 和普通的 Ethernet interfaces(RJ45) 不同，下图就是 serial interfaces

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-15_15-54.42sryj16qww0.webp)

长得非常像 VGA 口

常见的网络架构中已经不再使用 serial interfaces

假设现在 R1 和 R2 互联的接口都是 serial interface

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-15_15-56.4jh0h4fa9lhc.webp)

在 Serial connection 中，两端的角色不同

- One side of a serial connection functions as DCE(Data Communications Equipment). The DCE Side needs to specify the clock rate(speed) of the connection

  通过 `clock rate <number>` 的方式来调整

  > 只有 DCE 才可以使用这个命令，DCE 是由端口自己决定的，可以使用 `show controller <interface-id>` 来查看

  ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230613/2023-06-15_16-01.5yli32z720lc.webp)

  > Ethernet connection 使用 `speed` 命令来修改端口的速率，而 Serial connection 使用 `clock rate <number>` 来修改

- The other side functions as DTE(Data Terminal Equipment)

使用 `show interface` 命令看一下对应 serial interface 的信息

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230615/2023-06-15_16-04.1u7msnn2ijb4.png)

> The default encapsulation on a serial interface is HDLC

现在将默认的 encapsulation 从 HDLC 改为 PPP

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230615/2023-06-15_16-07.3q30mnvffykg.webp)

> 需要注意的是，如果需要将链路的改为 PPP，需要两端都使用 `encapsulation ppp`，否则 link state 就会显示为 down

R1/2 互联的端口配置如下

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230615/2023-06-15_16-09.6ewn5msydojk.webp)

可以使用 `show controllers <interface-id>` 来查看本端是 DCE 还是 DTE

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230615/2023-06-15_16-13.2t2gi7ezhtmo.webp)

#### Broadcast VS Point-to-Point

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230615/2023-06-15_16-24.1s9q20o3s18g.webp)

> 因为 Non-broadcast network type 不在 CCNA 的范围内这里做简单的介绍
>
> Non-broadcast network type default timers = Hello 30, Dead 120

## Configure the OSPF Network Type

除了使用对应的接口来改变 OSPF network type，还是通过 `ip ospf network <type>` 来手动设置 subnet network type

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230615/2023-06-15_16-20.48siem2w8m0w.webp)

那么为啥要手动来配置 subnet network type 呢？

在 Ethernet connection 中两台 Router 互联，都使用 OSPF，因为在 subnet 中必须要有一个 DR 和 BDR，所以两台 Router 都会发 SLAs，和 Router 的角色并不是强关联的。所以两台 Router 之间使用 point-to-point network 更加合理

> 并不是所有的 network type 都可以手动修改的，例如 serial interface 就不支持使用 broadcast network type

## OSPF Neighbor Requirements

Router 相互需要组成 OSPF neighbors 需要满足下面的条件

1. Area number must match
2. Interfaces must be in the same subnet
3. OSPF process must not be shutdown
4. OSPF Router IDs must be unique
5. Hello and Dead timers must match
6. Authentication settings must match
7. IP MTU settings must match
8. OSPF Network Type must match

### Area number must match

互联的接口 area number 需要一致

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230615/2023-06-15_16-27.291w0pvdyym.webp)

上图中 R1 G0/0 所在的 network 和 R2 G0/0 所在的 network area number 不一致，所以不能构成 OSPF neighbors，所以使用 `show ip ospf ne` 显示为空

修改成相同的 area number 后，就可以正常建立 OSPF neighbors

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230615/2023-06-15_16-30.16nd6406id1c.webp)

### Interfaces must be in the same subnet

互联的端口必须在一个子网

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230615/2023-06-15_16-32.4h4xaql3udq8.webp) 

上图中 R1 G0/0 192.168.1.1/30 通过 `network 192.168.1.0 0.0.0.3 area 0` 宣告出去，R2 G0/0 192.168.2.2/30 通过 `network 192.168.2.0 0.0.0.3 area 0` 宣告出去，因为两台 router 接口 IP 不在同一个 subnet 所以不能构成 OSPF neighbors

> 这里其实没意义，因为 IP 都不一样，即使链路 up/up，也不通，LSAs 肯定不能通过这条链路发出去

将两台 Router 的端口 IP 都配置在一个 subnet,后就可以正常构成 OSPF neighbors

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230615/2023-06-15_16-38.3cp8e46rpps0.webp)

> 这里即使 network 宣告的不一样，但是包含自己对应互联的端口 IP，也一样能构成 OSPF neighbors
>
> 例如  R1 `network 192.168.0.0 0.0.255.255 area 0`, R2 `network 192.168.1.0 0.0.0.3 area 0`

### OSPF process must not be shutdown

router OSPF 进程必须是正常的

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230615/2023-06-15_16-42.5ov9kmkapzi8.webp)

在 router 上 OSPF 进程是可以通过 `shutdown` 来关闭的，如果关闭后就不能正常构成 OSPF neighbors，可以使用 `no shutdown` 来重新启动 OSPF 进程

> 在使用 `router ospf <process-id>` 后，会默认启动 OSPF 进程，一般不会掉，只有在手动使用了 `shutdown` 后，才可能出现问题

### OSPF Router IDs must be unique

在 OSPF 中 Router ID 必须唯一

下图中并没有为 R1/2 设置 loopback，R1 router ID 使用 R0/0 192.168.1.1，R2 router ID 使用 G0/0 192.168.1.2

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230615/2023-06-15_16-52.y0ezsaapg8g.webp)

这是通过 `router-id 192.168.1.1` 来修改 R2 router id，并使用 `clear ip ospf process` 来让 router id 生效

这是可以看到 OSPF neighbor 表是空的，并不会构成 OSPF neighbors，同时可以看到 `detected duplicate router-id` 的信息

可以通过 `no router-id` 来删除手动设置的 router id，使用默认的规则来生成 router id

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230615/2023-06-15_16-58.o1pmxbpnlxc.webp)

这里可以看到删除了手动设置的 router id 后，使用默认的规则生成的 router 192.168.1.2 就可以和 R1 正常建立 OSPF neighbors

> 注意这里没有使用 `clear ip ospf process` 来从重置 OSPF 进程
>
> 因为在这时 R2 没有任何的 OSPF neighbor，所以完全不用关心 OSPF router id 冲突，所以设备也就不要求重置 OSPF process

### Hello and Dead timers must match

hello 和 dead timers 值必须要一样

> 虽然在 Broadcast network type 和 Point-to-Point network type 中 hello 和 dead timer 默认都为 10 和 40，但是也可以手动修改

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230615/2023-06-15_17-06.fpeaer8lylk.webp)

这里通过 `ip ospf hello-interval 5` 将 R2 的 hello timer 修改成 5，通过 `ip ospf dead-interval 20` 将 R2 的 dead timer 修改成 20；因为 R1 没有修改过，R1 hello timer 为 10，R1 dead timer 为 40；因为都不匹配，所以不能构成 OSPF neighbors

> 这里即使 dead timer 相同，也不能构成 OSPF neighbors，因为 hello timer 不一样
>
> 即只要 hello timer 或者是 dead timer 任意一个不一样就不能正常构成 OSPF neighbors

### Authentication settings must match

在 OSPF 中还可以配置 authentication password，只有 password 相同，才可以构成 OSPF neighbors

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230615/2023-06-15_17-12.6w5933glifg.webp)

- `ip osfp authentication-key jeremy`

  声明 OSPF 使用的 authentication password

- `ip ospf authentication`

  仅仅配置 authentication password 并不会直接生效，还需要使用改命令声明开启并使用 OSPF authentication

这里 R1 没有配置 authentication password，因为不匹配，所以就不能构成 OSPF neighbors，OSPF neighbor table 就为空

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230615/2023-06-15_17-16.2ixcqqkv0bsw.webp)

这里取消了 R2 OSPF authentication，因为都为空，所以可以构成 OSPF neighbors

### IP MTU settings must match

如果 router IP MTU 如果配置的不一样，OSPF 可能会有问题

> *Even if the IP MTU settings mismtach.These can become OSPF neighbors, but OSPF doesn’t operate properly*

可以通过 `ip mtu <number>` 来修改 router MTU，这里改成 1400 (Ethernet 默认 MTU 为 1500)

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230615/2023-06-15_17-21.50b24o2reyyo.webp)

这里即使修改了 MTU，OSPF neighor 状态还是会显示为 full state，但是其实 R1 和 R2 之间已经不能通过 OSPF 交换 LSAs 了。在 dead timer 后，状态会变为 down，后从 OSPF neighbor table 中删除对应的条目

如果使用 `clear ip ospf process` 重置了 OSPF，就会发现 R2 即使在 OSPF neighbor table 中有 R1 的信息，但是状态不能变为 full state，也就不能正常交换 LSAs

然后会一直重复输出下面的内容

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230615/2023-06-15_17-27.1dof250m75pc.webp)

可以使用 `no ip mtu` 来删除手动配置的 MTU，使用默认的 1500 MTU

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230615/2023-06-15_17-33.2qam6bq5eem8.webp)

在删除手动配置的 MTU 后，就可以看到 neighbors 变为 full state

### OSPF Network Type must match

互联端口 OSPF Network Type 必须要一样

这里配置了 R2 loopback0 2.2.2.2，并将 loopback0 加入到 OSPF 中，同时将 G0/0 从 broadcast 改为 point-to-point type；而 R1 G0/0 还是使用默认的 broacast type

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230615/2023-06-15_17-35.5p225o6tbv5s.webp)

这里可以可以看到，虽然 R2 显示 R1 还是 full state，但是因为使用了 point-to-point，这里 router 的角色变成了 `-` 表示空。在 R2 上看不到什么异常，上 R1 看一下

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230615/2023-06-15_17-41.1tdp7ex8f7c0.webp)

这里可以发现，R1 显示 R2 任然是 full state，逻辑上 OSPF 应该是正常的。但是仔细观察，可以发现 R1 并没有 2.2.2.2 的路由，即使 R2 宣告了

> 如果 OSPF neighbors 显示正常，但是没有路由，可以检查一下双方的 OSPF network type 是否一致

可以使用 `no ip ospf network` 来取消手动配置的 OSPF network type

## OSPF LSA Types

有路由拓扑

在之前的例子，R4 使用 `default-information originate` 宣告默认路由

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-16_13-43.601e57la860w.webp)

- The OSPF LSDBs is made up of LSAs

- There are 11 types of LSA, but there are only 3 you should be aware of for the CCNA

  Type1 (Router LSA)

  Type2 (Network LSA)

  Type5 (AS External LSA)

我们可以通过 `show  ip ospf database` 来查看 LSDB 中的 LSA

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-16_14-02.6cw695imud1c.webp)

> 在 OSPF network 中每台机器的 LSDB 都相同，所以无需区分那一台

- `Router Link state` 对应 Router LSA

  Router LSA 是 OSPF 中所有 router 都会发送的，所以这边包含 R1/2/3/4/5

- `Net Link state` 对应 Net LSA

  只有在 subnet 中的 DR 会发送，所以只有 R4 的

- `AS External link state` 对应 AS External LSA

  只有 R4 是 ASBR，所以只会有 R4 的

### Router LSA

- Every OSPF router generates this type of LSA
- It identifies the router using its router ID
- It also lists networks attached to the router’s OSPF-activated interfaces

### Network LSA

- Generated by the DR of each ‘multi-access’ network(ie. the broadcast network type)
- Lists the routers which are attached to the multi-access network

### AS External LSA

- Generated by ASBRs to describe routes to destinations outside of the AS(OSPF domain)

## LAB

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230616/2023-06-16_15-29.4yvnsl7v8nls.webp)

### 0x01

#### Configure the serial connection between R1 and R2(clock rate of 128000)

```
R1(config)#int serial 0/0/0
R1(config-if)#ip add 192.168.12.1 255.255.255.252
R1(config-if)#no shutdown

R2(config)#int serial 0/0/0
R2(config-if)#ip add 192.168.12.2 255.255.255.252
R2(config-if)#no shutdown
```

这里还需要先查看一下端口是否是 DCE，只有 DCE 端口才可以设置 clock rate

```
R1(config-if)#do sh controllers s0/0/0
Interface Serial0/0/0
Hardware is PowerQUICC MPC860
DCE V.35, clock rate 2000000
```

这里可以发现 R1 s0/0/0 是 DCE，所以使用下面的命令来为 s0/0/0 设备 clock rate

```
R1(config-if)#int s0/0/0
R1(config-if)#clock rate 128000
```

可以使用 `show int s0/0/0` 来检查

```
R1(config-if)#do sh  int s0/0/0
Serial0/0/0 is up, line protocol is up (connected)
  Hardware is HD64570
  Internet address is 192.168.12.1/30
  MTU 1500 bytes, BW 1544 Kbit, DLY 20000 usec,
     reliability 255/255, txload 1/255, rxload 1/255
  Encapsulation HDLC, loopback not set, keepalive set (10 sec)
```

Configure OSPF on R1 and R2

```
R1(config)#router ospf 1
R1(config-router)#network 192.168.12.0 0.0.0.3 area 0

R2(config)#router ospf 1
R2(config-router)#network 192.168.12.0 0.0.0.3 area 0
```

使用 `show ip ospf ne` 来检查

```
R2(config-router)#do sh ip ospf ne
Neighbor ID     Pri   State           Dead Time   Address         Interface
192.168.245.2     1   FULL/DR         00:00:33    192.168.245.2   GigabitEthernet0/0
192.168.12.1      0   FULL/  -        00:00:31    192.168.12.1    Serial0/0/0
```

或者 `sh ip ospf int s0/0/0` 来检查

```
R2(config-router)#do sh ip ospf int s0/0/0

Serial0/0/0 is up, line protocol is up
  Internet address is 192.168.12.2/30, Area 0
  Process ID 1, Router ID 192.168.245.1, Network Type POINT-TO-POINT, Cost: 64
```

### 0x02

#### Only R3 has a route to 10.0.2.0/24. Why?

```
R3#sh ip ospf ne
Neighbor ID     Pri   State           Dead Time   Address         Interface
192.168.245.2     0   FULL/  -        00:00:34    192.168.34.2    GigabitEthernet0/1
```

这里可以看见 R3 OSPF neighbor table 中只有 192.168.245.2(R4)，且 R4 没有任何角色(说明使用了 Point-to-Point network type)

可以使用 `show ip int ospf int g0/1` 来校验

```
R3#sh ip ospf int g0/1

GigabitEthernet0/1 is up, line protocol is up
  Internet address is 192.168.34.1/30, Area 0
  Process ID 1, Router ID 192.168.34.1, Network Type POINT-TO-POINT, Cost: 1
```

查看 R3 路由表

```
Gateway of last resort is not set

     10.0.0.0/8 is variably subnetted, 2 subnets, 2 masks
C       10.0.2.0/24 is directly connected, GigabitEthernet0/0
L       10.0.2.254/32 is directly connected, GigabitEthernet0/0
     192.168.34.0/24 is variably subnetted, 2 subnets, 2 masks
C       192.168.34.0/30 is directly connected, GigabitEthernet0/1
L       192.168.34.1/32 is directly connected, GigabitEthernet0/1
```

这里可以看见建立 full state 的 neighbor 的路由并没有通过 OSPF 学来，可能是 R3 G0/1 和 R4 G0/1 network type 设置的不一样导致的

查看 R4 G0/1 network type

```
R4#sh ip ospf int g0/1

GigabitEthernet0/1 is up, line protocol is up
  Internet address is 192.168.34.2/30, Area 0
  Process ID 1, Router ID 192.168.245.2, Network Type BROADCAST, Cost: 1
```

果然发现 R3 G0/1 使用 Point-to-Point network type，而 R4 G0/1 使用 Broadcast network type

#### Fix the problem

先查看一下 R3 OSPF 配置

```
R3#sh run | sec ospf
 ip ospf network point-to-point
 ip ospf priority 1
router ospf 1
 log-adjacency-changes
 passive-interface GigabitEthernet0/0
 network 10.0.2.254 0.0.0.0 area 0
 network 192.168.34.1 0.0.0.0 area 0
```

这里将 R3 的 G0/1 改为 Broadcast network type

```
R3(config-if)#no ip ospf network 
23:27:27: %OSPF-5-ADJCHG: Process 1, Nbr 192.168.245.2 on GigabitEthernet0/1 from FULL to DOWN, Neighbor Down: Interface down or detached

23:27:27: %OSPF-5-ADJCHG: Process 1, Nbr 192.168.245.2 on GigabitEthernet0/1 from LOADING to FULL, Loading Done
```

然后使用 `show ip ospf ne` 来查看邻居状态

```
R3(config-if)#do sh ip ospf ne
Neighbor ID     Pri   State           Dead Time   Address         Interface
192.168.245.2     1   FULL/DR         00:00:33    192.168.34.2   
```

这是如果使用 `show ip route` 就可以发现 R4 的 route

```
O       192.168.245.0/29 [110/2] via 192.168.34.2, 00:05:23, GigabitEthernet0/1
```

> 这里如果将 R3 R4 改为 Point-to-Point 也可以，且更加合理

```
R3(config)#int g0/1
R3(config-if)#ip ospf network point-to-point 

R4(config)#int g0/1
R4(config-if)#ip ospf network point-to-point 
```

### 0x03

#### R2 and R4 won’t become OSPF neighbors with R5. Why?

先看一下 R2 的 OSPF neighbor table

```
R2#sh ip ospf ne
Neighbor ID     Pri   State           Dead Time   Address         Interface
192.168.245.2     1   FULL/DR         00:00:31    192.168.245.2   GigabitEthernet0/0
192.168.12.1      0   FULL/  -        00:00:30    192.168.12.1    Serial0/0/0
```

这里确实没有看到 R5，但是可以看到和 R1 以及 R4 是 neighbors

看一下 R2 OSPF 配置

```
R2#sh run | sec ospf
router ospf 1
 log-adjacency-changes
 network 192.168.245.1 0.0.0.0 area 0
 network 192.168.12.0 0.0.0.3 area 0
```

看一下 R4 OSPF 配置

```
R4(config-if)#do sh run | sec ospf
 ip ospf network point-to-point
 ip ospf priority 1
router ospf 1
 log-adjacency-changes
 network 192.168.34.2 0.0.0.0 area 0
 network 192.168.245.2 0.0.0.0 area 0
```

看一下 R5 OSPF 配置

```
R5#sh run | sec ospf
 ip ospf hello-interval 5
 ip ospf dead-interval 20
router ospf 1
 log-adjacency-changes
 network 192.168.245.3 0.0.0.0 area 0
 default-information originate
```

这里可以很明显的看到 R5 hello timer 以及 dead timer 不是使用默认的，而 R2 和 R3 却使用默认的

#### Fix the problem

因为 timer 不是在全局修改的，所以看互联的 R5 G0/0

```
R5(config-router)#do sh ip ospf int g0/0

GigabitEthernet0/0 is up, line protocol is up
  Internet address is 192.168.245.3/29, Area 0
  Process ID 1, Router ID 203.0.113.1, Network Type BROADCAST, Cost: 1
  Transmit Delay is 1 sec, State DR, Priority 1
  Designated Router (ID) 203.0.113.1, Interface address 192.168.245.3
  No backup designated router on this network
  Timer intervals configured, Hello 5, Dead 20, Wait 20,
```

将 R5 G0/0 hello timer 以及 dead timer 改为默认的 10 和 40

```
R5(config-router)#int g0/0
R5(config-if)#ip ospf hello-interval 10
R5(config-if)#ip ospf dead-interval 40
```

这是使用 `show ip ospf ne` 就可以看到和 R5 正常建立 neighbor 了，并在 R2/R4 routing table 中出现 R5 宣告的默认路由

```
O*E2 0.0.0.0/0 [110/1] via 192.168.245.3, 00:00:03, GigabitEthernet0/0
```

> 原 LAB 中 R5 没有添加默认路由，这里为了方便检查手动添加了

### 0x04

#### PC1 and PC2 can’t ping the external server 8.8.8.8. Why?

先看 PC1, 使用 tracert 看断点

```
C:\>tracert 8.8.8.8
Tracing route to 8.8.8.8 over a maximum of 30 hops: 
  1   0 ms      0 ms      0 ms      10.0.1.254
  2   *         *         *         Request timed out.
  3   *         *         *         Request timed out
```

这里可以看见能到 R1 并回包，R1 到 8.8.8.8 可能没有路由，先在 R1 使用 ping 来测试

```
R1#ping 8.8.8.8
Type escape sequence to abort.
Sending 5, 100-byte ICMP Echos to 8.8.8.8, timeout is 2 seconds:
!!!!!
Success rate is 100 percent (5/5), round-trip min/avg/max = 6/17/40 ms
```

这里可以看到 R1 是可以正常访问 8.8.8.8 的，说明每台 router 都有往 8.8.8.8 的路由，但是因为从 PC1 到 8.8.8.8 源 IP 不同，所以可能是 router 没有往 PC1 的路由(R1 一定有，因为是直连的，Connected route)

这里看一下 R2

```
Gateway of last resort is 192.168.245.3 to network 0.0.0.0

     10.0.0.0/24 is subnetted, 1 subnets
O       10.0.2.0/24 [110/3] via 192.168.245.2, 4294967273:4294967257:4294967287, GigabitEthernet0/0
     192.168.12.0/24 is variably subnetted, 2 subnets, 2 masks
C       192.168.12.0/30 is directly connected, Serial0/0/0
L       192.168.12.2/32 is directly connected, Serial0/0/0
     192.168.34.0/30 is subnetted, 1 subnets
O       192.168.34.0/30 [110/2] via 192.168.245.2, 4294967273:4294967257:4294967287, GigabitEthernet0/0
     192.168.245.0/24 is variably subnetted, 2 subnets, 2 masks
C       192.168.245.0/29 is directly connected, GigabitEthernet0/0
L       192.168.245.1/32 is directly connected, GigabitEthernet0/0
O*E2 0.0.0.0/0 [110/1] via 192.168.245.3, 4294967273:4294967254:4294967237, GigabitEthernet0/0
```

这里可以发现 R2 是没有往 10.0.1.0/24 的路由，所以 8.8.8.8 回程会有问题的。这里可能是 R1 没有宣告对应的段，看一下 R1 OSPF 配置

```
R1#sh run | sec ospf
router ospf 1
 log-adjacency-changes
 network 192.168.12.0 0.0.0.3 area 0
```

这里果然没有宣告 10.0.1.0/24

#### Fix the prolem

只需要让 R1 宣告 10.0.1.0/24 PC1 就可以正常访问 8.8.8.8

```
R1(config-router)#network 10.0.1.0 0.0.0.255 area 0
```

使用 PC1 校验一下

```
C:\>ping 8.8.8.8

Pinging 8.8.8.8 with 32 bytes of data:

Reply from 8.8.8.8: bytes=32 time=21ms TTL=252
Reply from 8.8.8.8: bytes=32 time=16ms TTL=252
```

检查发现 PC2 已经通了

### 0x05

Examine the LSDB. What LSAs are present?

**referenes**

[^jeremy’s IT Lab]:https://www.youtube.com/watch?v=3ew26ujkiDI&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=53