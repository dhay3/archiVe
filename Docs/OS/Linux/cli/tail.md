# tail



## 0x01 Overview

syntax

```
tail [OPTION]... [FILE]...
```

`tail` 默认会打印出 the last 10 lines of ==each== FILE

## 0x02 Options

- `-f | --follow`

  输出 append 的内容，可以理解为监控文件的输出

- `-n | --lines=[NUM]`

  输出指定 NUM 行 last line 或者使用 `+NUM` 表示 skip NUM-1 lines(即从第几行开始输出)
