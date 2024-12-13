# Shell tput

参考：

https://www.cnblogs.com/technologylife/p/8275044.html

https://commandnotfound.cn/linux/1/8/tput-%E5%91%BD%E4%BB%A4

tput使用terminfo数据库，对终端会话进行初始化和操作，也可以对shell脚本输出的内容做一定的修改

## 光标属性

在shell脚本或命令行中，可以利用tput命令改变光标属性。

```shell
tput clear      # 清除屏幕
tput sc         # 记录当前光标位置
tput rc         # 恢复光标到最后保存位置
tput civis      # 光标不可见
tput cnorm      # 光标可见
tput cup x y    # 光标按设定坐标点移动
```

利用上面参数编写一个终端时钟

```shell
#!/bin/bash

for ((i=0;i<10;i++))
do
        tput sc; tput civis                     # 记录光标位置,及隐藏光标
        echo -ne $(date +'%Y-%m-%d %H:%M:%S')   # 显示时间
        sleep 1
        tput rc                                 # 恢复光标到记录位置
done

tput el; tput cnorm                             # 退出时清理终端,恢复光标显示
```

## 文本属性

```
tput blink      # 文本闪烁，退出脚本后复原
tput bold       # 文本加粗
tput el         # 清除到行尾
tput smso       # 启动突出模式
tput rmso       # 停止突出模式
tput smul       # 下划线模式
tput rmul       # 取消下划线模式
tput sgr0       # 恢复默认终端
tput rev        # 反相终端
```

此外，还可以改变文本的颜色

```
tput setb 背景色
tput setf 前景色
```

颜色代号为

```
0：黑色
1：蓝色
2：绿色
3：青色
4：红色
5：洋红色
6：黄色
7：白色
```

现在为"终端时钟"添加，变换颜色和闪烁功能

```shell
#!/bin/bash

for ((i=0;i<8;i++))
do
        tput sc; tput civis                     # 记录光标位置,及隐藏光标
        tput blink; tput setf $i                # 文本闪烁,更改文本颜色
        echo -ne $(date +'%Y-%m-%d %H:%M:%S')   # 显示时间
        sleep 1
        tput rc                                 # 恢复光标到记录位置
done

tput el; tput cnorm                             # 退出时清理终端,恢复光标显示
```