---
createTime: 2024-07-04 16:39
tags:
  - "#hash1"
  - "#hash2"
---

# SD 03 - WebUI txt2img

## 0x01 Overview

> [!NOTE]
> Wiki[^6] 中提供了一组教学使用文档，如果有文字描述不清楚的可以看 SECourses's Playlist
> hover on button 一般都会有提示


## 0x02 txt2img

txt2img/img2img 大体上参数差不多，区别就在于字面上的含义

![](https://github.com/dhay3/picx-images-hosting/raw/master/20240704/2024-07-04_10-00-18.9dcu5q4pjp.webp)

txt2txt/img2img 都有 2 个 prompt 框分别是
- 想要生成图片的 prompt
- 想要生成图片中不想要的效果的 prompt

在 textarea 下面有 2 个 向下箭头的按钮 可以列出一组按照类别分类的 prompt

#### txt2img

Generation 面板是 txt2img 对 TODO

##### Smapling method/Schedule type/Sampling steps

> [!important]
> 去噪的方法以及调度，要想弄明白具体是什么该怎么选可以看 [Stable Diffusion Samplers: A Comprehensive Guide - Stable Diffusion Art](https://stable-diffusion-art.com/samplers/)

为了方便记忆总结如下

**Smapling method**

- Old-School ODE samplers
	hundred years ago 就被发明的 sampler
	
	- Euler – The simplest possible solver.
	- Heun – A more accurate but slower version of Euler.
	- LMS (Linear multi-step method) – Same speed as Euler but (supposedly) more accurate.

- Ancestral samplers
	在 sampling 的过程中，还会随机增加 noise，就会导致生成的图片不会 converge(收敛，即每次生成的图片都不一样)
	
	==如果为了图片 converge，就不应该使用 Ancestral smaplers==
	
	通常 Ancestral sampler 会包含一个单独的 a （不绝对）
	
	- Euler a
	- DPM2 a
	- DPM++ 2S a
	- DPM++ 2S a Karras

- DDIM/PLMS
	DDIM(Denoising Diffusion Implicit Model)
	PLMS(Pseudo Linear Multi-Step method)
	为 SD1 设计的 sampler，因为 SD 版本迭代的原因，通常不选这两个

- DPM/DPM2/DPM++/DPM adaptive
	- DPM (Diffusion probabilistic model solver)
		为 SD2 设计的 sampler
	- DPM2
		可以生成比 DPM 更加准确的图片
	- DPM++
		对 DPM 的加强
	- DPM adaptive
		可以根据 sampling steps 自动调整？

- UniPC
	UniPC(Unified Predictor-Corrector)
	比较新的 sampler，在 5 - 10 sampling steps 就可以生成高画质的图片


**Schedule type**

- Karras
	生成的图片 noise 少，可以增强图片画质

**Sampling steps**

去噪的步长，步长约长 noise 越少，通常 15 以上几乎看不见噪点

	
##### How to choose samplers[^7]

如果不想看注脚，总结如下

**Image convergence**

1. 如果想要图片收敛，在不考虑图片生成时间的情况下选 DPM adaptive
2. 如果想要图片收敛，且考虑图片生成时间，但是效果不必特别强的情况下，可以选 DPM++ 2M Karras 或者 UniPC 或者 Heun

**Quality**

sampling steps 消耗的时间也越大

1. 如果想要图片质量好，在不考虑 smapling step 的情况下选 DPM2 或者 UniPC
2. 如果想要图片质量好，且 sampling step 少，可以选 DPM++ SDE Karras

有几条规则可控参考

> 1. If you want to use something fast, converging, new, and with decent quality, excellent choices are
>     - **DPM++ 2M Karras** with 20 – 30 steps
>     - **UniPC** with 20-30 steps.
> 2. If you want good quality images and don’t care about convergence, good choices are
>     - **DPM++ SDE Karras** with 10-15 steps (Note: This is a slower sampler)
>     - **DDIM** with 10-15 steps.
> 3. Avoid using any ancestral samplers if you prefer stable, reproducible images.
> 4. **Euler** and **Heun** are fine choices if you prefer something simple. Reduce the number of steps for Heun to save time.



##### Hires.fix

##### Refiner

##### seed

#### img2img



### 0x04b Extras

### 0x04c PNG Info

### 0x04d Checkpoint Merger

### 0x04e Train


---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^6]:[Guides and Tutorials · AUTOMATIC1111/stable-diffusion-webui Wiki · GitHub](https://github.com/AUTOMATIC1111/stable-diffusion-webui/wiki/Guides-and-Tutorials)
[^7]:[Stable Diffusion Samplers: A Comprehensive Guide - Stable Diffusion Art](https://stable-diffusion-art.com/samplers/#Evaluating_samplers)