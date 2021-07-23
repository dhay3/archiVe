# DISPLAY

参考：

https://askubuntu.com/questions/432255/what-is-the-display-environment-variable

DISPLAY环境变量是有x window system定义的，a display通常由a keyboard，a mouse，a screen组成，由 X server管理。X client 通过X network将graphical通过DISPLAY发送给X server。

DISPLAY环境变量通常由如下几部分组成

```
hostname:D.S
```

- hostname表示运行X server的hostname，如果没有指明表示当前主机
- D表示display的序列号，一般为0，如果有多个display接入到主机不一定是0
- S表示screen number，表示显示屏的number，如果没有指明表示0

```
cpl in ~ λ echo $DISPLAY
:0
```

如果DISPLAY的值为空，==表明你系统当前可能不存在display设备（云主机）==