---
createTime: 2025-02-06 16:50
license: cc by 4.0
tags: 
 - "#hash1" 
 - "#hash2"
---

# RAID 06 - How to Configure RAID

[RAID阵列配置-服务器RAID配置-服务器RAID配置方法-浪潮信息](www.ieisystem.com/qa/raid/)

## 0x01 Preface

Hardware RAID 和 Software RAID 管理 RAID 的方式不同，配置 RAID 的方式也不同

## 0x02 Hardware RAID

根据 RAID Card 的厂商不同，Hardware RAID 配置方式各有不同，支持的 RAID Level 也不同

> [!note]
> 大多服务器可以直接通过 Legacy BIOS 或者是 UEFI 的界面配置 RAID，但是也有一些服务器没有对应的选项。
> 
> 这时可以在 RAID Card 启动的界面，键入快捷键来进入 RAID 配置，常见的快捷键有
> 
> <kbd>ctrl</kbd> + <kbd>R</kbd>/<kbd>A</kbd>/<kbd>H</kbd>/<kbd>C</kbd>

### 0x02a ZTE R5300 G4

以 ZTE R5300 G4 使用 RM241B-18i 2G RAID Controller 为例

![](https://github.com/dhay3/picx-images-hosting/raw/master/Snipaste_2025-03-05_12-13-06.3nrrixta0a.webp)

在引导的过程中会出现如下界面

![](https://github.com/dhay3/picx-images-hosting/raw/master/Snipaste_2025-02-19_17-37-37.32i3jn5jgl.webp)

> [!note]
> 这里也表明了 R5300 G4 RAID Card 只支持 SAS/SATA 接口的磁盘，所以 M2/NVME 接口的磁盘并不会在配置界面中显示(这时通常需要使用 Software RAID 来对 SSD 配置 RAID)

按照提示 <kbd>ctrl</kbd> + <kbd>a</kbd> 进入 RAID controller 管理界面

![](https://github.com/dhay3/picx-images-hosting/raw/master/Snipaste_2025-02-20_09-37-12.7egwr7h2mc.webp)

- Controller Details

	查看 RAID Controller 详细的信息

- Configure Contoller Settings

	配置 RAID Controller 全局配置

- Array Configuration

	管理 RAID

- Disk Utilities

	磁盘管理工具

这里选择 Array Configuraton 进入 RAID 配置界面

![](https://github.com/dhay3/picx-images-hosting/raw/master/Snipaste_2025-02-20_09-37-41.60udn6ao3h.webp)

- Create Array

	创建 RAID

- Manage Arrays

	管理(查看/删除/修改) RAID

- Select Boot Device

	选择系统引导盘(或者 RAID)

这里选择 Manage Arrays 查看当前 RAID 配置

![](https://github.com/dhay3/picx-images-hosting/raw/master/Snipaste_2025-02-20_09-37-56.2rv9qio4tv.webp)

这里可以看到服务器已经有了一个 RAID

其中

- ARRAY-A 为 RAID 的标识符
- 002-PD(s) 表示 ARRAY 由 2 块 Physical Drives(PDs) 组成
- 01-LD(s) 表示 ARRAY 将 2 块 Physical Drives 组成 1 块 Logical Drives(LD)
- Boot Port Array 表示 ARRAY-A 是引导盘(系统盘)

回车后可以展示 ARRAY A 的详细配置






LSI MegaRAID Tri-Mode SAS3508 为例

```
[root@localhost ~]# lspci | grep -iP 'lsi|raid|adaptec'
17:00.0 RAID bus controller: Broadcom / LSI MegaRAID Tri-Mode SAS3508 (rev 01)
```


### 0x02b Dell R730 XD


## 0x03 Hardware RAID strocli

`storcli` 是一个专门为 MegaRAID 系列产品设计的 RAID 管理命令行工具，由于直接和 RAID Controller 通信，所以也属于 Hardware RAID（提供命令行工具的好处是无需重启操作系统，就可以配置 RAID）

可以通过如下命令来查看 RAID Card 是否为 MegaRAID 系列

```
[root@localhost ~]# lspci | grep -iP 'mega'
17:00.0 RAID bus controller: Broadcom / LSI MegaRAID Tri-Mode SAS3508 (rev 01)
```

> [!note]
> `dmidecode -t slot` 不一定会显示 RAID Card 信息，有些 motherboard OEM 不会将 RAID Card 信息写入 dmi table

具体使用方式参考

## 0x03 Software RAID

*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [Avago SAS3508 RAID控制卡 用户指南（Whitley平台）01](https://support.xfusion.com/support/#/zh/docOnline/DOC2020027330?path=zh-cn_topic_0000001153866551&relationId=DOC2020027329&mark=0&pid=23692812)
- [R5300 G4](https://enterprise.zte.com.cn/product-detail.html?id=74&lan=zh-CN)
- [Dell-PowerEdge-RAID-Controller-H730P.pdf](https://i.dell.com/sites/doccontent/shared-content/data-sheets/en/Documents/Dell-PowerEdge-RAID-Controller-H730P.pdf)

***References***





