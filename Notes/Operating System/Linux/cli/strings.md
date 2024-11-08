# strings

## 0x01 Overview

syntax

```
strings [-afovV] [-min-len]
[-n min-len] [--bytes=min-len]
[-t radix] [--radix=radix]
[-e encoding] [--encoding=encoding]
[-U method] [--unicode=method]
[-] [--all] [--print-file-name]
[-T bfdname] [--target=bfdname]
[-w] [--include-all-whitespace]
[-s] [--output-separator sep_string]
[--help] [--version] file...
```

`strings` 是 linux 一个用于打印出可识别文字的工具，常用于输出二进制文件中的文本信息

## 0x02 Examples

```
strings /usr/bin/echo
```

