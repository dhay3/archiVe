---
title: Conda 03 - Configuration
author: "0x00"
createTime: 2024-06-02
lastModifiedTime: 2024-06-02-16:32
draft: false
tags:
  - Python
---

# Conda 03 - Configuration

# 0x01 Overview


> [!tip]
> 如果需要看所有的配置项，看官方文档[^1]

Conda 默认会读取 `~/.condarc` 作为配置文件（使用 yaml）。主要有如下几个常用的配置项

1. channels [list]

	conda 使用的 channels

3. auto_activate_base [bool]

	在初始化 Shell 的时候，是否自动激活 base environment
2. changeps1 [bool]

	在切换或者自动激活 environment 时，是否修改 Shell PROMPT
3. env_prompt [str]

	使用 environments 时，Shell 关于 conda 部分的 PROMPT

## conda config

`conda config` 是修改 `.condarc` 的一个用户态工具，类似于 `git config`。可以和如下参数一起使用

- `--show`

	查看所有的配置
- `--set`

	设置一个 bool 或者 string 参数的值


## 0x02 Exmaple

```yaml
channels:
  - defaults
env_prompt: "╭─ conda on ({default_env})\n╰─ "
auto_activate_base: false
```

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[Configuration — conda 24.5.0 documentation](https://docs.conda.io/projects/conda/en/stable/configuration.html)
[^2]:[conda config — conda 24.5.0 documentation](https://docs.conda.io/projects/conda/en/stable/commands/config.html)

