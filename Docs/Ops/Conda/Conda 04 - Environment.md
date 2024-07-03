---
title: Conda 04 - Environment
author: "0x00"
createTime: 2024-06-02
lastModifiedTime: 2024-06-02-19:16
draft: true
tags:
  - Python
---

## 0x01 Overview

Environments 是 Conda 中的一个概念，是包和软件的集合。每个 Environment 之间互相隔离，Environments 中的包和软件只对当前 Environment 生效。类似于 Python 中的 venv 的概念，但不仅限于包

## 0x02 Create environment



## 

## 0x02 PATH env

在安装 Conda 后，Conda 会 hijack PATH env
每当 activate 或者是 deactivate environment 时， PATH 的值都会变化。会优先使用 Conda 的 bin 目录

例如

```shell
# condabin 目录下，只有一个 conda binary
cc in ~ λ echo ${PATH}
/home/cc/anaconda3/condabin:/home/cc/.local/bin:/usr/local/sbin:/usr/local/bin:/usr/bin:/opt/cuda/bin:/opt/cuda/nsight_compute:/opt/cuda/nsight_systems/bin:/var/lib/flatpak/exports/bin:/usr/lib/jvm/default/bin:/opt/nessus/bin:/opt/nessus/sbin:/usr/bin/site_perl:/usr/bin/vendor_perl:/usr/bin/core_perl:/opt/rocm/bin:/var/lib/snapd/snap/bin:/opt/nessus/bin:/opt/nessus/sbin

# 切换到 base environment
cc in ~ λ conda activate

# anaconda3/bin 包含所有 base environment 下安装的 binary
╭─ conda on (base)
╰─ cc in ~ λ echo ${PATH}
/home/cc/anaconda3/bin:/home/cc/anaconda3/condabin:/home/cc/.local/bin:/usr/local/sbin:/usr/local/bin:/usr/bin:/opt/cuda/bin:/opt/cuda/nsight_compute:/opt/cuda/nsight_systems/bin:/var/lib/flatpak/exports/bin:/usr/lib/jvm/default/bin:/opt/nessus/bin:/opt/nessus/sbin:/usr/bin/site_perl:/usr/bin/vendor_perl:/usr/bin/core_perl:/opt/rocm/bin:/var/lib/snapd/snap/bin:/opt/nessus/bin:/opt/nessus/sbin

# 切换到另外的 environment
╭─ conda on (base)
╰─ cc in ~ λ conda activate sd

# 当 environment 切换时 PATH 的值也会发生变化
╭─ conda on (sd)
╰─ cc in ~ λ echo ${PATH}
/home/cc/anaconda3/envs/sd/bin:/home/cc/anaconda3/condabin:/home/cc/.local/bin:/usr/local/sbin:/usr/local/bin:/usr/bin:/opt/cuda/bin:/opt/cuda/nsight_compute:/opt/cuda/nsight_systems/bin:/var/lib/flatpak/exports/bin:/usr/lib/jvm/default/bin:/opt/nessus/bin:/opt/nessus/sbin:/usr/bin/site_perl:/usr/bin/vendor_perl:/usr/bin/core_perl:/opt/rocm/bin:/var/lib/snapd/snap/bin:/opt/nessus/bin:/opt/nessus/sbin
```

另外需要注意的一点是，安装 Conda 后，如果切换用户， `PATH` 会被继承。因为生产是 subshell，在 conda init 脚本中 `PATH` 被 export，所以会被继承

```zsh
cc in ~ λ conda activate sd

╭─ conda on (sd)
╰─ cc in ~ λ su root
Password:
[Hedwig cc]# echo $PATH
/home/cc/anaconda3/envs/sd/bin:/home/cc/anaconda3/condabin:/home/cc/.local/bin:/usr/local/sbin:/usr/local/bin:/usr/bin:/opt/cuda/bin:/opt/cuda/nsight_compute:/opt/cuda/nsight_systems/bin:/var/lib/flatpak/exports/bin:/usr/lib/jvm/default/bin:/opt/nessus/bin:/opt/nessus/sbin:/usr/bin/site_perl:/usr/bin/vendor_perl:/usr/bin/core_perl:/opt/rocm/bin:/var/lib/snapd/snap/bin:/opt/nessus/bin:/opt/nessus/sbin
```


---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

