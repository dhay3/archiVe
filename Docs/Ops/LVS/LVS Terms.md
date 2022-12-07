# LVS Terms

ref

https://en.wikipedia.org/wiki/Linux_Virtual_Server

http://www.austintek.com/LVS/LVS-HOWTO/

## Roles

> 通常将 LVS director 和 Real servers 称为 LVS

### Client

computers requestiing services from the virtual server.

the client does not connect to these backend servers and they are not in the ipvsadm table

### LVS director

The director which runs the `ip_vs` code，receives the connect request from the client and then choses one of the avaiable backend servers（called real server in LVS-speak）, to services the client’s reuqest

通常也被称为 dispatcher, load balancer, redirector

### Real  servers

nodes that handle requests. The individual realservers can be added and removed from the LVS without the client being aware of the changes in the LVS hardware, by commands given to the director. This allows realservers to be failed out, for routine maintenance and upgrades, or on hardware failure and for realservers to be added ( or removed ) in times of high ( or low ) load

## IP

```
                        ________
                       |        |
                       | client | (local or on internet)
                       |________|
                          CIP
                           |
--                      (router)
                          DGW
                           | outside network
                           |
L                         VIP
i                      ____|_____
n                     |          | (director can have 1 or 2 NICs)
u                     | director |
x                     |__________|
                      DIP (and PIP)
V                          |
i                          | DRIP network
r         ----------------------------------
t         |                |               |
u         |                |               |
a        RIP1             RIP2            RIP3
l    _____________   _____________   _____________
    |             | |             | |             |
S   | realserver1 | | realserver2 | | realserver3 |
e   |_____________| |_____________| |_____________|
r
v
e
r
---
```

### VIP

virtual IP address, the IP address used by the director to provide services to client computers

### RIP

real IP address, the IP address used to connect to the cluster nodes

### DIP

directors IP address, the IP address used by the director to connect to network of real IP address

### CIP

client IP address, the IP address assigned to a client computer, that is uses as the source  IP address for requests being sent to the cluster

## miscellaneous

### forwarding method

currently [LVS-NAT](http://www.austintek.com/LVS/LVS-HOWTO/HOWTO/LVS-HOWTO.LVS-NAT.html#LVS-HOWTO.LVS-NAT), [LVS-DR](http://www.austintek.com/LVS/LVS-HOWTO/HOWTO/LVS-HOWTO.LVS-DR.html#LVS-HOWTO.LVS-DR), [LVS-Tun](http://www.austintek.com/LVS/LVS-HOWTO/HOWTO/LVS-HOWTO.LVS-Tun.html#LVS-HOWTO.LVS-Tun). The director is a router with somewhat different rules for forwarding packets than a normal router. The forwarding method determines how the director sends packets from the client to the realservers. 	



### schuduling

he algorithm the director uses to select a realserver to service a new connection request from a client. 		

