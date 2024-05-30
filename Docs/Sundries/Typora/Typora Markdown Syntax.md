---

layout: post
title: Blogging Like a Hacker
---
# Typora Markdown Syntax

## 0x01 Overview

Typora Markdown 语法大体上和 [GitHub Flavored Markdown](https://help.github.com/articles/github-flavored-markdown/) 类似，但有细微不同

## 0x02 Syntax

### TOC

目录

#### Syntax

```
[TOC]
```

#### Display

[TOC]

### YAML Frontmatter

一般用在博客,必须声明在文档的开头

#### Syntax

```
---
```

#### Display

看文档的开头

### Headings

标题

快捷键：<kbd>ctrl</kbd> + <kbd>n</kbd>

#### Syntax

```
# 一级标题 ctrl + 1

## 二级标题 ctrl + 2

### 三级标题 ctrl + 4

... 以此类推最多到 六级标题 ctrl + 6
```

#### Display

# 一级标题

## 二级标题

### 三级标题

### FontStyles

文字格式

#### Syntax

```
<u>下划线</u>  ctrl + u

*斜体*  ctrl + i

**加粗**  ctrl + b

~~删除线~~ 

==高亮==
```

#### Display

<u>下划线</u> 

*斜体* 

**加粗** 

~~删除线~~  

==高亮==

### Subscript

下标

#### Syntax

```
H~2~O
```

#### Display

H~2~O

### Superscript

上标

#### Syntax

```
X^2^
```

#### Display

X^2^

### Image

图片

快捷键：<kbd>ctrl</kbd> + <kbd>shift</kbd> + <kbd>i</kbd>

#### Syntax

```
<img src=""/>
```

#### Display

![]()

### Links

链接

快捷键：<kbd>ctrl</kbd> + <kbd>k</kbd>

#### Syntax

```
[链接](https://youtube.com) 
```

#### Display

[链接](https://youtube.com) 

### Blockquotes

引用

#### Syntax

```
> 文字引用  ctrl + shift +q
```

#### Display

> 文字引用

### Hrizontal Rules

分隔线

#### Syntax

```
---
```

#### Display

---

### List

列表

#### Syntax

```
# 无序列表
- order1 ctrl + shift + ]
- order2
  - order 2.1
  - order 2.2
# 有序列表
1. order1  ctrl + shift + [
2. order2
```

#### Display

- order1 
- order2 
  - order 2.1
  - order 2.2

1. order1
2. order2

### Task List

任务列表

#### Syntax

```
- [ ] to-do 
- [X] done
```

#### Display

- [ ] to-do 
- [x] done

### Code Block

代码块,还支持 mermaidjs

快捷键：<kbd>ctrl</kbd> + <kbd>shift</kbd> + <kbd>k</kbd>

#### Syntax

```
```javascript
function test() {
  console.log("notice the blank line before this function?");
}
```  ctrl + shift + k
```

#### Display

```javascript
function test() {
  console.log("notice the blank line before this function?");
}
```

### inline Code

行内代码

#### Syntax

```
`print()`
```

#### Display

`print()`

### Math Block

数学公式块

快捷键：<kbd>ctrl</kbd> + <kbd>shift</kbd> + <kbd>m</kbd>

#### Syntax

```
$$
\mathbf{V}_1 \times \mathbf{V}_2 =  \begin{vmatrix}
\mathbf{i} & \mathbf{j} & \mathbf{k} \\
\frac{\partial X}{\partial u} &  \frac{\partial Y}{\partial u} & 0 \\
\frac{\partial X}{\partial v} &  \frac{\partial Y}{\partial v} & 0 \\
\end{vmatrix}
$$
```

#### Display

$$
\mathbf{V}_1 \times \mathbf{V}_2 =  \begin{vmatrix}
\mathbf{i} & \mathbf{j} & \mathbf{k} \\
\frac{\partial X}{\partial u} &  \frac{\partial Y}{\partial u} & 0 \\
\frac{\partial X}{\partial v} &  \frac{\partial Y}{\partial v} & 0 \\
\end{vmatrix}
$$

### inline Math

行内数学公式

#### Syntax

```
$a\div\b=c$
```

#### Display

$a \div b=c$

### Tables

表格

快捷键：<kbd>ctrl</kbd> + <kbd>t</kbd>

#### Syntax

```
|A|A|A|
|-|-|-|
| | | |
```

#### Display

| A    | A    | A    |
| ---- | ---- | ---- |
|      |      |      |

### Reference Links

引用链接

#### Syntax

```
This is [an example][id] reference-style link.

Then, anywhere in the document, you define your link label on a line by itself like this:

[id]: http://example.com/  "Optional Title Here"
```

#### Display

This is [an example][id] reference-style link.

Then, anywhere in the document, you define your link label on a line by itself like this:

[id]: http://example.com/  "Optional Title Here"

### HTML

#### Syntax

```
<font color="red" size=7 face="黑体">html代码</font>

<font face= "微软雅黑">我是微软雅黑</font>

<table><tr><td bgcolor=orange>背景色是：orange</td></tr></table>
```

#### Display

<font color="red" size=7 face="黑体">html代码</font>

<font face= "微软雅黑">我是微软雅黑</font>

<table><tr><td bgcolor=orange>背景色是：orange</td></tr></table>

**references**

[typora]: https://support.typora.io/Markdown-Reference/