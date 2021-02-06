# bat脚本 显示与注释

参考：http://docs.30c.org/dosbat/chapter02/index.html

## 显示

- echo

  echo on缺省值，默认回显

  echo off，关闭回显

  如果需要关闭echo off 命令回显，请使用`@echo off`

  `>`, `>>`同linux中命令相同

- title

  更改bat调用Dos窗口的标题

- @

  不对命令回显

- pause

  运行完脚本后，默认不会挂起。使用pause命令挂起当前窗口。

## 注释

- `rem`

  与Java和C中的注释类似，但是如果没有关闭回显，运行脚本后会显示

- `::`

  与`rem`类似，但是不会在运行脚本后显示