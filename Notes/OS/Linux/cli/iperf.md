---
createTime: 2025-02-05 17:06
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# iperf3

## 0x01 Preface

iperf3 是一个网络吞吐测试工具，以 C/S 模式运行

## 0x02 Syntax

```

```

## iperf VS iperf3


## Optional Args

> [!note]
> 具体参数看 manual page

### General

- `-b`

- `-e`

- `-f | --format <[bkmgBKMG]>`

	指定输出 bandwidth report 中的单位

- `-i | --interval <n>`

	每隔 n seconds 后，再输出 bandwidht report

- `-l | --len <n[kmKM]>`

	指定 TCP/UDP 可以使用的 read/write buffer 大小，默认 TCP 128K，UDP 1470K

- `-m | --print-mss`

	打印 TCP MSS

- `-M | --mss <n>`

	设置 TCP MSS

- `-N | --nodelay`

	设置 TCP no delay

- `-o | --output <filename>`

	将 bandwidth report 和错误信息输出到指定 filename 中

- `-p | --port <n>`

	server 使用指定的端口 n，默认 5001

- `-u | --udp`

	使用 UDP 探测而不是 TCP

- `-w | --winidow <n[kmKM]>`

	指定 TCP 可以使用的 window size

- `-B | --bind <host>`

	使用指定的 NIC

### Server Related

- `-b | --bandwidth <n[kmgKMG]>`

	设置 server read bandwidth 最大到 n bps

- `-s | --server`

	iperf 以 server 模式运行

- `-B | --bind <ip | ip%deivce>`

	server 监听指定的 ip 或者是指定的 device

- `-D | --deamon`

	server 以 daemon 的形式运行

### Client Related

- `-c | --client <server>`

	iperf 以 client 模式运行，连接指定的 iperf server

- 

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [iPerf - The TCP, UDP and SCTP network bandwidth measurement tool](https://iperf.fr/)
- `man iperf`
- `man iperf3`

***References***


