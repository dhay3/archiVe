# firewalld - Overview

## Overview

> firewalld is a simple, stateful, zone-based firewall.

stateful

## Zones

*A firewall zone defines the trust level for a connection, interface or  source address binding. This is a one to many relation, which means that a connection, interface or source can only be part of one zone, but a  zone can be used for many network connections, interfaces and sources.*

可以将 zones 想象成防火墙策略的集合，规定了

Zones 严格遵循如下规则

1. traffic egress/ingress one and only one zone

   出/入向流量只进入一个 zone

2. a zone defines a level of trust
3. intra-zone(within in the same zone) is allow by default
4. inter-zone(zone to zone) is denied by default

firewalld 有入向 predefined zones

| Zone     |                                                              |
| -------- | ------------------------------------------------------------ |
| drop     | Any incoming network packets are dropped, there is no reply. Only outgoing network connections are possible. |
| block    | Any incoming network connections are rejected with an  icmp-host-prohibited message for IPv4 and icmp6-adm-prohibited for IPv6. Only network connections initiated within this system are possible. |
| public   | For use in public areas. You do not trust the other computers on  networks to not harm your computer. Only selected incoming connections  are accepted. |
| external | For use on external networks with IPv4 masquerading enabled  especially for routers. You do not trust the other computers on networks to not harm your computer. Only selected incoming connections are  accepted. |
| dmz      | For computers in your demilitarized zone that are publicly-accessible with limited access to your internal network. Only selected incoming  connections are accepted. |
| work     | For use in work areas. You mostly trust the other computers on  networks to not harm your computer. Only selected incoming connections  are accepted. |
| home     | For use in home areas. You mostly trust the other computers on  networks to not harm your computer. Only selected incoming connections  are accepted. |
| internal | For use on internal networks. You mostly trust the other computers on the networks to not harm your computer. Only selected incoming  connections are accepted. |
| truested | All network connections are accepted.                        |

**references**

[^1]:https://firewalld.org/documentation/concepts.html	