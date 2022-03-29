# Linux lshw

> 部分硬件信息不会显示可以使用`dmidecode`来查看

## digest

lshw 是一个用于显示系统hardware的工具，包扣 memory configuration, firmware version, mainboard configuration, cpu version and speed, cache configuration, bus speed 等等

## Optional args

### output

- `-X`

  使用X11 GUI 运用lshw

- `-json | -html | xml`

  输出指定给的内容

### filter

- `-businfo`

  显示和bus相关信息

- `-class CLASS | -C CLASS`

  只查看特定类型的 hardware，可以使用`--short`或`-businfo`来查看可以使用的class

  常用network来查看NIC的信息和logical name的关联

  ```
  cpl in /etc/systemd λ lshw -C network
  WARNING: you should run this program as super-user.
    *-network                 
         description: Wireless interface
         product: Wi-Fi 6 AX200
         vendor: Intel Corporation
         physical id: 0
         bus info: pci@0000:01:00.0
         logical name: wlp1s0
         version: 1a
         serial: 64:bc:58:bd:a6:19
         width: 64 bits
         clock: 33MHz
         capabilities: bus_master cap_list ethernet physical wireless
         configuration: broadcast=yes driver=iwlwifi driverversion=5.15.25-1-MANJARO firmware=63.c04f3485.0 cc-a0-63.ucode ip=192.168.124.9 latency=0 link=yes multicast=yes wireless=IEEE 802.11
         resources: irq:72 memory:fd500000-fd503fff
  ```

  

