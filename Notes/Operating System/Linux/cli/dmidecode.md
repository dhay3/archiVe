# dmidecode

## Overview

syntax

```
dmidecode [options]
```

`dmidecode` 是一个 Linux 上用于 decode DMI table(实际是 SMBIOS) 的工具，在使用 `dmidecode` 之前需要先了解一下什么是 DMI 和 SMBIOS

## DMI/SMBIOS

Desktop Management Interface(DMI) 是一个允许系统调用以及管理硬件设备的通用的标准框架

SMBIOS 是 DMI 的一个真子集($SMBIOS \subseteq DMI$)，规定了系统可以从 BIOS 中读取并调用的信息(SMBIOS table)。`demicode` 就是读取这些信息的一个工具

SMBIOS table 中的每条 entry 由下面格式组成

```
Handle 0x0002, DMI type 2, 8 bytes.  Base  Board  Information
	Manufacturer: Intel
  Product Name: C440GX+
  Version: 727281-001
  Serial Number: INCY92700942
```

- handle

  唯一标识符，十六进制

- type

  SMBIOS 定义了不同组件 type 不同

- size

  数据的大小，一般比实际要小

- decode value

  根据type不同，显示的内容不同


## Optional args

- `-s | --string KEYWORD`

  只显示 SMBIOS table 中指定 KEYWORD 的信息，KEYWORD 的值具体看 Man page

  ```
  (base) 0x00 in ~ λ sudo dmidecode -s bios-version
  GZCN27WW
  ```

  能用 `-s` 查看的都可以用 `-t` 查看

- `-t | --type TYPE`

  只显示 SMBIOS table 中指定 TYPE 的信息， Type 的值具体看 Man page
  
  ```
  (base) 0x00 in ~ λ sudo dmidecode -t memory      
  # dmidecode 3.5
  Getting SMBIOS data from sysfs.
  SMBIOS 3.3.0 present.
  
  Handle 0x0001, DMI type 16, 23 bytes
  Physical Memory Array
          Location: System Board Or Motherboard
          Use: System Memory
          Error Correction Type: None
          Maximum Capacity: 16 GB
          Error Information Handle: 0x0000
          Number Of Devices: 2
  ...
  ```

## Examples

- 查看机器型号，序列号等信息

  ```
  (base) 0x00 in ~ λ sudo dmidecode -t system
  # dmidecode 3.5
  Getting SMBIOS data from sysfs.
  SMBIOS 3.3.0 present.
  
  Handle 0x002F, DMI type 32, 11 bytes
  System Boot Information
          Status: No errors detected
  
  Handle 0x000E, DMI type 1, 27 bytes
  System Information
          Manufacturer: LENOVO
          Product Name: 82MS
          Version: Lenovo XiaoXinPro 14ACH 2021
          Serial Number: PF2PTW6L
          UUID: cb18d6e6-7fe9-11eb-80ec-38f3ab15da50
          Wake-up Type: Power Switch
          SKU Number: LENOVO_MT_82MS_BU_idea_FM_XiaoXinPro 14ACH 2021
          Family: XiaoXinPro 14ACH 2021
  ```

- 查看机器 BIOS 信息

  ```
  (base) 0x00 in ~ λ sudo dmidecode -t bios
  ...
  BIOS Information
          Vendor: LENOVO
          Version: GZCN27WW
          Release Date: 05/01/2022
          Address: 0xE0000
          Runtime Size: 128 kB
          ROM Size: 16 MB
          Characteristics:
  ...
  ```

- 查看机器内存信息

  ```
  (base) 0x00 in ~ λ sudo dmidecode -t memory
  # dmidecode 3.5
  Getting SMBIOS data from sysfs.
  SMBIOS 3.3.0 present.
  
  Handle 0x0001, DMI type 16, 23 bytes
  Physical Memory Array
          Location: System Board Or Motherboard
          Use: System Memory
          Error Correction Type: None
          Maximum Capacity: 16 GB
          Error Information Handle: 0x0000
          Number Of Devices: 2
  ```

  这里可以看出最大物理内存为16GB，由2条内存条组成双通。省略的信息里还可以看出内存的频率(对应 Speed 的字段)

- 查看机器 CPU 信息

  ```
  (base) 0x00 in ~ λ sudo dmidecode -t processor
  ...
  				Voltage: 1.2 V
          External Clock: 100 MHz
          Max Speed: 4450 MHz
          Current Speed: 3200 MHz
          Status: Populated, Enabled
          Upgrade: None
          L1 Cache Handle: 0x0003
          L2 Cache Handle: 0x0004
          L3 Cache Handle: 0x0005
          Serial Number: Null
          Asset Tag: Null
          Part Number: Null
          Core Count: 8
          Core Enabled: 8
          Thread Count: 16
  
  ```

  如果想要看 CPU Cache 可以使用 `dmidecode -t cache`

  当然你也可以使用 `lscpu` 来查看

- 查看机器 接口信息

  ```
  (base) 0x00 in ~ λ sudo dmidecode -t connector  
  ...
  Handle 0x0014, DMI type 8, 9 bytes
  Port Connector Information
          Internal Reference Designator: J49
          Internal Connector Type: None
          External Reference Designator: USB 3.1 A P0
          External Connector Type: Access Bus (USB)
          Port Type: USB
  
  Handle 0x0015, DMI type 8, 9 bytes
  Port Connector Information
          Internal Reference Designator: J70
          Internal Connector Type: None
          External Reference Designator: USB 3.1 A P1
          External Connector Type: Access Bus (USB)
          Port Type: USB
  ...
  ```

**references**

[^1]:https://en.wikipedia.org/wiki/System_Management_BIOS
[^2]:https://en.wikipedia.org/wiki/Desktop_Management_Interface
[^3]:https://www.vmware.com/topics/glossary/content/desktop-management.html

