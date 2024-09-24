# IOS安装

参考：

https://www.youtube.com/watch?v=ykF15xI44mE

https://www.youtube.com/watch?v=YQcWuWGjppY

https://docs.gns3.com/docs/emulators/cisco-ios-images-for-dynamips/

https://docs.gns3.com/docs/troubleshooting-faq/where-do-i-get-ios-images/

 c3640, c3660, c3725, c3745 and c7200官方推荐IOS

单个IOS可能会耗尽CPU，所以需要设置一个==Idle-PC value==，尽可能保证IOS使用的CPU在正常范围内。每个IOS使用的Idle-PC value值都不一样。每个IOS使用的RAM也不同，具体配置参考上面的链接。

由于产品的生命周期，除了c7200外都不能在Cisco官网下载。(==c7200系列GNS3只支持7206==)

[官网IOS下载地址](https://software.cisco.com/download/home/282188585/type?catid=268437899)

```mermaid
graph LR
a(IOS Software)-->b(all release)-->c(15.2M)-->d(15.2.4M10)
```

[非官方下载地址](https://mega.nz/folder/nJR3BTjJ#N5wZsncqDkdKyFQLELU1wQ)

导入IOS参考：https://www.kjnotes.com/devtools/53

1. Network Adapter：不同设备可以选择的network adapters(NIC)数量和类型不同
2. WIC modules：WAN Interface Card(WIC)广域网接口卡

### IOS下载地址

http://srijit.com/working-cisco-ios-gns3/

https://gist.github.com/takyon12/ec938f7dbd4d32d2ba3e

https://yaser-rahmati.gitbook.io/gns3/cisco-ios-images-for-dynamips