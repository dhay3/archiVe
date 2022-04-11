# WAN vs LAN

ref:

https://www.trentonsystems.com/blog/what-is-a-lan-port#:~:text=A%20LAN%20port%2C%20also%20known,other%20devices%20to%20the%20internet

https://en.wikipedia.org/wiki/Local_area_network

https://www.truecable.com/blogs/cable-academy/wan-vs-lan#

https://www.diffen.com/difference/LAN_vs_WAN

## 0x1 Terms

![2022-04-10_21-55](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20220410/2022-04-10_21-55.29s6s7hr7rk0.webp)

![2022-04-10_22-24](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20220410/2022-04-10_22-24.6vuml2tjej40.webp)

### LAN

local area netwrok just a collection of devices connected over a short distance（也可以直接理解为内网）

在LAN中使用最多的是 Ethernet 和 WIFI

### WAN

wide area network just a collection of devices connected over wide distance（可以直接理解为公网，因为internet就是最大的wan）

大陆通常通过ISP来连接internet

## 0x2 Port

### 电口 vs 光口

- 电口 LAN/WAN port：通常是RJ45头，使用 cat 5/6/7 线，多用于家用
- 光口 SFP port：通常是SFP(光模块)，使用fiber（光纤），多用与大型机房

### LAN port vs WAN port

1. 针对家用路由器，通常LAN口常用灰色或黑色表明，WAN口使用黄色或蓝色来区别
2. 在不考虑有硬件损耗（物理硬件会有大概10%的损耗），使用同等速率的物理硬件的情况下。LAN口和WAN理论上的速率有100mbps，1000mbps，10000mbps。但是WAN的速率一般受限于接入 ISP 的带宽（一般来说月费花的越多带宽越大）

## 0x3 Speed

影响网络速率的，有如下几个因素

1. 接入的ISP带宽
2. 通过无线连接的速率比有线的速率要慢
3. 硬件和接入带宽不匹配。例如接入的ISP带宽是1000mbps，使用1000mbps电口，但是使用了100mbps的电缆，所以改网络理论上的最大带宽为100mbps
4. 硬件性能，路由器的CPU和内存