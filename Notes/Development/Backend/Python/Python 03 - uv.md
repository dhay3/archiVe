---
createTime: 2025-10-24 10:09
license: cc by 4.0
tags: 
 - "#hash1" 
 - "#hash2"
---

# Python 03 - uv

## 0x01 Preface[^2]

> 号称比 pip 快 10 - 100 倍

`uv` 是一个 Rust 编写的 Python 项目以及包管理器，可以替代 `pip`, `pip-tools`, `pipx`, `poetry`, `pyenv`, `twine`, `virtualenv` 等工具

## 0x02 Usage[^1]

```
An extremely fast Python package manager.

Usage: uv [OPTIONS] <COMMAND>

Commands:
  auth     Manage authentication
  run      Run a command or script
  init     Create a new project
  add      Add dependencies to the project
  remove   Remove dependencies from the project
  version  Read or update the project's version
  sync     Update the project's environment
  lock     Update the project's lockfile
  export   Export the project's lockfile to an alternate format
  tree     Display the project's dependency tree
  format   Format Python code in the project
  tool     Run and install commands provided by Python packages
  python   Manage Python versions and installations
  pip      Manage Python packages with a pip-compatible interface
  venv     Create a virtual environment
  build    Build Python packages into source distributions and wheels
  publish  Upload distributions to an index
  cache    Manage uv's cache
  self     Manage the uv executable
  help     Display documentation for a command

```

### 0x02a init

创建一个新的 Python 项目

```
uv init example
Initialized project `example` at `/home/cc/PlayGround/example`

tree -a -L 1 example
example
├── .git
├── .gitignore
├── main.py
├── pyproject.toml
├── .python-version
└── README.md
```

### 0x02b add

几乎等价于 `pip install` 会将依赖写入 `pyproject.toml`

```
uv add requests==2.31.0
Resolved 6 packages in 524ms
Installed 5 packages in 10ms
 + certifi==2025.10.5
 + charset-normalizer==3.4.4
 + idna==3.11
 + requests==2.31.0
 + urllib3==2.5.0
```

### 0x02c remove

几乎等价于 `pip uninstall` 会将依赖从 `pyproject.toml` 中移除

```
uv remove requests
Resolved 1 package in 2ms
Uninstalled 5 packages in 2ms
 - certifi==2025.10.5
 - charset-normalizer==3.4.4
 - idna==3.11
 - requests==2.31.0
 - urllib3==2.5.0
```

### 0x02d tree

以树结构展示 `pip list`

```
uv tree
Resolved 6 packages in 0.58ms
example v0.1.0
└── requests v2.31.0
    ├── certifi v2025.10.5
    ├── charset-normalizer v3.4.4
    ├── idna v3.11
    └── urllib3 v2.5.0
```

### 0x02e run

> [!NOTE]
> 第一次运行会自动生成 venv，也可以是使用 `uv venv` 手动创建

在 `uv` 构建的环境中运行特定脚本，等价于 `.venv/bin/python <scripts>`

```
uv run main.py
Hello from example!
```

### 0x03f sync

`uv` 特有的指令，确保当前环境中的依赖符合 `requirements.txt` 或者 `pyproject.toml` 中的版本要求

```
uv sync
Resolved 2 packages in 0.53ms
Uninstalled 4 packages in 1ms
 - certifi==2025.10.5
 - charset-normalizer==3.4.4
 - idna==3.11
 - urllib3==2.5.0
```

### 0x03g pip

`uv` 也提供了一套兼容 `pip` 的接口

```
uv pip install
uv pip uninstall
uv pip list
uv pip show
uv pip freeze
```

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*



***See also***

- [uv](https://docs.astral.sh/uv/)
- [Working on projects \| uv](https://docs.astral.sh/uv/guides/projects/)
- [GitHub - astral-sh/uv: An extremely fast Python package and project manager, written in Rust.](https://github.com/astral-sh/uv)



***References***

[^1]:[Features \| uv](https://docs.astral.sh/uv/getting-started/features/)
[^2]:[uv](https://docs.astral.sh/uv/)


