# NetworkManager 配置文件

配置文件可以出现在如下几个地方，按照如下顺序读取，后者会覆盖前者

-  /etc/NetworkManager/NetworkManager.conf,
-  /etc/NetworkManager/conf.d/name.conf,
-  /run/NetworkManager/conf.d/name.conf,
-  /usr/lib/NetworkManager/conf.d/name.conf,
-  /var/lib/NetworkManager/NetworkManager-intern.conf

尽量不要修改全局配置文件`/etc/NetworkManager/NetworkManager.conf`，应该讲其他的配置文件写入到`/etc/NetworkManager/conf.d`下。

==NetworkManager(以下简称NC)可以对用户一些网络配置选项覆盖==，可以通过systemd关闭该service，或者修改被覆盖文件的attr

```
systemctl stop Networkmanager
chattr +i /etc/resolv.conf
```

## 配置格式

必须要要有一个`[mian]`

```
[mian]
plugins=keyfile
```

可以使用如下格式增加或减少，使用`,`分开多个

```
plugins+=another-plugins
plugins-=remove-me
plugins=a,b,c
```

## main

- plugins

  NC加载的plugins，keyfile一定会被加载

- no-auto-default

  指定NIC是否生成默认的无线连接，默认eth0

  ```
  no-auto-default=00:22:68:5c:5d:c4,00:1e:65:ff:aa:ee
  no-auto-default=eth0,eth1
  no-auto-default=*
  ```

- dhcp

  指定DHCP使用的客户端，默认dhclient

- dns

  If the key is unspecified, ==default is used==, unless `/etc/resolv.conf` is a symlink to `/run/systemd/resolve/stub-resolv.conf`, `/run/systemd/resolve/resolv.conf`,
     `/lib/systemd/resolv.conf` or `/usr/lib/systemd/resolv.conf`. In that case, ==systemd-resolved== is chosen automatically.

  有如下几个值：

  1. default：NC会更新`/etc/resolv.conf`使用当前连接网络的DNS配置，在`/etc/NetworkManager/system-connections/xxx.connection`中配置

  2. dnsmasq：使用dnsmasq做为本地的caching DNS server

     ```
     [main]
     dns=dnsmasq
     ```

     当使用`nmcli general reload`会重新读取NetworkManager配置，然后启动dnsmasq

  3. systemd-resolved：NC会将DNS配置推送给systemd-resolve

  4. none：NetworkManager不会修改`resolv.conf`配置。同时也会配置rc-manager=umanaged

- rc-manager

  告诉dns该如何生成文件，默认值有NetworkManager决定

  有如下几个值

  symilnk：如果`/etc/resolv.conf`是普通文件跟新该文件，如果是链接文件不不做任何操作

  file：NC会将`/etc/resolv.conf`写成普通文件，如果是链接更新指向的文件

  unmanaged：不修改`/etc/resolv.conf`

- systemd-resolved

  是否将DNS配置发送到systemd，默认true

## ifupdonw

- managed

  是否将`/etc/network/interfaces`中的NIC被NetworkManager管理，true or false

## device

针对单独NIC的配置



## 例子

