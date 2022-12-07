# LVS Schedulers

ref

https://en.wikipedia.org/wiki/Linux_Virtual_Server

http://www.austintek.com/LVS/LVS-HOWTO/HOWTO/LVS-HOWTO.introduction.html

## Digest

scheduling the algorithm the direcotr uses to select a realserver to service a *new connection request*（注意只对新建连接生效） from a client

schedulers 也被称为调度算法用于 balancing，主要支持以下几种算法

- Round-robin (`ip_vs_rr.c`)
- Weighted round-robin (`ip_vs_wrr.c`)
- Least-connection (`ip_vs_lc.c`)
- Weighted least-connection (`ip_vs_wlc.c`)
- Locality-based least-connection (`ip_vs_lblc.c`)
- Locality-based least-connection with replication (`ip_vs_lblcr.c`)
- Destination hashing (`ip_vs_dh.c`)
- Source hashing (`ip_vs_sh.c`)
- Shortest expected delay (`ip_vs_sed.c`)
- Never queue (`ip_vs_nq.c`)
- Maglev hashing (`ip_vs_mh.c`)

## Schedulers

### Round-robin

### Weightd round-robing

