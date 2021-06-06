# cisco command cheat sheet

参考：

https://www.netwrix.com/cisco_commands_cheat_sheet.html

https://www.pcwdld.com/cisco-commands-cheat-sheet

## featrue

- 在Cisco的设备中如果不清楚该如何使用命令，可以在命令后跟上`?`就会提示可以使用的参数和参数的含义。==如果显示的内容过多也可以通过`/keywords`来过滤==
- 和linux一样支持tab补全，==且支持缩写==
- 在`?`中大写显示的选项表示用户需要填写的内容，`<cr>`表示carriage return
- 命令的前面添加`no`来表示取消配置
- 配置完成后可以通过write命令加配置写入到RAM

## Modes

cisco router有六种模式，通过如下命令可以在指定mode下切换到指定mode。==每种模式下可以使用的命令不同==。

| **Command**                         | **mode**   | **Description**                                              |
| ----------------------------------- | ---------- | ------------------------------------------------------------ |
| enable                              | U          | Moves from User to ==Privileged mode(P)==.                   |
| logout                              | U          | Exit User mode(U).                                           |
| `configure <terminal>`              | P          | Moves from Privileged to ==Configure mode(G)==.              |
| disable                             | P          | Exit ==User mode(U)==.                                       |
| `Interface <interface description>` | G          | Enter ==interface configuration mode(I)==.                   |
| vlan vlan-id                        | G          | Moves to configure ==vlan mode(V)==.                         |
| Vlan database                       | P          | if U wanna move to vlan mode you should use this command before it. |
| line                                | G          | Enter ==line mode(L)== from Global configuration mode.       |
| exit/(end \|ctrl+z)                 | G, R, L, V | return to previous mode.end will back to the User mode       |

==如果想要在Configuration mode中使用Privileged mode中的命令可以在命令前加do==，例如：`R2(config-router)#do sh ip int br`

## Prompts

命令行prompt可以清楚区别当前所处的mode

| **Prompt**             | **PresentMode** | **Description**                                              |
| ---------------------- | --------------- | ------------------------------------------------------------ |
| Router>                | U               | User EXEC mode, is the first level of access.                |
| Router#                | P               | Privileged EXEC mode. The second level of access, accessible with the “enable” command. |
| Router(config)#        | G               | Configuration mode. Accessible only via the privileged EXEC mode. |
| Router(config-if)#     | I               | Interface mode. Level accessible via configuration mode.     |
| Router(config-router)# | R               | Routing mode. Level within configuration mode.               |
| Router(config-line)#   | L               | Line level (vty, tty, async). Accessed via the configuration mode |
| Router(config-vlan)#   | V               | Config-vlan, accessible via the global configuration mode.   |
| Switch(vlan)#          | VD              | Vlan database, accessible from the privileged EXEC mode.     |

## Iface

interface表示一个接口

- f代表fastethernet，快速以太网接口
- s代表serial，串行接口

例如表示`f0/0`，fastethernet 第一张NIC，第一个端口

## pipeline



## Basic configuration commands

mode表示该命令只能在指定的mode中执行

| **Command**                                    | **Mode** | **Description**                                              |
| ---------------------------------------------- | -------- | ------------------------------------------------------------ |
| show version                                   | U,P      | Display information about IOS and router.                    |
| show interfaces                                | U,P      | Display physical attributes of the router’s interfaces.      |
| show ip route                                  | U,P      | Display the current state of the routing table.              |
| show access-lists                              | P        | Display current configured ACLs and their contents.          |
| show ip interface brief                        | P        | Displays a summary of the status for each interface.         |
| show running-config                            | P        | Display the current configuration.                           |
| show startup-config                            | P        | Display the configuration at startup.                        |
| enable                                         | U        | ==Acces Privilege mode==                                     |
| configure terminal                             | P        | ==Access Configuration mode.==                               |
| `interface <iface>`                            | G        | ==Enter interface configuration.==                           |
| `ip address <ip address> <mask>`               | I        | Assign an IP address to the specified interface.             |
| shutdown/no shutdown                           | I        | Turn off or turn on an interface. Use both to reset.         |
| `description <name-string>`                    | I        | Set a description to the interface.                          |
| `show ip interface <type number>`              | U,P      | Displays the usability status of the protocols for the interfaces. |
| `show running-config interface  <slot/number>` | P        | Displays the running configuration for a specific interface. |
| `hostname <name>`                              | G        | Set a hostname for the Cisco device.                         |
| `enable secret <password>`                     | G        | Set an “enable” secret password.                             |
| copy running-config startup-config             | P        | Saves the current (running) configuration in the startup configuration into the NVRAM. The command saves the configuration so when the device reloads, it loads the latest configuration file. |
| copy startup-config running-config             | P        | It saves (overwrites) the startup configuration into the running configuration. |
| copy from-location to-location                 | P        | It copies a file (or set of files) from a location to another location. |
| erase nvram                                    | G        | Delete the current startup configuration files. ==The command returns the device to its factory default.== |
| reload                                         | G        | Reboot the device. The NVRAM will take the latest configuration. |
| erase startup-config                           | G        | Erase the NVRAM filesystem. The command achieves the similar outcome as “erase nvram” |

