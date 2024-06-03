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

Stable Diffusion WebUI 安装很简单，只需要一行命令[^1]

```shell
sudo pacman -S git python3 -y && git clone https://github.com/AUTOMATIC1111/stable-diffusion-webui && cd stable-diffusion-webui && ./webui.sh
```

下载会需要一点时间，默认会下载 Stablediffusion V1-5 safesensor

### xformers

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

## Usage



---

*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]: [GitHub - AUTOMATIC1111/stable-diffusion-webui: Stable Diffusion web UI](https://github.com/AUTOMATIC1111/stable-diffusion-webui)
[^2]: [Xformers · AUTOMATIC1111/stable-diffusion-webui Wiki · GitHub](https://github.com/AUTOMATIC1111/stable-diffusion-webui/wiki/Xformers)