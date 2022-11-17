# IPtables-extensions state Match

ref

https://www.linuxtopia.org/Linux_Firewall_iptables/x1347.html

## Digest

state 是 conntrack 的一个子集, 不涵盖 NAT 的场景

## Optional args

- `[!] --state state`

  where state is a comma separated list of the connection states to match. Only a subset of the states unterstood by “conntrack” are recognized: INVALID, ESTABLISHED, NEW, RELATED or UNTRACKED

## state

| State       | Explanation                                                  |
| ----------- | ------------------------------------------------------------ |
| NEW         | The **NEW** state tells us that the packet is ==the first packet(注意是第一个报文，并不只针对 SYN 报文)== that we see. This means that the first packet that the conntrack module sees, within a specific connection, will be matched. For example, if we see a SYN packet and it is the first packet in a connection that we see, it will match. However, the packet may as well not be a SYN packet and still be considered **NEW**. This may lead to certain problems in some instances, but it may also be extremely helpful when we need to pick up lost connections from other firewalls, or when a connection has already timed out, but in reality is not closed. |
| ESTABLISHED | The **ESTABLISHED** state has seen traffic in both directions and will then continuously match those packets. **ESTABLISHED** connections are fairly easy to understand. The only requirement to get into an **ESTABLISHED** state is that one host sends a packet, and that it later on gets a reply from the other host. The **NEW** state will upon receipt of the reply packet to or through the firewall change to the **ESTABLISHED** state. ICMP reply messages can also be considered as **ESTABLISHED**, if we created a packet that in turn generated the reply ICMP message. |
| RELATED     | The **RELATED** state is one of the more tricky states. A connection is considered **RELATED** when it is related to another already **ESTABLISHED** connection. What this means, is that for a connection to be considered as **RELATED**, we must first have a connection that is considered **ESTABLISHED**. The **ESTABLISHED** connection will then spawn a connection outside of the main connection. The newly spawned connection will then be considered **RELATED**, if the conntrack module is able to understand that it is **RELATED**. Some good examples of connections that can be considered as **RELATED** are the FTP-data connections that are considered **RELATED** to the FTP control port, and the DCC connections issued through IRC. This could be used to allow ICMP error messages, FTP transfers and DCC's to work properly through the firewall. Do note that most TCP protocols and some UDP protocols that rely on this mechanism are quite complex and send connection information within the payload of the TCP or UDP data segments, and hence require special helper modules to be correctly understood. |
| INVALID     | The **INVALID** state means that the packet can't be identified or that it does not have any state. This may be due to several reasons, such as the system running out of memory or ICMP error messages that do not respond to any known connections. Generally, it is a good idea to **DROP** everything in this state. |

### TCP connection

> 需要注意的一点是
>
> NEW state 并不仅仅只针对 SYN 报文，如果直接发送畸形的 TCP 报文也是会匹配的，例如 直接发送 ACK 或者 RST 报文。因为 IPtables 把第一个报文当做 NEW state

从上面的文字也可以看出 IPtables state machine 和 TCP state machine 是有很大的区别的

**TCP 3-way handshake**

![2022-11-17_12-31](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221117/2022-11-17_12-31.3iaao6cbo8zk.webp)

1. IPtables state machine 并没有 TCP state machine 中类似 SYN_RCD 的这种状态
2. 在 TCP state machine 中如果发送或者收到了一个 SYN 报文，会被认为是 NEW state
3. 如果回了 SYN-ACK 或者是 TCP 建连后的报文，会被认为 ESTABLISHED

可以思考一下为什么会出现这种情况：

如果我们要实现 Digest 部分的例子，按照 TCP state machine 需要设置什么样的规则呢

我们应该设置客户端出向放通 80 端口，入向只放通 80 端口的 SYN-ACK 报文和 TCP 建连后的报文，但是这样我们需要设置两条规则。这两条规则都有一点共性就是都作为服务端的回包。IPtables 为了方便管理将这两种状态的报文统称为 ESTABLISHED

**TCP 4-way handshake**

看了 TCP 建连的 3-way handshake，看一下 关闭连接的 4-way handshake 是否和 TCP state machine 也有区别呢，答案是肯定的

![2022-11-17_12-47](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221117/2022-11-17_12-47.2eoxj8esnpxc.webp)

1. FIN-WAIT. CLOSE-WAIT 在 IPtables state machine 中统称为 ESTABLISHED
2. TIME-WAIT 在 IPtables state machine 中被称为 CLOSED, 如果收到了 RST 报文同样会进入 CLOSED

### UDP connections
