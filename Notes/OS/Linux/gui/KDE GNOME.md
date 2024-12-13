# KDE/GNOME

> QT开发的软件可以在GNOME上运行，GTK开发的软件可以在KDE上运行

GTK是一款开源跨平台的GUI widget toolkit，常被用于GNOME的软件开发，可以理解为是SDK。常见的软件有`GNOME SHELL`

## QT

读作cute，和GTK一样一款开源跨平台的GUI widget toolkit。可以理解为是SDK。常见的软件有`KDE Plasma`

下面展示的是Display，GUI interface toolkit，Application之间的关系

![2021-07-05_22-26](https://github.com/dhay3/image-repo/raw/master/20210601/2021-07-05_22-26.amg7gf2tu0s.png)

## KDE

Kool Desktop Environment也被称为KDE，有很多的linux distro都用KDE Plasma 5作为默认的desktop enviroment，例如Kubuntu，Manjaro等等。也有一些其他的distro用gnome例如CentOS，ubuntu等等

KDE has different versions

- kde-full

  完整的KDE pakcage，包含plasma和其他一些软件

- kde-standard

  plasma和一些基础的软件，kate(text editor)，Konqueror(default web brower)，Kget (Download Manager), KMail (email client), Dolphin (File Manager) etc.

- kde-plasma-desktop

[KDE VS GNOME](https://linuxhint.com/comparing_kde_vs_gnome/#:~:text=KDE%20offers%20a%20fresh%20and,the%20needs%20of%20their%20users.)

### install KDE

1. 安装kde-full

2. 选择kde使用的display manager，查看默认使用的display manager命令如下

   ```
   cpl in /etc/netplan λ cat /etc/X11/default-display-manager 
   /usr/sbin/gdm3
   ```


### Error

#### 0x001

在使用KWIN时，如果新建规则匹配所有窗口并且尺寸大于屏幕所显示的尺寸登入后就会无法显示。可以按照如下步骤进行修复

1. choose recovery mode while booting
2. choose root shell
3. delete `~/.config/kwinrules`
4. resume reboot

## GNOME

GNOME3是大多数distro中默认的desktop environment