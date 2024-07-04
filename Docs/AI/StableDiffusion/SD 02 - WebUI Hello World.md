---
createTime: 2024-05-30
tags:
  - AI
  - Stable-Diffusion
---
# SD 02 - WebUI Hello World

## 0x01 Overview

[GitHub - AUTOMATIC1111/stable-diffusion-webui: Stable Diffusion web UI](https://github.com/AUTOMATIC1111/stable-diffusion-webui) 是一个 SD 高客制化开箱即用的 Web UI

在食用前先看 [SD 01 - General Terms](SD%2001%20-%20General%20Terms.md)，了解一些 Terms

## 0x02 Installation[^1]

Web UI 依赖 Python3.10，使用 Conda 创建一个 Python3.10 的虚拟环境
```shell
conda create -n sd "python=3.10 git" -y
```

切换到新建的 environment
```shell
conda activate sd
```

克隆 web-ui repository 
```shell
git clone https://github.com/AUTOMATIC1111/stable-diffusion-webui && cd stable-diffusion-webui
```

可以直接运行 `lanuch.py`，将依赖安装到 conda virtual environment 中
```shell
python lanuch.py
```

也可以选择运行 `webui.sh`，会自动创建一个 venv (具体可以看 `webui.sh` 和 `launch_utils.py`)
```shell
./webui.sh
```

这里更推荐通过 `webui.sh` 运行(==如果系统 python 在 3.11 之前，忘记切换 conda 环境就会在系统上安装 Web UI 的依赖==)

默认会下载 [stable-diffusion-v1-5 model](https://huggingface.co/runwayml/stable-diffusion-v1-5)。如果已经有模型了，可以直接将模型放到 `models/stable-diffusion` 下，会跳过下载 sd1.5

### 0x02a xformers[^2]

xformers 是 meta 的一个 library，可以加快图片的生成以及降低 VRAM 的使用

> [!important]
> 在 20230123 之后不需要手动编译 xformers 了，可以直接通过 `--xformers` 实现自动安装
> eg `./webui.sh --xformers`

1. go to the webui directory
2. `source ./venv/bin/activate` 如果使用来 Conda 来管理环境，可以跳过
3. `cd repositories`
4. `git clone https://github.com/facebookresearch/xformers.git`
5. `cd xformers`
6. `git submodule update --init --recursive`
7. `pip install -r requirements.txt`
8. `pip install -e .`

### 0x02b Trouble shooting

1. 运行 `webui.sh` 时出现 `version GLIBCXX_3.4.30 not found (required by /usr/lib/libtcmalloc_minimal.so.4)`[^4]
	建议直接将 `LD_PRELOAD=/lib/libstdc++.so.6` 写入到 `webui-user.sh`

2. 运行 `webui.sh` 时出现 `cannot import name 'Undefined' from 'pydantic.fields'` [^3]
	```
	pip install "pydantic==1.10.15" "albumentations==1.4.3" 
	```
	可以直接将其写入到 `requirements.txt` 中

## 0x03 ENV and Options[^5]

### 0x03a ENV

> 这里只列出常用的 ENV，具体可以看 wiki

Web UI 有如下几个常用的 ENV

- COMMANDLINE_ARGS
	Additional commandline arguments for the main program.
- SD_WEBUI_LOG_LEVEL
	Log verbosity. Supports any valid logging level supported by Python's built-in `logging` module. Defaults to `INFO` if not set.

这些 ENV 都可以被写入到 `webui-user.sh` 中，当 Web UI 启动时会从该文件中读取

eg
```
export COMMANDLINE_ARGS="--allow-code --xformers --skip-torch-cuda-test --no-half-vae --api --ckpt-dir A:\\stable-diffusion-checkpoints"
```

### 0x03b Options

> 这里只列出常用的 Options，具体以及默认值可以看 wiki

- `--ckpt-dir`
	Path to directory with Stable Diffusion checkpoints.
	推荐将所有的 checkpoints 放到 NAS 里
- `--lora-dir`
	Path to directory with Lora networks.
- `--hypernetwork-dir`
	hypernetwork directory.
- `--allow-code`
	Allow custom script execution from web UI.
- `--share`
	Use `share=True` for gradio and make the UI accessible through their site.
	当前的服务可以从 gradio 公网访问
- `--listen`
	Launch gradio with 0.0.0.0 as server name, allowing to respond to network requests.
	监听所有 interface，可以让 LAN 的机器访问到
- `--port`
	Launch gradio with given server port, you need root/admin rights for ports < 1024; defaults to 7860 if available.
- `--gradio-auth`
	Set gradio authentication like `username:password`; or comma-delimit multiple like `u1:p1,u2:p2,u3:p3`.
- `--gradio-auth-path`
	Set gradio authentication file path ex. `/path/to/auth/file` same auth format as `--gradio-auth`.
- `--api`
	Launch web UI with API.
- `--api-auth`
	Set authentication for API like `username:password`; or comma-delimit multiple like `u1:p1,u2:p2,u3:p3`.
- `--tls-keyfile`
	Partially enables TLS, requires `--tls-certfile` to fully function.
	TLS 私钥，建议不要设置 passphase
- `--tls-certfile`
	Partially enables TLS, requires `--tls-keyfile` to fully function.
	TLS 证书
- `--disable-tls-verify`
	When passed, enables the use of self-signed certificates.
	允许使用自己签名证书,
- `--xformers`
	Enable xformers for cross attention layers.
	启用 xformers 时，必须指定
- `--no-half`
	Do not switch the model to 16-bit floats.
- `--no-half-vae `
	Do not switch the VAE model to 16-bit floats.
- `--medvram`
	Enable Stable Diffusion model optimizations for sacrificing a some performance for low VRAM usage.
	针对 low VRAM card 非常有用，牺牲性能以降低 VRAM 的使用
- `--lowvram`
	Enable Stable Diffusion model optimizations for sacrificing a lot of speed for very low VRAM usage.
	针对 low VRAM card 非常有用，牺牲性能以降低 VRAM 的使用

可以在 `webui-user.sh` 添加常用参数

eg
```
export COMMANDLINE_ARGS="--xformers --medvram --no-half --no-half-vae --listen --tls-keyfile tls/stable-diffusion-webui.pem --tls-certfile tls/stable-diffusion-webui.crt --disable-tls-verify"
```

## 0x04 Directory

一些常用的目录如下
- `stable-diffusion-webui/models/Stable-diffusion`
	存放 SD 模型的目录
- `stable-diffusion-webui/models/Lora`
	存放 Lora 的目录

## 0x05 Custom

### 0x05a [Custom Images Filename Name and Subdirectory](https://github.com/AUTOMATIC1111/stable-diffusion-webui/wiki/Custom-Images-Filename-Name-and-Subdirectory)

### 0x05b [Change model folder location](https://github.com/AUTOMATIC1111/stable-diffusion-webui/wiki/Change-model-folder-location)

---

*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]: [Install and Run on NVidia GPUs · AUTOMATIC1111/stable-diffusion-webui Wiki · GitHub](https://github.com/AUTOMATIC1111/stable-diffusion-webui/wiki/Install-and-Run-on-NVidia-GPUs)
[^2]: [Xformers · AUTOMATIC1111/stable-diffusion-webui Wiki · GitHub](https://github.com/AUTOMATIC1111/stable-diffusion-webui/wiki/Xformers)
[^3]:[ImportError: cannot import name 'Undefined' from 'pydantic.fields' (D:\\a1111\\stable-diffusion-webui\\venv\\lib\\site-packages\\pydantic\\fields.py) · AUTOMATIC1111/stable-diffusion-webui · Discussion #15557 · GitHub](https://github.com/AUTOMATIC1111/stable-diffusion-webui/discussions/15557)
[^4]:[\[Bug\]: Using TCMalloc: libtcmalloc.so.4 python3: /home/carlosm/anaconda3/bin/../lib/libstdc++.so.6: version \`GLIBCXX\_3.4.30' not found (required by /usr/lib/libtcmalloc.so.4) · Issue #10208 · AUTOMATIC1111/stable-diffusion-webui · GitHub](https://github.com/AUTOMATIC1111/stable-diffusion-webui/issues/10208)
[^5]:[Command Line Arguments and Settings · AUTOMATIC1111/stable-diffusion-webui Wiki · GitHub](https://github.com/AUTOMATIC1111/stable-diffusion-webui/wiki/Command-Line-Arguments-and-Settings)
