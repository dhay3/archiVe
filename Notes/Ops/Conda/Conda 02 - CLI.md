---
title: Conda 02 - CLI
author: "0x00"
createTime: 2024-05-30
lastModifiedTime: 2024-05-30-09:03
draft: true
tags:
  - Python
---

# Conda 02 - CLI

## 0x01 Overview

Conda 大体上的命令和 `pip` 类似，可以参考对比下表[^1]

|Task|Conda package and environment manager command|Pip package manager command|Virtualenv environment manager command|
|---|---|---|---|
|Install a package|`conda install $PACKAGE_NAME`|`pip install $PACKAGE_NAME`|X|
|Update a package|`conda update --name $ENVIRONMENT_NAME $PACKAGE_NAME`|`pip install --upgrade $PACKAGE_NAME`|X|
|Update package manager|`conda update conda`|Linux/macOS: `pip install -U pip` Win: `python -m pip install -U pip`|X|
|Uninstall a package|`conda remove --name $ENVIRONMENT_NAME $PACKAGE_NAME`|`pip uninstall $PACKAGE_NAME`|X|
|Create an environment|`conda create --name $ENVIRONMENT_NAME python`|X|`cd $ENV_BASE_DIR; virtualenv $ENVIRONMENT_NAME`|
|Activate an environment|`conda activate $ENVIRONMENT_NAME`*|X|`source $ENV_BASE_DIR/$ENVIRONMENT_NAME/bin/activate`|
|Deactivate an environment|`conda deactivate`|X|`deactivate`|
|Search available packages|`conda search $SEARCH_TERM`|`pip search $SEARCH_TERM`|X|
|Install package from specific source|`conda install --channel $URL $PACKAGE_NAME`|`pip install --index-url $URL $PACKAGE_NAME`|X|
|List installed packages|`conda list --name $ENVIRONMENT_NAME`|`pip list`|X|
|Create requirements file|`conda list --export`|`pip freeze`|X|
|List all environments|`conda info --envs`|X|Install virtualenv wrapper, then `lsvirtualenv`|
|Install other package manager|`conda install pip`|`pip install conda`|X|
|Install Python|`conda install python=x.x`|X|X|
|Update Python|`conda update python`*|X|X|

也可以参考 Cheatsheet[^2]

## 0x02 Environment relative

### 0x02a conda create

用于创建 environment

```shell
conda create -n <ENVNAME>
```

还可以在创建 environment 时安装指定的包
```shell
conda create -n <ENVNAME> python=3.10.0 selenium=4.22.0
```

还可以使用 `--clone` 来克隆 environment
```shell
conda create -n <DSTENV> --clone <SRCENV>
```

### 0x02b conda activate

用于 activate 指定 environment
```shell
conda acitvate <ENVNAME>
```

### 0x02c conda deactivate

用于 deactivate 指定 environment 
```shell
conda deactivate <ENVNAME>
```
如果在一个 environment 中 activate 另外一个 environment，deactivate 不会直接退出到 shell
如果想要直接退出到 shell，建议直接关掉当前 shell

### 0x02d conda env

`conda evn` 由多个 subcommands 组成
1. config
2. create
3. export
4. list
5. remove
6. update

这里只挑几个介绍，具体看 conda 官方文档

**create**

根据 environment definition file(通常是一个 yaml 文件) 创建 environment
```shell
conda env create -f /path/to/environment.yml -n <ENVNAME>
```
如果在 environment definition file 定义了 `name: <ENVNAME>` 可以无需使用 `-n`

**export**

根据当前 environment 中安装的 package，导出 environment definition file
```shell
conda env export --file /path/to/environment.yml
```

**list**

展示所有的 environments
```shell
conda env list
```

**remove**

删除指定的 environment
```shell
conda env remove -n <ENVNAME>
```

**update**

根据 environment difinition file 更新当前 environment
```shell
conda env update -f=/path/to/environment.yml
```

### 0x02e conda rename

重命名 environment
```shell
conda rename -n OLDNAME NEWNAME
```

## 0x03 Package relative

### 0x03a conda install

从 channels 中安装包，可以安装多个包
```
conda install <mathspec[...]>
```

可以参考 [Conda 06 - Package Install](Conda%2006%20-%20Package%20Install.md)
eg.
```
conda install "conda-forge::selenium=4.1.0" "defaults::lxml=5.2.1"
conda install "selenium=4.1.0" "lxml=5.2.1"
```

### 0x03a conda search

从 channels 中搜索包
```
conda search <matchspec>
```

可以参考 [Conda 05 - Package Search](Conda%2005%20-%20Package%20Search.md)
eg.
```
conda search conda-forge::scapy
conda search conda-forge::scapy=2.4.3
conda search scapy=2.4.5
```

### 0x03c conda remove

删除安装的包，可以删除多个包
```
conda remove <mathspec[...]>
```

eg.
```
conda remove xz zlib
```

### 0x03d conda update

更新 package 到最新 compatible version

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]: [Commands — conda 24.5.0 documentation](https://docs.conda.io/projects/conda/en/stable/commands/index.html)
[^2]: [Cheat sheet — conda 24.5.0 documentation](https://docs.conda.io/projects/conda/en/stable/user-guide/cheatsheet.html)

