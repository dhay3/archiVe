# SSID，BSSID，ESSID

参考：

https://www.juniper.net/documentation/en_US/junos-space-apps/network-director3.7/topics/concept/wireless-ssid-bssid-essid.html

## SSID

> an ssid is the name of network

在一个区域内可以有多个WLAN共存，每个WLAN需要一个唯一的ID，service set ID(SSID)。SSID保证数据包在正确的WLAN中，即使当WLAN出现重叠。作为普通的WLAN user，你只需要关注SSID。提供authentication就可以连接到该WLAN

## BSSID

> bssid identify access points and their Client(AP MAC address)

在一个WLAN中可能存在多个access points(无线ap)，为了保证数据包发送到正确的access point，所以需要用一个ID来区分access point，basic service set identifier(BSSID)。可以理解为access point的MAC地址。





Each Access Point has its Own BSS

> tips
>
> 我们可以使用`nmcli`来查看wlp1s0(这是我的无线网卡)的SSID和BSSID

```
cpl in /tmp λ nmcli device wifi list wlp1s0
IN-USE  BSSID              SSID               MODE   CHAN  RATE        SIGNAL  BARS  SECURITY  
        EC:26:CA:F8:01:74  TP-LINK_401        Infra  11    405 Mbit/s  65      ▂▄▆_  WPA1 WPA2 
        F4:2A:7D:91:01:2F  春晓               Infra  1     405 Mbit/s  44      ▂▄__  WPA1 WPA2 
```

## 例子

一家公司面积比较大，安装了若干台无线接入点（AP或者无线路由器），公司员工只需要知道一个SSID就可以在公司范围内任意地方接入无线网络。**BSSID其实就是每个无线接入点的MAC地址**。当员工在公司内部移动的时候，SSID是不变的。但BSSID随着你切换到不同的无线接入点，是在不停变化的。