---
title: Conda 01 - Terms
author: "0x00"
createTime: 2024-05-30
lastModifiedTime: 2024-05-30-08:54
draft: true
tags:
  - Python
---

# Conda 01 - Terms

## 0x01 Overview

Conda 是一个包管理器，最早只被用在 Python 中，用于解决在不同 Data Science 项目中依赖的问题。发展至今，Conda 已经不仅限于 Python, 其他编程语言也可以使用 Conda 来管理[^1](你可以安装 JDK，Maven 等等)

## 0x02 Terms

### 0x02a Environments[^2]

*An environment is a directory that contains a specific collection of packages that you have installed. For example, you may have one environment with NumPy 1.7 and its dependencies, and another environment with NumPy 1.6 for legacy testing. If you change one environment, your other environments are not affected. You can easily activate or deactivate environments, which is how you switch between them.*

Environments 是 Conda 中的一个概念，是包和软件的集合。每个 Environment 之间互相隔离，Environments 中的包和软件只对当前 Environment 生效。类似与 Python 中的 venv 的概念

#### Environments VS venv

1. Conda environments 针对系统全局，但是 venv 通常只针对单项目

### 0x02b Channels[^3]

*Channels are the locations where packages are stored. They serve as the base for hosting and managing packages.*

Channels 类似与 PyPI, 是包仓库。Conda 会从这些 Channels 将包下载到本地
主要有如下 Channels
- Default Channel 
	默认使用的 channel(由多个 [channels](https://repo.anaconda.com/pkgs/) 组成)，包数量较小，但是包绝对安全
- Conda-forge
	社区驱动的 channel，包数量多(几乎等价于 PyPI，但是某些包版本可能不全)，但是包不一定绝对安全

不同 channels 中可以含有相同版本的包

#### Channels VS PyPI

1. Channels 中的包和 PyPI 的包格式不同，不通用
2. 在 PyPI 中的包，Channels 不一定有


---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]: [docs.conda](https://docs.conda.io/en/latest/)
[^2]: [Environments — conda 24.5.0 documentation](https://docs.conda.io/projects/conda/en/stable/user-guide/concepts/environments.html)
[^3]: [Channels — conda 24.5.0 documentation](https://docs.conda.io/projects/conda/en/stable/user-guide/concepts/channels.html)
