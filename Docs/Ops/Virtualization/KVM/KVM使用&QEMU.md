# KVM使用/QEMU概述

ref

https://wiki.archlinux.org/title/QEMU

> KVM一般使用 QEMU 来管理

## QEMU概述

qemu is a generic and open source machine emulator and virtualizer

qemu 和 vmware/virtualbox 一样都是一款虚拟化模块器

## QEMU 安装

qemu 有两种安转方式

1. qemu-full 有GUI
2. qemu-base 无GUI

可以通过`pacman -Ss qemu `来安转qemu-full，同时 qemu 支持如下几种运行方式

### full-system emulation

qemu emulates a full system, including one or serval processors and various peripherials

qemu 以这种方式方式运行不需要在Linux base OS 上运行，同时因为是 full-system emulation 相对应的速度也会比较慢。可以通过如下方式来启动 full-system emulation

通常启动命令都是以`qemu-system-target_architecture`。例如x86架构的，启动命令就是`qemu-system-x86_64`，如果是arm架构的，就是`qemu-system-arm`

同时还会提供如下特性

> 需要注意的是 non-headless 和 headless 都是使用 `qemu-system-x86_64`来运行的，所以不能同时安装即不兼容

1. non-headless(default)

   有GUI的特性

2. headless

   没有GUI的特性

### usermode emulation

//TODO

## libvrt

qemu没有和 vmware 或者 virtualbox 便捷的方式来管理 virtual machines。所有的运行参数需要在启动的命令行中指出。除非自己写了启动脚本，例如

```
#!/usr/bin/env bash
set -eo pipefail
PACMAN_PATH=$(readlink -f $(which pacman))
function check_qemu() {
  $PACMAN_PATH -Qs qemu >&/dev/null
  if (($? == 1)); then
    echo "install qemu first"
    exit 1
  fi
}
check_qemu
CORE=${CORE-2}
RAM=${RAM-4}
while [ $# -gt 0 ]; do
  param=$1
  case "${param}" in
  -c)
    CORE=$2
    shift
    shift
    ;;
  -r)
    RAM=$2
    shift
    shift
    ;;
  *)
    echo "usage lsw -c [core] -r [ram]G, default 2c4g"
    exit 1
    ;;
  esac
done
coproc qemu-system-x86_64 -enable-kvm \
  -smp cores="${CORE}",threads=2 \
  -boot order=a,menu=on \
  -drive file=win-disk,format=qcow2 \
  -m "${RAM}"G >&/dev/null || exit 1
```

否则一般通过 libvrt 来管理



