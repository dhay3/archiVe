---
createTime: 2024-07-02 09:16
tags:
  - "#hash1"
  - "#hash2"
---

# Conda 05 - Conda vs PIP

## 0x01 Overview

## 0x03 Conda VS PIP[^4]

> Conda 环境中安装任何需要放到 PATH (安装软件时没有使用绝对路径，例如 `/usr/local/bin/`，而是直接使用 `${PATH}` )目录下的软件，都会放到 Conda 的目录中

### 0x03a Similarities 

1. Conda 和 `pip` 都是包管理器

### 0x03b Differences

1. Conda 从 Channels 中获取包，而 `pip` 从 PyPI 获取包
2. `pip` 可以安装的包，Conda 不一定能安装(在 PyPI 中的包，Channels 不一定有)
3. Conda 使用 binary 做为 packages format，而 `pip` 使用 wheels 或者 sources。所以两者直接的包互不影响
4. Conda 可以安装 Python packages 和 Softwares，但是 `pip`  只能安装 Python packages。例如 Conda 可以安装 Python 本体，但是 `pip` 不能。如果你使用 Conda 安装了 Python，那么系统环境的 Python 就会被 Conda 接管(可以使用 `readlink -f $(which python)` 来校验)
5. `pip` 本身没有隔离环境的功能，需要依赖 venv 模块，而 Conda 通过 Environments 来隔离环境
6. Conda 支持命令行 search package，`pip` 不支持
7. 在未使用 venv 或者是 Conda 的情况下(`python`) `pip install` 安装的 packages 作用用全局，而 `conda install` 安装的 packages 只作用有当前 environment[^5]。

## 0x04 PIP in Conda[^6]

*Running conda after pip has the potential to overwrite and potentially break packages installed via pip. Similarly, pip may upgrade or remove a package which a conda-installed package requires.*

`pip` 和 Conda 一起使用，可能会出现 packages 覆盖的问题。所以为了避免这个问题的出现，应该尽量只使用 `pip` 或者是 Conda

默认不使用 Conda， 使用 Conda 通过命令来切换

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^4]: [Anaconda | Understanding Conda and Pip](https://www.anaconda.com/blog/understanding-conda-and-pip)
[^5]: [python - pip install vs conda install - Stack Overflow](https://stackoverflow.com/questions/65536064/pip-install-vs-conda-install)
[^6]: [Anaconda | Using Pip in a Conda Environment](https://www.anaconda.com/blog/using-pip-in-a-conda-environment)


