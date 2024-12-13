# NetworkManager概述

参考：

https://wiki.archlinux.org/index.php/NetworkManager_(%E7%AE%80%E4%BD%93%E4%B8%AD%E6%96%87)

NetworkManager是用于检测网络和自动连接网络的程序。无论是有线还是无线，同时提供可视化界面。

同时提供一个daemon(Networkmanager)，一个命令行界面(`nmcli`)和一个基于curses的界面(`nmtui`)

**禁用NetworkManager**

```
systemctl mask NetworkManager
systemctl mask NetworkManager-dispatcher

[root@chz ~]# systemctl mask NetworkManager
Created symlink from /etc/systemd/system/NetworkManager.service to /dev/null.
[root@chz ~]# systemctl start NetworkManager
Failed to start NetworkManager.service: Unit is masked.
```

