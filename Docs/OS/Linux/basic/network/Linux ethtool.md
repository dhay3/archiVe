# ethtool

参考：

https://zhuanlan.zhihu.com/p/146383216

ethtool用于控制NIC driver和硬件设置的功能，有一些具体的名词需要了解

- **半双工**：半双工模式允许设备一次只能发送或接收数据包。
- **全双工**：全双工模式允许设备可以同时发送和接收数据包。
- **自动协商**(autonegotiation)：自动协商是一种机制，允许设备自动选择最佳网速和工作模式（全双工或半双工模式）。
- **速度**：默认情况下，它会使用最大速度，你可以根据自己的需要改变它。
- **链接检测**：链接检测可以显示网卡的状态。如果显示为 `no`，请尝试重启网卡。如果链路检测仍显示 `no`，则检查交换机与系统之间连接的线缆是否有问题。

## options

- -i | --driver

  查看NIC使用的driver

  ```
  cpl in ~ λ ethtool -i wlp1s0 
  driver: iwlwifi
  version: 5.10.42-1-MANJARO
  firmware-version: 59.601f3a66.0 cc-a0-59.ucode
  expansion-rom-version: 
  bus-info: 0000:01:00.0
  supports-statistics: yes
  supports-test: no
  supports-eeprom-access: no
  supports-register-dump: no
  supports-priv-flags: no
  ```

- -S 

  NIC使用的静态数据

  ```
  cpl in ~ λ ethtool -S wlp1s0 
  NIC statistics:
       rx_packets: 4993
       rx_bytes: 1045456
       rx_duplicates: 0
       rx_fragments: 5044
       rx_dropped: 237
       tx_packets: 43880
       tx_bytes: 11677363
       tx_filtered: 0
       tx_retry_failed: 0
       tx_retries: 135
       sta_state: 4
       txrate: 866700000
       rxrate: 585000000
       signal: 210
       channel: 0
       noise: 18446744073709551615
       ch_time: 18446744073709551615
       ch_time_busy: 18446744073709551615
       ch_time_ext_busy: 18446744073709551615
       ch_time_rx: 18446744073709551615
       ch_time_tx: 18446744073709551615
  ```

  接收(rx)和传输(tx)的具体数据大小

- -s | --change

  修改NIC的一些行为

  ```
  cpl in ~ λ sudo ethtool -s wlp1s0 autoneg on 
  netlink error: failed to retrieve link settings
  netlink error: Operation not supported
  ```

