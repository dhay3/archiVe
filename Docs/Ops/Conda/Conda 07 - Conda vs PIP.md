---
createTime: 2024-07-02 09:16
tags:
  - "#hash1"
  - "#hash2"
---

# Conda 05 - Conda vs PIP

## 0x01 Overview

## 0x03 Conda VS PIP[^1]

> [!important]
> Conda 环境中安装任何需要放到 PATH 目录(安装软件时没有使用绝对路径，例如 `/usr/local/bin/`，而是直接使用 `${PATH}` )下的软件，都会放到 Conda 的目录中(具体可以查看 `${PATH}` 变量)

### 0x03a Similarities 

1. Conda 和 `pip` 都是包管理器

### 0x03b Differences

1. Conda 从 Channels 中获取包，而 `pip` 从 PyPI 获取包
2. `pip` 可以安装的包，Conda 不一定能安装(在 PyPI 中的包，Channels 不一定有)
3. Conda 使用 binary 做为 packages format，而 `pip` 使用 wheels 或者 sources
4. Conda 可以安装 Python packages 和 Softwares，但是 `pip`  只能安装 Python packages。例如 Conda 可以安装 Python 本体，但是 `pip` 不能。如果你使用 Conda 安装了 Python，那么系统环境的 Python 就会被 Conda 接管(可以使用 `readlink -f $(which python)` 来校验)
5. `pip` 本身没有隔离环境的功能，需要依赖 venv 模块，而 Conda 通过 Environments 来隔离环境
6. Conda 支持命令行 search package，`pip` 不支持
7. 在未使用 venv 或者是 Conda 的情况下(`python`) `pip install` 安装的 packages 作用用全局，而 `conda install` 安装的 packages 只作用有当前 environment[^2]。

## 0x04 pip install/conda install

`pip install` 和 `conda install` 一起使用

1. 当包只在 PyPI 中，即 Conda channels 没有对应的包时。可以使用 `pip install` 安装，包会在 `/path/to/env/lib/pythonx.y.z`
    例如
    
    ```shell
    ╭─ conda on (lab3)
    ╰─ cc in ~/anaconda3/envs/lab λ conda search shorten-url
    Loading channels: done
    No match found for: shorten-url. Search: *shorten-url*
    ....
     
    ╭─ conda on (lab3)
    ╰─ cc in ~/anaconda3/envs/lab λ pip install shorten-url
    ...
    
    ╭─ conda on (lab3)
    ╰─ cc in ~/anaconda3/envs/lab λ pip show shorten-url
    <frozen graalpy.pip_hook>:48: RuntimeWarning: You are using an untested version of pip. GraalPy provides patches and workarounds for a number of packages when used with compatible pip versions. We recommend to stick with the pip version that ships with this version of GraalPy.
    Name: shorten-url
    Version: 1.0.0
    Summary: Python Library to help you short and expand urls using https://rel.ink/
    Home-page: https://github.com/Quarantine-Projects/shorten_url
    Author: Salil Gautam
    Author-email:
    License: MIT
    Location: /home/cc/anaconda3/envs/lab3/lib/python3.10/site-packages
    Requires: requests
    Required-by:
    ```
2. 当包只在 Conda Channels 中，即 PyPI 中没有对应的包时(通常是软件)。可以使用 `conda install` 安装，包会在 `/path/to/env/bin`
    例如
    
    ```shell
    ╭─ conda on (lab3)
    ╰─ cc in ~/anaconda3/envs/lab λ conda install "openjdk=8.0.152"
    ...
    ╭─ conda on (lab3)
    ╰─ cc in ~/anaconda3/envs/lab λ which java
    /home/cc/anaconda3/envs/lab3/bin/java
    ```

3. 当包在 Conda Channels 和 PyPI 中， 通常都是 Python package。可以使用 `conda install` 或者是 `pip install` 安装，包会在 `/path/to/env/lib/pythonx.y.z`

   > [!important]
   >
   > 如果包是以 `pip install` 的方式安装的，`conda list <package>` 中 Channel 会以 pypi 标识

   ```shell
   ╭─ conda on (lab3)
   ╰─ cc in ~/anaconda3/envs/lab λ pip install "requests==2.32.1"
   ...
   ╭─ conda on (lab3)
   ╰─ cc in ~/anaconda3/envs/lab λ conda list requests
   # packages in environment at /home/cc/anaconda3/envs/lab3:
   #
   # Name                    Version                   Build  Channel
   requests                  2.32.1                   pypi_0    pypi
   
   ╭─ conda on (lab3)
   ╰─ cc in ~/anaconda3/envs/lab λ conda install "requests=2.28.0"
   ...
   
   ╭─ conda on (lab3)
   ╰─ cc in ~/anaconda3/envs/lab λ conda list requests
   # packages in environment at /home/cc/anaconda3/envs/lab3:
   #
   # Name                    Version                   Build  Channel
   requests                  2.32.1                   pypi_0    pypi
   
   ```
   
   conda 还是会显示 pip 安装的版本，这显然是不合符逻辑的
   ==正确的操作是，如果一个包在 Conda Channels 和 PyPI 中，应该只使用 `conda install` 或者 `pip install` ==

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]: [Anaconda | Understanding Conda and Pip](https://www.anaconda.com/blog/understanding-conda-and-pip)
[^2]: [python - pip install vs conda install - Stack Overflow](https://stackoverflow.com/questions/65536064/pip-install-vs-conda-install)