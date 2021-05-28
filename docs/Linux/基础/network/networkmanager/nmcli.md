# nmcli

参考：

https://wiki.archlinux.org/index.php/NetworkManager_(%E7%AE%80%E4%BD%93%E4%B8%AD%E6%96%87)

nmcli用于管理NetworkManager，同时也可以展示网络情况。

syntax：`nmcli [options] [command]`

## options

- `-f {field1,field2... | common | all}`

  显示指定的列，如果忽略使用common

  ```bash
  cpl in ~ λ nmcli -f wifi general status
  WIFI    
  enabled
  ```

  默认使用common

- -p | --pretty

  格式化输出

## command

### general command

- status

  如果没有command。默认使用status

  ```bash
  root in ~ λ nmcli general staus
  eth0: connected to Ifupdown (eth0)
          "Intel 82545EM"
          ethernet (e1000), 00:50:56:32:1B:8C, hw, mtu 1500
          ip4 default
          inet4 192.168.80.200/24
          route4 0.0.0.0/0
          route4 192.168.80.0/24
          inet6 fe80::250:56ff:fe32:1b8c/64
          route6 ff00::/8
          route6 fe80::/64
  ```

- hostname [newhostname]

  修改hostname。等价于修改`/etc/hostname`

- permissions

  ```bash
  root in ~ λ nmcli general permissions 
  PERMISSION                                                        VALUE 
  org.freedesktop.NetworkManager.checkpoint-rollback                yes   
  org.freedesktop.NetworkManager.enable-disable-connectivity-check  yes   
  org.freedesktop.NetworkManager.enable-disable-network             yes   
  org.freedesktop.NetworkManager.enable-disable-statistics          yes 
  ```

### networking command

- ==on | off==

  打开或关闭由NetworkManager管理的网络

  ```
  root in ~ λ nmcli networking off
  root in ~ λ ip a
  1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
      link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
      inet 127.0.0.1/8 scope host lo
         valid_lft forever preferred_lft forever
      inet6 ::1/128 scope host 
         valid_lft forever preferred_lft forever
  2: eth0: <BROADCAST,MULTICAST> mtu 1500 qdisc pfifo_fast state DOWN group default qlen 1000
      link/ether 00:50:56:32:1b:8c brd ff:ff:ff:ff:ff:ff
  3: eth1: <BROADCAST,MULTICAST> mtu 1500 qdisc pfifo_fast state DOWN group default qlen 1000
      link/ether 00:0c:29:a0:ef:a3 brd ff:ff:ff:ff:ff:ff
  4: eth2: <BROADCAST,MULTICAST> mtu 1500 qdisc pfifo_fast state DOWN group default qlen 1000
      link/ether 00:50:56:2b:2f:5c brd ff:ff:ff:ff:ff:ff
  5: br-74480271fc99: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
      link/ether 02:42:2c:72:62:e7 brd ff:ff:ff:ff:ff:ff
  6: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default 
      link/ether 02:42:f9:11:59:a6 brd ff:ff:ff:ff:ff:ff
  8: vethca9d8e2@if7: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master br-74480271fc99 state UP group default 
      link/ether 5e:d6:b9:90:4b:8f brd ff:ff:ff:ff:ff:ff link-netnsid 0
      inet6 fe80::5cd6:b9ff:fe90:4b8f/64 scope link 
         valid_lft forever preferred_lft forever
  ```

### radio command

- wifi [on | off]

  打开或关闭wifi探测，如果没有指定on或off，显示wifi的状态。

  ```
  root in ~ λnmcli radio wifi 
  enabled
  root in ~ λ nmcli radio wifi off
  root in ~ λ nmcli radio wifi    
  disabled
  ```

### connect command

- show

  显示连接的网络的UUID

  ```
  root in ~ λ nmcli connection show      
  NAME                UUID                                  TYPE      DEVICE 
  Ifupdown (eth0)     681b428f-beaf-8932-dce4-687ed5bae28e  ethernet  eth0   
  Ifupdown (eth1)     7b635ed6-2640-7ad8-675d-744db12dd9fa  ethernet  eth1   
  Wired connection 1  c3babedb-da1c-34a4-97ac-3efaa19c6e18  ethernet  -- 
  ```

- up | down

  打开或关闭指定uuid的网络

  ```
  root in ~ λ nmcli connection down uuid 681b428f-beaf-8932-dce4-687ed5bae28e
  Connection 'Ifupdown (eth0)' successfully deactivated (D-Bus active path: /org/freedesktop/NetworkManager/ActiveConnection/5)
  root in ~ λ nmcli connection show
  NAME                UUID                                  TYPE      DEVICE 
  Ifupdown (eth1)     7b635ed6-2640-7ad8-675d-744db12dd9fa  ethernet  eth1   
  Ifupdown (eth0)     681b428f-beaf-8932-dce4-687ed5bae28e  ethernet  --     
  Wired connection 1  c3babedb-da1c-34a4-97ac-3efaa19c6e18  ethernet  --   
  ```
  
- add

  添加一个iface，至少需要`connection.type`。具体使用查看`nm-settings`

  ```bash
  #nmcli connection add [property value...]
  cpl in ~ λ nmcli connection add type bridge ifname br0
  Connection 'bridge-br0' (e80f97c3-5417-4d63-be59-97693d8d354f) successfully added.
  ```

  ifname表示创建iface的name

- delete

  删除指定链接

  ```bash
  cpl in ~ λ nmcli connection delete bridge-br0 
  Connection 'bridge-br0' (e80f97c3-5417-4d63-be59-97693d8d354f) successfully deleted.
  ```

### ==device command==

- status

  缺省值，展示当前网络的状态

  ```
  root in ~ λ nmcli device
  DEVICE           TYPE      STATE        CONNECTION      
  eth0             ethernet  connected    Ifupdown (eth0) 
  eth1             ethernet  connected    Ifupdown (eth1) 
  eth2             ethernet  unavailable  --              
  br-74480271fc99  bridge    unmanaged    --              
  docker0          bridge    unmanaged    --              
  vethca9d8e2      ethernet  unmanaged    --              
  lo               loopback  unmanaged    --    
  ```

- show

  展示详细信息

  ```
  root in ~ λ nmcli device show eth0
  GENERAL.DEVICE:                         eth0
  GENERAL.TYPE:                           ethernet
  GENERAL.HWADDR:                         00:50:56:32:1B:8C
  GENERAL.MTU:                            1500
  GENERAL.STATE:                          100 (connected)
  GENERAL.CONNECTION:                     Ifupdown (eth0)
  GENERAL.CON-PATH:                       /org/freedesktop/NetworkManager/ActiveConnection/7
  WIRED-PROPERTIES.CARRIER:               on
  IP4.ADDRESS[1]:                         192.168.80.200/24
  IP4.GATEWAY:                            192.168.80.2
  IP4.ROUTE[1]:                           dst = 192.168.80.0/24, nh = 0.0.0.0, mt = 102
  IP4.ROUTE[2]:                           dst = 0.0.0.0/0, nh = 192.168.80.2, mt = 102
  IP4.DNS[1]:                             8.8.8.8
  IP6.ADDRESS[1]:                         fe80::67db:b119:4edf:4043/64
  IP6.GATEWAY:                            --
  IP6.ROUTE[1]:                           dst = fe80::/64, nh = ::, mt = 102
  IP6.ROUTE[2]:                           dst = ff00::/8, nh = ::, mt = 256, table=255
  
  ```

- connect

  连接指定iface

- disconnect

  断开指定的iface

- wifi list

  显示所有可用的wifi，`*`表示当前连接的网络















