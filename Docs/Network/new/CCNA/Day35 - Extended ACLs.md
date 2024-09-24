# Day35 - Extended ACLs

## Advantages of named of ACL Config mode

![](https://github.com/dhay3/image-repo/raw/master/20230626/2023-06-28_15-27.4bfq4v9n39hc.webp)

使用 `access-list <number> {permit | deny} <source>` 和 `ip access-list standard <number>` 虽然结果一样，但是使用 `ip access-list standard <number>` 有几个好处

1. You can easily delete individual entries in the ACL with `no <entry-number>`

   ![](https://github.com/dhay3/image-repo/raw/master/20230626/2023-06-28_15-32.3875t53apdq8.webp)

   例如我们想要删除配置的 `deny 192.168.3.0 0.0.0.255` 就可以使用 `no 30`，30 是使用 `show access-lists` 得到的

   如果是使用 `access-lists 1 deny 192.168.3.0 0.0.0.255` 的方式配置 ACE，你可能以为可以通过 `no access-lists 1 deny 192.168.3.0 0.0.0.255` 的方式来删除配置的 ACE

   ![](https://github.com/dhay3/image-repo/raw/master/20230626/2023-06-28_15-35.2v9pzqvm1jnk.webp)

   ***When configuring/editing numbered ACLs from global config mode, you can’t delete individual entries, you can only delete the entire ACL***

   但是实际并不会，而是会删除整条 ACL

   > 但是可以在 global config mode 中配置 ACL，然后在 named ACL config mode 中来修改 ACL

2. You can insert new entries in between other entries by specifying the  sequence(entry) number

   在 global config mode 中不能手动指定 sequence(entry) number，后配置的 ACE sequence 会递增 10，但是不支持中间插入一条，如果需要在中间插入一条 ACE，需要使用 named ACL 的方式

   例如添加一条 `30 deny 192.168.2.0 0.0.0.255` ACE

   ![](https://github.com/dhay3/image-repo/raw/master/20230626/2023-06-28_15-44.6ijlm2yyei2o.webp)

   因为 sequence number 30 在 20 和 40 之间，所以对应的 ACE 也会在 20 和 40 之间，按照 20 30 40 逻辑过滤报文 

## Resequencing ACL

> 不管是 Standard ACL 还是 Extended ACL，命令均相同

如果想要修改 ACL 初始以及递增的 sequence 值，可以使用 `ip access-list resequence <acl-number> <start-seq-num> <increment>` 命令

![](https://github.com/dhay3/image-repo/raw/master/20230626/2023-06-28_15-50.5mrubvnp4d1c.webp)

> 这里按照 1 3 2 4 5 的顺序显示 ACE，是因为 router 会重新将 ACE 排序以加快处理 ACL 的速率

如果我们这时使用 `ip access-list resequence 1 10 10`，ACL 起始的 sequence number 就会从 10 开始，然后以 10 递增

> 这里为什么需要修改 1 3 2 4 5 的这个 sequence number 呢？
>
> 如果我们想要对 1 3 2 4 5 的 ACL 增加 ACE，这个是不行的 

## Extended ACLs

Extended ACLs 功能和逻辑上和 Standard ACLs 类似，一样可以使用 numbered 或者是 named 方式来配置

*Numbered ACLs use the following ranges: 100 - 199, 2000 - 2699*

和 Standard ACLs 最大的区别就是，Extended ACLs 支持的过滤条件更加多，不仅仅支持过滤 Source IP address

在 CCNA 中只关注几个过滤条件 Layer 4 Protocol/Port, source address, and destination address

![](https://github.com/dhay3/image-repo/raw/master/20230626/2023-06-28_16-04.1k28edib0c1s.webp)

两种方式可以配置 extend ACLs

1. 在 global config mode 中使用 `access-list number [permit | deny] protocol src-ip dest-ip`

2. 使用 `ip access-list extended {name | number}` 来定义 ACL，使用 `[seq-num] [permit | deny] protocol src-ip dest-ip` 来定义 ACE


protocol 字段可以使用 protocol number 或者是 protocol name 来表示

![](https://github.com/dhay3/image-repo/raw/master/20230629/2023-06-29_11-11.5teblq34jj0g.webp)

这里需要注意一个选项 `ip`，如果使用该 protocol 只要使用了 3 层 IP 协议的报文都会被过滤

如果想要过滤 /32 source 或者是 destination，和 standard ACL 不一样，在 extended ACL 中必须使用 `host` 或者是指明 0.0.0.0 wildcard mask

![](https://github.com/dhay3/image-repo/raw/master/20230629/2023-06-29_11-18.5yrsy3gstvuo.webp)

上图中的 `deny tcp any 10.0.0.0 0.0.0.255` 会过滤所有从 10.0.0.0/24 过来的 4 层报文

几个例子加深一些影响

![](https://github.com/dhay3/image-repo/raw/master/20230629/2023-06-29_11-24.4u16ku3w2pds.webp)

1. 还可以使用 `permit ip any`
2. 还可以使用 `deny udp 10.0.0.0 0.0.0.255 192.168.1.1 0.0.0.0`
3. 还可以使用 `deny icmp 172.16.1.1 0.0.0.0 192.168.0.0 0.0.255.255`

### Matching the TCP/UDP port numbers

如果过滤的是 protocol 是 TCP/UDP，还可按照 source port/destination port 来过滤，有几个 keywords

- `eq 80 = equal to port 80`
- `gt 80 = greater than 80`
- `lt 80 = less than 80`
- `neq 80 = not equal 80`
- `range 80 100 = from port 80 to port 100`

![](https://github.com/dhay3/image-repo/raw/master/20230629/2023-06-29_11-30.67822zhyvny8.webp)

和 protocol 一样，port 也可以使用 number 或者是 protocol name(表示协议对应的默认端口)来表示

![](https://github.com/dhay3/image-repo/raw/master/20230629/2023-06-29_11-32.78owzgphyzy8.webp)

例如输入 `deny tcp any host 1.1.1.1 eq http`, 就会拒绝所有 4 层从 1.1.1.1:80 过来的报文

除此外在 port 后还可以更上一些过滤条件

1. ack: match the TCP ACK flag
2. fin: match the TCP FIN flag
3. syn: match the TCP SYN flag
4. ttl: match packets with a specific TTL value
5. dscp: match packets with a specific DSCP value

一些例子加深一下影响

![](https://github.com/dhay3/image-repo/raw/master/20230629/2023-06-29_12-06.7gi87dinmum8.webp)

## Example

有如下拓扑和需要满足的条件

![](https://github.com/dhay3/image-repo/raw/master/20230629/2023-06-29_12-08.18268jtu1isg.webp)

> *Extended ACLs should be applied **as close to the source as possible**,to limit how far the packets travel in the network before being denied*
>
> *Standard ACLs are less specific, so if they are appied close to the source there is a risk of blocking more traffic than intended*
>
> 在 extended ACL 中应该尽量靠近源，所以这个例子在 R1 上配置 ACL

先配置满足第一个条件的 ACL

![](https://github.com/dhay3/image-repo/raw/master/20230629/2023-06-29_12-11.3iwnbqu3x0w0.webp)

按照上面的规则需要配置在源就近的接口上，即 R1 G0/1 inbound

![](https://github.com/dhay3/image-repo/raw/master/20230629/2023-06-29_12-16.3s1bjr521eps.webp)

如果 PC1 想要访问 SRV1 443，入向报文到 R1 G0/1 时就会被直接丢弃

配置满足第二个条件的 ACL

![](https://github.com/dhay3/image-repo/raw/master/20230629/2023-06-29_12-20.20z4ulcoyd1c.webp)

同理按照就近规则需要配置在 R1 G0/2 inbound

配置满足第三个条件的 ACL

![](https://github.com/dhay3/image-repo/raw/master/20230629/2023-06-29_12-22.70peg5w9exz4.webp)

这里只需要配置 3 条规则，不需要配置 `deny icmp 192.168.2.0 0.0.0.255 10.0.2.0 0.0.0.255`，因为在第二个条件中过滤的 protocol ip 中已经包括了 ICMP，所以可以不用配置(但是也可以声明，同样不影响结果)

![](https://github.com/dhay3/image-repo/raw/master/20230629/2023-06-29_12-26.5tgzb6dk0t8g.webp)

按照就近规则需要配置在 R1 G0/0 outbound

如果 PC1/PC3 ping SRV1 报文到 R1 G0/0 出向就会被直接丢弃

使用 `show access-lists` 可以看到所有的配置的 ACL 如下

![](https://github.com/dhay3/image-repo/raw/master/20230629/2023-06-29_12-28.53z3ktkt5t6o.webp)

我们还可以使用 `show ip interface <interface-id>` 来查看端口上配置的 ACL

![](https://github.com/dhay3/image-repo/raw/master/20230629/2023-06-29_12-29.ij0f9usdr9c.webp)

这里可以看到 R1 G0/0 outbound 配置了 ACL BLOCK_ICMP, inbound 没有配置任何 ACL

> 这里再声明一点，一个端口一个方向只能有一个 ACL 生效，后配置的覆盖前面配置的

## LAB

![](https://github.com/dhay3/image-repo/raw/master/20230630/2023-06-30_16-02.12zf9lkyufk0.webp)

Confiugre extended ACLs to fulfill the following network policies

1. Hosts in 172.16.2.0/24 can’t communicate with PC1

   ```
   R1(config)#ip access-list extended 110
   R1(config-ext-nacl)#deny ip 172.16.2.0 0.0.0.255 172.16.1.1 0.0.0.0
   R1(config-ext-nacl)#permit ip any any
   R1(config-ext-nacl)#int g0/1
   R1(config-if)#ip access-group 110 in
   ```

2. Hosts in 172.16.1.0/24 can’t access the DNS service on SRV1

   > 这里需要过滤 tcp，因为 DNS 同样也可以通过 tcp 来查询，例如 Doh

   ```
   R1(config)#access-list 120 deny udp 172.16.1.0 0.0.0.255 host 192.168.1.100 eq 53
   R1(config)#access-list 120 deny tcp 172.16.1.0 0.0.0.255 host 192.168.1.100 eq 53
   R1(config-ext-nacl)#access-list 120 permit ip any any
   R1(config-ext-nacl)#int g0/0
   R1(config-if)#ip access-group 120 in
   ```

   PC1

   ```
   C:\>ping pc3
   Pinging 172.16.2.1 with 32 bytes of data:
   Request timed out.
   
   C:\>telnet 192.168.1.100 80
   Trying 192.168.1.100 ...Open
   [Connection to 192.168.1.100 closed by foreign host]
   ```

3. Hosts in 172.16.2.0/24 can’t access the HTTP or HTTPS services on SRV2

   ```
   R1(config)#ip access-list extended 110
   R1(config-ext-nacl)#15 deny tcp 172.16.2.0 0.0.0.255 host 192.168.2.100 eq 80
   R1(config-ext-nacl)#18 deny tcp 172.16.2.0 0.0.0.255 host 192.168.2.100 eq 443
   
   R1(config-ext-nacl)#do sh ac
   Extended IP access list 110
       10 deny ip 172.16.2.0 0.0.0.255 host 172.16.1.1 (2 match(es))
       15 deny tcp 172.16.2.0 0.0.0.255 host 192.168.2.100 eq www
       18 deny tcp 172.16.2.0 0.0.0.255 host 192.168.2.100 eq 443
       20 permit ip any any (5 match(es))
   ```

   > 在条件 1 中已经在 R1 G0/1 in 上配置了，所以这里不需要配置

   PC3

   ```
   C:\>telnet 192.168.2.100 90
   Trying 192.168.2.100 ...
   % Connection refused by remote host
   C:\>telnet 192.168.2.100 80
   Trying 192.168.2.100 ...
   % Connection timed out; remote host not responding
   C:\>telnet 192.168.2.100 443
   Trying 192.168.2.100 ...
   % Connection timed out; remote host not responding
   ```

**references**

1. [^jeremy’s IT Lab]:https://www.youtube.com/watch?v=dUttKY_CNXE&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=67