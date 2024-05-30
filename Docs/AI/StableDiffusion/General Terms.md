# General Terms

## Stablediffusion

SD1.5 是 Stablediffusion 1.5 checkpoints 的简称，同理 SD2,SD-XL

## Checkpoint

A checkpoint refers to a file that contains the weights and parameters of a trained neural network model

The key points about Stable Diffusion checkpoints are:

1. They store the trained state of the neural network model at a specific point during the training process,  allowing you to use that model for image generation.
2. They do not contain any actual  image data. The checkpoint only has the numerical weights/parameters  that define the behavior of the neural networks.
3. Different checkpoints can be  trained on different datasets, allowing them to specialize in generating particular styles, subjects, or types of images.
4. You can load different checkpoints  into the Stable Diffusion application to utilize the capabilities of  those trained models for your image generation tasks.
5. Checkpoints use file extensions like .ckpt, .safetensors, etc. to store the model weights in a compact format.

简而言之，可以将 checkpoints 认为是一个模型

## Safesensors

Safetensors in Stable Diffusion are a secure and  efficient format for storing and loading the trained weights and  parameters of the neural network model.

Safesensors 是 checkpoints 的一种安全高效的格式

## LoRA

LoRA(Low Rank Adaptation)

LoRA (Low-Rank Adaptation) is a technique used in Stable  Diffusion to efficiently fine-tune and adapt pre-trained diffusion  models to new concepts, styles, or subjects without modifying the  original large model weights.

可以将 LoRA 认为是 checkpoints 的一个插件，实现模型的微调

## Hypernetwork

A hypernetwork in Stable Diffusion is a technique to  fine-tune and adapt the pre-trained diffusion model to generate images  aligned with specific styles, concepts, or subjects without modifying  the original large model weights.

hypernetwork 功能和逻辑类似 LoRA，但是效率上没有 LoRA 高

**references**

[^1]:https://github.com/AUTOMATIC1111/stable-diffusion-webui/wiki