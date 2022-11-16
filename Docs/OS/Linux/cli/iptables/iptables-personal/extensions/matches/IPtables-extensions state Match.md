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
| NEW         | The **NEW** state tells us that the packet is the first packet that we see. This means that the first packet that the conntrack module sees, within a specific connection, will be matched. For example, if we see a SYN packet and it is the first packet in a connection that we see, it will match. However, the packet may as well not be a SYN packet and still be considered **NEW**. This may lead to certain problems in some instances, but it may also be extremely helpful when we need to pick up lost connections from other firewalls, or when a connection has already timed out, but in reality is not closed. |
| ESTABLISHED | The **ESTABLISHED** state has seen traffic in both directions and will then continuously match those packets. **ESTABLISHED** connections are fairly easy to understand. The only requirement to get into an **ESTABLISHED** state is that one host sends a packet, and that it later on gets a reply from the other host. The **NEW** state will upon receipt of the reply packet to or through the firewall change to the **ESTABLISHED** state. ICMP reply messages can also be considered as **ESTABLISHED**, if we created a packet that in turn generated the reply ICMP message. |
| RELATED     | The **RELATED** state is one of the more tricky states. A connection is considered **RELATED** when it is related to another already **ESTABLISHED** connection. What this means, is that for a connection to be considered as **RELATED**, we must first have a connection that is considered **ESTABLISHED**. The **ESTABLISHED** connection will then spawn a connection outside of the main connection. The newly spawned connection will then be considered **RELATED**, if the conntrack module is able to understand that it is **RELATED**. Some good examples of connections that can be considered as **RELATED** are the FTP-data connections that are considered **RELATED** to the FTP control port, and the DCC connections issued through IRC. This could be used to allow ICMP error messages, FTP transfers and DCC's to work properly through the firewall. Do note that most TCP protocols and some UDP protocols that rely on this mechanism are quite complex and send connection information within the payload of the TCP or UDP data segments, and hence require special helper modules to be correctly understood. |
| INVALID     | The **INVALID** state means that the packet can't be identified or that it does not have any state. This may be due to several reasons, such as the system running out of memory or ICMP error messages that do not respond to any known connections. Generally, it is a good idea to **DROP** everything in this state. |

