# Day34 - Standard ACLs

## What are ACLs

Access Control Lists(ACLs)

- ACLs function as a packet filer, instructing the router to permit or discard specific traffic

- ACLs can filter taffic based on source/destination IP addresses, source/destination Layer 4 ports, etc
- ACLs are configured globally on the router(global config mode. They are ordered sequence of ACEs(Access Control Entries)

## How ACLs works

- Configuring an ACL in global config mode will not make the ACL take effect

- The ACL must applied to an interface

- ACLs are applied either inbound or outbound

- **ACLs are made up of one or more ACEs**

- ***When router checkes a packet against the ACL, it processes the ACEs in order, from top to bottom. If the packet matches one of the ACEs in the ACL, the router takes the action and stops processing the ACL. All entries below the matching entry will be ignored***

  将该规则抽象成函数，即从上往下的顺序匹配 if ACE，如果有一条匹配就 return

  > 和规则匹配的详细程度没关系

- **A maximum of one ACL can be applied to a single interface per direction**

  > ==即一个端口一个方向只能应用一个 ACL，后配置的 ACL 会 override 之前配置的 ACL==

例如如下拓扑

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230626/2023-06-26_11-25.1p2wxj1n74yo.webp)

需要满足如下条件

1. Hosts in 192.168.1.0/24 can access the 10.0.1.0/24 network
2. Hosts in 192.168.2.0/24 cannot access the 10.0.1.0/24 network

逻辑上需要在将 ACL 配置成如下逻辑(这里先不管配置在那个 router 那个 interface 上)

```
if source IP == 192.168.1.0/24 then:
	permit 
if source IP == 192.168.2.0/24 then:
	deny
if source IP == any then:
	permit
```

配置了 ACLs 还不会直接生效，只有 ACL 被应用在端口上才会生效，可以应用在端口的两个方向上

1. inbound

   only take effect on traffic when coming to this interface

2. outbound

   only take effect on traffic when exiting this interface

假设现在在 R1 G0/2 outbound 上配置 ACL

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230626/2023-06-26_12-01.4v1f12ngasjk.webp)

1. PC3 访问 10.0.1.100 过 R1 G0/2 因为是 inbound 不会使用 ACL 过滤报文
2. R1 G0/0 outbound 到 SRV1 没有 ACL 所以也不会过滤报文
3. SRV1 到 R1 G0/0 inbound 回包没有 ACL 所以页不会过滤报文
4. R1 G0/2 到 PC3 回包因为是 outbound 所以会按照 ACL 过滤报文，source IP address 是 10.0.1.100 匹配 ACE3 所以放行

5. 所以在 R1 G0、2 outbound 上配置 ACL 不满足条件

假设现在在 R1 G0/2 inbound 是配置 ACL

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230626/2023-06-26_12-07.40j8xzz3rojk.webp)

1. PC3 访问 10.0.1.100 过 R1 G0/2 因为是 inbound 所以会使用 ACL 过报文，source IP address 192.168.2.1 匹配 ACE2 所以 deny 直接丢弃报文
2. 虽然 PC1/PC2 可以访问 SRV1 也是满足条件的，但是 PC1/2 到 SRV1 不经过 R1 G0/2 outbound，ACL 也就不会被使用，所以还不是最佳的

> 注意这里 PC3 和 PC4 还是可以正常通信的，因为实际是通过 switch 互联的，而不是 router, ACL 只能在 router 上

R2 G0/1 outbound 和 R2 G0/0 inbound 是最佳的选择

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230626/2023-06-26_12-15.zypwvflc9ps.webp)

1. PC3 访问 10.0.1.100 到 R2 G0/0 inbound 没有 ACL 所以不会使用 ACL 过滤报文
2. 通过 R2 G0/1 outbound 转发会按照 ACL 过滤报文，source IP 192.168.2.100 匹配 ACE2，所以 deny，报文被丢弃
3. PC1 访问 10.0.1.100 到 R2 G0/0 inbound 没有 ACL 所以不会使用 ACL 过滤报文
4. SRV1 回包到 R2 G0/1 inbound 没有 ACL 所以不会使用 ACL 过滤报文
5. R2 G0/0 outbound 到 PC1 没有 ACL 所以也不会过滤报文

## Implicit deny

如果 packet 没有匹配任何一条 ACE，会出现什么情况？

例如 router ACL 逻辑如下

```
if source IP == 192.168.1.0/24 then:
	permit
if source IP == 192.168.2.0/24 then:
	deny
```

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230626/2023-06-26_14-03.uc6gfw2ut9c.webp)

现在 router 收到 src 10.0.0.1 dst 1.1.1.1 的报文，没有匹配到 ACL 中的任意一条 ACE， router 会怎么做？

