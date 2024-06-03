---
title: Conda 01 - Overview
author: "0x00"
createTime: 2024-05-30
lastModifiedTime: 2024-05-30-08:54
draft: true
tags:
  - Python
---

# Conda 01 - Overview

## 0x01 Overview

Conda 是一个包管理器，最早只被用在 Python 中，用于解决在不同 Data Science 项目中依赖的问题。发展至今，Conda 已经不仅先于 Python, 其他编程语言也可以使用 Conda 来管理[^1]

## 0x02 Terms

### Environments[^2]

*An environment is a directory that contains a specific collection of packages that you have installed. For example, you may have one environment with NumPy 1.7 and its dependencies, and another environment with NumPy 1.6 for legacy testing. If you change one environment, your other environments are not affected. You can easily activate or deactivate environments, which is how you switch between them.*

Environments 是 Conda 中的一个概念，是包和软件的集合。每个 Environment 之间互相隔离，Environments 中的包和软件只对当前 Environment 生效。类似与 Python 中的 venv 的概念

### Channels[^3]

*Channels are the locations where packages are stored. They serve as the base for hosting and managing packages.*

Channels 类似与 Arch 中的 package mirror, 是包仓库。Conda 会从这些 Channels 将包下载到本地。主要有如下 Channels
- Default Channel
- Conda-forge(社区免费的)

Conda 默认会先从 Default Channel 下载包，如果没有就会从 Conda-forge 下载

## 0x03 Conda VS PIP[^4]

venv 更像只针对单项目，conda 针对系统全局

在 Conda 环境下安装任何需要放到 PATH 目录下的软件会放到 Conda 的目录中

Conda 大体和 `pip` 相同都是包管理器，主要有如下区别
1. Conda 可以安装 Python packages 和 Softwares，但是 `pip`  只能安装 Python packages。例如 Conda 可以安装 Python 本体，但是 `pip` 不能
2. `pip` 本身没有隔离环境的功能，需要依赖 venv 模块，而 Conda 通过更简单的方式来隔离环境
3. Conda 从 Channel 获取包，而 `pip` 从 PyPI 获取包
4. Conda 使用 binary 做为 packages format，而 `pip` 使用 wheels 或者 sources。所以两者直接的包互不影响
5. Conda 可以安装的包比 `pip` 少，但是可以通过 `conda build` 来构建
6. 在未使用 venv 或者是 Conda 的情况下(`python`) `pip install` 安装的 packages 作用用全局，而 `conda install` 安装的 packages 只作用有当前 environment[^5]。如果你使用 Conda 安装了 Python，那么系统环境的 Python 就会被 Conda 接管(具体可以看 PATH 的值变化)，在调用 `pip install` 几乎等价于 `conda install` 两者没有区别

## 0x04 PIP in Conda[^6]

*Running conda after pip has the potential to overwrite and potentially break packages installed via pip. Similarly, pip may upgrade or remove a package which a conda-installed package requires.*

`pip` 和 Conda 一起使用，可能会出现 packages 覆盖的问题。所以为了避免这个问题的出现，应该尽量只使用 `pip` 或者是 Conda

默认不使用 Conda， 使用 Conda 通过命令来切换

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]: [docs.conda](https://docs.conda.io/en/latest/)
[^2]: [Environments — conda 24.5.0 documentation](https://docs.conda.io/projects/conda/en/stable/user-guide/concepts/environments.html)
[^3]: [Channels — conda 24.5.0 documentation](https://docs.conda.io/projects/conda/en/stable/user-guide/concepts/channels.html)
[^4]: [Anaconda | Understanding Conda and Pip](https://www.anaconda.com/blog/understanding-conda-and-pip)
[^5]: [python - pip install vs conda install - Stack Overflow](https://stackoverflow.com/questions/65536064/pip-install-vs-conda-install)
[^6]: [Anaconda | Using Pip in a Conda Environment](https://www.anaconda.com/blog/using-pip-in-a-conda-environment)