# Linux md5sum

ref

https://www.ruanyifeng.com/blog/2019/11/hash-sum.html

## Digest

syntax:

```
md5sum [OPTION]... [FILE]...
```

`md5sum` 使用 md5 algo 生成 md5 加密字符串 或者 用来校验 md5 签名

## Optional args

- `-c | --check`

  从文件中读取 checksums 并校验

- `--quiet`

  如果校验通过不输出 OK

## Examples

- 从 Stdin 生成 md5 值

  ```
  [vagrant@localhost ~]$ echo 123 | md5sum 
  ba1f2511fc30423bdbb183fe33f3dd0f  -
  ```

- 校验 md5

  这里需要注意的一点是 md5 签名文件中的内容必须包含 md5值 和 文件名

  ```
  [vagrant@localhost ~]$ md5sum test > testmd5.txt
  [vagrant@localhost ~]$ md5sum -c testmd5.txt 
  test: OK
  [vagrant@localhost ~]$ cat testmd5.txt 
  e19c1283c925b3206685ff522acfe3e6  test
  ```

  

