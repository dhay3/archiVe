# Linux arp

- arp 

  查看主机的ARP table entry，使用`-n`参数以数字形式显示IP

  ```
  [root@chz ~]# arp 
  Address                  HWtype  HWaddress           Flags Mask            Iface
  gateway                  ether   00:50:56:e3:f8:b0   C                     ens33
  [root@chz ~]# arp -n
  Address                  HWtype  HWaddress           Flags Mask            Iface
  192.168.80.2             ether   00:50:56:e3:f8:b0   C                     ens33
  
  ```

- arp -d 

  删除指定IP的entry

  ```
  [root@chz ~]# arp -d 192.168.80.2
  [root@chz ~]# arp 
  Address                  HWtype  HWaddress           Flags Mask            Iface
  192.168.80.254           ether   00:50:56:eb:9a:f6   C                     ens33
  192.168.80.254           ether   00:50:56:eb:9a:f6   C                     ens33
  ```

- arp -s

  ```
  [root@chz ~]# arp -s 192.168.80.200 00:50:56:eb:9a:f6
  [root@chz ~]# arp
  Address                  HWtype  HWaddress           Flags Mask            Iface
  gateway                  ether   00:50:56:e3:f8:b0   C                     ens33
  192.168.80.200           ether   00:50:56:eb:9a:f6   CM                    ens33
  192.168.80.254           ether   00:50:56:eb:9a:f6   C                     ens33
  ```

- arp -a

  以BSD格式显示

  ```
  [root@chz ~]# arp -a
  gateway (192.168.80.2) at 00:50:56:e3:f8:b0 [ether] on ens33
  ? (192.168.80.200) at 00:50:56:eb:9a:f6 [ether] PERM on ens33
  ? (192.168.80.254) at 00:50:56:eb:9a:f6 [ether] on ens33
  ```

  

