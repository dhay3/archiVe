# lscpu

## Digest

syntax:`lscpu [optoins]`

lscpu 用于从 `/proc/cpuinfo` 以及其他的和 CPU architecture 有关的文件中读取信息并转成 human readable 的内容 

通常包含：

CPUs, thread, cores, sockets,  cache, Non-Uniform example

关于这些 CPU 术语可以参考 [cpu.md]()

例如：

```
CPU(s):                          16
On-line CPU(s) list:             0-15
Vendor ID:                       AuthenticAMD
Model name:                      AMD Ryzen 7 5800H with Radeon Graphics
CPU family:                      25
Model:                           80
Thread(s) per core:              2
Core(s) per socket:              8
Socket(s): 						 1
```

表示一共有 16 个逻辑 CPU，8 核 2 线程，1 socket

```
Virtualization features: 
  Virtualization:        AMD-V
```

从 virtualization 可以看出当前 CPU 支持虚拟化，使用 AMD-v 区别与 Intel-V

## Optional args

- `-C`

  查看 cache 信息

- `-b | --online`

  `-c | --offline`

  查看 online 和 offline 的 CPU

