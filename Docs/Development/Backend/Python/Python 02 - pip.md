---
createTime: 2024-07-02 16:52
tags:
  - "#Python"
---

# Python 02 - pip

## 0x01 Overview

`pip` 是 Python 的一个包管理器(其实就是一个 module)，默认会从 PyPI 中下载包

## 0x02 Installation/Upgradeing pip[^1]

> [!important] 
> 在 python 3.11.0 之后，Python 不允许直接将 packages 安装在非 virtual environment 中(出于安全，以及系统冲突的原因)，所以官方的大多数安装或者更新方式都会失效，同理 `pip install`。如果想要 packages 能在系统全局使用，就需要使用包管理器下载，在 Arch 中包通常以 `python-*` 命名,以 `pip` 为例即为 `python-pip`，或者删除 `/usr/lib/python3.x/EXTERNALLY-MANAGED` [^2]

## 0x02 Usage

> 基础使用方法直接可以查看官方文档[^3]

**requirement specifier**[^4]

在介绍使用方法之前，需要了解 requirement specifier，`pip` 大多 commands 都会用到 requirement specifier

为了方便记忆，定义如下 EBNF
```
requirement specifier=package-name[operator version]
package-name = [a-zA-z]*
operator = "==" | ">=" | "<=" | ">" | "<"
version = [1-0]*
```

例如
```
selenium == 4.12.0
selenium > 4.12.0
```

如果 requirement specifier 和 commands 一起使用，最好使用 quote(一些特殊符号在 Shell 中有特殊含义)
```
pip install "request>=2.8.1"
```
如果在 requirements file 中使用，就无需使用 quote

### 0x02a install/uninstall

#### install

安装包
```
python -m pip install [options] <requirement specifier> [package-index-options] ...
python -m pip install [options] -r <requirements file> [package-index-options] ...
```

eg
```
python -m pip install "requests>=2.8.1"
```

当使用 `-r` 时，requirements file 的格式参考 [Requirements File Format - pip documentation v24.1.1](https://pip.pypa.io/en/stable/reference/requirements-file-format/#requirements-file-format)

eg
```
pytest
docopt == 0.6.1
requests >= 2.8.1
```

#### uninstall

卸载包
```
python -m pip uninstall [options] <package> ...
python -m pip uninstall [options] -r <requirements file> ...
```

使用上和 `install` 类似

### 0x02b list

列出已经安装的包
```
python -m pip list [options]
```

### 0x02c show

查看包的详细信息，安装目录，依赖
```
python -m pip list [options]
```

eg
```
$ pip show trash-cli
Name: trash-cli
Version: 0.24.5.26
Summary: Command line interface to FreeDesktop.org Trash.
Home-page: https://github.com/andreafrancia/trash-cli
Author: Andrea Francia
Author-email: andrea@andreafrancia.it
License: GPL v2
Location: /usr/lib/python3.12/site-packages
Requires: psutil, six
Required-by:
```

### 0x02d freeze

将当前安装的包的列表，以 requirement files 的格式输出
```
python -m pip freeze [options]
```

### 0x02e search

PyPI 已经不支持直接使用 `pip search` 的功能了


## 0x03 Where to install[^5]

pip 安装包的位置，根据是否有 venv 或者是 conda 虚拟环境而不同

可以使用 `pip show <package>` 来校验

1. 如果没有在 venv 或者是 conda 虚拟环境中，地址为 `/usr/lib/pythonx.y.z/site-packages`
2. 如果在 venv 中，地址为 `/path/to/venv/lib/pythonx.y.z/site-packages`
3. 如果在 conda 中，地位为 `/path/to/env/lib/pythonx.y.z/site-packages`

## 0x04 pip VS pip3

在 Linux 上只和 Hashbage 有关

```
readlink -f $(which {pip,pip3})
       /usr/bin/pip
       /usr/bin/pip3

cat /usr/bin/pip && echo &&  cat /usr/bin/pip3
#!/usr/bin/python
# -*- coding: utf-8 -*-
import re
import sys
from pip._internal.cli.main import main
if __name__ == "__main__":
    sys.argv[0] = re.sub(r"(-script\.pyw|\.exe)?$", "", sys.argv[0])
    sys.exit(main())

#!/usr/bin/python
# -*- coding: utf-8 -*-
import re
import sys
from pip._internal.cli.main import main
if __name__ == "__main__":
    sys.argv[0] = re.sub(r"(-script\.pyw|\.exe)?$", "", sys.argv[0])
    sys.exit(main())


```

这里可以发现 `pip` 和 `pip3` 的 Hashbang 都一致，那么在 Linux 上如果就会直接使用 `/usr/bin/python` 来解析 `pip` 和 `pip3` ，那么结果都一样

## 0x05 python -m pip VS pip

> `python -m` 会从 `sys.path` 找到对应的 module，并执行对应 moudule 的 `main()` 函数

由 [0x04 pip VS pip3](#0x04%20pip%20VS%20pip3) 可知 `pip` 由 Hashbang 控制，所以

- 当 `pip` Hashbang 等于 `which python` 时两者等价
- 当 `pip` Hashbang 不等于 `which python` 时两者不同

在大多数情况下两者都是相等的，安装 Python 时，会安装对应版本的 `pip`

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[Installation - pip documentation v24.1.1](https://pip.pypa.io/en/stable/installation/)
[^2]:[python - How do I solve "error: externally-managed-environment" every time I use pip 3? - Stack Overflow](https://stackoverflow.com/questions/75608323/how-do-i-solve-error-externally-managed-environment-every-time-i-use-pip-3)
[^3]:[User Guide - pip documentation v24.1.1](https://pip.pypa.io/en/stable/user_guide/#)
[^4]:[Requirement Specifiers - pip documentation v24.1.1](https://pip.pypa.io/en/stable/reference/requirement-specifiers/#requirement-specifiers)
[^5]:[site — Site-specific configuration hook — Python 3.12.4 documentation](https://docs.python.org/3/library/site.html)