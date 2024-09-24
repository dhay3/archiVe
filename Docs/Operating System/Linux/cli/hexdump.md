# Linux hexdump

ref:

https://opensource.com/article/19/8/dig-binary-files-hexdump

https://stackoverflow.com/questions/23416762/commandline-hexdump-with-ascii-output

## Digest

> `hexdump` 不支持将 hexadecimal 内容转为 ascii
>
> 可以通过一定条件使用`xxd`来解析

syntax

```
hexdump options file ...
hd options file ...
```

以 hexadecimal( 16进制 ), decimal( 10进制 ), octal( 8进制 ) 或者 ascii 显示文件内容或者是 Stdout

`hexdump` 按照进制也可以细分成如下几个命令

1. `od` for octal
2. `xxd` for hexadecimal (安装了 `vim` 一般都含有该命令)

## Output

> 可以使用 `man ascii` 对照字符分析

`hexdump` 一般输出分为两部分，offset 和 interpret content

```
[vagrant@localhost ~]$ hexdump -b test
0000000 167 150 157 040 141 155 040 151 056 012 111 047 155 040 160 145
0000010 162 154 056 012                                                
0000014
```

offset 一般以 hexadecimal 显示，即例子中左半边部分 `0000000`, `0000010`

interpret content 为内容解析成指定进制后或格式的输出

## Optional args

需要注意的一点是 `hexdump` 不加参数的时候会以倒叙的 16进制输出，如果需要按照正常的顺序输出可以使用 `-C` 参数

- `-b | --one-byte-octal`

  将内容解析成 8 进制格式显示

- `-d | --two-bytes-decimal`

  将内容解析成 10 进制格式显示

- `-c | --one-byte-char`

  将内容解析成  ascii 格式显示

- `-C | --canonical`

  将内容解析成 16 进制格式显示，同时显示 ascii 内容

- `-n | --length lenght`

  只解析 length 长度的内容(长度的单位可以是 KB,MB,KiB,MiB,etc..)，一般用于校验文件格式是否符合文件头

  ```
  $ hexdump --length 8 pixel.png
  0000000 5089 474e 0a0d 0a1a
  0000008
  ```

- `-s | --skip offset`

  跳过 offset 长度的内容(长度的单位可以是 KB,MB,KiB,MiB,etc..)

