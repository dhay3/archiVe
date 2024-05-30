---
title: Conda 01 - Overview
author: "0x00"
createTime: 2024-05-30
lastModifiedTime: 2024-05-30-08:54
draft: true
tags:
  - Python
  - AI
---

# 0x01 Overview

Conda 是一个包管理器，最早只被用在 Python 中，用于解决在不同 Data Science 项目中依赖的问题。发展至今，Conda 已经不仅先于 Python, 其他编程语言也可以使用 Conda 来管理[^1]

# 0x02 Terms

## Environments[^2]

*An environment is a directory that contains a specific collection of packages that you have installed. For example, you may have one environment with NumPy 1.7 and its dependencies, and another environment with NumPy 1.6 for legacy testing. If you change one environment, your other environments are not affected. You can easily activate or deactivate environments, which is how you switch between them.*

Environments 是 Conda 中的一个概念，是包的集合。每个 Environment 之间互相隔离，Environments 中的包只对当前 Environment 生效。

## Channels[^3]

*Channels are the locations where packages are stored. They serve as the base for hosting and managing packages.*

Channels 类似与 Arch 中的 package mirror, 是包仓库。Conda 会从这些 Channels 将包下载到本地。主要有如下 Channels
- Default Channel
- Conda-forge(社区免费的)

Conda 默认会先从 Default Channel 下载包，如果没有就会从 Conda-forge 下载

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]: [docs.conda](https://docs.conda.io/en/latest/)
[^2]: [Environments — conda 24.5.0 documentation](https://docs.conda.io/projects/conda/en/stable/user-guide/concepts/environments.html)
[^3]: [Channels — conda 24.5.0 documentation](https://docs.conda.io/projects/conda/en/stable/user-guide/concepts/channels.html)