## Network Access

配置VLANs和trunks，以及在二层的协议配置。有些选项并不是所有的Cisco设备都有的。

==`switchport`命令只针对一些iface有效，可以通过`sh int <iface> switchport`来查看。==

可以使用virtual trunking protocol(VPT)，把一台交换机配置成VTP Server, 其余交换机配置成VTP Client,这样他们可以自动学习到server 上的VLAN 信息

| **Command**                                           | **Mode** | **Description**                                              |
| ----------------------------------------------------- | -------- | ------------------------------------------------------------ |
| cdp run/no cdp run                                    | P        | The “cdp run” command enables Cisco Discovery Protocol. The “no cdp run” disables it. |
| show cdp                                              | P        | Display global information for CDP.                          |
| show cdp neighbors                                    | P        | Display all CDP neighbors.                                   |
| lldp run/no lldp run                                  | P        | The “lldp run” command enables the LLDP Protocol. The “no lldp run” disables it. |
| show lldp                                             | P        | Displays global information for LLDP                         |
| show lldp neighbors                                   | P        | Show all LLDP neighbors.                                     |
| show mac-address-table                                | P        | Display all the MAC address entries in a table.              |
| spanning-tree mode rapid-pvst                         | G        | A global configuration command that configures the device for Rapid Per VLAN Spanning Tree protocol. |
| `spanning-tree vlan <1-4094> priority <0-61440>`      | G        | Manually set the bridge priority per vlan.                   |
| `spanning-tree vlan <1-4094> root primary`            | G        | Make the switch the root of the SP.                          |
| `no spanning-tree vlan <1-4094>`                      | G        | Disable SP on the specific VLAN.                             |
| show spanning-tree summary                            | P        | Show a summary of all SP instances and ports.                |
| show spanning-tree detail                             | P        | Show detailed information of each port in the spanning-tree process. |
| show vlans                                            | P        | Lists each VLAN and all interfaces assigned to that VLAN. The output does not include trunks. |
| show vlan brief                                       | P        | Displays vlan information in brief                           |
| show interfaces switchport                            | P        | Display configuration settings about all the switch port interfaces. |
| show interfaces trunk                                 | P        | Display information about the operational trunks along with their VLANs. |
| `vlan <1-4094>`                                       | G        | ==Enter VLAN configuration mode== and create a VLAN with an associated number ID. |
| `name <name>`                                         | V        | Within the VLAN configuration mode, assign a name to the VLAN |
| switchport mode access                                | I        | In the interface configuration mode, the command assigns the interface link type as an access link. |
| `switchport access vlan <vlan-number>`                | I        | Assign this interface to specific VLAN. vlan-number should in the vlan datatbase |
| `interface range <fnic-number/firstport - lastport> ` | G        | Use interface range configuration mode for batch command     |
| `channel-group <number>`                              | I        | Assign the Etherchannel. Set the interface range to a channel group. |
| no switchport access vlan                             | I        | Remove VLAN assignment from interface. It returns to default VLAN 1 |
| show vtp status                                       | P        | Display all vtp status                                       |
| `vtp mode <server | client | transparent>`            | G        | In the global configuration mode, set the device as server, client, or transparent vtp mode. |
| switchport mode trunk                                 | I        | An interface configuration mode. Set the interface link type as a trunk link. |
| `switchport trunk native vlan <vlan-number> `         | I        | Set native VLAN to a specific number.                        |
| `switchport trunk allowed vlan <vlan-numbers>`        | I        | Allow specific VLANs on this trunk.                          |
| switchport trunk encapsulation dot1q                  | I        | Sets the 802.1Q encapsulation on the trunk link.             |

## IP connectivity

用于配置route

| **Command**                                                  | **Mode** | **Description**                                              |
| ------------------------------------------------------------ | -------- | ------------------------------------------------------------ |
| Show ip route                                                | P        | Show the routing table.                                      |
| Show ip route ospf                                           | P        | Show routes created by the OSPF protocol.                    |
| ip default-gateway <ip_address>                              | G        | Set the default gateway for the router.                      |
| `ip route <destation-prefix> <mask> <next hop>`              | G        | Create a new static route                                    |
| `no ip route <destion-prefix> <destion-prefix-netmask> <next hop>` | G        | ==Remove a specific static route.==                          |
| `ip route <destion-prefix> <destion-prefix-netmask> <next hop>` | G        | ==Configure a default route. prefix stands for the left part of CIDR== |
| `router ospf <process ID>`                                   | G        | Enable OSPF with an ID. The command will open the router configuration mode. |
| show ip ospf interface                                       | P        | Display all the active OSPF interfaces                       |

