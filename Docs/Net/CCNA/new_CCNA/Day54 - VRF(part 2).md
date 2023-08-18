# Day54 - VRF(part 2)

*Virtual Routing & Forwarding is used to divide single router into multiple virtual routers*

> Similar to how VLANs are used to divide a single switch(LAN) into multiple virtual switches(VLANs)
>
> 类似与 VLAN 将 LAN 逻辑上划分成多个不同的 LAN

例如下图中黑框代表一个实际的 Router R1，可以将其划分成 3 个虚拟的 Router VRF1/2/3

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-17_17-34.4j3mebjwnoc0.webp)

假设没有 VRF，如果从 G0/0 收到报文，R1 可以将报文转发到其他所有的端口。而如果使用了 VRF，从 G0/0 收到报文只能转发到 G1/0，从那个 VRF 收到的报文，只能转发到那个 VRF 里的其他端口

It does this by allowing a router to build multiple separate routing tables

> 通过将划分路由表实现

- Interfaces(Layer 3 only) & routes are configured to be in a specific VRF(aka VRF instance)

- Router interfaces, SVIs & routed ports on multilayer switches can be configured in a VRF

- Traffic in one VRF cannot be forwarded out of an interface in another VRF

  *As an exception, VRF Leaking can be configured to allow traffic to pass between VRF’s*

- VRF is commonly used to facilitate MPLS

  VRF 通常和 MPLS 一起使用

VRF is commonly used by service providers to allow one device to carry traffic from multiple customers

- Each customer’s traffic is isolated from the others

- Customer IP addresses can overlap without issues

  在不同的 VRF IP 地址可以配置成相同的，不会影响正常的转发和使用

  ![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-17_17-55.7amj3btq8e00.webp)

## VRF Configuration

假设拓扑如下

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-17_17-59.f5i9e2i3zm0.webp)

先看一下不使用 VRF 的情况

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-17_18-01.164yb8d8g7ek.webp)

在不使用 VRF 的情况下，如果想要将 Router G0/2 配置成 G0/0 同段的 IP 是不支持的。因为逻辑上 Router 的两个接口配置成同段的没有意义，正常应该使用 Switch 来连接同段的地址，而不是用 Router 的接口来划分

想要配置 VRF 很简单

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-18_13-43.2qrcrzh8ok80.webp)

- `R1(config)#ip vrf <vrf-name>`

  创建 VRF

- `R1(config-if)#ip vrf forwarding <vrf-name>`

  将接口划分到 VRF

  > 如果接口之前有配置 IP address，在配置 VRF 后会将 IP address 移除，这一点可以从 Syslog 的输出中得出
  >
  > 所以使用该命令后，还需要为接口配置 IP address，才可以正常使用

- `R1#show ip vrf`

  查看 VRF 和 接口之间的关系

在配置完 VRF 后，如果使用 `show ip route` 来查看 route，这时是不会显示的，因为 `show ip route` 是查看 global routing table 的，但是所有的接口 IP address 都配置在 VRF 中了，所以需要使用 `show ip route vrf <vrf-name>` 来查看指定 VRF 中的 routing table

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-18_13-46.4pomi0xn03g0.webp)

校验一下配置，如果直接使用 `ping 192.168.1.2` 是不会成功的，因为 global routing table 中现在没有对应的路由。应该使用 `ping vrf <vrf-name> <address>` 

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230816/2023-08-18_13-49.4wd9r9xp0a00.webp)

**references**

1. ^https://www.youtube.com/watch?v=Ge4644KUvh4&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=107