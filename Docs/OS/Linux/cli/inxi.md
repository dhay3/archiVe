# inxi

## 0x01 Overview

syntax

```
inxi

inxi [-AbBCdDEfFGhiIjJlLmMnNopPrRsSuUwyYzZ]

inxi [-c -NUMBER] [--sensors-exclude SENSORS] [--sensors-use SENSORS] [-t [c|m|cm|mc][NUMBER]] [-v NUMBER] [-w [LOCATION]] [--weather-unit {m|i|mi|im}] [-y WIDTH]

inxi [--edid] [--memory-modules] [--memory-short] [--recommends] [--sensors-default] [--slots] [--version]  [--version-short]

inxi [-x|-xx|-xxx|-a] -OPTION(s)
```

`inxi` 是一个类似 `lshw` 的一个工具，用于主机上的设备信息

## 0x02 Optial args

- `-b | --basic`

  查看基本信息(输出所有)


- `-M | --machine`

  查看主机信息(Motherboard, BIOS, Device)

  ```
  (base) cc in ~ λ inxi -M
  Machine:
    Type: Laptop System: Micro-Star product: Alpha 17 C7VF v: REV:1.0
      serial: <superuser required>
    Mobo: Micro-Star model: MS-17KK v: REV:1.0 serial: <superuser required>
      UEFI: American Megatrends LLC. v: E17KKAMS.118 date: 01/11/2024
  ```

- `-m | --memory`

  查看内存信息

  ```
  (base) cc in ~ λ inxi -m
  Memory:
    System RAM: total: 32 GiB available: 30.53 GiB used: 7.08 GiB (23.2%)
    Array-1: capacity: 128 GiB slots: 2 modules: 2 EC: None
    Device-1: Channel-A DIMM 0 type: DDR5 size: 16 GiB speed: spec: 5600 MT/s
      actual: 5200 MT/s
    Device-2: Channel-B DIMM 0 type: DDR5 size: 16 GiB speed: spec: 5600 MT/s
      actual: 5200 MT/s
  
  ```

- `-C | --cpu`

  查看 CPU 信息


- `-G | --graphics`

  查看 GPU 信息

- `-D | --disk`

  查看硬盘信息

  ```
  )base) cc in ~ λ inxi -D
  Drives:
    Local Storage: total: 2.75 TiB used: 799.72 GiB (28.4%)
    ID-1: /dev/nvme0n1 vendor: Samsung model: MZVL21T0HCLR-00B00
      size: 953.87 GiB
    ID-2: /dev/nvme1n1 vendor: Smart Modular Tech. model: SHPP41-2000GM
      size: 1.82 TiB
  ```

- `-E | --bluetooth`

  查看蓝牙设备

  ```
  (base) cc in ~ λ inxi -E
  Bluetooth:
    Device-1: Foxconn / Hon Hai driver: btusb type: USB
    Report: btmgmt ID: hci0 rfk-id: 0 state: down bt-service: enabled,running
      rfk-block: hardware: no software: no address: A8:3B:76:1E:25:B2 bt-v: 5.3
  ```

- `-A | --audio`

  查看声卡设备

  ```
  (base) cc in ~ λ inxi -A
  Audio:
    Device-1: NVIDIA driver: snd_hda_intel
    Device-2: AMD Rembrandt Radeon High Definition Audio driver: snd_hda_intel
    Device-3: AMD ACP/ACP3X/ACP6x Audio Coprocessor driver: snd_rpl_pci_acp6x
    Device-4: AMD Family 17h/19h HD Audio driver: snd_hda_intel
    API: ALSA v: k6.7.7-1-MANJARO status: kernel-api
    Server-1: PulseAudio v: 17.0 status: active
  
  ```

- `-B | --battery`

  查看电池信息

  ```
  (base) cc in ~ λ inxi -B
  Battery:
    ID-1: BAT1 charge: 91.7 Wh (98.2%) condition: 93.4/95.0 Wh (98.3%)
  ```

- `-N | --network`

  查看网络设备信息

- `-J | --usb`

  查看 USB 信息

- `-v <number>`

  详细程度

