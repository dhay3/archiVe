# switchport

reference：

https://ipwithease.com/switchport-trunk-mode-vs-access-mode/

switchport 顾名思义就是交换机的端口

在cisco switches中大多数都被配置成 dynamic desirable mode。意味着如果是交换机互联就会从access切换成trunk

## Trunk ports

通常被用在switch to switch 或 switch to router

trunk carry multiple vlans across devices and maintain VLAN

tags in Ethernet frames for receiving directly connected device differentiates between different Vlans

使用`switchport mode trunk`强制将端口该为trunk mode

## Access ports

Access ports 只在一个VLAN中，通常被用在PC，laptop和printer

使用`switchport mode access`强制将端口改为access mode

## Trunk mode VS Access mode

| parameter             | trunk mode                                                   | access mode                                                  |
| --------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| Terminology           | A trunk port can carry traffic in one or more VLANs on the same physical link. Trunked ports differentiate Vlans by either adding a tag to the packet (802.1Q) or encapsulation the packet (ISL). | Access ports are part of only one VLAN and normally used for terminating end devices likes PC, Laptop and printer. |
| Default Behavior      | By default, a trunk interface can carry traffic for all VLANs. | By default, an access port carries only one VLAN             |
| Configuration         | To designate a port to a trunk mode -“Switchport mode trunk” | To designate a port to access mode -“Switchport mode access” |
| Use case              | Switch to Switch connectivity<br/>Switch to Router (When using Router on a Stick or dot1q trunk)<br/>Switch to Server (specific cases only especially in VM technology) | Switch to PC/laptop<br/>Switch to Printer<br/>Switch to Router<br/>Note - This is typical standard procedure when such switch port serves end users such as PC, printer, or servers. |
| VLAN Tags             | Port configured in Trunk mode will carry VLAN tags           | Port configured in Access mode will not carry VLAN tags (stripped of VLAN tags) |
| Verification commands | Show Vlan brief<br/>show interface x/x switchport            | Show interface trunk<br/>show interface x/x switchport       |









