---
createTime: 2025-04-17 10:57
license: cc by 4.0
tags: 
 - "#hash1" 
 - "#hash2"
---

# RAID - ZTE R5300 G4X

## 0x01 Preface

以 ZTE5300 G4X 使用 RM241B-18i 2G RAID Controller，将两块 960G 的盘组成 Hardware RAID 1 用作系统盘为例

| Zone | Slot | Location | Disk Locater                        | Size (GB) | Media Type | Vendor Id | Product Id | Serial Num     | Interface Type | Max Speed | Negotiation Speed | Controller Name |
| ---- | ---- | -------- | ----------------------------------- | --------- | ---------- | --------- | ---------- | -------------- | -------------- | --------- | ----------------- | --------------- |
| --   | 50   | N/A      | OnBoard Slot1 : CN4 : EID0 : Slot16 | 960       | SSD        | Samsung   | MZ-7L39600 | S6KNNN0X257541 | SATA           | 6.0Gb/s   | 6.0Gb/s           | RM241B-18i 2G   |
| --   | 51   | N/A      | OnBoard Slot1 : CN4 : EID0 : Slot17 | 960       | SSD        | Samsung   | MZ-7L39600 | S6KNNN0X257537 | SATA           | 6.0Gb/s   | 6.0Gb/s           | RM241B-18i 2G   |


## 0x02 How to Configure RAID

R5300 可以通过两种方式来配置 RAID

1. Web UI
2. Legacy Boot SAS/SATA Utilities

### 0x02a Web UI

登入 IPMI Web 界面，选择 “设置”-“RAID 管理”-“逻辑设备信息”

![](https://github.com/dhay3/picx-images-hosting/raw/master/2025-04-17_13-18-26.m9kx0qx9.webp)

在这个界面即可管理配置 RAID

> [!NOTE]
> Web 界面并不可靠，经常会出现 “未找到可用的控制器”，但是 RAID Card 实际仍然正常。这时可以使用 [0x02b Legacy Boot](#0x02b%20Legacy%20Boot)

![](https://github.com/dhay3/picx-images-hosting/raw/master/2025-04-17_18-28-42.4joao72i6h.webp)

### 0x02b Legacy Boot

> [!NOTE]
> RM241B-18i RAID Card 只兼容 Legacy Boot，如果设置为 UEFI Boot，在引导界面就不会加载 RAID Controller 配置项

登入虚拟控制台（ZTE KVM/HTML 不支持挂载盘，所以通常使用 KVM/JAVA），重启服务器进入 BIOS 设置界面，将 Boot Mode 从 UEFI 设置为 Legacy

![](https://github.com/dhay3/picx-images-hosting/raw/master/2025-04-17_15-32-39.m9l1mhmk.webp)

<kbd>f10</kbd> 保存并退出会重新引导

出现如下界面按照提示键入 <kbd>ctrl</kbd> + <kbd>a</kbd> 进入 RAID 管理界面

![](https://github.com/dhay3/picx-images-hosting/raw/master/2025-04-17_15-49-47.102cy8bdx2.webp)

会有 4 个选项

- Controller Details (查看 Controller 详细信息，固件版本、序列号等等)
- Configure Controller Settings（配置 Contrller，修改性能模式）
- Array Configuration（配置 RAID）
- Disk Utilities（查看 RAID Card 可以识别的硬盘信息）

这里选中 Array Configuration 进入配置 RAID 界面

![](https://github.com/dhay3/picx-images-hosting/raw/master/2025-04-17_15-54-13.41y8zgi9dt.webp)

会有 3 个选项

- Create Array（创建 RAID）
- Manage Arrays（查看删除 RAID）
- Select Boot Device（选择引导盘）

![](https://github.com/dhay3/picx-images-hosting/raw/master/2025-04-17_16-06-32.41y8zgy4ue.webp)

#### Query RAID

选择 Manage Arrays 可以查看当前的 RAID 配置

其中

```
ARRAY-A - 002-PD(s), 01-LD(s) (Boot Port Array)
```

- ARRAY-A 为 RAID 的标识符(不是 Logical)
- 002-PD(s) 表示 ARRAY-A 由 2 块 Physical Drives(PDs) 组成
- 01-LD(s) 表示 ARRAY-A 将 2 块 PDs 组成 1 块 Logical Drives(LD)
- Boot Port Array 表示 ARRAY-A 是引导盘

![](https://github.com/dhay3/picx-images-hosting/raw/master/2025-04-17_16-12-02.sz52thocx.webp)

键入 <kbd>enter</kbd> 可以查看详细信息

![](https://github.com/dhay3/picx-images-hosting/raw/master/2025-04-17_16-18-28.64e1njc03m.webp)

#### Delete RAID

选择 Manage Arrays，可以看到提示使用 <kbd>ctrl</kbd> + <kbd>d</kbd> ==来删除所有的 RAID 配置==

![](https://github.com/dhay3/picx-images-hosting/raw/master/2025-04-17_16-26-40.7snekqcu70.webp)

会出现如下界面，键入 <kbd>Y</kbd> 即可删除

![](https://github.com/dhay3/picx-images-hosting/raw/master/2025-04-17_16-32-47.2ks3xqrujp.webp)

也可以进入详情页使用 <kbd>del</kbd> 来删除指定 LD RAID 配置，同理键入 <kbd>Y</kbd> 即可删除

![](https://github.com/dhay3/picx-images-hosting/raw/master/2025-04-17_16-30-53.4g4oqd1wq6.webp)

#### Create RAID

选择 Create Array

使用 <kbd>ins</kbd>/<kbd>space</kbd> 来选中需要组成 LD 的 PD

使用 <kbd>del</kbd> 来取消选中不需要组成 LD 的 PD

可以在 Selected Drives 中看到选中的 PD

![](https://github.com/dhay3/picx-images-hosting/raw/master/2025-04-17_16-41-08.3yen1sdipg.webp)

当选中后键入 <kbd>enter</kbd> 即可进入创建 RAID 的配置界面

![](https://raw.githubusercontent.com/dhay3/picx-images-hosting/refs/heads/master/2025-04-17_16-44-43.26lo6vyqqx.webp)

- RAID Level
- Logical Drive Name （LD 标识符，装系统时识别的盘符）
- Strip/Full Stripe Size （stripping 的块大小）
- Parity Group Count（每个 parity group 的盘数量）
- Build Method（RAID 重建的）
- Size（组成 RAID 后的盘大小）
- SSD OverProvisiong（SSD 防坏盘策略）
- Acceleration Method（缓存策略）

因为需要组 RAID1 系统盘，所以按照下图配置

![](https://github.com/dhay3/picx-images-hosting/raw/master/2025-04-17_16-56-34.6pnp9vk4f3.webp)


---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [中兴服务器R5300 G4 Raid配置\_中兴r5300g4服务器raid配置-CSDN博客](https://blog.csdn.net/m0_46354425/article/details/127537161)

***References***


