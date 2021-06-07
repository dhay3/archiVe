# Linux pstree

以树形结构展示进程和子进程

pattern：`pstree [options] [pid|uid]`

使用`-p`参数进程的pid关系列表

```
[root@cyberpelican \]# pstree | more
systemd-+-ModemManager---2*[{ModemManager}]
        |-NetworkManager---2*[{NetworkManager}]
        |-agetty
        |-apache2---5*[apache2]
        |-blueman-tray---2*[{blueman-tray}]
...

[root@cyberpelican \]# pstree -p 1576
xfconfd(1576)─┬─{xfconfd}(1577)
              └─{xfconfd}(1578)

```

使用`-g`参数展示gid

```
systemd(1)-+-ModemManager(498)-+-{ModemManager}(498)
           |                   `-{ModemManager}(498)
           |-NetworkManager(455)-+-{NetworkManager}(455)
           |                     `-{NetworkManager}(455)
           |-agetty(931)
           |-apache2(19160)-+-apache2(19160)
           |                |-apache2(19160)
           |                |-apache2(19160)
           |                |-apache2(19160)
           |                `-apache2(19160)
           |-blueman-tray(1524)-+-{blueman-tray}(1524)
           |                    `-{blueman-tray}(1524)
           |-colord(1710)-+-{colord}(1710)
```

