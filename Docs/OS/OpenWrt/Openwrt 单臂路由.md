# Openwrt 单臂路由

ref

https://www.youtube.com/watch?v=OPG47_wT3XI

https://post.smzdm.com/p/a7dg5pko/

## Digest

单臂路由 router-on-a-stick 这个实际是被用在 VLAN 中的一个术语，通过逻辑口来使用 VLAN 路由

在 Openwrt 中单臂路由指 WAN 和 LAN 同时作用在一个物理端口上

## Topology

https://post.smzdm.com/p/a7dg5pko/

单臂路由的接线方法如图所示。光猫设置为桥接模式，然后将所有有线设备接入一个[交换机](https://www.smzdm.com/fenlei/jiaohuanji/)。此时Wan口和Lan口会占用同一个端口，信号会通过广播的方式发送出去。

[![单网口软路由配置方法，几分钟搞定单臂路由和旁路由](https://qnam.smzdm.com/202108/17/611b5e022ca1a6348.jpg_e1080.jpg)](https://post.smzdm.com/p/a7dg5pko/pic_12/)

我们也可以利用无线[路由器](https://www.smzdm.com/fenlei/luyouqi/)的Lan口来做数据交换。这样就不需要额外的交换机了。将所有有线设备都接入无线路由器的Lan口（所有lan口需要支持千兆）。通过单臂路由进行拨号。

[![单网口软路由配置方法，几分钟搞定单臂路由和旁路由](https://qnam.smzdm.com/202108/17/611b5e02a2f077333.jpg_e1080.jpg)](https://post.smzdm.com/p/a7dg5pko/pic_13/)

## Configuration

 按照上图把线接好之后，在open WRT里选择网络中的接口选项，点击WAN口的修改。

[![单网口软路由配置方法，几分钟搞定单臂路由和旁路由](https://qnam.smzdm.com/202108/17/611b5e0384078970.jpg_e1080.jpg)](https://post.smzdm.com/p/a7dg5pko/pic_14/)

在wan的基本设置里，将协议改为PPPOE拨号模式，并且输入宽带的账号密码。

[![单网口软路由配置方法，几分钟搞定单臂路由和旁路由](https://qnam.smzdm.com/202108/17/611b5e03e6bc92155.jpg_e1080.jpg)](https://post.smzdm.com/p/a7dg5pko/pic_15/)

 在物理设置里，把以太网适配器改到与LAN同一端口。比如我这里Lan在“ETH0”，那就把WAN也选到“ETH0”，然后点击保存设置。

[![单网口软路由配置方法，几分钟搞定单臂路由和旁路由](https://am.zdmimg.com/202108/17/611b5e03e4c0f2021.jpg_e1080.jpg)](https://post.smzdm.com/p/a7dg5pko/pic_16/)

然后在选择“LAN”口设置。在物理接口这里取消“桥接接口”的选项。这一步很关键，它起到了隔绝WAN口和LAN口的作用。点击保存。

[![单网口软路由配置方法，几分钟搞定单臂路由和旁路由](https://qnam.smzdm.com/202108/17/611b5e03e9a293424.jpg_e1080.jpg)](https://post.smzdm.com/p/a7dg5pko/pic_17/)

最后把WAN6的接口也选择到“ETH0”里，点击保存并应用让所有设置生效。

[![单网口软路由配置方法，几分钟搞定单臂路由和旁路由](https://qnam.smzdm.com/202108/17/611b5e042b0ee2074.jpg_e1080.jpg)](https://post.smzdm.com/p/a7dg5pko/pic_18/)

设置完成后，WAN口会自动拨号。拨号成功会它会获得公网的IP地址。

[![单网口软路由配置方法，几分钟搞定单臂路由和旁路由](https://am.zdmimg.com/202108/17/611b5e04b73421288.jpg_e1080.jpg)](https://post.smzdm.com/p/a7dg5pko/pic_19/)