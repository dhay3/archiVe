---
createTime: 2025-10-30 15:32
license: cc by 4.0
tags: 
 - "#hash1" 
 - "#hash2"
---

# Pandoc 01 - Overview

## 0x01 Preface

pandoc 是 haskell 的一个库，也是一个命令行工具，用于将 markup 的文档转为其他格式的文档

例如

Markdonw 转 norff，Markdown 转 html

这里主要介绍命令行工具的使用

## 0x02 CLI

### 0x02a syntax

> [!note]
> 如果没有指定文件 pandoc 默认会从 stdin 读取，如果没有使用 `-o` 参数默认会输出到 stdout

```
pandoc [options] [input-file]...
```

pandoc 默认只会生成文档片段

例如源文件为 markdown

```
## Test

- test1
- test2
- test3
- test4
```

pandoc 转换 html 后并不是一个有效的 HTML 文件

```
$  pandoc -w html test.md
<h2 id="test">Test</h2>
<ul>
<li>test1</li>
<li>test2</li>
<li>test3</li>
<li>test4</li>
</ul>
```

如果想要生成一个有效的文档（standalone document），需要配合使用 `-s` 或者 `--standalone` 会输出 header 以及 footer

```
pandoc -s -w html test.md
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml" lang="" xml:lang="">
<head>
  <meta charset="utf-8" />
  <meta name="generator" content="pandoc" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0, user-scalable=yes" />
  <title>test</title>
</head>
<body>
<h2 id="test">Test</h2>
<ul>
<li>test1</li>
<li>test2</li>
<li>test3</li>
<li>test4</li>
</ul>
</body>
</html>
```

此外还需要注意的是，pandoc 默认使用 utf-8 读取和输出，如果文件不是使用 utf-8 格式编写的，需要先使用`iconv`转化，例如

```
iconv -t utf-8 input.txt | pandoc
```

### 0x02b Specifying Formats

pandoc 可以通过 `-f/--from` 和 `-t/--to` 来指定 input-file 和 output-file 的格式，所有可用的 input-file 格式可以参考 `--list-input-formats`，所有可用的 output-file 格式可以参考 `--list-output-formats`

当没有指定 input-file 和 output-file 的格式，pandoc 会根据文件后缀名来猜测（如果没有指定文件后缀默认为 markdown），例如

```
pandoc -o hello.tex hello.md
```

pandoc 就会推测 input-file 格式为 markdown，需要将其转为 latex 格式 output-file

### 0x02c Optional Args

#### General Args

- `-f | -r | --from | --read FORMAT`

	input-files 的格式

- `-t | -w | --to | --write FORMAT`

	output-files 的格式

- `-o | --output FILE`

	将 output-files 写入文件，而不是输出到 stdout


#### General Write Args

- `-s | --standlone`

	为文档生成有效的 header 和 footer

- `--sandbox`

	dry run

- `--toc | --table-of-contents`

	自动生成 TOC，需要和 `-s` 一起使用


## 0x03 Quick start

将 test.md 转成 html 格式

```
pandoc test.md -o test.html
```

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [Title Unavailable \| Site Unreachable](https://pandoc.org/MANUAL.html)


***References***


