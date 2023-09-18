# Day55 - Wireless Fundamentals

## Wireless Networks

> Wirelss LAN are defined in ==802.11==

- All devices within range receive all frames, like devices connected to an Ethernet Hub

  > 和 Hub 一样，在 wireless 中的所有设备可以接受所有的报文，所以安全系数不高

  - Privacy of data within the LAN is a greater concern

  - CSMA/CA(Carrier Sense Multiple Access with Collision Avoidance) is used to facilitate half-duplex communications

    > CSMA/CD is used in wired networks to detect and recover from collisions
    >
    > CSMA/CA is used in wireless networks to avoid collisions

- Wireless communications are regulated by varous international and national bodies

- Wireless signal coverage area must be considered

  - signal range

  - signal absorption,reflection,refraction,diffraction and scattering

    这些都是导致无线信号衰减的因子

- Other devices using the same channels can cause interference(干扰)

  例如邻居加的 WIFI 信号

### Absorption

吸收

例如无线设备和终端设备中间隔了一堵墙，墙会吸收部分无线信号

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-18_14-12.83rk3jb4ihk.webp)

### Reflection

反射

无线信号通常会被金属反射一部分，这也是在电梯里无线信号弱的原因

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-18_14-15.574q62dn1aw0.webp)

### Refraction

折射

和光的折射一样，无线信号经过不同的传输介质发生信号偏移

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-18_14-17.1sjo2mstzew0.webp)

### Diffraction

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-18_14-19.37h2gpx4ieo0.webp)

### Scattering

漫反射

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-18_14-20.61r87y9jhk80.webp)

## Radio Frequencey Bands

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-18_14-28.6wnyx7k015o0.webp)

通常 WIFI 会使用 2.4GHz, 5 GHz band, 一个 band 会划分成多个 channel

- In the 2.4 GHz band, it is recommended to use channels 1,6 and 11

  例如 3 个 AP，分别使用 1，6，11 channels

  ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-18_14-36.l4jw7le6onk.webp)

  > 2.4 GHz band 是存在 overlapping 的

- The 5 GHz band consists of non-overlapping channels, so it is much easier to avoid interference between adjacent APs

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-18_14-38.2yuex80e2bc0.webp)

## Service sets

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-18_14-41.44txocg7fe20.webp)

### IBSS

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-18_14-44.4b9wzz16clc0.webp)

### BSS

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-18_14-47.2zqihvc3nts0.webp)

### ESS

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-18_14-50.3urzcdkysr80.webp)

### MBSS

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-18_14-52.19fela3vr4tc.webp)

## AP

### repeater

repeater 中文也叫做 中继

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-18_14-56.5eq9cu764zk0.webp)

### WGB

和 repeater 的区别就在于，PC1 是有线连接到 WGB 的

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-18_14-58.342h1a1suy4.webp)

### Outdoor bridge

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-18_15-00.3m7e8oca07c0.webp)

**references**

1. ^https://www.youtube.com/watch?v=zuYiktLqNYQ&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=109