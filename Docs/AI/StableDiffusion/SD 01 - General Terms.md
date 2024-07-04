---
createTime: 2024-07-04 10:16
tags:
  - "#hash1"
  - "#hash2"
---

# SD 01 - General Terms

## 0x01 SD

SD 是 Stablediffusion 的简写，大多数模型或者是插件都会以 SD 命名，例如

SD1.5 是 Stablediffusion 1.5 checkpoints 的简称，同理 SD2,SD-XL

## 0x02 Checkpoint

> A checkpoint refers to a file that contains the weights and parameters of a trained neural network model
> 
> The key points about Stable Diffusion checkpoints are:
> 
> 1. They store the trained state of the neural network model at a specific point during the training process,  allowing you to use that model for image generation.
> 2. They do not contain any actual  image data. The checkpoint only has the numerical weights/parameters  that define the behavior of the neural networks.
> 3. Different checkpoints can be  trained on different datasets, allowing them to specialize in generating particular styles, subjects, or types of images.
> 4. You can load different checkpoints  into the Stable Diffusion application to utilize the capabilities of  those trained models for your image generation tasks.
> 5. Checkpoints use file extensions like .ckpt, .safetensors, etc. to store the model weights in a compact format.

简而言之，可以将 checkpoints 认为是一个模型

## 0x03 Safesensors

> Safetensors in Stable Diffusion are a secure and  efficient format for storing and loading the trained weights and  parameters of the neural network model.

Safesensors 是 checkpoints 的一种安全高效的格式

## 0x04 LoRA

LoRA(Low Rank Adaptation)

> LoRA (Low-Rank Adaptation) is a technique used in Stable  Diffusion to efficiently fine-tune and adapt pre-trained diffusion  models to new concepts, styles, or subjects without modifying the  original large model weights.

可以将 LoRA 认为是 checkpoints 的一个插件，实现模型的微调

## 0x05 Hypernetwork

> A hypernetwork in Stable Diffusion is a technique to  fine-tune and adapt the pre-trained diffusion model to generate images  aligned with specific styles, concepts, or subjects without modifying  the original large model weights.

hypernetwork 功能和逻辑类似 LoRA，但是效率上没有 LoRA 高

## 0x05 Sampling[^2]

> This **denoising process** is called **sampling** because Stable Diffusion generates a new sample image in each step. The method used in sampling is called the **sampler** or **sampling method**.

简单的说 Sampling 就是 denoising process(去噪的过程)，SD 通过多次 Sampling 生成一张清晰的照片

![](https://stable-diffusion-art.com/wp-content/uploads/2022/12/image-84.png)

## 0x06 Schedule[^2]

> You must have noticed the noisy image gradually turns into a clear one. The **noise schedule** controls the **noise level at each sampling step**. The noise is highest at the first step and gradually reduces to zero at the last step.

Smapling 是一个渐进的过程，而 Schedule 就是控制每一步 denoising 的调度器


---

*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:https://github.com/AUTOMATIC1111/stable-diffusion-webui/wiki
[^2]:[Fetching Title#luge](https://stable-diffusion-art.com/samplers/)