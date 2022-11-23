# Linux xxd

## Digest

syntax:

```
xxd -h[elp]
xxd [options] [infile [outfile]]
xxd -r[evert] [options] [infile [outfile]]
```

`xxd` 是 `vim` 中包含的命令，用于将文件或者 Stdout 转为 hexadecimal  dump 或者将 hexadecimal dump 转为原始的 binary form

## Examples

- convert content into hexdump

  ```
  [vagrant@localhost ~]$ cat test
  who am i.
  I'm perl.
  [vagrant@localhost ~]$ xxd test
  0000000: 7768 6f20 616d 2069 2e0a 4927 6d20 7065  who am i..I'm pe
  0000010: 726c 2e0a
  ```

- convert hexdump into binary

  ```
  [vagrant@localhost ~]$ cat t.hex
  ELF>J@@�y@8
  [vagrant@localhost ~]$ xxd t.hex | xxd -r
  ELF>J@@�y@8
  ```

  