router 会 deny 对应的报文，这一动作被称为 Implicit deny

*The implicit deny tells the router to deny all traffic that doesn’t match any of the configured entries in the ACL*

## ACL Types

按照过滤的条件 ACL 主要分为两大类

1. Standard ACLs

   match based on ==Source IP address only==

   按照 ACL 的名字可以分为两类

   - Standard Numbered ACLs

     > Statndard ACLs can use 1-99 and 1300-1999

     例如 ACL 1, ACL 2

   - Standard Named ACLs

     例如 ACL to-internet

2. Extended ACLs

   match based on Source/Destination IP, Souce/Destination port,etc

   同样的按照名字也可以分为两类

   - Extended Numbered ACLs
   - Extended Named ACLs

### Configure Standard ACLs

#### standard numbered ACLs

如果需要配置 Standard ACLs 可以使用 `access-list <number> {deny | permit} <ip> <wildcard-mask>`

> 注意 number 只能是 1-99 或者是 1300-1999

例如 `access-list 1 deny 1.1.1.1 0.0.0.0`

就表示只 deny 1.1.1.1

如果你想配置 32 bit mask IP 的规则，也可以不用 wildcard-mask， router 会默认识别成 32 bit mask

例如 `access-list 1 deny 1.1.1.1` 等价与上面的命令

除此外，还可以使用 `access-list 1 deny host 1.1.1.1` 来表示匹配 32 bit mask IP，等价于上面的 2 个命令

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230626/2023-06-26_14-19.4nw168fo0xkw.webp)

> 如果想要配置非 32 bit mask 的就只能用第一种的方式

如果想要一条规则配置所有的 IP，可以使用例如 `access-list 1 permit any` 等价与 `access-list 1 permit 0.0.0.0 255.255.255.255`

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230626/2023-06-26_14-26.68h5k2c6r2bk.webp)

假设使用了如下配置

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230626/2023-06-26_14-29.5jnhtlxsmj5s.webp)

我们就可以使用 `show access-lists` 来查看

这里可以看到只有第一条和第二条 ACE 加入了 ACL 1 ，并且将 1.1.1.1 0.0.0.0 直接转为 1.1.1.1，0.0.0.0 255.255.255.255 转为 any，因为两者都是等价的

10 和 20 分别为对应 ACE 在 ACL 中判断的先后顺序，即先判断值小的

还可以使用 `show ip access-lists` 来查和 IP 相关的 ACL

> `show access-lists` 可以查看所有的 ACL，而 `show ip access-lists` 只能查看和 IP 相关的 ACL，例如不能查看 协议、端口、标志位相关的 ACL

除上面两个命令外来查看 ACL，还可以使用 `show running-config | include access-list` 查看，这里也可以观察到 router 将输入的命令自动转为等价的 1.1.1.1 和 any

如果需要将 ACL 应用到 interface 上可以使用 `ip access-group <number> {in | out}`

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230626/2023-06-26_14-55.3szl3syefzy8.webp)

##### ACL remark

ACL remark 类似注释，没有实际的效果，只为 ACL 其注释提示的功能，在 Standard Named ACLs 中同样适用

可以使用 `access-list <number> remark <remark>` 命令来配置

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230626/2023-06-26_14-25.4j387tqjsjcw.webp)

> remark 部分的 \# hash 并不是并要的

##### Example

例如下面拓扑

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230626/2023-06-26_15-04.dron03a9mhs.webp)

需要满足如下条件

1. PC1 can access 192.168.2.0/24
2. Other PCs in 192.168.1.0/24 can’t access 192.168.2.0/24

逻辑上 ACL 如下

```
if source IP == 192.168.1.1 then:
	permit 192.168.2.0 => R1 G0/2
if source IP == 192.168.1.0/24 then
	deny 192.168.2.0 => R1 G0/2
```

因为 192.168.2.0/24 和 192.168.1.0/24 都通过 R1 互联，所以配置 R1

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230626/2023-06-26_15-07.650zoob8s51c.webp)

1. `access-list 1 permit 192.168.1.1`

   满足条件1 的 ACE

2. `access-list 1 deny 1921.168.1.0 0.0.0.255`

   满足条件2 的 ACE

3. `access-list 1 permit any`

   因为默认会有 implicit deny，所以需要配置这条规则，保证其他流量不被过滤

4. `ip access-group 1 out`

   将 ACL 1 应用到 R1 g0/2 outbound

> 在 ACL 需要应用到那个端口上有一个小技巧
>
> *Standard ACLs should be applied as close to the destination as possible*
>
> 个人觉得其实配置在 inbound 更合理，因为既然要 deny traffic 那可以完全没必要让 traffic 再经过 router，可能会增加网络拥塞的几率

