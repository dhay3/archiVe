---
createTime: 2025-02-06 13:39
license: cc by 4.0
tags: 
 - "#hash1" 
 - "#hash2"
---

# storcli

## 0x01 Preface

Storage Command Line Tool(`strocli`) 是一个为 MegaRAID 系列产品专门设计的 RAID 管理命令行工具，这么做的好处是无需重启操作系统，就可以配置 RAID

可以通过如下命令来查看 RAID Card 是否为 MegaRAID 系列

```
[root@localhost ~]# lspci | grep -iP 'mega'
17:00.0 RAID bus controller: Broadcom / LSI MegaRAID Tri-Mode SAS3508 (rev 01)
```

> [!note]
> `dmidecode -t slot` 不一定会显示 RAID Card 信息，有些 motherboard OEM 不会将 RAID Card 信息写入 dmi table

或者通过带外



## 0x02 Installation

MegaRAID 是 Broadcom 的产品，所以我们要到 [Broadcom 的官网](https://www.broadcom.com/)去下载 `storcli`

找到 Support and Services 然后点击 Documents, Download and Support 选择 Support Documents and Downloads

![](https://github.com/dhay3/picx-images-hosting/raw/master/2025-02-07_10-55-23.2rv928yl4n.webp)

然后点击 Refine Your Search 填写表单

```
Product Family: RAID Controller Cards
Keyword: storcli
```

![](https://github.com/dhay3/picx-images-hosting/raw/master/2025-02-07_13-27-25.102a7hudpn.webp)

会有几个版本

- StorCLI

	适用于 MegaRAID 96xx RAID Card 之前的型号，Controller 通常对应 SAS4xx 之前

- StorCLI2

	适用于 MegaRAID 96xx RAID Card 以及之后的型号，Contorller 通常对应 SAS4xx 以及之后的型号，和 StorCLI 不兼容

按需选择下载，解压后路径大概如下

```
tree -L 1
.
├── ARM
├── EFI
├── FreeBSD
├── JSON-Schema
├── Linux
├── Linux_Lite
├── Linux-PPC
├── readme.txt
├── storcliconf.ini
├── ThirdPartyLicenseNotice.pdf
├── Ubuntu
├── VMware
└── Windows
```

操作系统不同，安装的方式也不同（具体参考安装包中的 `readme.txt`），Boardcom 按照目录来区分操作系统，但是命名几乎就是乱来的

- Linux

	适用于 rhel based

	```
	#install
	rpm -ivh <storcli-x.xx-x.noarch.rpm>
	#upgrade
	rpm -Uvh <storcli-x.xx-x.noarch.rpm> 
	```

- Linux_Lite

	提供 binary，适用于 x86 架构（通过 `readelf -h storcli64 | grep Machine` 确认）

- Linux-PPC

	提供 binary，适用于 PowerPC 架构（通过 `readelf -h storcli64 | grep Machine` 确认）

- ARM

	适用于 ARM 架构

- Ubuntu

	适用于 debian based

	```
	dpkg -i <storcli-x.xx-x.deb> 
	```

- VMware

	

- Windows

	


## 0x02 Syntax

```

```



---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [stroCLI Reference Manual](https://docs.broadcom.com/doc/12352476)
- [StorCLI2-Intf-UG101.pdf](https://techdocs.broadcom.com/content/dam/broadcom/techdocs/data-center-solutions/tools/generated-pdfs/StorCLI2-Intf-UG101.pdf)
- [Avago SAS3508 RAID控制卡 用户指南（Whitley平台）01](https://support.xfusion.com/support/#/zh/docOnline/DOC2020027330?path=zh-cn_topic_0000001153867443&relationId=DOC2020027329&mark=82&pid=23692812)


***References***

[^1]:[stroCLI Reference Manual](https://docs.broadcom.com/doc/12352476)
