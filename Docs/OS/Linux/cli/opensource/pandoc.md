# pandoc

https://pandoc.org/

https://pandoc.org/getting-started.html

## Digest

> pandoc 也会自动按照文件后缀识别格式，即使没有使用 `-r` 或者 `-w` 来标识

syntax：

`pandoc [options] [input-file]...`

pandoc 是一个 Haskell 写的格式转换工具，也是 Haskell 内嵌的 lib 。支持例如 Markdown 转 norff，Markdown 转 html 等。具体可以参考官网

pandoc 默认会将转换的文件分成 fragment，如果需要将文档转成 standalone ，需要使用 `-s` 或者 `--standalone`

pandoc 默认使用 utf-8 读取和输出，如果文件不是使用 utf-8 格式编写的，需要先使用`iconv`转化，例如

```
iconv -t utf-8 input.txt | pandoc | iconv -f utf-8
```

## Postional args

如果没有指定  input-file，pandoc 就会从 stdin 中读取，然后输出到 stdout

## Optional args

### general args

- `--list-input-formats | output-formats`

  pandoc 支持转换的输入和输出格式

- `-f | -r | --from | --read FORMAT`

  input files 的格式

- `-t | -w | --to | -w FORMAT`

  output files 的格式

- `-o | --output FILE`

  write output to FILE instead of stdout

- `-d FILE`

  FILE contains a set of default options which wrote in YAML

- `--verbose`

- `--log=FILE`

  日志模式

### general write args

- `-s | --standlone`

  instead a single file of fragment

- `--sandbox`

  dry run

- `--toc | --table-of-contents`

  automatically generated table of contents

## Exmaple

将从 stdin 读取的 html 格式的内容，转换成 markdown 并在 stdout 输出。如果在 Linux，需要使用 ctrl + D 发送指定的 POSIX signale 告诉 pandoc 停止并执行转化。如果在 windows 上可以使用 ctrl + Z

```
pandoc -r html -w markdown
<p>Hello <em>pandoc</em>!</p>
<ul>
<li>one</li>
<li>two</li>
</ul>
```

读取 pandoc.md 文件并转成 html 格式输出到 stdout 

```
pandoc -w html pandoc.md
```

可以使用 `-o` 写到指定文件

```
pandoc -w html -o pandoc.html pandoc.md
```

也可以不使用 `-r` 或者 `-w` 来标识文件类型，pandoc 会自动识别

```
pandoc  -o pandoc.html pandoc.md
```



