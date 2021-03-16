# Linux zero文件

`/dev/zero`，会提供无限的空字符(NULL，0x00)，可以使用该文生成一个特定大小的空白文件。

```
dd if=/dev/zero of=foobar count=1024 bs=1024
```