## IP services

用于配置NAT，DHCP和DNS

| **Command**                                                  | **Mode** | **Description**                                              |
| ------------------------------------------------------------ | -------- | ------------------------------------------------------------ |
| `ip nat <inside | outside>`                                  | I        | Specific whether the interface is the inside or outside of NAT. |
| `ip nat inside source <ACL No.> <pool | static IP> <overload>` | G        | Configure dynamic NAT. It instructs the router to translate all addresses identified by the ACL on the pool. To configure Port Address Translation (PAT) use the “overload” at the end. |
| `ip nat inside source static <local IP> <global IP>`         | G        | Create a static NAT from inside (local IP) to outside (global IP) |
| `ip nat outside source static <ACL No.> <pool | static IP>`  | G        | Create a static NAT from outside (ACL) to inside (IP pool)   |
| `ntp peer <ip-address>`                                      | G        | Configure the time by synchronizing it from an NTP server.   |
| `ip dhcp excluded-address <first-ip-address> <last-ip-address>` | G        | The IP addresses that the DHCP server should not assign to the DHCP client. |
| `ip dhcp pool <name>`                                        | G        | Enters the ==DHCP pool configuration mode== and creates a new DHCP pool. |
| `network <network ID> <mask>`                                | G – DHCP | Inside the DHCP configuration mode. Define the address pool for the DHCP server. |
| `default-router <IP address>`                                | G – DHCP | Set the default gateway IP address for the DHCP clients.     |
| `dns-server <IP address>`                                    | G – DHCP | Set the DNS server IP address for the DHCP clients.          |
| `ip helper-address <ip address>`                             | I        | Turns an interface into a DHCP bridge. The interface redirects DHCP broadcast packets to a specific IP. |
| show ip dhcp pool                                            | P        | Display information about the DHCP pool                      |
| show ip dhcp binding                                         | P        | Display information about all the current DHCP bindings.     |
| ip dns server                                                | G        | Enable DNS service.                                          |
| ip domain-lookup                                             | G        | Enable domain lookup service. DNS client                     |
| ip name-server <IP address \| domain name>                   | G        | Set a public DNS server.                                     |
| `snmp-server community <community-string> ro`                | G        | Enable SNMP Read-Only public community strings.              |
| `snmp-server community <community-string> rw`                | G        | Enable SNMP Read-Only private community strings.             |
| `snmp-server host <ip-address> version <community-string>`   | G        | Specific the hosts to receive the SNMP traps                 |
| `logging <ip address>`                                       | G        | Determines the Syslog server to send log messages.           |
| logging trap level                                           | G        | Limit Syslog messages based on severity level                |
| show logging                                                 | P        | Shows the state logging (syslog). Shows the errors, events, and host addresses. It also shows SNMP configuration and activity. |
| terminal monitor                                             | P        | Enables debug and system’s error messages for the current terminal. |
| sh ip ssh                                                    | P        | Verify SSH access into the device.                           |

## security

配置ACLs，port security，basic AAA configuration

==wildcard-bits表示net-mask子网掩码的反码，选项host表示net-mask 255.255.255.255==

使用`no access-list <access-list-number>`来取消指定acl的配置

| **Command**                                                  | **Mode** | **Description**                                              |
| ------------------------------------------------------------ | -------- | ------------------------------------------------------------ |
| `enable secret <password>`                                   | G        | Set an “enable” secret password. Enable secret passwords are hashed via the ==MD5 algorithm==. |
| line vty 0 4                                                 | G        | A global configuration command to access the ==virtual terminal configuration==. VTY is a virtual port used to access the device via SSH and Telnet. 0 4 to allow five simultaneous virtual connections |
| line console 0                                               | G        | A global configuration command to access the ==console configuration==. |
| `password <password>`                                        | L        | ==Once in line mode, set a password for those remote sessions== with the “password” command. |
| Login local                                                  | L        | The authentication uses only locally configured credentials. |
| `username <username> privilege <level> secret <password>`    | G        | Require a username with a specific password. Also configure different levels of privilege. |
| service password-encryption                                  | G        | Makes the device encrypt all passwords saved on the configuration file. |
| crypto key generate rsa                                      | G        | Generate a set of RSA key pairs for your device. These keys may be used for remote access via SSH. |
| access-list                                                  | G        | Defined a numbered ACL                                       |
| ip access-list                                               | G        | Defined an IPv4 ACL.(config ACL interactivly)                |
| `access-list <1-99> <deny | permit>  <host> <wildcard-bits>[log]` | G        | Create a standard ACL.                                       |
| `access-list <100-199> <deny | permit>  <protocol-number> <source> <source-wildcard-bits> <destination> <destination-wildcard-bits> [Options] ` | G        | Create an extended ACL.                                      |
| `ip access-class <access-list-name> <in | out> no ip access-group <access-list-name> <in | out>` | L        | A line configuration command mode. It restricts incoming and outgoing connections to a particular vty line. Use “no” to remove the restriction. |
| show ip access-list                                          | P        | Show all IPv4 ACLs                                           |
| switchport mode access                                       | I        | From the interface configuration mode, this command assigns the interface link type as an access link. |
| switchport port-security                                     | I        | enable dynamic port security on the specific interface.      |
| `switchport port-security maximum <max value>`               | I        | Specify the maximum number of secure MAC addresses on the specific interface. |
| `switchport port-security mac-address <mac-address | sticky [mac-address]>` | I        | Force a specific mac-address to the interface. Also use the “sticky” option to make the interface remember the first mac-address connected to the interface. |
| `switchport port-security violation <shutdown | restrict | protect> ` | I        | Define the action to be taken when a violation is detected on the port. |
| show port security                                           | P        | Display the port security configuration on each interface.   |