现在 PC1 ping PC3

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230626/2023-06-26_15-27.10i7rivq42ww.webp)

1. 入报文到 R1 G0/1 inbound 因为没有应用 ACL，所以不会检查 ACL
2. R1 查看路由表转发到 G0/2
3. 因为 R1 G0/2 outbound 应用了 ACL，所以会检查 ACL，匹配 `10 permit 192.168.1.1`，所以正常通过 R1 G0/2 转发到 PC3
4. PC3 回包到 R1 G0/2 inbound 因为没有应用 ACL，所以不会检查 ACL，同理 R1 G0/1

现在 PC2 ping PC3

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230626/2023-06-26_15-36.70mg3vons1z4.webp)

1. 入报文到 R1 G0/1 inbound 因为没有应用 ACL，所以不会检查 ACL
2. R1 查看路由表转发到 G0/2
3. 因为 R1 G0/2 outbound 应用了 ACL，所以会检查 ACL，匹配 `10 deny  192.168.1.0 wildcard bits 0.0.0.255`，所以 R2 会将报文直接丢弃，PC2 也就不能 ping 通 PC3

#### Standard Named ACLs

如果需要使用 Standard Named ACLs 可以使用 `ip access-list standard <acl-name>` 会进入 standard name ACL config mode

> Standard Named ACLs 也可以通过这种方式来配置
>
> `ip access-list standard <acl-number>`

然后可以使用 `[entry-number] {deny | permit} ip wildcard-mask` 来配置 ACE，entry-number 可以不指定会自动按照输入的命令先后顺序赋值，对应判断 ACE 的先后顺序

> entry-number，默认以 10 递增，从 10 开启，也可以手动指定 ACE 使用的具体值

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230626/2023-06-26_15-47.499u9vcr2b9c.webp)

和 Standard Numbered ACLs 一样，也可以使用 `show access-lists`，`show ip access-lists` 来查看配置的 ACL，另外也可以使用 `show running-config | section access-list` 来查看 ACL 整个部分

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230626/2023-06-26_15-49.d8hf6gjehrk.webp)

##### Example

例如如下拓扑

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230626/2023-06-26_15-54.35bks3qaxtkw.webp)

需要满足如下条件

1. PCs in 192.168.1.0/24 can’t access 10.0.2.0/24
2. PC3 can’t access 10.0.1.0/24
3. Other PCs in 192.168.2.0/24 can access 10.0.1.0/24
4. PC1 can access 10.0.1.0/24
5. Other PCs in 192.168.1.0/24 can’t access 10.0.1.0/24

逻辑上 ACL 如下

```
if source IP == 192.168.1.0/24 then
	deny 10.0.2.0/24 => R2 G0/2
if source IP == 192.168.2.1 then
	deny 10.0.1.0/24 => R2 G0/1
if source IP == 192.168.2.0/24 then
	permit 10.0.1.0/24 => R2 G0/1
if source IP == 192.168.1.1 then
	permit 10.0.1.0/24 => G2 G0/1
if source IP == 192.168.1.0/24 then
	deny 10.0.1.0/24 => G2 G0/1
```

因为 10.0.2.0/24 和 10.0.1.0/24 都通过 R2 互联，所以配置 R2

条件 1 配置在 R2 G0/2, 条件 2345 配置在 R2 G0/1

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230626/2023-06-26_16-08.1vk2sgvhu4u8.webp)

1. `ip access-list standard TO_10.0.2.0/24`

   定义一个 Standard Named ACL

2. `deny 192.1681.1.0 0.0.0.255`

   `permit any`

   满足条件 1 的 ACE

3. `interface g0/2`

   `ip access-group TO_10.0.2.0/24 out`

   将 TO_10.0.2.0/24 Standard Named ACl 应用在 R2 g0/2 outbound

4. `ip access-list standard TO_10.0.1.0/24`

   定义另外一个 Standard Name ACL

5. `deny 192.168.2.1`

   满足条件 2 的 ACE

6. `permit 192.168.2.0 0.0.0.255`

   满足条件 3 的 ACE

7. `permit 192.168.1.1`

   满足条件 4 的 ACE

8. `deny 192.168.1.0 0.0.0.255`

   满足条件 5 的 ACE

8. `permit any`

   防止 implicit deny 过滤其他所有的流量

9. `interface g0/1`

   `ip access-group TO_10.0.1.0/24 out`

   将 TO_10.0.1.0/24 Standard Named ACL 应用在 R2 g0/1 outbound

如果这时使用 `show ip access-lists` 来查看配置的 ACL

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230626/2023-06-26_16-24.a1ruaceryc8.webp)

这里可以看到命令输入的顺序和 `show ip access-lists` 显示的不一样

- The router may re-order the /32 entries

