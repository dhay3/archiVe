---
createTime: 2024-12-13 14:55
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# uptime

## 0x01 Preface

`uptime` 主要用于查看系统已经运行了多长时间，但是也包括其他的一些信息。

例如 `uptime` 结果如下

```
14:51:39 up  5:52,  2 users,  load average: 1.28, 0.99, 0.85
```

- uptime since last boot

	`14:51:39 up  5:52`，表示当前时间为 14:51:39 已经运行了 `5:52`

- count of logged in users currently 

	`2 users`，表示当前主机有 2 个登入的用户(*可以通过 `w` 来验证*)

- CPU load average for the past 1, 5, 15 mins

  > load average 是一个比较特殊的指标，需要单独拿出来

## 0x02 Syntax

```
uptime [options]
```

## 0x03 Optional Args

> [!note] 
> 具体看 man page

## 0x04 Load Average

> [!note]
> load average is the average number of processes that are either in a runnable or uninterruptable state.
> 
> A process in a runnable state is either using the CPU or waiting to use the CPU.
> 
> A process in uninterruptable state is waiting for some I/O access, eg waiting for disk

在 Linux 中通过单位时间内(通常为 1/5/15 mins，瞬时的进程数并不能有效表示出现负载) runnable(正在使用 CPU 或者等待使用 CPU) 或者是 uninterruptable(等待 I/O) 的进程数均值来描述 CPU 的负载。即 load average 表示单位时间内占用 CPU 的进程数

例如 `load average: 1.28, 0.99, 0.85` 表示

- 1 min 内平均有 `1.28` 个进程处于 runnable 或者是 uninterruptable
- 5 mins 内平均有 `0.99` 个进程处于 runnable 或者是 uninterruptable
- 15 mins 平均有 `0.85` 个进程处于 runnable 或者是 uninterruptable 

我们都知道 CPU 同一时间只能处理一个进程发出的一个指令，所以为了提高效率将 CPU 虚拟化成了 Core 的逻辑。也就出现了 single-core processor 和 multi-core processor

而 load average 的含义在 single-core processor 和 multi-core processor 的场景下差异很大

### 0x04a single-core

在 single-core 的场景下

- 1 min 内平均有 `1.28` 个进程都占用了 CPU，那么同一时间应该需要 1.28 个 CPU 来处理，得出 1 min CPU 负载为 128%
- 5 mins 内平均有 `0.99` 个进程都占用了 CPU，那么同一时间应该需要 0.99 个 CPU 来处理，得出 5 mins CPU 负载为 99%
- 15 mins 内平均有 `0.85` 个进程都占用了 CPU，那么同一时间应该需要 0.85 个 CPU 来处理，得出 15 mins CPU 负载为 85%

同理可以得出 idel 的比率

### 0x04b multi-core

在 multi-core 的场景下(假设 4 core)

- 1 min 内平均有 `1.28` 个进程都占用了 CPU，那么同一时间应该需要 1.28 个 CPU 来处理，得出 1 min CPU 负载为 $1.28/4$
- 5 mins 内平均有 `0.99` 个进程都占用了 CPU，那么同一时间应该需要 0.99 个 CPU 来处理，得出 5 mins CPU 负载为 $0.99/4$
- 15 mins 内平均有 `0.85` 个进程都占用了 CPU，那么同一时间应该需要 0.85 个 CPU 来处理，得出 15 mins CPU 负载为 $0.85/4$

同理可以得出 idel 的比率


---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- `man uptime`
- [Understanding the Load Average on Linux and Other Unix-like Systems](https://www.howtogeek.com/194642/understanding-the-load-average-on-linux-and-other-unix-like-systems/)
- [Understanding Linux Load Average: A Beginner's Guide with Real-World Analogies - Linodelinux](https://linodelinux.com/understanding-linux-load-average/)

***References***


