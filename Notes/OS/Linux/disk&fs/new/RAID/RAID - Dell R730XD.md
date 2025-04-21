---
createTime: 2025-03-05 12:49
license: cc by 4.0
tags: 
 - "#hash1" 
 - "#hash2"
---

# RAID - Dell R730XD

## 0x01 Preface

以 Dell R730 XD 使用 PERC H730P Mini RAID Controller 为例

![](https://github.com/dhay3/picx-images-hosting/raw/master/Snipaste_2025-02-25_16-22-28.60ue05rhdm.webp)

需求：

将两块小盘组成 RAID 1，用作系统盘

## 0x02 How to Enter RAID Configuration Interface

有 2 种方式进入 RAID Controller 配置界面

1. 通过快捷键进入
2. 通过 UEFI 界面的选项进入

### 0x02a Shortcut

在引导的过程中会出现如下界面

![](https://github.com/dhay3/picx-images-hosting/raw/master/Snipaste_2025-02-25_16-15-12.8ojuaip7ir.webp)

按照提示 <kbd>ctrl</kbd> + <kbd>r</kbd> 可以进入 RAID controller 管理界面

#### VD Mgmt

Virtual Disk 管理界面

![](https://github.com/dhay3/picx-images-hosting/raw/master/Snipaste_2025-02-25_16-01-30.1ovksmhb9x.webp)

#### PD Mgmt

#### Ctrl Mgmt

#### Properties



这里可以看见

### 0x02b UEFI



## How to Check RAID

这里可以看到当前没有任何 RAID 配置，有 7 块盘，按照需求我们需要将 `00:01:00` 和 `00:01:01` 组成 RAID1 用作系统盘

选择 No Configuration Present 可以进入创建 RAID 界面

![](https://github.com/dhay3/picx-images-hosting/raw/master/Snipaste_2025-02-25_16-03-40.b91olayk1.webp)



![](https://github.com/dhay3/picx-images-hosting/raw/master/Snipaste_2025-02-25_16-05-02.77dp8rusf0.webp)



## How to Create RAID



## How to Delete RAID


---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***



***References***


