---
createTime: 2024-07-02 09:29
tags:
  - "#hash1"
  - "#hash2"
---

# Conda 04 - Package Search

## 0x01 Overview

Conda 提供了一个接口 `conda search` 用于查询 Channels 中的包的(这是 `pip` 不具备的)

## 0x02 conda search

conda package 根据 5 元组(channel/subdir/name/version/build)来精确到一个包

![](https://docs.conda.io/projects/conda/en/stable/_images/conda_search.png)

其中 name 是必需要提供的，其余的可以选择性提供
1. channel
	(Optional) Can either be a channel name or URL. Channel names may include letters, numbers, dashes, and underscores.
	可以使用按照如下方式过滤 channel
1. subdir
	(Optional) A subdirectory of a channel. Many subdirs are used for architectures, but this is not required. Must have a channel and backslash preceding it. For example: `main/noarch`
	通常 subdir 对应 OS/CPU architect，可以按照如下方式过滤 subdir
1. name
	(Required) Package name. May include the `*` wildcard. For example, `*py*` returns all packages that have "py" in their names, such as "numpy", "pytorch", "python", etc.
	可以按照如下方式过滤 name
1. version
	(Optional) Package version. May include the `*` wildcard or a version range(s) in single quotes. For example: `numpy=1.17.*` returns all numpy packages with a version containing "1.17." and `numpy>1.17,<1.19.2` returns all numpy packages with versions greater than 1.17 and less than 1.19.2.
	可以按照如下方式过滤 version
1. build
	(Optional) Package build name. May include the `*` wildcard. For example, `numpy 1.17.3 py38*` returns all version 1.17.3 numpy packages with a build name that contains the text "py38".
	一般用不到，可以按照如下方式过滤 build

按照上面的条件你可以使用如下方式(也被称为 standard specification)来过滤包，其中 space 等价于 `=`

> [!NOTE]
> conda 只有一个 positional args 需要使用 double quotes 将 condition 包裹

```shell
conda search "[[channel][/subdir]][::]<name>[ |=|>|<|>=|<=][version] [build]"
```

当指定了 channel 或者是 subdir 时，必须要使用 `::`
如果没有指定 channel 或者是 subdir，则可以不使用 `::`
例如

```shell
conda search "conda-forge::selenium>=4.21.0"
conda search "selenium>=4.21.0"
conda search "::selenium>=4.21.0"
```

或者使用 key-value pair notaion

```shell
conda search "<name>\[[channel=chanel],[subdir=subdir],[name=name],[version=version],[build=build]\]"
```

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[Package search and install specifications — conda 24.5.0 documentation](https://docs.conda.io/projects/conda/en/stable/user-guide/concepts/pkg-search.html)z