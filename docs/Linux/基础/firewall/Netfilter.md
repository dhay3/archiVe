# Netfilter

## 概述

Netfilter是linux内核中一个软件框架，用于管理网络数据包。具有NAT，修改数据包以及过滤数据包的功能。用户可以通过iptables，ufw等软件来控制Netfilter。

Netfilter 内置了5个数据包的挂载点(Hook)：

1. PRE_ROUTING
2. INPUT
3. OUTPUT
4. FORWARD
5. POST_ROUTING