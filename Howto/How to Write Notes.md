---
createTime: 2024-11-08 10:27
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# How to Write Notes

## 0x01 Preface

> [!IMPORTANT]
> 本文没有具体说明使用的格式都使用 Markdown

PM 中有一个逻辑，一个公司如果需要扩大规模，就需要提高效率，而提高效率就需要标准化。记笔记也一样，Your Konwledge 就是公司，要想扩大 Konwledge 的规模，就需要提高记笔记的效率，而提高效率就需要一个标准的逻辑和结构

在 [How to Ask Questions](How%20to%20Ask%20Questions.md) 中我们把 Ask Questions 这一动作拆分成 3 部分。相似地 Write Notes 也可以拆分成

1. Before You Write
2. When You Write
3. Post You Write

## 0x02 Before You Write

### 0x02a Use Markdown

Markdown[^1] 是一个轻量级的 markup language 可以快捷生成格式化的内容，使用 Markdown 可以大大提高写作的效率

### 0x02b Better Editor

好马还得配好鞍，选择一个称手的 Markdown Editor

推荐

- [NeoVim](https://neovim.io/)
- [Obsidian](https://obsidian.md/)
- [Typora](https://typora.io/)

## 0x03 When You Write

> [!note]
> [How to Ask Questions](How%20to%20Ask%20Questions.md) 中的 Write in Clear 同样适用于本文

这部分主要从结构和格式出发，在 [APA Style](https://apastyle.apa.org/) 的基础上，结合一些自己的实践，按照 HTML 中 header/body/footer 的结构划分

### 0x03a Header

#### Title

Title 指 filename 也是 Note 唯一的 Heading 1，且使用 Title Case Capitalization[^2]

例如 本文的 How to Write Notes

同时 Title 需要精简明了地展示 Main Topic，但是如果 Main Topic 是某个 Subject 下的，且你需要对这个 Subject 记录很多不同的 Topic 时

那么你需要以 `{Subject} {Serial Number} - {Main Topic}` 的格式命名

例如 `CCNP 01 - RIP`，`CCNP 02 - BGP`

#### Front Matter

Front Matter 是一篇 Note 最开始的部分。应该包括文章的 metadata，可以是

- createTime
- modifiedTime
- author
- license
- tags
- etc.

例如

```
---
createTime: 2024-11-08 10:27
license: cc by 4.0
tags:
  - "#Games"
  - "#Emulator"
---
```

### 0x03b Body

#### Heading

> [!NOTE]
> 这里的 `0x` 并不表示 hex decimal，只是一个方便区分 Heading 的标识符

每一个 Heading 对应一个 Sub Topic，从 Heading 2 - Heading 5 中选择，且使用 Title Case Capitalization[^2]

Heading 2 要以 `0x[0-9][0-9] {Sub Topic}` 的格式命名，例如 `0x03 When You Write`

Heading 3 要以 `0x[0-9][0-9][a-z] {Sub Topic}` 的格式命名，例如 `0x03b Body`

Heading 4/5 直接使用 Sub Topic 命名，且要尽量避免使用 Heading 5

例如

```
## 0x01 How to Kill a Mockingbrid

### 0x01a Harper Lee

### 0x01b Stroy Overview

#### Background

#### Main Roles
```

#### Paragraphs

首行无需按照中文的规则缩进 2 个字符，直接顶格书写

例如

```
Cemu 是一个开源的 Wii Emulator
```

paragraphs 之间需要额外的 `\n`

例如

```
你可以在 Cemu 上玩 botw

也可以在 Ryujinx 上玩 botw
```

##### Punctuation

在连续的中文中使用全角符号，而在连续的英文内容中使用半角符号

> [!TIP]
> 可以使用 rime 的 key_binder/bindings 实现快速切换

例如

```
流行的 terminal emulator 有 alacritty, kitty, wezterm, ghost，其中 kitty, ghost 支持在命令行通过 kitty protocol 直接浏览图片 
```

#### Qutations/Callout/Italic/Bold

- 如果内容是引用自 secondary sources 的，且需要强调的，请使用 [Block Quotations](#Block%20Quotations)
- 如果内容不是引用自 secondary sources 的，且需要强调的，请使用 [Callout](#Callout)
- 如果内容是引用自 secondary sources 的，且不需要强调的，请使用 [Italic](#Italic)
- 如果仅仅只需要表示强调(通常在 paragraph 内使用)，请使用 [Bold](#Bold)

##### Block Quotations

Block Quotations 用于**强调**某一段**引用**的内容，需要单独为一段落

例如

```
> Buster Keaton 因为其个性的表情也被称为 "The Great Stone Face"

如果你是一个默片影迷，一定听说过 Buster Keaton，也一定看过 Sherlock Jr. 和 The General
```

此外，使用 Block Quotations 时，不要额外使用 quotation marks

例如下面就是一个错误示例

```
> "Buster Keaton 因为其个性的表情也被称为 "The Great Stone Face""
```

##### Callout

> [!NOTE]
> 在 Obsidian flavor Markdown style 中叫作 callout，而 Github 将其称为 Alerts[^3]

Callout 用于**强调**某一段内容不是来自 secondary source 但是**需要注意**的

例如

```
> ![note]
> Obsidian 支持的 callout types 也和 Github 不同
```

如例子中所述的，Github 支持 5 种

```
> [!NOTE]
> Useful information that users should know, even when skimming content.

> [!TIP]
> Helpful advice for doing things better or more easily.

> [!IMPORTANT]
> Key information users need to know to achieve their goal.

> [!WARNING]
> Urgent info that needs immediate user attention to avoid problems.

> [!CAUTION]
> Advises about risks or negative outcomes of certain actions.
```

而 Obsidian 支持多种 callouts[^4]，为了两者的兼容尽量只使用 Github 支持的格式

##### Italic

Italic 用于某一段**引用**，但是不做强调

例如

```
*Italics used to draw attention to text*
```

Italics 可以在段落内使用，也可以单独用作段落

##### Bold

Bold 只用于强调

例如

```
GNU is an operating system that is free software—that is, it respects users' **freedom**.
```

#### List

list content 和 list title 需要以空行分隔，例如 

```
- list a title

	list content

- list b title

	list content

```

否则 Github flavor Markdown 会错误渲染，下面就是一个错误的实例

```
- list a title
	list content
- list b title
	list content
```

#### Code Block

code block 中的内容如果是代码或者脚本，需要标明，例如

```python
def __run__():
	import asyncio
	runtask()
```

### 0x03c Footer

#### See also

延展阅读

例如

```
- [Daring Fireball: Markdown](https://daringfireball.net/projects/markdown/)
- [Style and Grammar Guidelines](https://apastyle.apa.org/style-grammar-guidelines/)
```

#### References

引用对应的原文地址

例如

```
[^1]:[Daring Fireball: Markdown](https://daringfireball.net/projects/markdown/)
[^2]:[Title case capitalization](https://apastyle.apa.org/style-grammar-guidelines/capitalization/title-case)
```

## 0x04 Post You Write

### 0x04a Proofread

把你的 Notes 想象成需要出版的刊物，在“出版”前需要做校对

- 确保没有 mis-spelling
- 确保没有歧义的地方

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [Daring Fireball: Markdown](https://daringfireball.net/projects/markdown/)
- [Style and Grammar Guidelines](https://apastyle.apa.org/style-grammar-guidelines/)
- [Techniques and Tips for Listening and Note Taking | UNSW Current Students](https://www.student.unsw.edu.au/notetaking-tips)

***References***

[^1]:[Daring Fireball: Markdown](https://daringfireball.net/projects/markdown/)
[^2]:[Title case capitalization](https://apastyle.apa.org/style-grammar-guidelines/capitalization/title-case)
	- **major words:** Nouns, verbs (including linking verbs), adjectives, adverbs, pronouns, and all words of four letters or more are considered major words.
		名词，动词，形容词，副词，代词 首字母大写
	- **minor words:** Short (i.e., three letters or fewer) conjunctions, short prepositions, and all articles are considered minor words.
		连词，介词，冠词 首字母不用大小
[^3]:[Basic writing and formatting syntax - GitHub Docs](https://docs.github.com/en/get-started/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax#alerts)
[^4]:[Callouts - Obsidian Help](https://help.obsidian.md/Editing+and+formatting/Callouts)


