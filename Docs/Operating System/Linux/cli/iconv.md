# iconv

## Digeset

syntax:

```
iconv [options] [-f from-encoding] [-t to-encoding] [inputfile]...
```

the `iconv` program reads in text in one encoding and outputs the text in another encoding

`iconv` 可以将文件的一个 encode 转成指定的 encode

## Optional args

- `-f | --from-code from-encoding`

  use from-encoding for output characters

- `-t | --to-code to-encoding`

  use to-encoding for output characters

  if the string `//IGNORE` is appended to to-encoding, characters that cannot be converted are discarded and an error is printed after conversion

  if the string `//TRANSLIT`（转义） is appended to to-encoding, characters that cannot be converted are transliterated when needed and posible

- `-o | --output outputfile`

  输出到指定文件

- `-c`

  silently discard character set encodings

## Example

将 stdin 重定向到 input.txt 转换成 utf-8 输出到 output.txt

```
$ iconv -f ISO-8859-15 -t UTF-8 < input.txt > output.txt
```