- This improves the efficiency of processing the ACL

  > 主要就是为了提高处理 ACL 的效率，逻辑上按照 10 20 30 40 50 来过滤规则，但是实际上按照 30 10 20 40 50 来过滤规则

- It does not change the effect of the ACL

  这里可以通过黄框中 ACE 对应的 sequence number(entry-number)得出

- This applies to both standard named and standard numbered ACLs

现在 PC2 ping SRV1

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230626/2023-06-26_16-35.622qy6qtgl4w.webp)

1. 入报文到 R2 G0/1 inbound 因为没有应用 ACL，所以不会检查 ACL
2. R2 查看路由表转发到 G0/1
3. 因为 R2 G0/1 outbound 应用了 ACL，所以会检查 ACL，匹配 `40 deny 192.168.1.0, wildcard bits 0.0.0.255`，所以 R2 不会通过 R2 G0/1 转发到 SRV1

## Quiz

### quiz3

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230626/2023-06-26_16-43.3jrrgaomig5c.webp)

==每一个接口只一个方向只能应用一个 ACL，后面应用的 ACL 会覆盖之前的==，所以这里实际使用的是 ACL 10

## LAB

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230626/2023-06-26_17-02.29yhnoyx29q8.webp)

### 0x01

Configure OSPF on R1 and R2 to allow full connectivity between the PCs and server

配置前先看一下 `show running-config`

R1

```
R1(config)#int lo0
R1(config-if)#ip add 1.1.1.1 255.255.255.255
R1(config)#router ospf 1
R1(config-router)#network 172.16.1.0 0.0.0.255 area 0
R1(config-router)#network 172.16.2.0 0.0.0.255 area 0
R1(config-router)#network 1.1.1.1 0.0.0.0 area 0
R2(config-router)#network 203.0.113.0 0.0.0.3 area 0
R1(config-router)#passive-interface lo0
R1(config-router)#passive-interface g0/0
R1(config-router)#passive-interface g0/1
```

R2

```
R1(config)#int lo0
R1(config-if)#ip add 2.2.2.2 255.255.255.255
R2(config)#router ospf 1
R2(config-router)#network 192.168.1.0 0.0.0.255 area 0
R2(config-router)#network 192.168.2.0 0.0.0.255 area 0
R2(config-router)#network 2.2.2.2 0.0.0.0 area 0
R2(config-router)#network 203.0.113.0 0.0.0.3 area 0
R2(config-router)#passive-interface lo0
R2(config-router)#passive-interface g0/0
R2(config-router)#passive-interface g0/1
```

配置完可以使用 `show ip ospf ne` 和 `show ip protocol` 来检查

### 0x02

Configure standard numbered ACLs on R1 and standard named ACLs on R2 to fullfill the following network policies

1. Only PC1 and PC3 can access 192.168.1.0/24

   ```
   R2(config)#ip access-list standard TO_SRV1
   R2(config-std-nacl)#permit 172.16.1.1
   R2(config-std-nacl)#permit 172.16.2.1
   R2(config-std-nacl)#do sh access-list 
   Standard IP access list TO_SRV1
       10 permit host 172.16.1.1
       20 permit host 172.16.2.1
   R2(config-std-nacl)#int g0/0
   R2(config-if)#ip access-group TO_SRV1 out
   ```

   > 虽然有 implicit deny，但是也可以手动添加一条规则 `deny any` 来显式声明

2. Hosts in 172.16.2.0/24 can’t access 192.168.2.0/24

   ```
   R2(config)#access-list 1 deny 172.16.2.0 0.0.0.255 
   R2(config)#access-list 1 permit any
   R2(config)#int g0/1
   R2(config-if)#ip access-group 1 out
   ```

3. 172.16.1.0/24 can’t access 172.16.2.0/24

   ```
   R1(config)#ip access-list standard TO_172.16.2.0
   R1(config-std-nacl)#deny 172.16.1.0 0.0.0.255
   R1(config-std-nacl)#permit any
   R1(config-std-nacl)#int g0/1
   R1(config-if)#ip access-group TO_172.16.2.0 out
   ```

4. 172.16.2.0/24 can’t access 172.16.1.0/24

   ```
   R1(config)#access-list 1 deny 172.16.2.0 0.0.0.255
   R1(config)#access-list 1 permit any
   R1(config)#int g0/0
   R1(config-if)#ip access-group 1 out
   ```

注意这里如果测试都会回 ICMP Destination host unreachable，因为对应的报文已经到了 router，但是如果 ACL 配置在 inbound 方向上，就会显式 asterisk 表示丢包

如果配置完不匹配现象，需要考虑 implicit deny，回向的报文同样会经过 ACL

**references**

1. [^jeremy’s IT Lab]:https://www.youtube.com/watch?v=z023_eRUtSo&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=65