# Linux modprobe

## digest

> kernel 在2.4 之前的不支持modprobe

syntax：`modprobe [options] [modulename][module parameters]`

modprobe 用于从linux kernel添加或删除mod，为了方便模块名中的`-`和`_`无区别。

modprobe 默认会读取`/lib/modules/$(uname -r)`下的所有mod 和 文件。除了`modprobe.d`下的可选配置文件。

modprobe不处理module是否正确，所以module可能会有问题。这是可以使用`dmesg`来查看ring buffer中kernel的信息

modprobe使用`modules.dep.bin`文件来查找module的依赖