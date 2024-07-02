---
createTime: 2024-07-02 13:00
tags:
  - "#hash1"
  - "#hash2"
---

# Conda 06 - Package Install

## 0x01 Overview

Conda 按装包非常简单，可以直接通过 `conda install` 实现。过滤包的格式和 `conda search` 相同[Conda 05 - Package Search](Conda%2005%20-%20Package%20Search.md)

## 0x02 Order

Conda 可以使用不同的 channels，但是不同的 channels 中会包含同版本的包。这就涉及到安装 packages 时，需要使用那个包的问题
Conda 会按照如下顺序

1. Sorts packages from highest to lowest channel priority.
    选中在 `condarc` 中出现的第一个 channel(可以使用 `conda config --show channels` 来查看)
2. Sorts tied packages---packages with the same channel priority---from highest to lowest version number. For example, if channelA contains NumPy 1.12.0 and 1.13.1, NumPy 1.13.1 will be sorted higher.
    同 channel 内使用 version 高的 package
3. Sorts still-tied packages---packages with the same channel priority and same version---from highest to lowest build number. For example, if channelA contains both NumPy 1.12.0 build 1 and build 2, build 2 is sorted first. Any packages in channelB would be sorted below those in channelA.
    同 channel 同 version 使用 build 高的 package
4. Installs the first package on the sorted list that satisfies the installation specifications.

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

