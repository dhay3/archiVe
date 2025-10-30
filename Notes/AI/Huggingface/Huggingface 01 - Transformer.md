---
createTime: 2025-10-25 16:28
license: cc by 4.0
tags: 
 - "#hash1" 
 - "#hash2"
---

# Huggingface 01 - Transformer

## 0x01 Preface

在介绍 Transformer 前需要了解 2 个概念

- NLP

	Natrual Language Processing is the broader field focused on enabling computers to understand, interpret, and generate human language. NLP encompasses many techniques and tasks such as sentiment analysis, named entity recognition, and machine translation.

	NLP 就是让计算机理解并解释人类语言这一行为的统称

> NLP 的核心就是 classifying(分词)，将 sentences 拆分然后理解每一个词语

- LLMs

	Large Language Models are a powerful subset of NLP models characterized by their massive size, extensive training data, and ability to perform a wide range of language tasks with minimal task-specific training. Models like the Llama, GPT, or Claude series are examples of LLMs that have revolutionized what’s possible in NLP.

	LLM 是 NLP 的子集，指主要处理语言的模型

## 0x02 Transformer Quick Start

> [!NOTE]
> Transformers is the pivot across frameworks

Transformers 是 Huggingface 提供的一个通用模型接口，可以将其想象成 AI 界的 openAPI，通过 transformers 可以调用各种各样的模型

### 0x02a Installation[^1]

推荐使用 uv 来搭建环境

```
uv init tf
cd tf
uv add transformers torch
```

测试

```
python -c "from transformers import pipeline; print(pipeline('sentiment-analysis')('hugging face is the best'))"
[{'label': 'POSITIVE', 'score': 0.9998704791069031}]
```

## 0x03 Pipeline

transformers 最核心的对象就是 `pipeline()` 函数，可以输入文本调用模型输出结果

一些常用的 pipeline

> [!NOTE]
> 所有可用的 [pipelines](https://huggingface.co/models?sort=trending)
> 
> 调用 pipelines 时会自动下载相关的模型，默认存储在 `~/.cache/huggingface/hub/`

**Text pipelins**

- `text-generation`: Generate text from a prompt
- `text-classification`: Classify text into predefined categories
- `summarization`: Create a shorter version of a text while preserving key information
- `translation`: Translate text from one language to another
- `zero-shot-classification`: Classify text without prior training on specific labels
- `feature-extraction`: Extract vector representations of text

 **Image pipelines**

- `image-to-text`: Generate text descriptions of images
- `image-classification`: Identify objects in an image
- `object-detection`: Locate and identify objects in images

 **Audio pipelines**

- `automatic-speech-recognition`: Convert speech to text
- `audio-classification`: Classify audio into categories
- `text-to-speech`: Convert text to spoken audio

 **Multimodal pipelines**

- `image-text-to-text`: Respond to an image based on a text prompt

### 0x03a Quick start

以 text-generation pipeline 为例，会为输入的嗯文本润色

```
from transformers import pipeline

generator = pipeline("text-generation")
print(generator(
    "In this course, we will teach you how to",
))
```

默认调用的模型会在运行脚本后提示

```
No model was supplied, defaulted to openai-community/gpt2 and revision 607a30d (https://huggingface.co/openai-community/gpt2).
```

，如果想要调用特定的模型可以使用 model 入参

```
from transformers import pipeline

generator = pipeline('text-generation', model='HuggingFaceTB/SmolLM2-360M')
print(
    generator(
        'In this course, we will teach you how to',
    )
)
```

上述代码会返回 text 关联每个标签的比分，sequence 和 education 关联较大，所以 scores 最高

```
{'sequence': 'This is a course about the Transformers library',
 'labels': ['education', 'business', 'politics'],
 'scores': [0.8445963859558105, 0.111976258456707, 0.043427448719739914]}
```

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [Introduction - Hugging Face LLM Course](https://huggingface.co/learn/llm-course/chapter1/1)
- [Transformers](https://huggingface.co/docs/transformers/index)

***References***

[^1]:[Installation](https://huggingface.co/docs/transformers/installation)
