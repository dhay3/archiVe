# KDE connect

ref：

https://kdeconnect.kde.org/

https://userbase.kde.org/KDEConnect#Overview

## Digest

KDE Connect 是 KDE destktop environment 中的一个组件，通过他可以连接不同终端来控制某些功能。有如下几种特性

- receive your phone notifications on your desktop computer and reply to messages
- control music playing on your desktop from your phone
- use your phone as a remote control for your desktop
- run predefined commands on your PC from connected devices
- check your phone from the desktop
- ring your phone to help finding it
- share files and links between devices
- browse your phone from the desktop
- control the desktop’s volume from the phone

为了使用这些功能，有如下几点限制

- implements a secure communication protocol over the network and allows any developeer to create plugins on top of it
- has a component that you install on your desktop
- has a KDE connect client app you run on your phone

## IOS

https://github.com/KDE/kdeconnect-ios

这里只介绍IOS的安装，桌面版KDE connect 在安装 KDE plasma 是时会自动安装

IOS 版本需要大于 14，且安装了 TestFilght

通过如下链接下载

https://www.google.com.hk/url?sa=t&rct=j&q=&esrc=s&source=web&cd=&cad=rja&uact=8&ved=2ahUKEwiDuLucrPn2AhXxglYBHSHWAUYQFnoECAoQAQ&url=https%3A%2F%2Ftestflight.apple.com%2Fjoin%2FvxCluwBF&usg=AOvVaw0tYc_CEEOc2Jh-08IjrSha

IOS版本目前bug比较多，常会出现crashed

## How to use

### pairing

手机端和电脑端的KDE Connect 同时开启且在同一个LAN中就会开始配对，如果无法配对尝试阅读如下链接

https://userbase.kde.org/KDEConnect#I_have_two_devices_running_KDE_Connect_on_the_same_network.2C_but_they_can.27t_see_each_other

如果并未使用防火墙，但无法链接可以看一下手机端上的权限是否开齐，如果开齐了但是未能解决可以通过手动添加IP来解决（这种方式可能在DHCP租期过期后失效，需要重新指定）

### Send files

如果从手机往电脑传，默认会将文件放在`~/Downloads`下（可以在share and receive 中设置），这时可以使用`gwenview`来查看



### Remote input

在手机上可以模拟电脑的鼠标

### Run command

在手机上运行电脑上预设的命令