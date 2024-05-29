# Obsidian 02 - Markdown Syntax

Obsidian 使用的 Markdown 语法整体和 [Daring Fireball][1] 的类似, 这里只记录不同或者是模糊的语法

### Fontstyles

文字格式

| Style                  | Syntax                 | Example                                  | Output                                 |
| ---------------------- | ---------------------- | ---------------------------------------- | -------------------------------------- |
| Bold                   | `** **` or `__ __`     | `**Bold text**`                          | **Bold text**                          |
| Italic                 | `* *` or `_ _`         | `*Italic text*`                          | *Italic text*                          |
| Strikethrough          | `~~ ~~`                | `~~Striked out text~~`                   | ~~Striked out text~~                   |
| Highlight              | `== ==`                | `==Highlighted text==`                   | ==Highlighted text==                   |
| Bold and nested italic | `** **` and `_ _`      | `**Bold text and _nested italic_ text**` | **Bold text and _nested italic_ text** |
| Bold and italic        | `*** ***` or `___ ___` | `***Bold and italic text***`             | ***Bold and italic text***             |

### [Internal Link][4]

用于超链接自己写的文档或者文档中的内容，区别与 Externel

#### Syntax

支持 2 种格式

```
Wikilink: [[Three laws of motion]]
Markdown: [Three laws of motion](Three%20laws%20of%20motion.md)
```

如果想要链接特定文档中的特定 heading 可以使用如下格式，在输入 `#` 时会自动提示可以链接的 heading

```
[[Three laws of motion#Second law]].
```

如果想要链接特定文档中的 blockquote 可以使用如下格式，在输入 `^` 时会自动提示可以链接的 blockquote

```
[[2023-01-01#^quote-of-the-day]].
```

默认链接显示的文本就是 Internal Link 中输入的内容，可以通过如下格式修改文本

```
[[Internal links|custom display text]]
```

### [Embed Files][5]

用于引用(直接插入到当前文档)自己写的文档、文档中的内容，或者 Vault 下内容

#### Syntax

就是在 Internal Link 的基础上添加了 exclamation

```
![[Three laws of motion]]
```

如果想要引用特定文档中的特定 heading 可以使用如下格式，在输入 `#` 时会自动提示可以链接的 heading

```
[[Three laws of motion#Second law]].
```

如果想要引用特定文档中的 blockquote 可以使用如下格式，在输入 `^` 时会自动提示可以链接的 blockquote

```
[[2023-01-01#^quote-of-the-day]].
```

除此外还可以引入图片(图片必须在 Vault 下)

```
![[Engelbart.jpg]]
```

可以如下方式来修改图片大小，如果没有指定 height，就会按照比例缩放图片

```
![[Engelbart.jpg|100x145]]
![[Engelbart.jpg|100]]
```

同样的可以引入音频，具体所有支持的格式查看 [Accepted file formats][6]

```
![[Excerpt from Mother of All Demos (1968).ogg]]
```

### Embed web pages

通过 HTML iframe 可以直接将网页插入，但是大多数站点不支持

#### Syntax

```
<iframe src="INSERT YOUR URL HERE"></iframe>
```

### Embed video

和 external link 的语法相同，可以将视频直接插入

#### Syntax

```
![](https://www.youtube.com/watch?v=NnTvZWp5Q7o)
```

### [Callouts][7]

特殊含义的 blockquote

#### Syntax

在 blockquote 的第一行添加 `[!xxx]`, xxx 的值具体参考 [Callouts Supported types][3]

```
> [!info]
> Here's a callout block.
> It supports **Markdown**, [[Internal link|Wikilinks]], and [[Embed files|embeds]]!
> ![[Engelbart.jpg]]
```

### Footnotes

注脚,这里和 Typora 中的格式完全不相同

#### Syntax

```
This is a simple footnote[^1].

[^1]: This is the referenced text.
[^2]: Add 2 spaces at the start of each new line.
  This lets you write footnotes that span multiple lines.
[^note]: Named footnotes still appear as numbers, but can make it easier to identify and link references.
```

### Tags

标签，用于 Search Pulgin

#### Syntax

直接在 Hash 后面带上 tags 即可

```
#DevOps
```

### Comments

注释，只会在 Editing view 中显示

#### Syntax

```
This is an %%inline%% comment.

%%
This is a block comment.

Block comments can span multiple lines.
%%
```

### [Frontmatter][8]

Frontmatter 在 Obsidian 中被称为 Properties, 同样必须出现在文档的开头

通用模板如下

```
---
title: "tensorflow"
Author: "0x00"
createTime: 2024-05-27
lastModifiedTime: 2024-05-29
draft: true
Tags: 
 - "Linux"
 - "Python"
---
```

**referneces**

[1]:https://daringfireball.net/projects/markdown/syntax
[2]:https://help.obsidian.md/
[3]: https://help.obsidian.md/Editing+and+formatting/Callouts#Supported+types
[4]:https://help.obsidian.md/Linking+notes+and+files/Internal+links
[5]:https://help.obsidian.md/Linking+notes+and+files/Embed+files
[6]:https://help.obsidian.md/Files+and+folders/Accepted+file+formats
[7]:https://help.obsidian.md/Editing+and+formatting/Callouts
[8]:https://help.obsidian.md/Editing+and+formatting/Properties