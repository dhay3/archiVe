# Broadcast

## What is Broadcast?

*Broadcast is a method delivers messages to all nodes in a range using one-to-all association*

简而言之在网络中 Broadcast(广播) 就是一种 “一对所有” 的通信模型，这里借用 wikipedia 中的图

![](https://upload.wikimedia.org/wikipedia/commons/d/dc/Broadcast.svg)

“一” 指的是发送报文的 host(router/switch)

“所有” 指的是特定范围中的所有 hosts(router/switch)

## Why Broadcast?

> Broadcasts are useful when a host needs to find information without knowing exactly what other host can supply it, or when a host wants to provide information to a large set of hosts in a timely manner.
>
> cite from [RFC919](https://datatracker.ietf.org/doc/html/rfc919#section-3)

广播的出现主要因为两种场景

1. host 需要特定主机的信息，但是不知道特定主机的标识符，例如 ARP
2. host 需要推送消息到指定的 hosts，例如 Gratuitous ARP

## Broadcast Range

> When a datagram is broadcast, it imposes a cost on every host that hears it. Therefore, broadcasting should not be used indiscriminately, but rather only when it is the best solution to a problem.
>
> cite from [RFC919](https://datatracker.ietf.org/doc/html/rfc919#section-3)

因为“一对所有”的特性，Broadcast 是有缺点的，会对范围内的网络和设备都产生消耗

所以针对所有的 Internet

当然 Broadcast 的范围不是无限的

```
   Local Hardware Network

      The physical link to which the host is attached.

   Remote Hardware Network

      A physical network which is separated from the host by at least
      one gateway.

   Collection of Hardware Networks

      A set of hardware networks (transitively) connected by gateways.

   The IP world includes several kinds of logical network.  To avoid
   ambiguity, we will use the following terms:

   Internet

      The DARPA Internet collection of IP networks.

   IP Network

      One or a collection of several hardware networks that have one
      specific IP network number.
```

## Broadcast Address

在网络中如何将 messages(报文) 广播出去呢？

需要通过特定规则的 Broadcast Address，分为 3 类

- Layer 2 Broadcast Address
- Layer 3 Limited Broadcast Address
- Layer 3 Broadcast Address

### Layer 2 Broadcast

FFFF.FFFF.FFFF

### Layer 3 Limited Broadcast

255.255.255.255

### Layer 3 Broadcast

192.168.1.255

#### How to calculate

$$
layer\ 3\ broadcast\ address = net\ portion \ OR\ (\ NOT \ subnetmask)
$$

## Directly Broadcast



**references**

1. ^https://en.wikipedia.org/wiki/Broadcasting_(networking)
2. ^https://datatracker.ietf.org/doc/html/rfc919
3. ^https://www.youtube.com/watch?v=HZs93eNHyaU
4. ^https://community.cisco.com/t5/switching/255-255-255-255-vs-subnet-broadcast-address-amp-ip-direct/td-p/3215598