## troubleshooting commands

| **Command**                                                  | **Mode** | **Description**                                              |
| ------------------------------------------------------------ | -------- | ------------------------------------------------------------ |
| ping <target IP \| hostname> <repeat Count [5]> <source [IP \| interface] | P        | Diagnose connectivity with extended ping. Check reachability, RRTs, and packet loss. |
| traceroute <target IP \| hostname><source [IP \| interface]  | P        | Use traceroute to diagnose connectivity on a hop by hop basis. |
| telnet                                                       | P        | Use Telnet to check for listening ports (1 to 65535) on a remote device. |
| show interface                                               | P        | Use this command to discover the physical attributes; find duplex, link types, and speed mismatches. Both ends must match. Also use this command to find errors. |
| speed <10 \| 100 \| 1000 \| auto>                            | I        | Set the speed of an interface. Or configure it as auto.      |
| duplex <auto \| full \| half>                                | I        | Set the interface duplex.                                    |
| show interface \| include fastethernet \| input errors       | P        | This command searches across all interfaces and outputs the ones that include input errors. |
| show ip interface                                            | P        | Use this command to discover the status for all the protocols on that interface. |
| shutdown/no shutdown                                         | I        | Interface configuration mode. Restart an interface           |
| show ip route                                                | P        | This command is useful for determining the route of ip packets. |
| show cdp neighbors                                           | P        | Discover basic information about neighboring Cisco’s routers and switches |
| show mac address-table                                       | P        | Display the contents of the mac-address table.               |
| Show vlanShow vlan brief                                     | P        | Find vlan status and interfaces assigned to the vlans.       |
| show vtp status                                              | P        | Use this command to discover the current VTP mode of the device. |
| show interfaces trunk                                        | P        | Check the allowed VLANs on both ends of the trunk.           |
| show ip flow top-talkers                                     | P        | If Netflow is enabled, this command is very useful to troubleshoot top talk |

## sh int VS sh ip int

show interface 和 show ip interface的区别

show interface会显示OSI layer的详细信息，但是show ip  interface只会显示OSI layer 3层的信息

```
router-02#show ip int f3/0 
FastEthernet3/0 is up, line protocol is down
  Internet protocol processing disabled

router-02#sh int f3/0
FastEthernet3/0 is up, line protocol is down 
  Hardware is Fast Ethernet, address is c402.1d0e.f300 (bia c402.1d0e.f300)
  MTU 1500 bytes, BW 100000 Kbit/sec, DLY 100 usec, 
     reliability 255/255, txload 1/255, rxload 1/255
  Encapsulation ARPA, loopback not set
  Keepalive set (10 sec)
  Auto-duplex, Auto-speed
  ARP type: ARPA, ARP Timeout 04:00:00
  Last input never, output never, output hang never
  Last clearing of "show interface" counters never
  Input queue: 0/75/0/0 (size/max/drops/flushes); Total output drops: 0
  Queueing strategy: fifo
  Output queue: 0/40 (size/max)
  5 minute input rate 0 bits/sec, 0 packets/sec
  5 minute output rate 0 bits/sec, 0 packets/sec
     0 packets input, 0 bytes, 0 no buffer
     Received 0 broadcasts, 0 runts, 0 giants, 0 throttles
     0 input errors, 0 CRC, 0 frame, 0 overrun, 0 ignored
     0 input packets with dribble condition detected
     0 packets output, 0 bytes, 0 underruns
     0 output errors, 0 collisions, 2 interface resets
     0 unknown protocol drops
     0 babbles, 0 late collision, 0 deferred
     0 lost carrier, 0 no carrier
     0 output buffer failures, 0 output buffers swapped out
```

