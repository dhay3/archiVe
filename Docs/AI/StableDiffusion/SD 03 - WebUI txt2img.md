---
createTime: 2024-07-04 16:39
tags:
     - AI
     - Stable-Diffusion
---

# SD 03 - WebUI txt2img

> [!NOTE]
> Wiki[^6] 中提供了一组教学使用文档，如果有文字描述不清楚的可以看 SECourses's Playlist
> 
> hover on button or text 一般都会有提示

## 0x01 Overview

txt2img 顾名思义就是文本转图片

## 0x02 Main

先看主面板即(txt2img)

![](https://github.com/dhay3/picx-images-hosting/raw/master/20240704/2024-07-04_10-00-18.9dcu5q4pjp.webp)

主要由 2 部分组成

**Prompt part**
- 第 1 个 textarea 用于填写想要生成图片的 prompt
- 第 2 个 textarea 用于填写想要生成图片中不想要效果的 prompt

在 textarea 下面有 2 个 向下箭头的按钮 可以列出一组按照类别分类的 prompt


**Generate**


## 0x03 Generation

generation 面板是对图片的一些预设参数，对图片的生成起到至关重要的影响

### 0x03a Sampling method/Schedule type/Sampling steps

> [!important]
> 去噪的方法以及调度，要想弄明白具体是什么该怎么选可以看 [Stable Diffusion Samplers: A Comprehensive Guide - Stable Diffusion Art](https://stable-diffusion-art.com/samplers/)

Sampling method/Schedule type/Sampling steps 决定了图片的 quality(清晰度) 以及 convergence(收敛度，图片收敛度高则生成的图片类似，图片收敛度底则生成的图片差异大)

为了方便记忆总结如下

#### Sampling method

- **Old-School ODE samplers**
	hundred years ago 就被发明的 samplers
	
	- Euler – The simplest possible solver.
	- Heun – A more accurate but slower version of Euler.
	- LMS (Linear multi-step method) – Same speed as Euler but (supposedly) more accurate.

- **Ancestral samplers**
	在 sampling 的过程中，还会随机增加 noise，就会导致生成的图片收敛度低
	
	==如果为了图片 converge，就不应该使用 Ancestral smaplers==
	
	通常 Ancestral samplers 名称中会包含一个单独的 a （不绝对）
	
	- Euler a
	- DPM2 a
	- DPM++ 2S a
	- DPM++ 2S a Karras

- **DDIM/PLMS**
	DDIM(Denoising Diffusion Implicit Model)
	PLMS(Pseudo Linear Multi-Step method)
	为 SD1 设计的 sampler，因为 SD 版本迭代的原因，通常不选这两个

- **DPM/DPM2/DPM++/DPM adaptive**
	- DPM (Diffusion probabilistic model solver)
		为 SD2 设计的 sampler
	- DPM2
		可以生成比 DPM 更加准确的图片
	- DPM++
		对 DPM 的加强
	- DPM adaptive
		可以根据 sampling steps 自动调整？

- **UniPC**
	UniPC(Unified Predictor-Corrector)
	比较新的 sampler，在 5 - 10 sampling steps 就可以生成高画质的图片

#### Schedule type

- Karras
	生成的图片 noise 少，可以增强图片画质

#### Sampling steps

去噪的步长，步长约长 noise 越少，但是消耗的时间也越大，通常 15 以上几乎看不见噪点

#### **How to choose samplers**[^7]

具体可以看注脚，如果不想看注脚，总结如下

有几条规则可供参考

> 1. If you want to use something fast, converging, new, and with decent quality, excellent choices are
>     - **DPM++ 2M Karras** with 20 – 30 steps
>     - **UniPC** with 20-30 steps.
> 2. If you want good quality images and don’t care about convergence, good choices are
>     - **DPM++ SDE Karras** with 10-15 steps (Note: This is a slower sampler)
>     - **DDIM** with 10-15 steps.
> 3. Avoid using any ancestral samplers if you prefer stable, reproducible images.
> 4. **Euler** and **Heun** are fine choices if you prefer something simple. Reduce the number of steps for Heun to save time.

### 0x03c Hires.fix

### Refiner

### seed


---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^6]:[Guides and Tutorials · AUTOMATIC1111/stable-diffusion-webui Wiki · GitHub](https://github.com/AUTOMATIC1111/stable-diffusion-webui/wiki/Guides-and-Tutorials)
[^7]:[Stable Diffusion Samplers: A Comprehensive Guide - Stable Diffusion Art](https://stable-diffusion-art.com/samplers/#Evaluating_samplers)