---
title: Stable Diffusion WebUI
author: "0x00"
createTime: 2024-05-30
lastModifiedTime: 2024-05-30-14:05
draft: true
tags:
  - AI
  - Stable Diffusion
---
# Stable Diffusion WebUI


## Installation

### stable-diffusion-webui[^1]

> [!NOTE]
> ~~不推荐使用 conda 安装，defaults 以及 conda-forge channel 中缺少很多对应版本的包~~
> 推荐直接使用 conda 安装！！！

安装环境

```shell
sudo pacman -S git python3 -y
```

克隆 repository 

```shell
git clone https://github.com/AUTOMATIC1111/stable-diffusion-webui && cd stable-diffusion-webui
```

直接运行 `webui.sh` 启动 Web UI 即可
无需使用 venv 单独创建一个环境，启动脚本 webui.sh 中已实现隔离，以及依赖的安装(具体看源码 `webui.sh`，`launch_utils.py`)

```shell
./webui.sh
```

默认会下载 [stable-diffusion-v1-5 model](https://huggingface.co/runwayml/stable-diffusion-v1-5)。如果已经有模型了，可以直接将模型放到 `models/stable-diffusion` 下，会跳过下载 sd1.5

### xformers

> 可选

为了加速图片的生成，还可以安装 meta xformers[^2]
这里使用 `conda`
1. go to the webui directory
2. `source ./venv/bin/activate` 如果使用来 Conda 来管理环境，可以跳过
3. `cd repositories`
4. `git clone https://github.com/facebookresearch/xformers.git`
5. `cd xformers`
6. `git submodule update --init --recursive`
7. `pip install -r requirements.txt`
8. `pip install -e .`

### Trouble shooting

1. 运行 `webui.sh` 时出现 `version GLIBCXX_3.4.30 not found (required by /usr/lib/libtcmalloc_minimal.so.4)`
	建议直接将 `LD_PRELOAD=/lib/libstdc++.so.6` 写入到 `webui-user.sh`

2. 运行 `webui.sh` 时出现 `cannot import name 'Undefined' from 'pydantic.fields'` [^3]

## Usage

### Low VRAM Video Cards

当 VRAM 小于 4GB 时，可以通过 `--lowram`，`--medram`，`--xformers` 来降低速度以换取空间


---

*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]: [GitHub - AUTOMATIC1111/stable-diffusion-webui: Stable Diffusion web UI](https://github.com/AUTOMATIC1111/stable-diffusion-webui)
[^2]: [Xformers · AUTOMATIC1111/stable-diffusion-webui Wiki · GitHub](https://github.com/AUTOMATIC1111/stable-diffusion-webui/wiki/Xformers)
[^3]:[ImportError: cannot import name 'Undefined' from 'pydantic.fields' (D:\\a1111\\stable-diffusion-webui\\venv\\lib\\site-packages\\pydantic\\fields.py) · AUTOMATIC1111/stable-diffusion-webui · Discussion #15557 · GitHub](https://github.com/AUTOMATIC1111/stable-diffusion-webui/discussions/15557)