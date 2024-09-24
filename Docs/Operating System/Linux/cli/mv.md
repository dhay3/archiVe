# Linux mv

ref:

https://pimylifeup.com/mv-command-linux/

> 可以完全使用 `rsync` 替代

## Digest

syntax:

```
mv [OPTION]... [-T] SOURCE DEST
mv [OPTION]... SOURCE... DIRECTORY
mv [OPTION]... -t DIRECTORY SOURCE...
```

move or rename files

## Optional args

和 `rm` 不同，`mv` 没有递归操作

- `-v | --verbose`

  verbose out

- `-f | --fore`

  do not prompt before overwriting

- `-i | --interactive`

  prompt before overwriting

- `-n | --no-clobber`

  等价于 `set -o noclobber`

### back up args

- `-b`

  在 destination 对 source 文件备份

- `-S | --suffix=SUFFIX`

  使用默认的 suffix 替代默认的`~`

  提一嘴 POSIX 中一般将`~`都是识别成备用文件

## Examples

### 0x01 moving multiple files

file1,file2,file3 都被移动到了`~/example`

```
mv file1.txt file2.txt file3.txt ~/example
```

### 0x02 wildcard

example 下的文件除隐藏文件外 都会被移动到`~`

```
mv example/* ~
```

