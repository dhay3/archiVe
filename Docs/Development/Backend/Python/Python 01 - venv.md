---
createTime: 2024-07-01 14:05
tags:
  - "#Python"
---

# Python 01 - venv

## 0x01 Overview

venv 是 Python 中一个用于创建 virtual environment 的 **built-in module**

- Used to contain a specific Python interpreter and software libraries and binaries which are needed to support a project (library or application). These are by default isolated from software in other virtual environments and Python interpreters and libraries installed in the operating system.
    
- Contained in a directory, conventionally either named `venv` or `.venv` in the project directory, or under a container directory for lots of virtual environments, such as `~/.virtualenvs`.
    
- Not checked into source control systems such as Git.
    
- Considered as disposable – it should be simple to delete and recreate it from scratch. You don’t place any project code in the environment
    
- Not considered as movable or copyable – you just recreate the same environment in the target location.

## 0x02 Usage

venv 使用的方法很简单

使用 venv module 创建一个 virtual environment

```shell
python -m venv /path/to/venv
```

activate virtual environment

```shell
source /path/to/venv/bin/activate
```

激活后会接管 PATH env，因为接管 PATH env，当调用 python 时会直接使用 venv 里的 python，而非系统的

```
(venv) cc in ~/AI/stable-diffusion-webui on master ● λ readlink -f $(which python)
/home/cc/anaconda3/bin/python3.11
```

同时 `pip install` 的包也只会在当前环境中生效，因为 `sys.prefix` 为 venv 的地址

如果想要 deactivate virtual environment，直接输入 `deactivate` 即可

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[venv — Creation of virtual environments — Python 3.12.4 documentation](https://docs.python.org/3/library/venv.html)
[^2]:[Python Virtual Environments: A Primer – Real Python](https://realpython.com/python-virtual-environments-a-primer/)